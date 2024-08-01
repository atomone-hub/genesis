# ATOMONE CONSTITUTION

_This document is a work in progress. This document assumes familiarity with
the current workings of cosmoshub-4 as of Oct 11th 2022. What is described here
are modifications to what already is. This clause will be removed with future
revisions, and the corresponding parts of the document updated with a full
description of the constitution of the hub._


## Preamble

> TO DO: 
> - Add Vision and Mission statement for AtomOne Hub and a bit of history for
>   context for newcomers; a “Declaration of intent” (note however that we have
>   Article 1.3 for stating the mission)


We, the people of the Cosmos, in order to create a free world, enable voluntary
and borderless transactions, facilitate permissionless innovation, ensure
economic security, resist censorship, cater to economic and technological
development, allow for the creation of sovereign zones, and maintain order
among sovereign zones, do ordain and establish this Constitution for the
AtomOne Hub.

### Declaration of Intent

The vision of AtomOne Hub is to be a minimal, secure, and resilient fork of the
Cosmos Hub (Gaia), providing a reliable alternative that champions sovereignty,
security, and decentralization. AtomOne aims to complement the broader Cosmos
Ecosystem while introducing innovative solutions and perspectives, fostering a
diverse ecosystem of specialized zones in cooperation and competition.

By providing a secure and minimal alternative to the existing Cosmos Hub,
AtomOne addresses the diverse and evolving challenges within the Cosmos
ecosystem, stemming from differences about core tokenomics, monetization
strategies, and the balance between profit interests and community
responsibilities. Proposal #848 on the Cosmos Hub, which narrowly passed,
highlighted the deep divide within the community regarding fundamental
tokenomics and security design.

The AtomOne Hub is established to serve as the canonical minimal IBC/ICS token
hub and a base for a more intelligent and unified voting bloc to safeguard the
Cosmos Ecosystem. By aligning with these ideals and fostering a diverse and
cooperative environment, AtomOne Hub strives to set a new precedent for
adaptive and responsive self-organization in the multichain, multitoken
universe of the Cosmos.


# Article 1: General Provisions

### Section 1: Fundamental Principles

> TO DO: 
> - How to Enforce Constitution
> - Add something like: “This provision does not apply to Constitutional
>   amendments which follow procedure described in Section 4, Article 4"


This Constitution of the AtomOne Hub, hereinafter “the Constitution” hereby
establishes the foundations of the governance model, economical model, and
operating system of the AtomOne Hub.

All subsequent governance proposals shall align with the provisions of the
Constitution, and proponents of each proposal, along with all active governance
voters, are required to ensure consistency between such proposals and the
Constitution. 

However, this provision does not apply to constitutional amendments, which
shall follow the procedures described in Section 4, Article 4. While amendments
are permitted to innovate and adapt the Constitution, they must still respect
and adhere to the fundamental principles upon which this Constitution is
founded, and cannot violate any explicit restriction.

### Section 2: Sovereignty of the AtomOne Hub

> TO DO: 
> - Detail what the AtomOne Hub offers to consumer chains (but perhaps not
>   here, where sovereignty is discussed)


The AtomOne Hub is one of many sovereign zones that compose the Cosmos
interchain network.

The AtomOne Hub is composed of many chains: the root hub chain, where staking
and governance transactions are committed and executed, but also other chains
(core shards) secured by Interchain Security (ICS) that are subservient to the
governance of the root hub chain, hereinafter referred to as "Hub Governance".

Other sovereign zones (consumer shards) that are completely or partially
secured by AtomOne Hub ICS, by definition, have their own governance mechanism.
The AtomOne Hub principally enables and follows the will of the governance of
such sovereign zones regarding the pegged tokens originating from said zones,
except in well-defined exceptional circumstances involving bugs, theft, or harm
to the Cosmonauts of the Cosmos Ecosystem.

### Section 3: General Mission and Objectives

> TO DO: 
> - evaluate an alternative to a “new world” because it potentially sounds too
>   “fantastic” (proposed: “better world”, “fair and free digital
>   state/nation”)


The mission of the AtomOne Hub is to create a new world, enabling
permissionless yet secure interactions between the Cosmonauts of Cosmos in the
Solar System.

