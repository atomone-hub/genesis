package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Account struct {
	Address      string
	Type         string
	LiquidAmount sdk.Dec
	StakedAmount sdk.Dec
	Vote         govtypes.WeightedVoteOptions
	Delegations  []Delegation
}

type Delegation struct {
	Amount           sdk.Dec
	ValidatorAddress string
	Vote             govtypes.WeightedVoteOptions
}

// getAccounts returns the list of all account with their vote and
// power, from direct or indirect votes.
func getAccounts(
	delegsByAddr map[string][]stakingtypes.Delegation,
	votesByAddr map[string]govtypes.WeightedVoteOptions,
	valsByAddr map[string]govtypes.ValidatorGovInfo,
	balancesByAddr map[string]sdk.Coin,
	accountTypesPerAddr map[string]string,
) []Account {
	accountsByAddr := make(map[string]Account, len(delegsByAddr))
	// Feed delegations
	for addr, delegs := range delegsByAddr {
		account := Account{
			Address:      addr,
			Type:         accountTypesPerAddr[addr],
			LiquidAmount: sdk.ZeroDec(),
			StakedAmount: sdk.ZeroDec(),
			Vote:         votesByAddr[addr],
		}
		for _, deleg := range delegs {
			// Find validator
			val, ok := valsByAddr[deleg.ValidatorAddress]
			if !ok {
				// Validator isn't in active set or jailed, ignore
				continue
			}

			// Compute delegation voting power
			delegVotingPower := deleg.GetShares().MulInt(val.BondedTokens).Quo(val.DelegatorShares)
			account.StakedAmount = account.StakedAmount.Add(delegVotingPower)

			// Populate delegations with validator votes
			account.Delegations = append(account.Delegations, Delegation{
				ValidatorAddress: val.Address.String(),
				Amount:           delegVotingPower,
				Vote:             val.Vote,
			})
		}
		accountsByAddr[addr] = account
	}
	// Feed balances
	for addr, balance := range balancesByAddr {
		acc, ok := accountsByAddr[addr]
		if ok {
			acc.LiquidAmount = balance.Amount.ToDec()
			accountsByAddr[addr] = acc
		} else {
			accountsByAddr[addr] = Account{
				Address:      addr,
				Type:         accountTypesPerAddr[addr],
				LiquidAmount: balance.Amount.ToDec(),
				StakedAmount: sdk.ZeroDec(),
			}
		}
	}
	// Map to slice
	var accounts []Account
	for _, a := range accountsByAddr {
		accounts = append(accounts, a)
	}
	return accounts
}
