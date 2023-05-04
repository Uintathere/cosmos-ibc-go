package keeper_test

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"

	transferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
)

// TestMigratorMigrateParams tests successful parameter migration from x/params to self store
func (suite *KeeperTestSuite) TestMigratorMigrateParams() {
	testCases := []struct {
		msg            string
		malleate       func()
		expectedParams transfertypes.Params
	}{
		{
			"success: false-false params",
			func() {
				err := suite.chainA.GetSimApp().TransferKeeper.SetParams(
					suite.chainA.GetContext(),
					transfertypes.NewParams(false, false),
				)
				suite.Require().NoError(err)
			},
			transfertypes.NewParams(false, false),
		},
		{
			"success: false-true params",
			func() {
				err := suite.chainA.GetSimApp().TransferKeeper.SetParams(
					suite.chainA.GetContext(),
					transfertypes.NewParams(false, true),
				)
				suite.Require().NoError(err)
			},
			transfertypes.NewParams(false, true),
		},
		{
			"success: true-false params",
			func() {
				err := suite.chainA.GetSimApp().TransferKeeper.SetParams(
					suite.chainA.GetContext(),
					transfertypes.NewParams(true, false),
				)
				suite.Require().NoError(err)
			},
			transfertypes.NewParams(true, false),
		},
		{
			"success: true-true params",
			func() {
				err := suite.chainA.GetSimApp().TransferKeeper.SetParams(
					suite.chainA.GetContext(),
					transfertypes.NewParams(true, true),
				)
				suite.Require().NoError(err)
			},
			transfertypes.NewParams(true, true),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate() // explicitly set up params

			migrator := transferkeeper.NewMigrator(suite.chainA.GetSimApp().TransferKeeper, suite.chainA.GetSimApp().GetSubspace(transfertypes.ModuleName))
			err := migrator.MigrateParams(suite.chainA.GetContext())
			suite.Require().NoError(err)

			params := suite.chainA.GetSimApp().TransferKeeper.GetParams(suite.chainA.GetContext())
			suite.Require().Equal(tc.expectedParams, params)
		})
	}
}

func (suite *KeeperTestSuite) TestMigratorMigrateTraces() {
	testCases := []struct {
		msg            string
		malleate       func()
		expectedTraces transfertypes.Traces
	}{
		{
			"success: two slashes in base denom",
			func() {
				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(
					suite.chainA.GetContext(),
					transfertypes.DenomTrace{
						BaseDenom: "pool/1", Path: "transfer/channel-0/gamm",
					})
			},
			transfertypes.Traces{
				{
					BaseDenom: "gamm/pool/1", Path: "transfer/channel-0",
				},
			},
		},
		{
			"success: one slash in base denom",
			func() {
				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(
					suite.chainA.GetContext(),
					transfertypes.DenomTrace{
						BaseDenom: "0x85bcBCd7e79Ec36f4fBBDc54F90C643d921151AA", Path: "transfer/channel-149/erc",
					})
			},
			transfertypes.Traces{
				{
					BaseDenom: "erc/0x85bcBCd7e79Ec36f4fBBDc54F90C643d921151AA", Path: "transfer/channel-149",
				},
			},
		},
		{
			"success: multiple slashes in a row in base denom",
			func() {
				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(
					suite.chainA.GetContext(),
					transfertypes.DenomTrace{
						BaseDenom: "1", Path: "transfer/channel-5/gamm//pool",
					})
			},
			transfertypes.Traces{
				{
					BaseDenom: "gamm//pool/1", Path: "transfer/channel-5",
				},
			},
		},
		{
			"success: multihop base denom",
			func() {
				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(
					suite.chainA.GetContext(),
					transfertypes.DenomTrace{
						BaseDenom: "transfer/channel-1/uatom", Path: "transfer/channel-0",
					})
			},
			transfertypes.Traces{
				{
					BaseDenom: "uatom", Path: "transfer/channel-0/transfer/channel-1",
				},
			},
		},
		{
			"success: non-standard port",
			func() {
				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(
					suite.chainA.GetContext(),
					transfertypes.DenomTrace{
						BaseDenom: "customport/channel-7/uatom", Path: "transfer/channel-0/transfer/channel-1",
					})
			},
			transfertypes.Traces{
				{
					BaseDenom: "uatom", Path: "transfer/channel-0/transfer/channel-1/customport/channel-7",
				},
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate() // explicitly set up denom traces

			migrator := transferkeeper.NewMigrator(suite.chainA.GetSimApp().TransferKeeper, suite.chainA.GetSimApp().GetSubspace(transfertypes.ModuleName))
			err := migrator.MigrateTraces(suite.chainA.GetContext())
			suite.Require().NoError(err)

			traces := suite.chainA.GetSimApp().TransferKeeper.GetAllDenomTraces(suite.chainA.GetContext())
			suite.Require().Equal(tc.expectedTraces, traces)
		})
	}
}

