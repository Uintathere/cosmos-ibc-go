package keeper

import (
	"fmt"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/ibc-go/v3/modules/apps/nft-transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	coretypes "github.com/cosmos/ibc-go/v3/modules/core/types"
)

// SendTransfer handles nft-transfer sending logic.
// A sending chain may be acting as a source or sink zone.
//
// when a chain is sending tokens across a port and channel which are
// not equal to the last prefixed port and channel pair, it is acting as a source zone.
// when tokens are sent from a source zone, the destination port and
// channel will be prefixed onto the classId (once the tokens are received)
// adding another hop to the tokens record.
//
// when a chain is sending tokens across a port and channel which are
// equal to the last prefixed port and channel pair, it is acting as a sink zone.
// when tokens are sent from a sink zone, the last prefixed port and channel
// pair on the classId is removed (once the tokens are received), undoing the last hop in the tokens record.
//
// For example, assume these steps of transfer occur:
// A -> B -> C -> A -> C -> B -> A
//
//|                    sender  chain                      |                       receiver     chain              |
//| :-----: | -------------------------: | :------------: | :------------: | -------------------------: | :-----: |
//|  chain  |                    classID | (port,channel) | (port,channel) |                    classID |  chain  |
//|    A    |                   nftClass |    (p1,c1)     |    (p2,c2)     |             p2/c2/nftClass |    B    |
//|    B    |             p2/c2/nftClass |    (p3,c3)     |    (p4,c4)     |       p4/c4/p2/c2/nftClass |    C    |
//|    C    |       p4/c4/p2/c2/nftClass |    (p5,c5)     |    (p6,c6)     | p6/c6/p4/c4/p2/c2/nftClass |    A    |
//|    A    | p6/c6/p4/c4/p2/c2/nftClass |    (p6,c6)     |    (p5,c5)     |       p4/c4/p2/c2/nftClass |    C    |
//|    C    |       p4/c4/p2/c2/nftClass |    (p4,c4)     |    (p3,c3)     |             p2/c2/nftClass |    B    |
//|    B    |             p2/c2/nftClass |    (p2,c2)     |    (p1,c1)     |                   nftClass |    A    |
//
func (k Keeper) SendTransfer(
	ctx sdk.Context,
	sourcePort,
	sourceChannel,
	classID string,
	tokenIDs []string,
	sender sdk.AccAddress,
	receiver string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	// See spec for this logic: https://github.com/cosmos/ibc/blob/master/spec/app/ics-721-nft-transfer/README.md#packet-relay
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	class, exist := k.nftKeeper.GetClass(ctx, classID)
	if !exist {
		return sdkerrors.Wrap(types.ErrInvalidClassID, "classId not exist")
	}

	var tokenURIs []string
	for _, tokenID := range tokenIDs {
		tokenURI, err := k.handleToken(ctx,
			sourcePort,
			sourceChannel,
			classID,
			tokenID,
			sender,
		)
		if err != nil {
			return err
		}
		tokenURIs = append(tokenURIs, tokenURI)
	}

	labels := []metrics.Label{
		telemetry.NewLabel(coretypes.LabelDestinationPort, destinationPort),
		telemetry.NewLabel(coretypes.LabelDestinationChannel, destinationChannel),
	}

	packetData := types.NewNonFungibleTokenPacketData(
		classID, class.GetUri(), tokenIDs, tokenURIs, sender.String(), receiver,
	)

	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ics4Wrapper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	defer func() {
		telemetry.SetGaugeWithLabels(
			[]string{"tx", "msg", "ibc", "nft-transfer"},
			float32(len(tokenIDs)),
			[]metrics.Label{telemetry.NewLabel("class_id", class.GetUri())},
		)

		telemetry.IncrCounterWithLabels(
			[]string{"ibc", types.ModuleName, "send"},
			1,
			labels,
		)
	}()

	return nil
}

// OnRecvPacket processes a cross chain fungible token transfer. If the
// sender chain is the source of minted tokens then vouchers will be minted
// and sent to the receiving address. Otherwise if the sender chain is sending
// back tokens this chain originally transferred to it, the tokens are
// unescrowed and sent to the receiving address.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, data types.NonFungibleTokenPacketData) error {

	// defer func() {
	// 	if transferAmount.IsInt64() {
	// 		telemetry.SetGaugeWithLabels(
	// 			[]string{"ibc", types.ModuleName, "packet", "receive"},
	// 			float32(transferAmount.Int64()),
	// 			[]metrics.Label{telemetry.NewLabel(coretypes.LabelDenom, data.Denom)},
	// 		)
	// 	}

	// 	telemetry.IncrCounterWithLabels(
	// 		[]string{"ibc", types.ModuleName, "receive"},
	// 		1,
	// 		append(
	// 			labels, telemetry.NewLabel(coretypes.LabelSource, "false"),
	// 		),
	// 	)
	// }()

	return nil
}

// OnAcknowledgementPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain. If the acknowledgement
// was a success then nothing occurs. If the acknowledgement failed, then
// the sender is refunded their tokens using the refundPacketToken function.
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, data types.NonFungibleTokenPacketData, ack channeltypes.Acknowledgement) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return k.refundPacketToken(ctx, packet, data)
	default:
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	}
}

// OnTimeoutPacket refunds the sender since the original packet sent was
// never received and has been timed out.
func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, data types.NonFungibleTokenPacketData) error {
	return k.refundPacketToken(ctx, packet, data)
}

// refundPacketToken will unescrow and send back the tokens back to sender
// if the sending chain was the source chain. Otherwise, the sent tokens
// were burnt in the original send so new tokens are minted and sent to
// the sending address.
func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.NonFungibleTokenPacketData) error {
	return nil
}

// handleToken will escrow the tokens to escrow account
// if the token was away from origin chain . Otherwise, the sent tokens
// were burnt in the sending chain and will unescrow the token to receiver
// in the destination chain
func (k Keeper) handleToken(ctx sdk.Context,
	sourcePort,
	sourceChannel,
	classID,
	tokenID string,
	sender sdk.AccAddress,
) (string, error) {
	nft, exist := k.nftKeeper.GetNFT(ctx, classID, tokenID)
	if !exist {
		return "", sdkerrors.Wrap(types.ErrInvalidTokenID, "tokenId not exist")
	}

	owner := k.nftKeeper.GetOwner(ctx, classID, tokenID)
	if !sender.Equals(owner) {
		return "", sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not token owner")
	}

	isAwayFromOrigin := k.isAwayFromOrigin(sourcePort, sourceChannel, classID)
	if !isAwayFromOrigin {
		return nft.GetUri(), k.nftKeeper.Burn(ctx, classID, tokenID)
	}

	// create the escrow address for the tokens
	escrowAddress := types.GetEscrowAddress(sourcePort, sourceChannel)
	if err := k.nftKeeper.Transfer(ctx, classID, tokenID, escrowAddress.String()); err != nil {
		return "", err
	}
	return nft.GetUri(), nil
}

func (k Keeper) isAwayFromOrigin(sourcePort, sourceChannel, classID string) bool {
	prefixClassID := fmt.Sprintf("%s/%s", sourcePort, sourceChannel)
	return classID[:len(prefixClassID)] != prefixClassID
}