The objective of the AtomOne Hub is to provide a classical Byzantine Fault
Tolerant (BFT) proof-of-stake multi-token payment and transfer system; and to
scale the security of the platform to many applications hosted in other zones
via the AtomOne Hub ICS.

### Section 4: Cosmonauts

> TO DO: 
> - consider a more direct title that clarifies this section is about the basic
>   rights of cosmonauts
> - consider clearly defining how domestic and foreign jurisdictions interact
>   so that limits and scope are defined
> - potentially add language about not excluding instead of there is a right
> - define “active in governance” (i.e we need to say 50% over X period of time
>   or 100%, etc)
> - consider whether “at least one address” is necessary 
> - Add “Shall allow every person (to own an address) without limits”


Every person has the right to become a Cosmonaut and may freely engage on the
AtomOne Hub. As such, every Cosmonaut has the right to own at least one address
on the AtomOne Hub.

Any Cosmonaut may also become a Citizen of the AtomOne Hub by using their
address to stake ATONEs toward the AtomOne Hub and participate actively in
governance. The status of citizenship is granted in an autonomous manner.


### Section 5: Rights, Liberties and Obligations in the AtomOne Hub

> TO DO: 
> - Potentially detail mechanisms for conflict resolution, such as the proposed
>   “Decentralized Arbitration Courts”


The Liberty and Property of all Cosmonauts engaging in the AtomOne Hub is
hereby guaranteed. Any restriction to the Liberty and Property of Cosmonauts on
the AtomOne Hub shall be done only through the AtomOne Hub's governance and
requires Constitutional Majority.

Every Cosmonaut has the right to receive benefits from their engagement in the
AtomOne Hub, including rights derived from held or staked ATONEs in line with
the provisions of this Constitution.

Every Citizen (including Validators) allows the Governance of the AtomOne Hub
to restrict their staked ATONE property by partial or full slashing according
to their consensus voting activity.


## Article 2: Governance

### Section 1: The AtomOne Hub Chain

> TO DO: 
> - consider adding $photon to the functions
> - consider making the chainid parametric with respect to version, i.e.
>   atomonehub-1, atomonehub-2, etc..


The root hub chain of the AtomOne Hub is uniquely identified by chainid
"atomonehub". This chain commits and executes transactions that serve the
following functions:

- governance voting
- intra-hub token transfers
- IBC token transfers
- ICS1 and ICS2 management

The root hub chain shall not be used for experimentation. Experimentation
should occur in other zones.

### Section 2: AtomOne Hub Governance

> TO DO: 
> - fill in rules of cosmoshub-4 governance (including different types of props
>   and majority consensus requirements) accounting also for any change to the
>   governance design and process resulting from the constitutional group work
> - consider adding conditions for expedited proposals (if we end up
>   implementing them)
> - define relationship between quorum and passing thresholds for different
>   types of proposal passing thresholds (simple proposals, ⅔’s majority and
>   Constitutional Majority)


The working language of AtomOne Hub governance is English.

The quorum necessary for a proposal to be valid shall depend solely on the
number of independently bonded ATONEs, meaning any ATONE tokens bonded in a
different method shall not count towards quorum.

The governance process must extend the voting deadline to ensure a minimum of 2
weeks of voting after the minimum quorum has been met.

User-facing interfaces that present the results of voting on governance
proposals should also display the content of the memo field of each voting
transaction such that the reason for the vote can be seen.

#### The Supermajority of Two Thirds

The Supermajority is defined to be exactly "more than two thirds" (+2/3, or at
least one iota more than two thirds) and cannot change even by a Constitutional
Majority.

#### The Constitutional Majority

The Constitutional Majority shall be initially set at 90%. 
The Constitutional Majority shall not be set lower than 90% even with a
Constitutional Majority, but it may be set to any value between 90% and 100%.

### Section 3: Constitution Enforcement Processes

> TO DO: 
> - Should this be its own article or in a different article? 
> - Define Constitution Enforcement, ensure parameter changes/proposals can’t/won’t break Constitutional law. 


### Section 4: Airdrops and forks

Every Cosmonaut allows any other Cosmonaut to create full or partial airdrops
of new tokens to any chain using the distribution of any token on the AtomOne
Hub at any time.

