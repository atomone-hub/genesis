### **GovGen Townhall Meeting 16 Notes**

### Date: November 6th, 9am PST
<br> 

- Link to the [full transcript](https://docs.google.com/document/d/1ElevRoW7twLPQu0tc_fngdqQqerhuwqEwMOW43t1vI8/view)
- Link to the [full recording](https://drive.google.com/file/d/1Pqn14v7l2KQ3hE4x-4pGxW0Leq3L6mab/view?usp=sharing)
- Link to the [full message log](https://drive.google.com/file/d/1J3yiHN0W1W5DaTWBRlz8vy-xQ9FV5YJr/view?usp=drive_web)




**Main action items**

- Consider checking your allocation on the [distribution tracker.](https://govgen.io/#trackers)

- Read about the [distribution.](https://x.com/_atomone/status/1852103987950162034)

- Read the ["The Journey to AtomOne Launch: The Future of Decentralized Blockchain Governance"](https://allinbits.com/blog/atomone/)

- Delegate your tokens using any of the methods mentioned in the "The Journey to AtomOne Launch"

- Explore the [AiB Grants and Bounties Program](https://github.com/allinbits/grants/tree/main/AiB-BUIDL-Grants-and-Bounties-program)

<BR>





**Brief Summary of Discussion**

The first Community Townhall after AtomOne mainnet focused on updates, governance mechanisms, technical developments, and community contributions.

- The call started with the suggested talking points (as posted below) where the recent token distribution details were explained and there was clarification on voting multipliers and addressing questions about allocation discrepancies. Community-driven proposals were encouraged, including enabling IBC transfers through governance.
- Engineers presented updates on governance improvements, introducing a “governor” role for delegating voting power based on political alignment. Upcoming updates include preventing proposal spam and introducing Photon as the chain’s primary transaction fee token.
- A new staking dashboard was demonstrated, showcasing its simplicity and compatibility with Kepler, Leap, and Cosmos Station wallets. It will be open-sourced soon for wider community use.
- AIB expresses it's commitment to security, with ongoing audits, bug bounty programs, and third-party security reviews planned.
- Concerns were raised about the initial selection of validators. A community member clarified that the choice was based on technical requirements at launch and noted plans for a transparent delegation program to decentralize voting power further amongst other factors that will develop over time.
- There was discussion about expanding marketing efforts. A decentralized, community-led marketing initiative was proposed, to be funded by the AtomOne community treasury DAO.
- There was extensive discussion on identifying official validator representatives on Discord, with suggestions like using security contact emails or a JSON schema for validator properties.
- The current AtomOne constitution, initially voted on by the GovGen community, allows for future amendments through community proposals and high-approval votes. Discussions and suggestions for improvements were encouraged.
- The townhall concluded with thanks for community engagement and a reminder of upcoming updates.


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
   - if the number of active proposals exceeeds a target number, the deposit increases exponentially until it is less than or equal to that target number.
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
