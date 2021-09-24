package fee_test

import (
	"fmt"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	"github.com/cosmos/ibc-go/modules/apps/29-fee/types"
	transfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/testing"
)

// Tests OnChanOpenInit on ChainA
func (suite *FeeTestSuite) TestOnChanOpenInit() {
	testCases := []struct {
		name    string
		version string
		expPass bool
	}{
		{
			"valid fee middleware and transfer version",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			true,
		},
		{
			"fee version not included, only perform transfer logic",
			transfertypes.Version,
			true,
		},
		{
			"invalid fee middleware version",
			channeltypes.MergeChannelVersions("otherfee28-1", transfertypes.Version),
			false,
		},
		{
			"invalid transfer version",
			channeltypes.MergeChannelVersions(types.Version, "wrongics20-1"),
			false,
		},
		{
			"incorrect wrapping delimiter",
			fmt.Sprintf("%s//%s", types.Version, transfertypes.Version),
			false,
		},
		{
			"transfer version not wrapped",
			types.Version,
			false,
		},
		{
			"hanging delimiter",
			fmt.Sprintf("%s:%s:", types.Version, transfertypes.Version),
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// reset suite
			suite.SetupTest()
			suite.coordinator.SetupClients(suite.path)
			suite.coordinator.SetupConnections(suite.path)

			suite.path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID)
			channel := &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{suite.path.EndpointA.ConnectionID},
				Version:        tc.version,
			}

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			chanCap, err := suite.chainA.App.GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, suite.path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			err = cbs.OnChanOpenInit(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				suite.path.EndpointA.ChannelConfig.PortID, suite.path.EndpointA.ChannelID, chanCap, counterparty, channel.Version)

			if tc.expPass {
				suite.Require().NoError(err, "unexpected error from version: %s", tc.version)
			} else {
				suite.Require().Error(err, "error not returned for version: %s", tc.version)
			}
		})
	}
}

// Tests OnChanOpenTry on ChainA
func (suite *FeeTestSuite) TestOnChanOpenTry() {
	testCases := []struct {
		name      string
		version   string
		cpVersion string
		crossing  bool
		expPass   bool
	}{
		{
			"valid fee middleware and transfer version",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			false,
			true,
		},
		{
			"valid transfer version on try and counterparty",
			transfertypes.Version,
			transfertypes.Version,
			false,
			true,
		},
		{
			"valid fee middleware and transfer version, crossing hellos",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			true,
			true,
		},
		{
			"invalid fee middleware version",
			channeltypes.MergeChannelVersions("otherfee28-1", transfertypes.Version),
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			false,
			false,
		},
		{
			"invalid counterparty fee middleware version",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			channeltypes.MergeChannelVersions("wrongfee29-1", transfertypes.Version),
			false,
			false,
		},
		{
			"invalid transfer version",
			channeltypes.MergeChannelVersions(types.Version, "wrongics20-1"),
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			false,
			false,
		},
		{
			"invalid counterparty transfer version",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			channeltypes.MergeChannelVersions(types.Version, "wrongics20-1"),
			false,
			false,
		},
		{
			"transfer version not wrapped",
			types.Version,
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			false,
			false,
		},
		{
			"counterparty transfer version not wrapped",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			types.Version,
			false,
			false,
		},
		{
			"fee version not included on try, but included in counterparty",
			transfertypes.Version,
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			false,
			false,
		},
		{
			"fee version not included",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			transfertypes.Version,
			false,
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// reset suite
			suite.SetupTest()
			suite.coordinator.SetupConnections(suite.path)
			suite.path.EndpointB.ChanOpenInit()

			var (
				chanCap *capabilitytypes.Capability
				ok      bool
				err     error
			)
			if tc.crossing {
				suite.path.EndpointA.ChanOpenInit()
				chanCap, ok = suite.chainA.GetSimApp().ScopedTransferKeeper.GetCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, suite.path.EndpointA.ChannelID))
				suite.Require().True(ok)
			} else {
				chanCap, err = suite.chainA.App.GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, suite.path.EndpointA.ChannelID))
				suite.Require().NoError(err)
			}

			suite.path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID)
			channel := &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{suite.path.EndpointA.ConnectionID},
				Version:        tc.version,
			}

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			err = cbs.OnChanOpenTry(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				suite.path.EndpointA.ChannelConfig.PortID, suite.path.EndpointA.ChannelID, chanCap, counterparty, tc.version, tc.cpVersion)

			if tc.expPass {
				suite.Require().NoError(err, "unexpected error from version: %s", tc.version)
			} else {
				suite.Require().Error(err, "error not returned for version: %s", tc.version)
			}
		})
	}
}

// Tests OnChanOpenAck on ChainA
func (suite *FeeTestSuite) TestOnChanOpenAck() {
	testCases := []struct {
		name      string
		cpVersion string
		malleate  func(suite *FeeTestSuite)
		expPass   bool
	}{
		{
			"success",
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			func(suite *FeeTestSuite) {},
			true,
		},
		{
			"invalid fee version",
			channeltypes.MergeChannelVersions("fee29-A", transfertypes.Version),
			func(suite *FeeTestSuite) {},
			false,
		},
		{
			"invalid transfer version",
			channeltypes.MergeChannelVersions(types.Version, "ics20-4"),
			func(suite *FeeTestSuite) {},
			false,
		},
		{
			"previous INIT set without fee, however counterparty set fee version", // note this can only happen with incompetent or malicious counterparty chain
			channeltypes.MergeChannelVersions(types.Version, transfertypes.Version),
			func(suite *FeeTestSuite) {
				// do the first steps without fee version, then pass the fee version as counterparty version in ChanOpenACK
				suite.path.EndpointA.ChannelConfig.Version = transfertypes.Version
				suite.path.EndpointB.ChannelConfig.Version = transfertypes.Version
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()
			suite.coordinator.SetupClients(suite.path)
			suite.coordinator.SetupConnections(suite.path)

			// malleate test case
			tc.malleate(suite)

			suite.path.EndpointA.ChanOpenInit()
			suite.path.EndpointB.ChanOpenTry()

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			err = cbs.OnChanOpenAck(suite.chainA.GetContext(), suite.path.EndpointA.ChannelConfig.PortID, suite.path.EndpointA.ChannelID, tc.cpVersion)
			if tc.expPass {
				suite.Require().NoError(err, "unexpected error for case: %s", tc.name)
			} else {
				suite.Require().Error(err, "%s expected error but returned none", tc.name)
			}
		})
	}
}
