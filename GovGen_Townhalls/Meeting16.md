### **GovGen Townhall Meeting 16 Notes**

### Date: November 6th, 9am PST
<br> 


**Main action items**

- Consider checking your allocation on the [distribution tracker.](https://govgen.io/#trackers)

- Read about the [distribution.](https://x.com/_atomone/status/1852103987950162034)

- Read the ["The Journey to AtomOne Launch: The Future of Decentralized Blockchain Governance"](https://allinbits.com/blog/atomone/)

- Delegate your tokens using any of the methods mentioned in the "The Journey to AtomOne Launch"

- Explore the [AiB Grants and Bounties Program](https://github.com/allinbits/grants/tree/main/AiB-BUIDL-Grants-and-Bounties-program)

<BR>

**(Brief Summary of Discussion goes here)**

**Suggested Talking Points**

- Chain Launch
  - On October 18th, a group of community validators organized came together, submitted their genTXs, created a new genesis, and successfully launched the AtomOne chain.   
- [On-chain data](https://www.mintscan.io/atomone)
  - 60 validators
  - ~14,000 accounts
  - ~300,000 blocks
  - ~140,000 transactions
- [Atone Distribution](https://x.com/_atomone/status/1852103987950162034)
  - 40.56% vote turnout approved the ATONE token distribution on GovGen community chain. 
  - 96,997,800 ATONE tokens allocated to 1,128,299 Cosmos Hub (ATOM) addresses.
  - 5,388,766 ATONE were allocated to the Community Pool
  - 5,388,766 ATONE were allocated to a reserved address for the future funding of the AtomOne Treasury DAO
  - 107,775,332 tokens in the total supply of ATONE genesis.   
- Engineering updates
  - Governors
    - Implementation https://github.com/atomone-hub/atomone/pull/16
  - Proposal deposit auto throttling
    - https://forum.cosmos.network/t/governance-proposal-deposit-auto-throttler/10121
    - Currently evaluation the following formula:
    ```math
    f(n) = \begin{cases} D * (1+\alpha)^{(n - N)} & n \gt N \\ D & n \leq N \end{cases}
    ```
    Where:
      - `n` is the number of active proposals
      - `D` is the initial minimum deposit required to activate a proposal
      - `α` is the rate change
      - `f(n)` is the actual value of the minimum deposit
  - PHOTON implementation
    - Meta issue https://github.com/atomone-hub/atomone/issues/44
    - ADR https://github.com/atomone-hub/atomone/pull/34
    - New `x/photon` module https://github.com/atomone-hub/atomone/pull/45
- Staking Dashboard Updates
  - Status of staking and governance dApp
- Security Updates
  - Audits
- Next Steps
  - Ratification of Tokenomics and Constitution documents
  - IBC Transfer Enablement
- [AiB Grants and Bounties Program](https://github.com/allinbits/grants/tree/main/AiB-BUIDL-Grants-and-Bounties-program)
- [Validator Verification for Discord Roles](https://github.com/atomone-hub/atomone-validator-community/blob/main/validator-info.md)
- Q&A