Every Citizen allows any Cosmonaut to modify their pro-rata airdrop portion by
partial or full slashing based on their cryptographic voting activity according
to well defined principles at any time.

All forks that lose consensus continuity must change their token ticker symbol
to be distinct from $ATONE. If there are competing chains with comparably
similar continuity, then the fork that has a higher market cap (as measured
after both tokens have discovered fair market value with sufficient liquidity
for at least one week) should retain the name while other forks change their
    token ticker symbol.

### Section 5: Treasury DAOs

> TO DO: 
> - Add how members or daos can be removed and dissolved
> - Same person should use same address across multiple DAOs or have a
>   disclosure process to associate multiple addresses to the same identity
> - Consider, do we have anything that deals with the potential for another ICF
>   emerging?
> 	- Do we want to regulate that in any way? 
> - Consider adding a legal wrapper per DAO. Consider how or if Treasury DAOs
>   will be held accountable (i.e reporting back to the community). 
> - Consider adding terms, and restrictions for participating members, (i.e no
>   more than 3 at the same time and no more than 10 total per person). 
> - Suggestion: Each DAO could be overseen by the DAO above in a tree structure
>   where for a top level treasury DAO the oversight committee would be onchain
>   governance
> - Consider whether, or at what point, someone may say that the
>   structure/Treasury DAO itself looks like an entity (i.e., board of
>   directors, etc.). This could be relevant for purposes of potential
>   lawsuits/enforcement, etc


The AtomOne Hub governance may establish one or more transparent and
accountable Treasury Decentralized Autonomous Organizations (DAOs) by simple
majority vote. However, the funding of Treasury DAOs shall be subject to a
supermajority vote.

The operations of the Treasury DAO must operate on an ICS-secured zone outside
of the AtomOne Hub. As an exception, the tokens of the Treasury DAO may reside
on the AtomOne Hub as a m-of-n multisig account where n is at least 3 and m is
at least 1/2 of n, where each signatory is an authorized member of the Treasury
DAO and Citizen of the AtomOne Hub.

The Treasury DAO shall be composed of Cosmonauts, and at the top Executive
Board level be composed of one or more Citizens. All Cosmonauts and Citizens of
Treasury DAOs including their Oversight Committee must have public and known
real personal identities.

To enable the well-functioning of Treasury DAOs and the separation of powers in
the utmost interest of the AtomOne Hub, each member can hold just one type of
role within each Treasury DAO. All treasury DAO’s should have a predetermined
oversight committee that is approved by and subject to AtomOne Hub’s
Governance. The Oversight committee members cannot be paid by the Treasury DAO
and is independently compensated.

The members of the DAO must perform in function of the specifications outlined
in the origination proposal of each Treasury DAO to ensure that they fill their
role in line with their job description. They are accountable to the DAO’s
Oversight Committee and the Hub Governance. They can be dismissed from their
functions by a supermajority vote by the DAO’s Oversight Committee or the
Governance Hub.

Each Treasury DAO's Oversight Committee shall be composed of any number of
Cosmonauts. The Oversight Committee must have at least the right to freeze all
transfers of tokens from the Treasury DAO or its designated m-of-n multisig
account on the AtomOne Hub.

The role and functions of the Executive Board shall be further developed
through governance proposals within each DAO, unless the DAO chooses to defer
to AtomOne Hub governance by Citizens.


## Article 3: Economics

### Section 1: Economic Model

The one and only economic incentive model of the AtomOne Hub is the collection
of market-based transaction fees from a large number of transactions across all
the chains secured by the staking of ATONEs on the AtomOne Hub, including
ICS-secured chains.

### Section 2: The ATONE Token

The ATONE functions as voting shares, economic incentive shares, and security
bonds for the AtomOne Hub.

To preserve the security and identity of the acting governance and validator
set, the inflation rate of the ATONE is made to vary over time to target 2/3 of
all ATONEs. The maximum inflation rate is 20% non-compounded per year. There is
no minimum inflation rate, and it can even be negative (deflationary).

> TO DO:
> - Consider removing negative inflation
> - Or take extra care about how negative inflation is mentioned, i.e the burn
>   to come from the stakers, some additional framing to keep in mind prop848
>   and its origin, mention what happened with Cosmos Hub and its dangers
> - More context for newcomers ("ATOM must not be money"). 
> - Potentially host a working group to make the dynamic rate of inflation more
>   efficient (PID Controller). 


