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

### Get the list of voters

#### Get All voters

```sh
$ jq '.app_state.gov.votes[] | select(.proposal_id == "848")'  cosmoshub-4-export-18010657.json 
[...]
```

## TODO

- [ ] independently verify the snapshot data with another full node. each validator that wants to participate can start by attesting to the generated export.
- [ ] is there a schema for the exported json? how is it used? 
