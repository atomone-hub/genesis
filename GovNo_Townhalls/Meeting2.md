### **GovNo Townhall Meeting 2 Notes**

### Date: February 9th, 2024

<br>

- Link to the [full transcript here](https://docs.google.com/document/d/12yRU9Lt1MUIsbwdUBmb6jXZY-AHyjagtqRiGlXuiu1o/edit?usp=sharing)
- Link to the [full recording here](https://drive.google.com/file/d/1wrAhCuxrPDGcc8aRBOZUxQC2VcLsG0TN/view)
[![Recording](https://res.cloudinary.com/marcomontalbano/image/upload/v1707724221/video_to_markdown/images/google-drive--1wrAhCuxrPDGcc8aRBOZUxQC2VcLsG0TN-c05b58ac6eb4c4700831b2b3070cd403.jpg)](https://drive.google.com/file/d/1wrAhCuxrPDGcc8aRBOZUxQC2VcLsG0TN/view "Recording")

<br>

**Main action items**

- If you're interested in becoming a validator, make a PR on the [validator directory](https://github.com/atomone-hub/validator) following the process listed on there.

- Take a look at this code for [GovGen Genesis Creation](https://github.com/atomone-hub/govgen-genesis) and the [GovGen Chain Software](https://github.com/atomone-hub/govgen), submit an issue if you have something you want to bring up.

- [Join our Discord](https://discord.gg/atomone)

<br> 

**Brief Summary of Discussion**

Technical Updates and Governance Discussions:
- Distribution is 1:1 with No and NWV voters, initial distribution will probably be ~ 68m million GovGen tokens
- GovGen will have a few changes from the starting point of the fork of Cosmos Hub:

   - GovGen is non-transferable
   - Inflation Disabled
   - Community tax, proposer reward and bonus all set to 0
   - Deposit raised to 5000 GovGen
   - Voting period suggested to be extended to 365 days, discussion followed on different voting periods for different types of proposals, specifically for software upgrades, changing parameters. Conversation followed some consensus with talks about voting incentives then shut down.
   - Vote threshold set to 2/3
   - Quorum raised to 50%
   - Validators tentatively reduced to 30
   - Standard SDK v14.1 without the liquid-staking module
   - IBC module removed
   - ICS module removed
   - Spam prevention changed to 1% of initial deposit by proposer
   - GovGen will probably skip the testnet phase.

- Validator Onboarding Process
Part of the conversation focused on onboarding validators, discussing criteria, documentation and verification steps.
  - Some discussion on how to onboard validators that won't receive GovGen
   - Removing self-stake
   - Remove delegation potentially or whether validators should be able to vote.

- Governance Mechanisms and Validator Requirements
   - Discussion on how to proceed with multiple choice proposals
   - Only allowing tokens staked at the time of a proposal
   - Topic focused on the roles of validators and their requirements

- Chain Architecture and Documentation
   - Discussion on methods to improve accessibility and clarity.

- Governance Model and Community Involvement
   - Emphasis on the need for transparency, accountability, and inclusivity in decision-making processes. Discussion on general incentive for early participation in GovGen, potentially being called a founder.   -      - Talking through gamification of incentives.

- Validator Rewards
   - Discussion focused on whether it makes sense for GovGen or not

- Covered the Condorcet Paradox 
   - https://en.wikipedia.org/wiki/Condorcet_paradox

- Conclusion
   - Covered Action Items
   - Covered scheduling next meeting

<br>

**Main talking points**

- Validator Onboarding Process
- Technical Updates and Governance Discussions
- Governance Mechanisms and Validator Requirements
- Chain Architecture and Documentation
- Governance Model and Community Involvement
- Technical Updates and Validator Rewards
- Governance Policies and Tokenomics
- KYC Processes and Regulatory Considerations
- Governance Structures and Meeting Scheduling

<br>

<br>


 _This was the original Proposed Agenda for Townhall for Feb 6th 2024, 9 AM PDT/ 6 PM CET_ 

**Community growth and Ecosystem Development**

- GovGen and AtomOne progress (5 minutes)

  - Communication channels:

    - Discord is [set up](https://github.com/atomone-hub/genesis/pull/110) and moderation policy - please join! <https://discord.gg/atomone>

    - Twitter, the community has taken the @\_atomone handle and offered that to AiB for moderation, thank you so much @wnmnsr <https://twitter.com/_atomone>

  - FAQ

    - Add FAQ to read.me <https://github.com/atomone-hub/genesis/pull/109>

  - Community town halls and get togethers:  (5 minutes)

    - Structure of town halls, alternating between engineering and community calls.Tentative agendas to be published on GitHub at <https://github.com/atomone-hub/genesis/tree/main/GovNo_Townhalls> 2-3 days in advance of the each call - make sure you make suggestions for topics that you would like to discuss

    - Meeting notes and recordings will be published after each community call on GitHub 

  - Community growth:

    - AiB AtomOne Grant Program (10 minutes)

      - Anyone who needs financial assistance to help, make sure you’ll apply w/ CV and the area of contribution you are interested in etc. We will publish the grant application process and recommendations on Github soon. Stay tuned!

      - Brainstorm technical needs and other project ideas - let’s find the top three ideas!

        - Hub Minimalism

        - Scaled Security

        - Self-Reinforcing Constitution

        - Tendermint2

    - International Ambassador program ideation 

- Branding, naming and visuals (10 minutes)

  - Discuss the name change - [proposing](https://github.com/atomone-hub/genesis/pull/108) GovNo becomes GovGen due to community feedback

  - @wnmnsr would you want to come to the townhall to discuss AtomOne branding ideas?

    - <https://github.com/atomone-hub/genesis/issues/14>

  - Discussing the best way to decentralize voting on brand and visuals while not creating a governance spam scenario

- Validator directory and onboarding (15 minutes)

  - Give a brief overview of the directory, what it means, and what you can anticipate as a validator 

  - Forbole team has already signed up: https://github.com/atomone-hub/validator. @kwunyeung Would you like to join the Townhall and share with the community what inspired you to apply and what does your alignment with GovGen and AtomOne look like? 

**Engineering updates:**

- What’s happening in GovGen engineering? (20 to 30 minutes) 

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


* Building govgen chain: (<https://github.com/atomone-hub/govgen> )

  - Starting from Gaia 14.1.0

  - Revert to standard sdk 0.45.16 without LSM

  - Remove globalfee module and revert to old mempool decorator

  - Removed IBC & ICS and related modules (e.g. ICA)

  - Set bech32 prefix to govgen

  - Reduced gov initial deposit to 1% to better fit the 5000 $GOVGEN deposit threshold (initial deposit 50 $GOVGEN)

- What happens next? (5 minutes)
  
  - Testing, we are ready for launch!
  - Discuss initial parameters with the community
  - Build the validator set
  - Coordinate chain launch