func (suite *KeeperTestSuite) TestMigratorMigrateTracesCorruptionDetection() {
	// IBCDenom() previously would return "customport/channel-0/uatom", but now should return ibc/{hash}
	corruptedDenomTrace := transfertypes.DenomTrace{
		BaseDenom: "customport/channel-0/uatom",
		Path:      "",
	}
	suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(suite.chainA.GetContext(), corruptedDenomTrace)

	migrator := transferkeeper.NewMigrator(suite.chainA.GetSimApp().TransferKeeper, suite.chainA.GetSimApp().GetSubspace(transfertypes.ModuleName))
	suite.Panics(func() {
		migrator.MigrateTraces(suite.chainA.GetContext()) //nolint:errcheck // we shouldn't check the error here because we want to ensure that a panic occurs.
	})
}

func (suite *KeeperTestSuite) TestMigrateTotalEscrowForDenom() {
	var (
		path  *ibctesting.Path
		denom string
	)

	testCases := []struct {
		msg               string
		malleate          func()
		expectedEscrowAmt math.Int
	}{
		{
			"success: one native denom escrowed in one channel",
			func() {
				denom = sdk.DefaultBondDenom
				escrowAddress := transfertypes.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				coin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))

				// funds the escrow account to have balance
				suite.Require().NoError(banktestutil.FundAccount(suite.chainA.GetSimApp().BankKeeper, suite.chainA.GetContext(), escrowAddress, sdk.NewCoins(coin)))
			},
			math.NewInt(100),
		},
		{
			"success: one native denom escrowed in two channels",
			func() {
				denom = sdk.DefaultBondDenom
				extraPath := NewTransferPath(suite.chainA, suite.chainB)
				suite.coordinator.Setup(extraPath)

				escrowAddress1 := transfertypes.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				escrowAddress2 := transfertypes.GetEscrowAddress(extraPath.EndpointA.ChannelConfig.PortID, extraPath.EndpointA.ChannelID)
				coin1 := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))
				coin2 := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))

				// funds the escrow accounts to have balance
				suite.Require().NoError(banktestutil.FundAccount(suite.chainA.GetSimApp().BankKeeper, suite.chainA.GetContext(), escrowAddress1, sdk.NewCoins(coin1)))
				suite.Require().NoError(banktestutil.FundAccount(suite.chainA.GetSimApp().BankKeeper, suite.chainA.GetContext(), escrowAddress2, sdk.NewCoins(coin2)))
			},
			math.NewInt(200),
		},
		{
			"success: valid ibc denom escrowed in one channel",
			func() {
				escrowAddress := transfertypes.GetEscrowAddress(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
				trace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.DefaultBondDenom))
				coin := sdk.NewCoin(trace.IBCDenom(), sdk.NewInt(100))
				denom = trace.IBCDenom()

				suite.chainA.GetSimApp().TransferKeeper.SetDenomTrace(suite.chainA.GetContext(), trace)

				// funds the escrow account to have balance
				suite.Require().NoError(banktestutil.FundAccount(suite.chainA.GetSimApp().BankKeeper, suite.chainA.GetContext(), escrowAddress, sdk.NewCoins(coin)))
			},
			sdk.NewInt(100),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			path = NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			tc.malleate() // explicitly fund escrow account

			migrator := transferkeeper.NewMigrator(suite.chainA.GetSimApp().TransferKeeper, suite.chainA.GetSimApp().GetSubspace(transfertypes.ModuleName))
			suite.Require().NoError(migrator.MigrateTotalEscrowForDenom(suite.chainA.GetContext()))

			// check that the migration set the expected amount for both native and IBC tokens
			amount := suite.chainA.GetSimApp().TransferKeeper.GetTotalEscrowForDenom(suite.chainA.GetContext(), denom)
			suite.Require().Equal(tc.expectedEscrowAmt, amount.Amount)
		})
	}
}
