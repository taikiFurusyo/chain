package types_test

import (
	"testing"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	// For the test, use "cosmos" prefix
	testAddr        = "cosmos1nyd8wdqyrnjfwfnfysv6t0rrpcj4pmzkykytjh"
	testAddr2       = "cosmos1chjjqrherp2lgmj9wsqwe6leercydncqx2v209"
	testClassName   = "UnUniFi"
	testUri         = "ipfs://test/"
	testTokenSupply = 10000
	testSymbol      = "TEST"
	testDescription = "This description is for the valdation uni test"
)

func TestValidateMintingPermission(t *testing.T) {
	// OnlyOwner case
	owner, _ := sdk.AccAddressFromBech32(testAddr)
	classAttirbutes := types.ClassAttributes{
		Owner:             owner.Bytes(),
		MintingPermission: 0,
	}
	err := types.ValidateMintingPermission(classAttirbutes, owner)
	require.NoError(t, err)

	falseCase, _ := sdk.AccAddressFromBech32(testAddr2)
	err = types.ValidateMintingPermission(classAttirbutes, falseCase)
	require.Error(t, err)

	// AnyOne case
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 1,
	}
	err = types.ValidateMintingPermission(classAttirbutes, owner)
	require.NoError(t, err)

	// In case of now allowed option
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 3,
	}
	err = types.ValidateMintingPermission(classAttirbutes, owner)
	require.Error(t, err)
}

func TestValidateClassName(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateClassName(params.MinClassNameLen, params.MaxClassNameLen, testClassName)
	require.NoError(t, err)

	// invalid case which name is shorter than the default MinClassNameLen
	invalidClassNameShort := testClassName[:2]
	err = types.ValidateClassName(params.MinClassNameLen, 10000, invalidClassNameShort)
	require.Error(t, err)

	// invalid case which name is longer than the default MaxClassNameLen
	var invalidClassNameLong string
	for i := 0; i <= (int(params.MaxClassNameLen) / 7); i++ {
		invalidClassNameLong += testClassName
	}
	err = types.ValidateClassName(0, params.MaxClassNameLen, invalidClassNameLong)
	require.Error(t, err)
}

func TestValidateUri(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateUri(params.MinUriLen, params.MaxUriLen, testUri)
	require.NoError(t, err)

	// invalid case which uri is shoter than the default MinUriLen
	invalidUriShort := testUri[:4]
	err = types.ValidateUri(params.MinUriLen, 10000, invalidUriShort)
	require.Error(t, err)

	// invalid case which uri is longer than the default MaxUriLen
	var invalidUriLong string
	for i := 0; i <= (int(params.MaxUriLen) / len(testUri)); i++ {
		invalidUriLong += testUri
	}
	err = types.ValidateUri(0, params.MaxUriLen, invalidUriLong)
	require.Error(t, err)
}

func TestValidateTokenSupplyCap(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateTokenSupplyCap(params.MaxNFTSupplyCap, testTokenSupply)
	require.NoError(t, err)

	// invalid case which token supply cap is bigger than the default MaxTokenSupplyCap
	invalidTokenSupply := testTokenSupply * ((params.MaxNFTSupplyCap)/testTokenSupply + 1)
	err = types.ValidateTokenSupplyCap(params.MaxNFTSupplyCap, invalidTokenSupply)
	require.Error(t, err)
}

func TestValidateSymbol(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateSymbol(params.MaxSymbolLen, testSymbol)
	require.NoError(t, err)

	// invalid case which symbol is longer that the default MaxSymbolLen
	var invalidSymbol string
	for i := 0; i <= (int(params.MaxSymbolLen) / len(testSymbol)); i++ {
		invalidSymbol += testSymbol
	}
	err = types.ValidateSymbol(params.MaxSymbolLen, invalidSymbol)
	require.Error(t, err)
}

func TestValidateDescription(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateDescription(params.MaxDescriptionLen, testDescription)
	require.NoError(t, err)

	// invalid case which description is longer than the default MaxDescriptionLen
	var invalidDescription string
	for i := 0; i <= (int(params.MaxDescriptionLen) / len(testDescription)); i++ {
		invalidDescription += testDescription
	}
	err = types.ValidateDescription(params.MaxDescriptionLen, invalidDescription)
	require.Error(t, err)
}
