package token

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	wasmtypes "github.com/Finschia/wasmd/x/wasm/types"
)

type (
	// ClassKeeper defines the contract needed to be fulfilled for class dependencies.
	ClassKeeper interface {
		NewID(ctx sdk.Context) string
		HasID(ctx sdk.Context, id string) bool

		InitGenesis(ctx sdk.Context, data *ClassGenesisState)
		ExportGenesis(ctx sdk.Context) *ClassGenesisState
	}

	WasmKeeper interface {
		Execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
		GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
	}
)
