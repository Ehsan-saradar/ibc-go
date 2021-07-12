package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
)

// NewInitialChildGenesisState returns a child GenesisState for a completely new child chain.
// TODO: Include chain status
func NewInitialChildGenesisState(cs *ibctmtypes.ClientState, consState *ibctmtypes.ConsensusState) ChildGenesisState {
	return ChildGenesisState{
		NewChain:             true,
		ParentClientState:    cs,
		ParentConsensusState: consState,
	}
}

// NewRestartChildGenesisState returns a child GenesisState that has already been established.
func NewRestartChildGenesisState(channelID string, unbondingSequences []*UnbondingSequence) ChildGenesisState {
	return ChildGenesisState{
		ParentChannelId:    channelID,
		UnbondingSequences: unbondingSequences,
		NewChain:           false,
	}
}

// DefaultGenesisState returns a default disabled child chain genesis state. This allows the module to be hooked up to app without getting use
// unless explicitly specified in genesis.
func DefaultChildGenesisState() ChildGenesisState {
	return ChildGenesisState{
		Disabled: true,
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs ChildGenesisState) Validate() error {
	if gs.Disabled {
		return nil
	}
	if gs.NewChain {
		if gs.ParentClientState == nil {
			return sdkerrors.Wrap(ErrInvalidGenesis, "parent client state cannot be nil for new chain")
		}
		if err := gs.ParentClientState.Validate(); err != nil {
			return sdkerrors.Wrapf(ErrInvalidGenesis, "parent client state invalid for new chain %s", err.Error())
		}
		if gs.ParentConsensusState == nil {
			return sdkerrors.Wrap(ErrInvalidGenesis, "parent consensus state cannot be nil for new chain")
		}
		if err := gs.ParentConsensusState.ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(ErrInvalidGenesis, "parent consensus state invalid for new chain %s", err.Error())
		}
		if gs.ParentChannelId != "" {
			return sdkerrors.Wrap(ErrInvalidGenesis, "parent channel id cannot be set for new chain. must be established on handshake")
		}
		if gs.UnbondingSequences != nil {
			return sdkerrors.Wrap(ErrInvalidGenesis, "unbonding sequences must be nil for new chain")
		}
	} else {
		if gs.ParentChannelId == "" {
			return sdkerrors.Wrap(ErrInvalidGenesis, "parent channel id must be set for a restarting child genesis state")
		}
		if gs.ParentClientState != nil || gs.ParentConsensusState != nil {
			return sdkerrors.Wrap(ErrInvalidGenesis, "parent client state and consensus states must be nil for a restarting genesis state")
		}
		for _, us := range gs.UnbondingSequences {
			if us == nil {
				return sdkerrors.Wrap(ErrInvalidGenesis, "unbonding sequence cannot be nil in genesis")
			}
			if err := us.Validate(); err != nil {
				return sdkerrors.Wrap(err, "invalid genesis")
			}
		}
	}
	return nil
}

func (us UnbondingSequence) Validate() error {
	if us.UnbondingTime == 0 {
		return sdkerrors.Wrap(ErrInvalidUnbondingSequence, "cannot have 0 unbonding time")
	}
	if err := us.UnbondingPacket.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "invalid unbonding packet")
	}
	if us.UnbondingPacket.Sequence != us.Sequence {
		return sdkerrors.Wrapf(ErrInvalidUnbondingSequence, "unbonding sequence %d must match packet sequence %d", us.Sequence, us.UnbondingPacket.Sequence)
	}
	return nil
}