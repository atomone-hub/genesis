_Proposed Agenda for Townhall for Feb 6th 2024, 9 AM PDT/ 6 PM CET_ 

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

  - Invite Forbole, who has already signed up: <https://github.com/atomone-hub/validator>

  - What inspired you to align with GovGen and AtomOne? 

**Engineering updates:**

- What’s happening in GovGen engineering? (20 to 30 minutes) 

  - Building genesis from a snapshot of CosmosHub block [18010657](https://www.mintscan.io/cosmos/block/18010657)

    - Extract relevant data using jq commands <https://github.com/atomone-hub/genesis/pull/96>

    - Validate data against prop 848 final tally result, consolidate data with all accounts, balances (liquid and staked) and votes (direct & indirect). <https://github.com/atomone-hub/genesis/pull/103>

    - Created python script to get initial balances. 1:1 GOVGEN distribution to No and NWV voters (only staked amount and only the weight that voted either No or NWV) (<https://github.com/giunatale/govgen-genesis> )

* Building govgen chain: (<https://github.com/giunatale/govgen> )

  - Starting from Gaia 14.1.0

  - Revert to standard sdk 0.45.16 without LSM

  - Remove globalfee and revert to old mempool decorator

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

  - Removed IBC & ICS

  - Set bech32 prefix to govgen

- What happens next? (5 minutes)
