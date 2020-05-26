package magpie

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/simulation"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
)

const (
	EventTypeCreateSession    = types.EventTypeCreateSession
	AttributeKeySessionID     = types.AttributeKeySessionID
	AttributeKeyNamespace     = types.AttributeKeyNamespace
	AttributeKeyExternalOwner = types.AttributeKeyExternalOwner
	AttributeKeyExpiry        = types.AttributeKeyExpiry
	AttributeValueCategory    = types.AttributeValueCategory
	ModuleName                = types.ModuleName
	RouterKey                 = types.RouterKey
	StoreKey                  = types.StoreKey
	ActionCreationSession     = types.ActionCreationSession
	QuerySessions             = keeper.QuerySessions
	OpWeightMsgCreatePost     = simulation.OpWeightMsgCreatePost
)

var (
	// functions aliases
	NewHandler               = keeper.NewHandler
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RandomSessionData        = simulation.RandomSessionData
	DecodeStore              = simulation.DecodeStore
	RandomizedGenState       = simulation.RandomizedGenState
	WeightedOperations       = simulation.WeightedOperations
	SimulateMsgCreateSession = simulation.SimulateMsgCreateSession
	RegisterCodec            = types.RegisterCodec
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	SessionStoreKey          = types.SessionStoreKey
	NewMsgCreateSession      = types.NewMsgCreateSession
	ParseSessionID           = types.ParseSessionID
	NewSession               = types.NewSession

	// variable aliases
	RandomNamespaces      = simulation.RandomNamespaces
	ModuleCdc             = types.ModuleCdc
	SessionLengthKey      = types.SessionLengthKey
	LastSessionIDStoreKey = types.LastSessionIDStoreKey
	SessionStorePrefix    = types.SessionStorePrefix
)

type (
	Keeper           = keeper.Keeper
	SessionData      = simulation.SessionData
	GenesisState     = types.GenesisState
	MsgCreateSession = types.MsgCreateSession
	SessionID        = types.SessionID
	Session          = types.Session
	Sessions         = types.Sessions
)
