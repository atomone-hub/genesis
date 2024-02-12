### **GovNo Townhall Meeting 2 Notes**

### Date: February 13th, 2024

<br> 

**Suggested talking points**

- Validator Onboarding Process
- Technical Updates and Governance Discussions
- Governance Mechanisms and Validator Requirements
- Chain Architecture and Documentation
- Governance Model and Community Involvement
- Validator Rewards
- Governance Policies and Tokenomics
- KYC Processes and Regulatory Considerations
- Governance Structures and Meeting Scheduling

<br>

**Engineering Updates**

Testing GovGen did not show any issue
- Forked x/gov to implement discussed modifications
    - https://github.com/atomone-hub/govgen/pull/4
    - https://github.com/atomone-hub/govgen/pull/9
- Remove ability for validators to vote
    - https://github.com/atomone-hub/govgen/pull/10
- Allow different voting periods for different types of proposals (WIP)

<br>

**Previous Engineering Updates**

- Building genesis from a snapshot of CosmosHub block [18010657](https://www.mintscan.io/cosmos/block/18010657)

  - Extract relevant data using jq commands <https://github.com/atomone-hub/genesis/pull/96>

   - Validate data against prop 848 final tally result, consolidate data with all accounts, balances (liquid and staked) and votes (direct & indirect). <https://github.com/atomone-hub/genesis/pull/103>

   - Created python script to get initial balances. 1:1 GOVGEN distribution to No and NWV voters (only staked amount and only the weight that voted either No or NWV) (<https://github.com/atomone-hub/govgen-genesis> )

   - Chain params forked from Cosmso Hub, with the following changes:

     - Bank module:

       - Disabled sendTx

     - Distribution module:

       - community tax, proposer reward and bonus all set to 0

     - Mint module:

       - inflation disabled, no new minting of tokens

     - Gov module:

       - Deposit amount raised to 5000 GOVGEN

       - Voting period extended to 365 days

       - Vote threshold for pass raised to ⅔

       - Quorum raised to 50%

     - Staking module:

       - reduced validators to 30 tentatively

<br>

* Building govgen chain: (<https://github.com/atomone-hub/govgen> )

  - Starting from Gaia 14.1.0

  - Revert to standard sdk 0.45.16 without LSM

  - Remove globalfee module and revert to old mempool decorator

  - Removed IBC & ICS and related modules (e.g. ICA)

  - Set bech32 prefix to govgen

  - Reduced gov initial deposit to 1% to better fit the 5000 $GOVGEN deposit threshold (initial deposit 50 $GOVGEN)
