package main

import (
	"fmt"
	"os"

	h "github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func tally(
	votesByAddr map[string]govtypes.WeightedVoteOptions, valsByAddr map[string]govtypes.ValidatorGovInfo,
	delegsByAddr map[string][]stakingtypes.Delegation,
) (map[govtypes.VoteOption]sdk.Dec, sdk.Dec) {
	var (
		results = map[govtypes.VoteOption]sdk.Dec{
			govtypes.OptionYes:        sdk.ZeroDec(),
			govtypes.OptionAbstain:    sdk.ZeroDec(),
			govtypes.OptionNo:         sdk.ZeroDec(),
			govtypes.OptionNoWithVeto: sdk.ZeroDec(),
		}
		totalVotingPower = sdk.ZeroDec()
	)
	for voterAddr, vote := range votesByAddr {
		// Check voter delegations
		dels := delegsByAddr[voterAddr]
		// Initialize voter balance
		// balance := sdk.NewDec(0)
		for _, del := range dels {
			val, ok := valsByAddr[del.ValidatorAddress]
			if !ok {
				// Validator isn't in active set or jailed, ignore
				continue
			}
			// Reduce validator voting power with delegation that has voted
			val.DelegatorDeductions = val.DelegatorDeductions.Add(del.GetShares())
			valsByAddr[del.ValidatorAddress] = val

			// delegation shares * bonded / total shares
			votingPower := del.GetShares().MulInt(val.BondedTokens).Quo(val.DelegatorShares)
			// Iterate over vote options
			for _, option := range vote {
				subPower := votingPower.Mul(option.Weight)
				results[option.Option] = results[option.Option].Add(subPower)
			}
			totalVotingPower = totalVotingPower.Add(votingPower)
		}
	}
	// iterate over the validators again to tally their voting power
	for _, val := range valsByAddr {
		if len(val.Vote) == 0 {
			continue
		}
		sharesAfterDeductions := val.DelegatorShares.Sub(val.DelegatorDeductions)
		votingPower := sharesAfterDeductions.MulInt(val.BondedTokens).Quo(val.DelegatorShares)

		for _, option := range val.Vote {
			subPower := votingPower.Mul(option.Weight)
			results[option.Option] = results[option.Option].Add(subPower)
		}
		totalVotingPower = totalVotingPower.Add(votingPower)
	}
	return results, totalVotingPower
}

func printTallyResults(results map[govtypes.VoteOption]sdk.Dec, totalVotingPower sdk.Dec, prop govtypes.Proposal) {
	fmt.Println("Computed total voting power", h.Comma(totalVotingPower.TruncateInt64()))
	yesPercent := results[govtypes.OptionYes].
		Quo(totalVotingPower.Sub(results[govtypes.OptionAbstain]))
	fmt.Println("Yes percent:", yesPercent)
	tallyResult := govtypes.NewTallyResultFromMap(results)

	fmt.Println("--- TALLY RESULT ---")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Yes", "No", "NoWithVeto", "Abstain", "Total"})
	M := sdk.NewInt(1_000_000)
	appendTable := func(source string, t govtypes.TallyResult) {
		total := t.Yes.Add(t.No).Add(t.Abstain).Add(t.NoWithVeto)
		table.Append([]string{
			source,
			h.Comma(t.Yes.Quo(M).Int64()),
			h.Comma(t.No.Quo(M).Int64()),
			h.Comma(t.NoWithVeto.Quo(M).Int64()),
			h.Comma(t.Abstain.Quo(M).Int64()),
			h.Comma(total.Quo(M).Int64()),
		})
	}
	appendTable("computed", tallyResult)
	appendTable("from prop", prop.FinalTallyResult)
	diff := govtypes.NewTallyResult(
		tallyResult.Yes.Sub(prop.FinalTallyResult.Yes),
		tallyResult.Abstain.Sub(prop.FinalTallyResult.Abstain),
		tallyResult.No.Sub(prop.FinalTallyResult.No),
		tallyResult.NoWithVeto.Sub(prop.FinalTallyResult.NoWithVeto),
	)
	appendTable("diff", diff)
	table.Render()
}
