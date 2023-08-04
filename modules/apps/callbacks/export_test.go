package ibccallbacks

/*
	This file is to allow for unexported functions and fields to be accessible to the testing package.
*/

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v7/modules/apps/callbacks/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
)

// ProcessCallback is a wrapper around processCallback to allow the function to be directly called in tests.
func (im IBCMiddleware) ProcessCallback(
	ctx sdk.Context, packet channeltypes.Packet, callbackType types.CallbackType,
	callbackDataGetter func() (types.CallbackData, error),
	callbackExecutor func(sdk.Context, string, string) error,
) error {
	return im.processCallback(ctx, packet, callbackType, callbackDataGetter, callbackExecutor)
}

// GetICS4Wrapper is a getter for the IBCMiddleware's ICS4Wrapper.
func (im *IBCMiddleware) GetICS4Wrapper() porttypes.ICS4Wrapper {
	return im.ics4Wrapper
}
