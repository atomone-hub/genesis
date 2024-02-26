### **GovGen Townhall Meeting 5 Notes**

### Date: February 26th, 2024, 6:00 AM PST

<br> 

- Link to the [full transcript here] will go here
- Link to the [full recording here] will go here

**Main action items**

- Launch
- [Join our Discord](https://discord.gg/atomone)

<br> 

**Brief Summary of Discussion will go here**

You can read a synopsis of the existing state of GovGen in this blog post, including that the term for applications has been slightly extended. 

https://allinbits.com/blog/shape-the-future-of-atomone-with-govgen/

**Suggested Talking Points**
- Validator Onboarding Process & Timeline
  - The initial communicated deadline for submission was Friday, February 23, at 5 p.m. PST (Saturday, 1 a.m. UTC). However, the term has been extended until Saturday, February 24, at 5 p.m. PST (Sunday, 1 a.m. UTC). 
- Technical Updates and Governance Discussions
  - errata corrige: [PR 7](https://github.com/atomone-hub/govgen/issues/7) is superseded by [PR 20](https://github.com/atomone-hub/govgen/issues/20): validators are allowed to vote in governance, but delegations won't count there, and validators can vote on proposals only using self-stake. Block production/tendermint voting power distribution/staking will work regularly though.
  - GovGen v1.0.0 released: https://github.com/atomone-hub/govgen/releases/tag/v1.0.0
  - PR for final genesis that collects `gentxs` ready: https://github.com/atomone-hub/govgen-genesis/pull/3
    - validator accounts that do not exist in genesis distribution are funded with 25 $GOVGEN (self-stake is 1 $GOVGEN)
    - validator accounts that have balance less than 25 $GOVGEN are also topped off to fund operations
    - final genesis link will be available both at https://github.com/atomone-hub/govgen-genesis and https://github.com/atomone-hub/validator/tree/main/govgen-1 as soon as complete.
  - Initial staking distribution: we want validators to have some delegations on block 0. For this we decided on the following numbers, which can be discussed:
    - select accounts with balance > 25 $GOVGEN
    - stake 50% of that balance to validators with a fair distribution
    
    With the current genesis, that will stake 33,890,002 govgen, which makes an initial staking ratio of ~49% (total supply is 68,336,631 govgen). We are working on the algorithm to distribute this stake evenly among the validators.
  - questions:
    - fee and validator setup discussion and recommendations
    - final review of genesis parameters:
      - x/gov: are pass threshold, quorum, deposits and voting periods OK?
      - any other param to discuss/change?
- First Proposals for GovGen
- Cover any questions
