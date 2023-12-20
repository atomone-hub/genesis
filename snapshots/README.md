# Snapshots

## Block [18010657](https://www.mintscan.io/cosmos/block/18010657)

This block is selected to be pre-tally for [cosmoshub-4 proposal 848](https://www.mintscan.io/cosmos/proposals/848)

Block [18010657](https://www.mintscan.io/cosmos/block/18010657) is the last block where vote are present in the state.
In order to get all votes, we need to add the votes present in block [18010658](https://www.mintscan.io/cosmos/block/18010658) where the tally happens.
This is because after the tally is completed, the votes a removed from state.

### Download

An export is stored in S3 accessible here: https://atomone.fra1.digitaloceanspaces.com/cosmoshub4/cosmoshub-4-export-18010657.json

### How the block has been exported

On a stopped blockchain node containing the block 18010657

```sh
$ gaiad export --height 18010657 > cosmoshub-4-export-18010657.json 2>&1
```

### Verify the export

```sh
$ md5sum cosmoshub-4-export-18010657.json
84cdf83c11c7a88e0cf60070391a2519  cosmoshub-4-export-18010657.json
```

### Get direct & indirect voters

#### Get all direct voters

```sh
$ jq '[.app_state.gov.votes[] | select(.proposal_id == "848")]'  cosmoshub-4-export-18010657.json > votes.json
$ md5sum votes.json
a9782883000b3064e22d2200ea9cbdca  votes.json
```

Returns 173,165 votes (41Mb).

#### Get all delegations

```sh
$ jq jq '.app_state.staking.delegations' cosmoshub-4-export-18010657.json >
delegations.json
$ md5sum delegations.json
be316ecfb9d5853ffcb65b29cf1ddd8d  delegations.json
```

Returns 1,061,423 delegations (238Mb). If not found in direct voters, any
delegation address will inherit validator's vote.

#### Get all validators

```sh
$ jq '.app_state.staking.validators' cosmoshub-4-export-18010657.json >
validators.json
$ md5sum validators.json
16cb26b14afb4799b5c2504285b2cc14  validators.json
```

Returns 531 validators (610Kb). Useful for determining which votes belong to
which validators, and also for subtracting the voting power of direct voters.

To have only the active set, we need to get the `max_validator` parameters:

```sh
$ jq '.app_state.staking.params.max_validators' cosmoshub-4-export-18010657.json
180
```

Then we need to sort the list according to the `tokens` field, and limit to
180.


## TODO

- [ ] independently verify the snapshot data with another full node. each validator that wants to participate can start by attesting to the generated export.
- [ ] is there a schema for the exported json? how is it used? 