Inflated ATONEs are paid to bonded ATONE holders in proportion to each
delegator's staking amount, and staked ATONEs are converted to Bonded Share
Units.

> TO DO:
> - reassess the flow for delegators->governor-> validator (could cause more
>   troubles than it solves)
> - generic tax considerations

The ATONE Unbonding Period shall be a minimum of 3 weeks, and redelegation is
allowed twice per ATONE Unbonding Period.

Double signing at any height/round/step results in slashing penalty that is
proportional to the total amount of double signing by all validators for that
height/round/step, with evidence collected during the ATONE Unbonding Period;
the penalty shall range from +0% to 100% of the Upper Slashing Limit in linear
proportion, the latter when 1/3 of voting power double-signs.

Complex signing failures (those that require +1/3 to coordinate) shall result
in slashing at the Upper Slashing Limit of the offenders only after the
Tendermint accountability seizure.

> TO DO:
> - Consider revision of the above statement for better clarity or explain the
>   procedure in more details, going on and explaining that
>  
> 	“There are two cases of slashing:
> 	- Simple slashing: 0%-50%
> 	- Complex slashing: 50%”


The Upper Slashing Limit for both instances shall be 50%. This parameter may be
increased by a two thirds majority of the AtomOne Hub. In cases where there is
sufficient evidence of malice and intent, this parameter may be overruled by a
simple majority of the AtomOne Hub on a per-case basis up to 100%.

> TO DO:
> - Consider what happens if validator abuses their voting power to self harm
>   with the intention of harming delegators/governors 
> - Potentially add an exception for mis-slashing in case there is an issue in
>   the "official software", 
> - Define different limits/ percentages for different cases of double-slashing
> - Define procedures and processes for sharing logs and provide information. 
> - Define checklist requirements for validators to provide in case of
>   double-signing


### Section 3: “Liquid Staking”

For purposes of this Constitution we classify Liquid Staking as a protocol or
application that allows users to receive a tradeable derivative token in
exchange for native staking tokens that are autonomously bonded by the protocol
or by delegated entity. These derivative tokens represent some claim on the
underlying bonded tokens and may accrue rewards.

Liquid staking may only be supported through interchain accounts (aka
non-native liquid staking).

To limit the amount of liquid staked tokens so as to reduce systemic risk from
liquid staking, there shall be imposed a 5% tax (the "Liquid Staking Tax") on
all rewards paid out to staked interchain accounts at time of reward
withdrawal. This Liquid Staking Tax parameter may be adjusted by a two-thirds
supermajority vote of the AtomOne Hub.

If the amount of ATONEs staked using interchain accounts exceeds 20% (the
"Liquid Staking Factor") of the total staked ATONEs, the Liquid Staking Tax
shall automatically increase by 1% per x weeks. The Liquid Staking Factor
parameter may be adjusted by a two-thirds supermajority vote of the AtomOne
Hub.

> TO DO:
> - simplify the above two rules
> - Port PHOTON info from readme
> - If Interchain accounts are used, “Interchain accounts are not allowed on
>   the Hub, except the cases that it supports the official liquid staking
>   protocol.
> - Use weeks or days to measure the time period, since the months do not
>   contain the same number of days.
> - Perhaps we say "IF it is enabled by the hub by supermajority, it must
>   adhere to the following limitations". and otherwise not allowing it. and
>   explaining how we need an "official" way that is limited, for security.
>   pressure valve justification; intended to discourage rogue LS protocols
>   that are unsafe. 
> - Consider disabling ICA interchain accounts, for external chains to delegate
>   via logic such as via Stride. but still would need limitations; see
>   Celestia multisig?


