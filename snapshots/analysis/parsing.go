package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gogo/protobuf/jsonpb"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
)

var (
	registry    = codectypes.NewInterfaceRegistry()
	unmarshaler jsonpb.Unmarshaler
)

func init() {
	cryptocodec.RegisterInterfaces(registry)
	govtypes.RegisterInterfaces(registry)
	sdk.RegisterInterfaces(registry)
	proposaltypes.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)
	vestingtypes.RegisterInterfaces(registry)
	icatypes.RegisterInterfaces(registry)
	unmarshaler = jsonpb.Unmarshaler{AnyResolver: registry}
}

func parseAccountTypesPerAddr(path string) (map[string]string, error) {
	f, err := os.Open(filepath.Join(path, "auth_genesis.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var genesis authtypes.GenesisState
	err = unmarshaler.Unmarshal(f, &genesis)
	if err != nil {
		return nil, err
	}
	accountTypesPerAddr := make(map[string]string)
	for i, any := range genesis.Accounts {
		var acc authtypes.GenesisAccount
		registry.UnpackAny(any, &acc)
		accountTypesPerAddr[acc.GetAddress().String()] = genesis.Accounts[i].GetTypeUrl()
	}
	return accountTypesPerAddr, nil
}

func parseVotesByAddr(path string) (map[string]govtypes.WeightedVoteOptions, error) {
	f, err := os.Open(filepath.Join(path, "votes.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// XXX workaround to unmarshal votes because proto doesn't support top-level array
	dec := json.NewDecoder(f)
	_, err = dec.Token()
	if err != nil {
		return nil, err
	}
	votesByAddr := make(map[string]govtypes.WeightedVoteOptions)
	for dec.More() {
		var vote govtypes.Vote
		err := unmarshaler.UnmarshalNext(dec, &vote)
		if err != nil {
			return nil, err
		}
		votesByAddr[vote.Voter] = vote.Options
	}
	return votesByAddr, nil
}

func parseDelegationsByAddr(path string) (map[string][]stakingtypes.Delegation, error) {
	f, err := os.Open(filepath.Join(path, "delegations.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var delegs []stakingtypes.Delegation
	err = json.NewDecoder(f).Decode(&delegs)
	if err != nil {
		return nil, err
	}
	delegsByAddr := make(map[string][]stakingtypes.Delegation)
	for _, d := range delegs {
		delegsByAddr[d.DelegatorAddress] = append(delegsByAddr[d.DelegatorAddress], d)
	}
	return delegsByAddr, nil
}

func parseValidatorsByAddr(path string, votesByAddr map[string]govtypes.WeightedVoteOptions) (map[string]govtypes.ValidatorGovInfo, error) {
	f, err := os.Open(filepath.Join(path, "active_validators.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// XXX workaround to unmarshal validators because proto doesn't support top-level array
	dec := json.NewDecoder(f)
	_, err = dec.Token()
	if err != nil {
		return nil, err
	}
	valsByAddr := make(map[string]govtypes.ValidatorGovInfo)
	for dec.More() {
		var val stakingtypes.Validator
		err := unmarshaler.UnmarshalNext(dec, &val)
		if err != nil {
			return nil, err
		}

		valAddr, err := sdk.ValAddressFromBech32(val.OperatorAddress)
		if err != nil {
			panic(err)
		}
		accAddr := sdk.AccAddress(valAddr.Bytes()).String()
		valsByAddr[val.OperatorAddress] = govtypes.NewValidatorGovInfo(
			val.GetOperator(),
			val.GetBondedTokens(),
			val.GetDelegatorShares(),
			sdk.ZeroDec(),
			votesByAddr[accAddr],
		)
	}
	return valsByAddr, nil
}

func parseProp(path string) govtypes.Proposal {
	f, err := os.Open(filepath.Join(path, "prop.json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var prop govtypes.Proposal
	err = unmarshaler.Unmarshal(f, &prop)
	if err != nil {
		panic(err)
	}
	return prop
}

func parseBalancesByAddr(path, denom string) (map[string]sdk.Coin, error) {
	f, err := os.Open(filepath.Join(path, "balances.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var balances []banktypes.Balance
	err = json.NewDecoder(f).Decode(&balances)
	if err != nil {
		return nil, err
	}
	balancesByAddr := make(map[string]sdk.Coin)
	for _, b := range balances {
		for _, c := range b.Coins {
			// Filter denom
			if c.Denom == denom {
				balancesByAddr[b.Address] = c
				break
			}
		}
	}
	return balancesByAddr, nil
}
