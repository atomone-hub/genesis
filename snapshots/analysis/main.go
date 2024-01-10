package main

import (
	"encoding/json"
	"fmt"
	"os"

	h "github.com/dustin/go-humanize"
)

func main() {
	if len(os.Args) != 3 || (os.Args[1] != "tally" && os.Args[1] != "accounts") {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [tally|genesis] [datapath]\n", os.Args[0])
		os.Exit(1)
	}
	//-----------------------------------------
	// Read data from files

	var (
		command  = os.Args[1]
		datapath = os.Args[2]
	)
	votesByAddr, err := parseVotesByAddr(datapath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s votes\n", h.Comma(int64(len(votesByAddr))))
	valsByAddr, err := parseValidatorsByAddr(datapath, votesByAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d validators\n", len(valsByAddr))
	delegsByAddr, err := parseDelegationsByAddr(datapath)
	if err != nil {
		panic(err)
	}
	var numDeleg int
	for _, d := range delegsByAddr {
		numDeleg += len(d)
	}
	fmt.Printf("%s delegations for %s delegators\n", h.Comma(int64(numDeleg)),
		h.Comma(int64(len(delegsByAddr))))

	balancesByAddr, err := parseBalancesByAddr(datapath, "uatom")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s account balances\n", h.Comma(int64(len(balancesByAddr))))

	switch command {
	case "tally":
		results, totalVotingPower := tally(votesByAddr, valsByAddr, delegsByAddr)
		// Optionnaly print and compare tally with prop data
		printTallyResults(results, totalVotingPower, parseProp(datapath))

	case "accounts":
		accounts := getAccounts(delegsByAddr, votesByAddr, valsByAddr, balancesByAddr)

		bz, err := json.MarshalIndent(accounts, "", "  ")
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile("accounts.json", bz, 0o666); err != nil {
			panic(err)
		}
	}
}