> [!NOTE]
> The text below on $PHOTON is for now just a copy-paste of relevant text from
> the [README.md](https://github.com/atomone-hub/genesis/blob/main/README.md), it requires merging.


By default, Consumer chains can accept any fees token, it’s possible for the
$photon token to be a required fee token on one or all consumer chains by
supermajority, but no other fee token besides the $photon token may be made to
be required by other consumer chains.

The only fee token required to be accepted by all shards shall be the $PHOTON
token. This must not change even with a Constitutional Majority as a matter of
trust of a preagreed transaction declared in these Founding Documents, except
to better serve this invariant such as by allowing for an alias or by
supporting different denominations of the same underlying $PHOTON. AtomOne will
not promote the $ATONE token to be used as a fee token directly, even though it
must be supported as a bootstrapping and recovery measure.

While the convertibility from $PHOTON to the underlying $ATONE may be managed,
paused, or throttled by governance of $ATONE with a Constitutional Majority of
the $ATONE stakers not including the $ATONE of the $PHOTON bonders, all the
underlying $ATONE must be distributable back to $PHOTON holders through a fair
system and all of the $ATONE withdrawn within 20 years starting at any given
moment.

The intent is to throttle PHOTON to ATONE to make sure there is no take over of
ATONE security.  Rewards from the $ATONE tokens bonded to $PHOTON tokens shall
be distributed back to $PHOTON as if they were any other $ATONE staked tokens,
but they shall not exercise their voting power and instead yield entirely to
the other $ATONE staked tokens.

Tax will be deducted from these $PHOTON bonded $ATONE rewards as usual just
like regular validator staked $ATONE tokens, but unlike the tax burden for
validator staked $ATONE tokens, the tax burden for $PHOTON bonded $ATOM tokens
shall be capped at 10%. This cap cannot be changed even with a Constitutional
Majority except by also a two-thirds supermajority from the $PHOTON holders
with a prominently announced vote put forth by the AtomOne hub with a voting
period of at least one year, and a quorum threshold of at least 10% of the
total supply of $PHOTON tokens by direct participation where the increased tax
burden above the 10% must be used for common goods purposes on transparent and
accountable DAO systems

The $ATONEs bonded toward auto-staking do not count toward calculating the
bonding ratio target of 2/3 in either the numerator or denominator--they are
ignored.

### Section 4: Inflation

Any inflation of ATONEs to the Community Pool or a Designated Treasury DAO
beyond the default inflation rate described in the Constitution shall require a
constitutional majority vote of a special inflation governance proposal type.

> TO DO:
> - Potentially restrict limits to (for example) a 20% limit per year, cannot
>   be modified “even with Constitutional Majority”?


The special inflation proposal can include a description of the purpose of the
inflation, **but it cannot include any other modifications to the AtomOne Hub
or its Constitution**, nor the adoption of any new Treasury DAOs.

> TO DO:
> - This provision which prohibits bundled “omnibus” proposals should be placed
>   in other relevant sections


### Section 5: The Community Pool

> TO DO:
> - Define limitations on what Community Pool can be used for.
> - Explore: define threshold depending on spending amount
> - Define: framework for funds requests
> - Define: describe how transaction fees and inflationary ATONEs make it to
>   the community pool
> - Supermajority can adopt new framework for funding


The Community Pool tax proceeds shall apply to transaction fees and
inflationary ATONEs, and shall be sent to the Community Pool.

The Community Pool Tax rate shall initially be 5%, but can be increased up to
10% by two-thirds supermajority of the AtomOne Hub governance and limited to
33% increase per year.


### Section 6: Validators

> TO DO:
> - Governor, Validator and Citizen relationships
> - Definite profile of validator, not necessarily politicians or lobbyists
> - Actual provisions on commission and rewarding mechanisms
> - Address validator incentives and ICS remuneration (from readme)


## Article 4: Final Dispositions

### Section 1: Updates to the AtomOne Hub

> TO DO:
> - Extending what expectations are desired.


New updates to the AtomOne Hub shall be broken down into independent components
and discussed/proposed separately with adequate time between, regardless of any
omnibus whitepaper.


### Section 2: The Implementation

> TO DO:
> - Define base software (which version of Tendermint/CometBFT, SDK…)

The AtomOne Hub shall not have any VM functionality, but shall be plainly
implemented in a single garbage-collected language as reference (namely Go);
and other clients may implement all or portions of the stack in another
language like Rust.

The only cryptographic assumptions allowed to be used by the AtomOne Hub,
including its consensus protocol shall be Ed25519 and Secp256k1 elliptic
curves, and RIPEMD160 and SHA256 hash functions.

> TO DO:
> - potentially add Merkle proofs.

No zero-knowledge proof systems shall be adopted on the AtomOne Hub even if
they are composed of the approved primitives.

> TO DO:
> - Potentially explain why no ZK, 
> - alternatively, use a more general phrasing to prohibit “experimental” or
>   “novel cryptographic technologies” with ZK as a mention. 
> - Consider adding “No new cryptographic assumptions are allowed to be used
>   including those of “diffygap” assumptions. 
> - Definitely add, “Note that these restrictions concern only the main chain,
>   not the core or consumer shards,” 
> - Consider adding, “Quantum proofing is allowed if it’s additive,” and needs
>   to be clarified and 
> - Decide if there needs to be an expedited/emergency approval mechanism

Mixing implementations across validators is also to be avoided so as to prevent
a failure arising from a low Nakamoto coefficient among the types of
implementations. Instead AtomOne shall always support one standard
implementation.

The rules of this section may only be changed by two-thirds supermajority vote
of the AtomOne Hub.

> TO DO:
> - Determine whether this should be a constitutional majority or we should
>   consider omitting because section 4 (below) makes this redundant. 
> - There will potentially be other places to see if this language could be
>   refined when we add readme content.


### Section 3: Compute/storage/memory limitations

> TO DO:
> - To consider revising? Light clients should be able to run on any hardware,
>   full nodes might need more than a commodity computer.


For the sake of decentralization, accessibility, accountability, and security,
the AtomOne Hub and each ICS zone shall be restricted so that each can run on a
commodity computer.


### Section 4: Amendment of the Constitution

> TO DO:
> - Evaluate relation with Article 1, Section 1 and if and how amendments are
>   exempt from certain restrictions or need to adhere to other provisions


This constitution may be modified or additional parts and rules appended by
two-thirds supermajority of AtomOne Hub governance.


## Defined Terms

- **$ATONE** - the primary staking token for AtomOne. The ATONE token functions
  as voting shares, economic incentive shares, and security bonds for the
  AtomOne Hub.
- **$PHOTON** - the liquid staking token for AtomOne.
- **$phATOM** - the liquid staking token for Cosmos Hub Gaia offered on AtomOne
  for $ATOM (not $ATONE).
- **Additive** - when used in the context of software, it means that it does
  not require modifying existing primitives.
- **Airdrop** - a distribution of tokens to some accounts based on predefined
  criterias. May be distributed automatically at genesis or post-launch, and
  may require one or more claim transactions from the accounts, or some other
  predefined action.
- **AtomOne Hub** - an opinionated fork of the Cosmos Hub Gaia. It is a minimal
  IBC-token-pegging and ICS hosting hub.
- **BFT** - Byzantine Fault Tolerant, property of distributed systems that
  brings resilience over Byzantine faults.
- **Bonding** - in delegated Proof of Stake, Bonding is the process of locking
  one’s token to a validator node in order to help secure the network and
  attain governance voting rights. Also known as “staking”.
- **Bonded Share Unit** - a static representation of a delegator bonded stake
  in relation to the total stake of the validators to whom he/she has
  delegated. The Bonded Share Unit quantifies the delegator's claim on the
  rewards and risks associated with the validator's performance. The total
  number of Bonded Share Units held by a delegator is fixed unless changes
  occur due to redelegation or slashing events. It remains unaffected by the
  validator's total stake unless these specific actions are taken.
- **Citizen** - Cosmonaut that has staked his ATONEs toward the AtomeOne Hub,
  participating in the security of the network and in governance.
- **Community Pool** - a community pool is an allocation set aside by the
  protocol that is continuously growing through the means of a fixed percentage
  of fees on transactions executed throughout the network and inflationary
  rewards. They can be used to fund community-driven initiatives, such as
  development projects, marketing campaigns, or community events. It can be
  considered as the budget of the network.
- **Constitutional Majority** - see **Majority**.
- **Consumer Shards** - see **Shard**.
- **Core Shards** - see **Shard**.
- **Cosmonaut** - a participant of the Cosmos Ecosystem, incumbent and token
  holder of any Cosmos Ecosystem token.
- **Cosmos** - the interchain network composed of many sovereign zones
  connected by IBC.
- **DAO** - Decentralized Autonomous Organization
- **Economic incentive shares** - units of ownership or participation in an
  organization designed to align the financial interests of individuals (such
  as employees, managers, or stakeholders) with the long-term goals and
  performance of the organization. These shares provide economic benefits, such
  as dividends or profit-sharing, to incentivize behavior that promotes the
  organization's success, productivity, and profitability. They serve as a tool
  for motivating and retaining key individuals by linking their financial
  rewards to the organization's performance.
- **Fork** - in the context of blockchain, a fork is a copy of an existing
  blockchain’s distribution and software, with potential modifications and
  usually a different chain identifier than the original.
- **Governor** - a type of account that can have tokens’ governance voting
  power delegated to them.
- **IBC** - short for Inter-Blockchain Communication, is a protocol that
  enables communication between different blockchain networks using Byzantine
  Fault Tolerant (BFT) light-client proofs. It allows for the transfer of
  assets and information across independent blockchains, fostering
  interoperability and flexibility in the blockchain ecosystem. IBC is a
  cornerstone of the Cosmos network's architecture, enabling its vision of an
  'Internet of Blockchains'.
- **ICS1** - ICS1 short for Inter-Blockchain Security version one, also known
  as Simple Replicated Security, includes all protocols where the validator set
  is simply replicated across multiple blockchains, and slash conditions are
  always submitted to a root chain. 
- **ICS2** - ICS2 short for Inter-Blockchain Security version two, includes all
  protocols where slash conditions for complex failure scenarios of one
  validator set are handled by another validator set, where slashing affects
  tokens on the latter validator set.
  - **ICS2A** - includes all protocols of ICS2 where stake is entirely managed
    by the AtomOne Hub (in the form of ATONEs or other derivatives).
  - **ICS2B** - includes all protocols of ICS2 where stake is entirely managed
    by the logic of the other chain.
- **ICS-secured chain** - a chain that borrows security from the validator set
  of the AtomOne Hub by replicating it either fully or partially.
- **Interchain Accounts (if we include it)** - a Cosmos-SDK module that allows
  an IBC-connected chain to hold an account on another chain and perform
  transactions on this other chain.
- **Liberty** - “the condition of being free from oppressive restriction or
  control by a government or other power.”
- **Liquid Staking** - a protocol that allows users to receive a tradeable
  derivative token in exchange for native staking tokens that are autonomously
  bonded by the protocol or by delegated entity. These derivative tokens
  represent some claim on the underlying bonded tokens and may accrue rewards.
- **Majority** - the specified percentage of affirmative votes (YES votes)
  among bonded ATONEs required for a proposal to be approved at the end of the
  voting period. This approval is contingent upon the voting participation
  meeting or exceeding a predefined quorum value.
  - **Constitutional majority** - the Constitutional Majority is initially set
    at 90%. The Constitutional Majority cannot be made lower than 90% even with
    a Constitutional Majority, but it may be set to any value between 90% and
    100%. This elevated threshold aims to ensure broader agreement and
    inclusivity in critical decision-making processes. It reflects a commitment
    to achieving near-unanimous consensus on essential governance decisions,
    enhancing the legitimacy and stability of the outcomes.
  - **Simple majority** - a simple majority is achieved when more than 50% of
    total votes is in favor (YES votes). This threshold definition cannot be
    changed even with a Constitutional Majority.
  - **Supermajority of Two-Thirds** - the Supermajority is defined to be
    exactly "more than two thirds" (+2/3, or at least one iota more than two
    thirds) and cannot change even by a Constitutional Majority.

> TO DO:
> - Decide whether ABSTAIN votes do not count toward Majority

- **Person** - in juridical terms, a person is an entity that has legal rights
  and obligations. The concept of a person includes both natural persons and
  juridical (or legal) persons. A natural person is a human being with legal
  capacity to have rights and duties. A juridical person, also known as an
  artificial or legal person, is an entity created by law that is recognized as
  having legal rights and duties similar to those of a natural person.
- **Property** - “the exclusive right to possess, enjoy, and dispose of a
  thing.”
- **Proposal** - in the context of blockchain, a proposal refers to a document
  that requires deliberation in the form of voting from a blockchain governance
  structure. It is a formal request for funding, resources, or support to bring
  a specific project or idea to life or to make changes or updates to the
  existing software, parameters and foundational documents. A proposal may have
  a voting period in which any entity with voting power can participate, and
  after which it is considered approved or rejected depending on the tally of
  votes and some predefined thresholds. 
- **Quantum Computing** - quantum computing is a multidisciplinary field
  comprising aspects of computer science, physics, and mathematics that
  utilizes quantum mechanics to solve complex problems faster than on classical
  computers. This may have implications for blockchain technology that are
  worth exploring.
- **Quorum** - Percentage that represents the vote participation threshold on a
  proposal over the total existing voting power. . Whatever the final vote
  results are, a proposal is automatically rejected if the quorum is not
  reached. 
- **Root Hub Chain** - see **Root Shard**.
- **Root Shard** - see **Shard**.
- **Security bonds** - assets that are at stake and collateralized that are
  made non transferable for the purpose of slashing or use as collateral for
  paying victims in the event of faulty consensus votes and qualified
  governance faults or other provable misbehavior.
- **Shard** - a partition of the blockchain network that allows for
  parallel processing of transactions. In the context of Cosmos, a shard is
  simply a blockchain.
  - **Root Shard** - it is the chain where $ATONE staking and AtomOne Hub
    Governance transactions are committed and executed. It is also the chain
    acting as ICS provider.
  - **Core Shards** - ICS-secured chains that function as an extension of
    functionality of the AtomOne Hub, these consumer chains are completely
    controlled by the AtomOne Hub.
  - **Consumer Shards** - ICS-secured chains that retain full sovereignty due
    to their own governance and they are in a client relationship with AtomOne
    Hub.
- **Simple Majority** - See **Majority**.
- **Slashing** - in Proof of Stake, slashing refers to the process of
  penalizing validators and delegators for misbehaving or engaging in malicious
  activities.
- **Supermajority** - See **Majority**.
- **TreasuryDAO** -

> TO DO:
> - to clarify TreasuryDAO

- **Unbonding** - in Proof of Stake, Unbonding is the process of withdrawing
  one’s stake from its bonded state to receive fungible tokens. Unbonding is
  also the action performed in order to initiate this process, which includes
  the “unbonding period.”
- **Unbonding Period** - the unbonding period is a time frame during which a
  staker’s coins are locked on the blockchain after being unbonded and cannot
  be used in any transaction or earn further staking rewards. This is the
  estimated maximum amount of time it will take to detect misbehavior and still
  be able to slash the author of such misbehavior. The unbonding period should
  be longer than Governance voting periods in order to allow Governance to
  react in time whenever required.
- **Validator** - a validator is a crucial part of a blockchain network,
  particularly in Proof-of-Stake (PoS) consensus mechanisms. They play a vital
  role in verifying transactions and adding them to the blockchain. Validators
  are responsible for confirming the authenticity and accuracy of transaction
  records, ensuring the integrity and security of the network.
- **VM** - a virtual machine (VM) in the context of programming is an abstract
  computing machine that enables a computer program to be executed in a
  consistent, platform-independent manner. This VM provides a layer of
  abstraction between the executing program and the underlying hardware,
  typically interpreting or compiling the program's code at runtime, thus
  ensuring portability and compatibility across different hardware and
  operating systems.
- **Voting shares** - units of ownership in an organization that confer the
  right to vote on organizational matters, such as making decisions on policies
  and other significant issues that require stakeholder approval. These shares
  are typically distributed among stakeholders (e.g., shareholders in a
  corporation or members of a cooperative) and serve as a mechanism for
  democratic participation and governance within the organization.
- **ZK** - a zero-knowledge proof or zero-knowledge protocol is a method by
  which one party (the prover) can prove to another party (the verifier) that a
  given statement is true, while avoiding conveying to the verifier any
  information beyond the mere fact of the statement's truth. The intuition
  underlying zero-knowledge proofs is that it is trivial to prove the
  possession of certain information by simply revealing it; the challenge is to
  prove this possession without revealing the information, or any aspect of it
  whatsoever.
- **(Sovereign) Zone** - a (sovereign) zone is an independent blockchain (or an
  application hosted on a parent chain) with a well-defined governing body or
  bodies that dictate the governance and economic rules internal to that zone.
  By definition, a zone is sovereign or partially sovereign.
