# Snapshot analysis

Tool to validate governance data from a snapshot and make it usable to for a
blockchain genesis.

```
Usage: go run . [tally|accounts] PATH
```

Where PATH is a directory containing the following files:
- [`votes.json`][1]
- [`delegations.json`][2]
- [`active_validators.json`][3]
- [`prop.json`][4]

> [!NOTE]
> The links to these files corresponds to data extracted for the proposal 848,
> the extraction method is detailed [here](../README.md).
> The prop 848 is taken as an example, but any other proposal can be used to
> build the same result.

## Verification

Considering all these files downloaded in the `data/prop848` directory, you can
compute the tally and compare it to the prop `FinalTallyResult` field.

```
$ go run . tally data/prop848

173,165 votes
180 validators
1,061,423 delegations for 765,656 delegators
Computed total voting power 177,825,601,877,018
Yes percent: 0.517062127500689774
--- TALLY RESULT ---
+-----------+------------+------------+------------+------------+-------------+
|           |    YES     |     NO     | NOWITHVETO |  ABSTAIN   |    TOTAL    |
+-----------+------------+------------+------------+------------+-------------+
| computed  | 73,165,203 | 56,667,011 | 11,669,549 | 36,323,836 | 177,825,601 |
| from prop | 73,165,203 | 56,667,011 | 11,669,549 | 36,323,836 | 177,825,601 |
| diff      |          0 |          0 |          0 |          0 |           0 |
+-----------+------------+------------+------------+------------+-------------+
```

which shows that the tally calculated from these files is exactly the same as
the tally from the prop stored in the blockchain data.

## List of accounts

Then, to build the genesis, we need the list of accounts with their liquid and
staked balances, their delegations and their votes. This is doable thanks to
the `accounts` command:

```
$ go run . accounts data/prop848
```

This command creates a [`accounts.json`][5] file, which can be used to build
the `bank` module genesis, once the account balances are updated according to
their vote on prop 848.

[1]: https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/votes.json
[2]: https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/delegations.json
[3]: https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/active_validators.json
[4]: https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/prop.json
[5]: https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/accounts.json 
