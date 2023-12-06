# Snapshots

Following information in issue [atomone-hub/genesis#18](https://github.com/atomone-hub/genesis/issues/18)

## Block [18010659](https://www.mintscan.io/cosmos/block/18010659)

This block is selected to be pre-tally for [cosmoshub-4 proposal 848](https://www.mintscan.io/cosmos/proposals/848)

### Download

An export is stored in S3 accessible here: https://allinbits.fra1.digitaloceanspaces.com/cosmoshub4/cosmoshub-4-export-18010659.json

### How the block has been exported

On a stopped blockchain node containing the block 18010659

``` sh
$ gaiad export --height 18010659 > cosmoshub-4-export-18010659.json 2>&1
```

### Verify the export

``` shsh
$ md5sum cosmoshub-4-export-18010659.json
769441139d802f95c39a2ce0cf8fe857  cosmoshub-4-export-18010659.json
```
