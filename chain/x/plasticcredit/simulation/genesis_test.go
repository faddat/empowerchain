package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/assert"

	"github.com/EmpowerPlastic/empowerchain/x/plasticcredit"
	"github.com/EmpowerPlastic/empowerchain/x/plasticcredit/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		BondDenom:    sdk.DefaultBondDenom,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var plasticcreditGenesis plasticcredit.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[plasticcredit.ModuleName], &plasticcreditGenesis)

	assert.NotEmpty(t, plasticcreditGenesis.Params.IssuerCreator)
	assert.NotEmpty(t, plasticcreditGenesis.Issuers)
	assert.NotEmpty(t, plasticcreditGenesis.Applicants)
	assert.NotEmpty(t, plasticcreditGenesis.CreditClasses)
	assert.NotEmpty(t, plasticcreditGenesis.Projects)
	assert.NotEmpty(t, plasticcreditGenesis.CreditCollections)
	assert.NotEmpty(t, plasticcreditGenesis.CreditBalances)

	assert.NoError(t, plasticcreditGenesis.Validate())
}
