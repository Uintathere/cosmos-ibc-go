package client_test

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	client "github.com/cosmos/ibc-go/v3/modules/core/02-client"
	"github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

type ClientTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
}

func (suite *ClientTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)

	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) TestBeginBlocker() {
	for i := 0; i < 10; i++ {
		// increment height
		suite.coordinator.CommitBlock(suite.chainA, suite.chainB)

		suite.Require().NotPanics(func() {
			client.BeginBlocker(suite.chainA.GetContext(), suite.chainA.App.GetIBCKeeper().ClientKeeper)
		}, "BeginBlocker shouldn't panic")
	}
}

func (suite *ClientTestSuite) TestBeginBlockerConsensusState() {
	plan := &upgradetypes.Plan{
		Name:   "test",
		Height: suite.chainA.GetContext().BlockHeight() + 1,
	}
	// set upgrade plan in the upgrade store
	store := suite.chainA.GetContext().KVStore(suite.chainA.GetSimApp().GetKey(upgradetypes.StoreKey))
	bz := suite.chainA.App.AppCodec().MustMarshal(plan)
	store.Set(upgradetypes.PlanKey(), bz)

	nextValsHash := []byte("nextValsHash")
	newCtx := suite.chainA.GetContext().WithBlockHeader(tmproto.Header{
		Height:             suite.chainA.GetContext().BlockHeight(),
		NextValidatorsHash: nextValsHash,
	})

	err := suite.chainA.GetSimApp().UpgradeKeeper.SetUpgradedClient(newCtx, plan.Height, []byte("client state"))
	suite.Require().NoError(err)

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	suite.chainA.App.BeginBlock(req)

	// plan Height is at ctx.BlockHeight+1
	consState, found := suite.chainA.GetSimApp().UpgradeKeeper.GetUpgradedConsensusState(newCtx, plan.Height)
	suite.Require().True(found)
	bz, err = types.MarshalConsensusState(suite.chainA.App.AppCodec(), &ibctmtypes.ConsensusState{Timestamp: newCtx.BlockTime(), NextValidatorsHash: nextValsHash})
	suite.Require().NoError(err)
	suite.Require().Equal(bz, consState)
}

func (suite *ClientTestSuite) TestBeginBlockerWithUpgradePlan_EmitsUpgradeChainEvent() {
	plan := &upgradetypes.Plan{
		Name:   "test",
		Height: suite.chainA.GetContext().BlockHeight() + 1,
	}
	// set upgrade plan in the upgrade store
	store := suite.chainA.GetContext().KVStore(suite.chainA.GetSimApp().GetKey(upgradetypes.StoreKey))
	bz := suite.chainA.App.AppCodec().MustMarshal(plan)
	store.Set(upgradetypes.PlanKey(), bz)

	nextValsHash := []byte("nextValsHash")
	newCtx := suite.chainA.GetContext().WithBlockHeader(tmproto.Header{
		Height:             suite.chainA.GetContext().BlockHeight(),
		NextValidatorsHash: nextValsHash,
	})

	err := suite.chainA.GetSimApp().UpgradeKeeper.SetUpgradedClient(newCtx, plan.Height, []byte("client state"))
	suite.Require().NoError(err)

	client.BeginBlocker(suite.chainA.GetContext(), suite.chainA.App.GetIBCKeeper().ClientKeeper, suite.requireContainsEvent(types.EventTypeUpgradeChain, true))
}

func (suite *ClientTestSuite) TestBeginBlockerWithoutUpgradePlan_DoesNotEmitUpgradeChainEvent() {
	client.BeginBlocker(suite.chainA.GetContext(), suite.chainA.App.GetIBCKeeper().ClientKeeper, suite.requireContainsEvent(types.EventTypeUpgradeChain, false))
}

// requireContainsEvent returns a callback function which can be used to verify an event of a specific type was emitted.
func (suite *ClientTestSuite) requireContainsEvent(eventType string, shouldContain bool) func(events sdk.Events) {
	return func(events sdk.Events) {
		found := false
		var eventTypes []string
		for _, e := range events {
			eventTypes = append(eventTypes, e.Type)
			if e.Type == eventType {
				found = true
			}
		}
		if shouldContain {
			suite.Require().True(found, "event type %s was not found in %s", eventType, strings.Join(eventTypes, ","))
		} else {
			suite.Require().False(found, "event type %s was found in %s", eventType, strings.Join(eventTypes, ","))
		}
	}
}
