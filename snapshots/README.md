# Snapshots

## Block [18010657]

This block is selected to be pre-tally for [cosmoshub-4 proposal 848][prop848]

> [!NOTE]
> Block [18010658] is where the tally takes place, and where the proposal 848
> is finally passed. But since the votes are removed from the state during the
> tally, we can only read the votes from the previous block, hence block
> [18010657]. The votes added in the block [18010658] will be added manually, as
> explained below.

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

While direct voters are easy to extract, indirect voters must be determined by
iterating over delegations and correlating them with validator votes.

#### Get all direct voters

```sh
$ jq '[.app_state.gov.votes[] | select(.proposal_id == "848")]' cosmoshub-4-export-18010657.json > votes.json

$ md5sum votes.json
a9782883000b3064e22d2200ea9cbdca  votes.json
```
Returns 173,165 votes (41Mb).

We need to manually add the last votes from block [18010658], there's only
[one][votes18010658] actually:

```sh
$ jq '. += [{
  "option": "VOTE_OPTION_YES",
  "options": [
    {
      "option": "VOTE_OPTION_YES",
      "weight": "1.000000000000000000"
    }
  ],
  "proposal_id": "848",
  "voter": "cosmos1jq6rpkf233jq9h98tlarzk8w3pl3lx87sv3t28"
}]' votes.json > votes_final.json

$ md5sum votes_final.json
3c8af5f1a3e00b02de6e819cee6f3638  votes_final.json
```

With that last vote, we have 173,166 votes.

The file is available here https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/votes_final.json

#### Get all delegations

> [!NOTE]
> Delegations and validators should be taken from block [18010658] where the
> tally takes place. But since there's no delegation or redelegation in this
> block, we'll continue to extract from block [18010657] for convenience.

```sh
$ jq '.app_state.staking.delegations' cosmoshub-4-export-18010657.json > delegations.json

$ md5sum delegations.json
be316ecfb9d5853ffcb65b29cf1ddd8d  delegations.json
```

Returns 1,061,423 delegations (238Mb). If not found in direct voters, any
delegation address will inherit validator's vote.

The file is available here https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/delegations.json

#### Get active bonded validators

```sh
$ jq '.app_state.staking.validators' cosmoshub-4-export-18010657.json > validators.json

$ md5sum validators.json
16cb26b14afb4799b5c2504285b2cc14  validators.json
```

Returns 531 validators (610Kb).

To have the active set, we need to:
- Get the `max_validator` parameters:
```sh
$ jq '.app_state.staking.params.max_validators' cosmoshub-4-export-18010657.json
180
```
- Filter out bonded validators
- Sort by the `tokens` field (descending)
- Limit to `max_validators`

```sh
$ jq '[.[] | select(.status == "BOND_STATUS_BONDED")] | sort_by(.tokens|tonumber) | reverse | .[:180]' validators.json > active_validators.json

$ md5sum active_validators.json
d3c09490eba24a1c0ec52fa9af3f28ac active_validators.json
```

Now we have only the 180 active validators.

This procedures follows the code of the [`staking.Keeper.IterateBondedValidatorsByPower()`][code-validators]
function, which is used in the [`x/gov.Keeper.Tally()`][code-tally] function.

The file is available here https://atomone.fra1.digitaloceanspaces.com/cosmoshub-4/prop848/active_validators.json

## TODO

- [ ] independently verify the snapshot data with another full node. each
  validator that wants to participate can start by attesting to the generated
  export.
- [ ] is there a schema for the exported json? how is it used? 

[18010657]: https://www.mintscan.io/cosmos/block/18010657
[18010658]: https://www.mintscan.io/cosmos/block/18010658
[prop848]: https://www.mintscan.io/cosmos/proposals/848
[votes18010658]: https://www.mintscan.io/cosmos/tx/9E0250C856A9F3B369A5C85BAA07C5F7284C8466EA7F15AACCA5F0F3C99F59A4?height=18010658
[code-validators]: https://github.com/cosmos/cosmos-sdk/blob/9abd946ba0cdc6d0e708bf862b2ca202b13f2d7b/x/staking/keeper/alias_functions.go#L33
[code-tally]: https://github.com/cosmos/cosmos-sdk/blob/9abd946ba0cdc6d0e708bf862b2ca202b13f2d7b/x/gov/keeper/tally.go#L13
