package keeper_test

import (
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (suite *KeeperTestSuite) TestOnChanOpenInit() {
	var (
		channel  *channeltypes.Channel
		path     *ibctesting.Path
		chanCap  *capabilitytypes.Capability
		metadata icatypes.Metadata
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success",
			func() {
				path.EndpointA.SetChannel(*channel)
			},
			true,
		},
		{
			"invalid order - UNORDERED",
			func() {
				channel.Ordering = channeltypes.UNORDERED
			},
			false,
		},
		{
			"invalid port",
			func() {
				path.EndpointA.ChannelConfig.PortID = "invalid-port-id"
			},
			false,
		},
		{
			"invalid counterparty port ID",
			func() {
				path.EndpointA.SetChannel(*channel)
				channel.Counterparty.PortId = "invalid-port-id"
			},
			false,
		},
		{
			"invalid metadata bytestring",
			func() {
				path.EndpointA.SetChannel(*channel)
				channel.Version = "invalid-metadata-bytestring"
			},
			false,
		},
		{
			"connection not found",
			func() {
				channel.ConnectionHops = []string{"invalid-connnection-id"}
				path.EndpointA.SetChannel(*channel)
			},
			false,
		},
		{
			"invalid controller connection ID",
			func() {
				metadata.ControllerConnectionId = "invalid-connnection-id"

				bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
				suite.Require().NoError(err)

				channel.Version = string(bz)
				path.EndpointA.SetChannel(*channel)
			},
			false,
		},
		{
			"invalid host connection ID",
			func() {
				metadata.HostConnectionId = "invalid-connnection-id"

				bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
				suite.Require().NoError(err)

				channel.Version = string(bz)
				path.EndpointA.SetChannel(*channel)
			},
			false,
		},
		{
			"invalid version",
			func() {
				metadata.Version = "invalid-version"

				bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
				suite.Require().NoError(err)

				channel.Version = string(bz)
				path.EndpointA.SetChannel(*channel)
			},
			false,
		},
		{
			"channel is already active",
			func() {
				suite.chainA.GetSimApp().ICAControllerKeeper.SetActiveChannelID(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = NewICAPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			// mock init interchain account
			portID, err := icatypes.NewControllerPortID(TestOwnerAddress)
			suite.Require().NoError(err)

			portCap := suite.chainA.GetSimApp().IBCKeeper.PortKeeper.BindPort(suite.chainA.GetContext(), portID)
			suite.chainA.GetSimApp().ICAControllerKeeper.ClaimCapability(suite.chainA.GetContext(), portCap, host.PortPath(portID))
			path.EndpointA.ChannelConfig.PortID = portID

			// default values
			metadata = icatypes.Metadata{
				Version:                icatypes.Version,
				ControllerConnectionId: ibctesting.FirstConnectionID,
				HostConnectionId:       ibctesting.FirstConnectionID,
			}

			bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
			suite.Require().NoError(err)

			counterparty := channeltypes.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.ORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        string(bz),
			}

			chanCap, err = suite.chainA.App.GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			tc.malleate() // malleate mutates test data

			err = suite.chainA.GetSimApp().ICAControllerKeeper.OnChanOpenInit(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, channel.Version,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestOnChanOpenAck() {
	var (
		path     *ibctesting.Path
		metadata icatypes.Metadata
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"invalid port - host chain",
			func() {
				path.EndpointA.ChannelConfig.PortID = icatypes.PortID
			},
			false,
		},
		{
			"invalid port - unexpected prefix",
			func() {
				path.EndpointA.ChannelConfig.PortID = "invalid-port-id"
			},
			false,
		},
		{
			"invalid metadata bytestring",
			func() {
				path.EndpointA.Counterparty.ChannelConfig.Version = "invalid-metadata-bytestring"
			},
			false,
		},
		{
			"invalid account address",
			func() {
				metadata.Address = "invalid-account-address"

				bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
				suite.Require().NoError(err)

				path.EndpointA.Counterparty.ChannelConfig.Version = string(bz)
			},
			false,
		},
		{
			"invalid counterparty version",
			func() {
				metadata.Version = "invalid-version"

				bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
				suite.Require().NoError(err)

				path.EndpointA.Counterparty.ChannelConfig.Version = string(bz)
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = NewICAPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			err := InitInterchainAccount(path.EndpointA, TestOwnerAddress)
			suite.Require().NoError(err)

			err = path.EndpointB.ChanOpenTry()
			suite.Require().NoError(err)

			metadata = icatypes.Metadata{
				Version:                icatypes.Version,
				ControllerConnectionId: ibctesting.FirstConnectionID,
				HostConnectionId:       ibctesting.FirstConnectionID,
				Address:                TestAccAddress.String(),
			}

			bz, err := icatypes.ModuleCdc.MarshalJSON(&metadata)
			suite.Require().NoError(err)

			path.EndpointB.ChannelConfig.Version = string(bz)

			tc.malleate() // malleate mutates test data

			err = suite.chainA.GetSimApp().ICAControllerKeeper.OnChanOpenAck(suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointA.Counterparty.ChannelConfig.Version,
			)

			if tc.expPass {
				suite.Require().NoError(err)

				activeChannelID, found := suite.chainA.GetSimApp().ICAControllerKeeper.GetActiveChannelID(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID)
				suite.Require().True(found)

				suite.Require().Equal(path.EndpointA.ChannelID, activeChannelID)

				interchainAccAddress, found := suite.chainA.GetSimApp().ICAControllerKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID)
				suite.Require().True(found)

				suite.Require().Equal(metadata.Address, interchainAccAddress)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnChanCloseConfirm() {
	var (
		path *ibctesting.Path
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = NewICAPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			err := SetupICAPath(path, TestOwnerAddress)
			suite.Require().NoError(err)

			tc.malleate() // malleate mutates test data

			err = suite.chainB.GetSimApp().ICAControllerKeeper.OnChanCloseConfirm(suite.chainB.GetContext(),
				path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)

			activeChannelID, found := suite.chainB.GetSimApp().ICAControllerKeeper.GetActiveChannelID(suite.chainB.GetContext(), path.EndpointB.ChannelConfig.PortID)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().False(found)
				suite.Require().Empty(activeChannelID)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}
