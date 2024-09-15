# ATOM ONE CONSTITUTION 

_TODO: add phATOM from README.md_

_This document is a work in progress. This document assumes familiarity with
the current workings of cosmoshub4 as of Oct 11th, 2022. What is described here
are modifications to what already is. This clause will be removed with future
revisions, and the corresponding parts of the document updated with a full
description of the constitution of the hub._

## Preamble

We, the people of the Cosmos, in order to create a free world, enable voluntary
and borderless transactions, facilitate permissionless innovation, ensure
economic security, cater to economic and technological development, allow for
the creation of sovereign zones, and maintain order among sovereign zones, do
ordain and establish this Constitution for the AtomOne Hub.

## Part 0 Definitions

Cosmos is the interchain network composed of many sovereign zones connected by
IBC.

**A zone** is an independent chain (or an application hosted on a parent chain)
with a well-defined governing body or bodies that dictate the governance and
economic rules internal to that zone. By definition, a zone is sovereign or
partially sovereign. A treasury DAO of the AtomOne Hub is partially sovereign.

**IBC** is short for Interchain Blockchain Communication, and includes all
protocols that allow one chain to communicate state or messages with another
chain by tracking the consensus state of the other.

The ATOM is the primary staking token of the AtomOne Hub.

**ICS** is short for Interchain Security and includes all protocols that allow the
consensus of one chain to be partially or wholly secured by mechanisms on
another chain.

**ICS1** also known as Simple Replicated Security, includes all protocols where the
validator set is simply replicated across multiple blockchains, and slash
conditions are always submitted to a root chain.

**ICS2** includes all protocols where slash conditions for complex failure
scenarios of one validator set are handled by another validator set, where
slashing affects tokens on the latter validator set.

**ICS2A** includes all protocols of ICS2 where stake is entirely managed by the
AtomOne Hub (in the form of ATOMs or other derivatives).

**ICS2B** includes all protocols of ICS2 where stake is entirely managed by
the logic of the other chain.

**Auto-staking** is staking across all current validators in proportion to their
voting power. For example, if a validator that had 10% of the voting power were
to get slashed 30% on the AtomOne Hub, and 50% of ATOMs were either staked onto
the AtomOne Hub or free (not auto-staked), everyone who auto-staked ATOMs
on the AtomOne Hub would get slashed 1.5%. Inflationary ATOMs are paid to
auto-stakers such that they do not suffer from the inflation rate of ATOMs.

**A two-thirds supermajority** is where more than 2/3rd of all participating staked
ATOMs vote YES and the vote is above the current governance quorum value and
the voting period has concluded. ABSTAIN votes do not count toward the 2/3.

## Part 1 General Provisions

### Article 1.A. Fundamental Principles

This Constitution of the AtomOne Hub, hereinafter “the Constitution” hereby
establishes the foundations of the governance model, economic model, and
operating system of the AtomOne Hub.

All subsequent governance proposals must align with the provisions of
the Constitution, and proponents of each proposal, along with all active
governance voters, are required to ensure consistency between such proposals
and the Constitution.

### Article 1.B. Sovereignty of the AtomOne Hub

The Cosmos interchain network is composed of many sovereign zones, such as the
AtomOne Hub.

The AtomOne Hub is composed of many chains including the root hub chain where
staking and governance transactions are committed and executed, but also other
chains secured by ICS that are subservient to the governance of the root hub
chain, hereinafter referred to as "Hub Governance".

Other sovereign zones that are completely or partially secured by AtomOne Hub
ICS, by definition, have their own governance mechanism. The AtomOne Hub
principally enables and follows the will of the governance of such sovereign
zones regarding the pegged tokens originating from said zones, except in
well-defined exceptional circumstances involving bugs, theft, or harm to the
Cosmonauts of the Cosmos ecosystem.

### Article 1.C. General Mission and Objectives

The mission of the AtomOne Hub is to create a new world, enabling permissionless
yet secure interactions between the Cosmonauts of Cosmos in the Solar system.

The objective of the AtomOne Hub is to provide a classical BFT proof-of-stake
multi-token payment and transfer system; and to scale the security of the
platform to many applications hosted in other zones via AtomOne Hub ICS.

### Article 1.D. Cosmonauts

Every person has the right to become a Cosmonaut and can freely engage on the
AtomOne Hub. As such, every Cosmonaut has the right to own at least one address
on the AtomOne Hub.

Any Cosmonaut can also become a Citizen of the AtomOne Hub by using their
address to stake ATOMs toward the AtomOne Hub and participate actively in
governance. The status of citizenship is granted autonomously.

### Article 1.E. Rights, Liberties, and Obligations in the AtomOne Hub

The Liberty and Property of all Cosmonauts engaging in the AtomOne Hub is hereby
guaranteed. Any restriction to the Liberty and Property of Citizens on the
AtomOne Hub can be done only through the AtomOne Hub's governance.

Every Cosmonaut has the right to receive benefits from their engagement in the
AtomOne Hub, including rights derived from held or staked ATOMS in line with the
provisions of this Constitution.

Every Citizen allows the Governance of the AtomOne Hub to restrict their staked
ATOM property by partial or full slashing according to their voting activity.

## Part 2 Governance

### Article 2.A. The AtomOne Hub Chain

The root hub chain of the AtomOne Hub is uniquely identified by chainid
"cosmoshub". This chain commits and executes transactions that serve the
following functions:

 * governance voting
 * intra-hub token transfers
 * IBC token transfers
 * ICS1 and ICS2 management

### Article 2.B. AtomOne Hub governance

The working language of AtomOne Hub governance is English.

The quorum necessary for a proposal to be valid shall depend solely on the
number of bonded ATOMs.

The governance process must extend the voting deadline to ensure a minimum of 2
weeks of voting after the minimum quorum has been met.

UX interfaces that present the results of voting on governance proposals should
also display the content of the memo field of each voting transaction such that
the reason for the vote can be seen.

TODO: fill in the rules of cosmoshub4 governance.

### Article 2.C. Air-drops and forks

Every Cosmonaut allows any other Cosmonaut to create full or partial airdrops
of new tokens to any chain using the distribution of any token on the Cosmos
Hub at any time.

Every Citizen allows any Cosmonaut to modify their pro-rata airdrop portion
by partial or full slashing based on their cryptographic voting activity
according to well defined principles at any time.

### Article 2.D. Treasury DAOs

The AtomOne Hub governance may establish one or more transparent and accountable
Treasury DAOs by simple majority vote.

The operations of the Treasury DAO must operate on an ICS secured zone outside
of the hub. As an exception, the tokens of the Treasury DAO may reside on the
the AtomOne Hub as a m-of-n multisig account where n is at least 3 and m is at
least 1/2 of n, where each signatory is an authorized member of the Treasury
DAO and Citizen of the AtomOne Hub.

The Treasury DAO shall be composed of Cosmonauts, and at the top Executive
Board level be composed of one or more Citizens. All Cosmonauts and Citizens of
Treasury DAOs including their Oversight Committee must have public and known
real personal identities. 

To enable the well-functioning of Treasury DAOs and the separation of powers in
the utmost interest of the AtomOne Hub, each member can hold just one type of
role within each Treasury DAO.

The members of the DAO must perform efficiently in their roles in line with
their job description. They are accountable to each DAO’s Oversight Committee
and the Hub Governance. They can be dismissed from their functions by a
two-thirds majority vote by the DAO’s Oversight Committee or the Governance
Hub.

The role and functions of the Executive Board shall be further developed
through governance proposals within each DAO, unless the DAO chooses to defer
to AtomOne Hub governance by Citizens.

Each Treasury DAO shall have an Oversight Committee composed of any number of
Cosmonauts. The Oversight Committee must have at least the right to freeze all
transfers of tokens from the Treasury or its designated m-of-n multisig account
on the AtomOne Hub.

## Part 3 Economics

### Article 3.A. Economic model

The one and only economic incentive model of the AtomOne Hub is the collection
of market-based transaction fees from a large number of transactions across all
the chains secured by the staking of ATOMs on the AtomOne Hub, including ICS1
hosted blockchains.

### Article 3.B. The ATOM Token

The ATOM functions as voting shares, economic incentive shares, and security
bond for the AtomOne Hub.

To preserve the security and identity of the acting governance and validator
set, the inflation rate of the ATOM is made to vary over time to target 2/3 of
all ATOMs. The maximum inflation rate is 20% non-compounded per year. There is
no minimum inflation rate, and it can even be negative (deflationary).

Inflated ATOMs are paid to bonded ATOM holders in proportion to each
delegator's staking amount, and staked ATOMs are converted to Bonded Share
Units. 

The Atom Unbonding Period shall be 3 weeks, and redelegation is allowed twice
per Atom Unbonding Period.

Double signing at any height/round/step results in slashing penalty that is
proportional to the total amount of double signing by all validators for that
height/round/step, with evidence collected during the Atom Unbonding Period;
the penalty shall range from +0% to 100% of the Upper Slashing Limit in linear
proportion, the latter when 1/3 of voting power double-signs.

Complex signing failures (those that require +1/3 to coordinate) shall result
in slashing the Upper Slashing Limit.

The Upper Slashing Limit shall be 50%. This parameter may be increased by a two
thirds majority of the AtomOne Hub. In cases where there is sufficient evidence
of malice and intent, this parameter may be overruled by a simple majority of
the AtomOne Hub on a per-case basis up to 100%.

Liquid staking may only be supported through interchain accounts (aka
non-native liquid staking).

To limit the amount of liquid staked tokens so as to reduce systemic risk from
liquid staking, there shall be imposed a 5% tax (the "Liquid Staking Tax") on
all rewards paid out to staked interchain accounts at time of reward
withdrawal. This Liquid Staking Tax parameter may be adjusted by a two-thirds
supermajority vote of the AtomOne Hub. 

If the amount of ATOMs staked using interchain accounts exceeds 20% (the
"Liquid Staking Factor") of the total staked ATOMs, the Liquid Staking Tax
shall automatically increase by 1% per month. The Liquid Staking Factor
parameter may be adjusted by a two-thirds supermajority vote of the AtomOne Hub.

TODO: simplify the above two rules.

### Article 3.C. Intentionally left empty

(note: formerly an article on the PHOTON token)

### Article 3.D. Inflation

Any inflation of ATOMs to the Community Pool or a designated Treasury DAO
beyond the default inflation rate described in the Constitution shall require a
two-thirds supermajority vote of a special inflation governance proposal type.

The special inflation proposal can include a description of the purpose of the
inflation, but it cannot include any other modifications to the AtomOne Hub or
its Constitution, nor the adoption of any new Treasury DAOs.

### Article 3.E. The Common Pool 

The Common Pool tax proceeds shall apply to transaction fees and inflationary
ATOMs, and shall be sent to the Community Pool.

The Common Pool Tax rate shall initially be 2%, but can be increased up to 50%
by two-thirds supermajority of the AtomOne Hub governance.


## Part 4 Final Dispositions

## Article 4.A. Updates to the AtomOne Hub

New updates to the AtomOne Hub should be broken down into independent components
and discussed/proposed separately with adequate time between, regardless of any
omnibus whitepaper. 

## Article 4.B. The Implementation

The AtomOne Hub shall not have any VM functionality, but shall be plainly
implemented in a single garbage-collected language as reference (namely Go);
and other clients may implement all or portions of the stack in another
language like Rust.

The only cryptographic assumptions allowed to be used by the AtomOne Hub,
including its consensus protocol shall be Ed25519 and Secp256k1 elliptic
curves, and RIPEMD160 and SHA256 hash functions.

No zero-knowledge proof systems may be adopted on the AtomOne Hub even if they
are composed of the approved primitives.

The rules of this article may only be changed by two-thirds supermajority vote
of the AtomOne Hub.

## Article 4.C. Compute/storage/memory limitations

For the sake of decentralization, accessibility, accountability, and security,
the AtomOne Hub and each ICS zone shall be restricted so that each can run on
a commodity computer.

## Article 4.D. Amendment of the Constitution

This constitution may be modified or additional parts and rules appended by 
two-thirds supermajority of AtomOne Hub governance.

<hr />

# COMMENTARY

_This is not part of the Constitution_

## Comments from Jae Kwon

### About the economic model

The notion that ATOM is a "memecoin" ignores the obvious and original business
model for the hub: token transfer fees. Bitcoin and Ethereum gas transaction
fees are in the 10s/100s of millions, and we haven't even gotten to VISA scale
yet. ATOM is not money, it's VISA shares, IBM shares, and FED shares (but where
ATOM stakers are general partners rather than limited partners).

It's an alternative to the status quo that Bitcoin originally wanted to be, but
more. Well, imagine what kind of social manipulation we must be under, to be
pursuing such a dream.

The best part is we've done most of the work already. With minimal ICS the
simple-transfer-zones are already more or less done. We're 90% done with
massive scale MVP, and after that scaling will be relatively easy. AND this
AtomOne hub is a minimal hub that zones will want to use. The product market fit
is already there. It's simple, and we are already positioned for it. It is
neutral to application zones that provide more functionality than token
transfers.

As IBM's CEO once said, the secret cash cow of IBM is transaction processing.
> IBM Mainframe=FT, Tendermint=BFT

Cosmos is the VISA network built upon this decentralized BFT mainframe system.
Always was, and should remain.

New functionality can always be permissionlessly added on top of this base
AtomOne framework. The gno.land prop69 #exitdrop is a demonstration of value-add
to the AtomOne Hub, as it will provide Gnolang smart contracts while IBC pegging
to the AtomOne Hub for tokens.

Using ICS, it should be possible to run new Gnolang VM powered zones secured by
the AtomOne Hub, but also IBC connected to the gno.land chain for importing
logic hosted on the gno.land "github" (and paying gas fees & license fees to
each).

### The need for hubs

Imagine there are 10,000 zones. Now, if a zone fails and requires manual
intervention, having 10,000 IBC connections means needing agreement from all
10,000 zones on the recovery procedure, which will never happen. But, a zone
connected to a more secure hub will be protected when it needs intervention.

Another need for hubs; uniformity of guarantees. You need a hub to coordinate
shard zones where governance/policy and staking are applied to all shards.
Otherwise, you don’t have one system of guarantees, you have many independent
chains. Need to scale sendtx, might as well ICS.

### About security, and the need for ATOM/PHOTON separation

_UPDATE: My thinking on this section has evolved since PHOTON can be
implemented as auto-staking on top of ICS1_.

The staking ratio today on Ethereum PoS is 12%. With massive adoption, unless
we have complete laymen involved in staking, and with ETH becoming money, the
stake ratio should fall even lower, perhaps even to <1%. At that point it
becomes easy to coordinate a fund to take over the consensus process of the
chain. PoW networks have two “tokens” the mining infrastructure (which can be
bought or sold, and also is “inflationary”) and the coins themselves. This
separation allows Bitcoin to become widely adopted without worrying about
security vulnerabilities, because even the largest of whales cannot simply buy
1/2 of mining infrastructure. It isn’t a superfluid market, which makes it more
secure.

(In biology, it’s the difference between Eukaryotic (cell nucleus) and
Procaryotic (no nucleus) cells. Evolution has proven that multicellular
(inter-cellular) systems like us are generally Eukaryotic. They both exist, but
complexity demands more intracellular security.)

Imagine how easy it would be to create a fund and simply buy VISA… well not
even buy, but simply bond the capital of the market cap of VISA, of $391B.
That’s a lot of money, but if bonding that capital means one can take control
of the financial system, people would lend their money in a heartbeat. But
thankfully VISA shares are not money, and there are probably plenty of
shareholders who don’t want to sell.

Once there are ATOMs and PHOTONs, and the hub finds good ways to incentive
PHOTON usage, we will end up promoting PHOTON more than ATOMs, and the market
cap of PHOTON should theoretically eclipse that of the market cap of ATOM.
There’s much more money in circulation than the market cap of VISA/IBM/FED
combined.

Finally, with the rollout of ICS1 and ICS2, simple-transfer zones and other
application zones, the inflation rate of the ATOM may even become negative to
maintain the target of 2/3 bonding. 

Liquid staking somewhat usurps the point of bonded staking, and thus by nature
its utility is limited. It is already supported through the usage of interchain
accounts, so that is all that needs to be done to support it. Rather we should
limit liquid staking and other systemic risks by limiting how much can be
bonded through interchain accounts.

### About consensus-driven investments

Global governance consensus driven investments will fail worse than local
competition-driven investments, like central planning fails. Wherever
possible, we should ensure that intelligence is preserved or amplified in
decisions. The way to ensure that good decisions are rewarded and bad decisions
punished, is to require individual decision makers to put skin in the game.
This is why innovation happens in the private sector, and why governance
funding is seen as a corruption of private sector innovation, and why central
government planning historically has led to failure. It turns the incentive
model of individual merit, into a game of politics. This is true even when
decision making is weighted by relative capital.

We can see this clearly in the private sector investment world. The best
performing funds do not have their decisions made by weighted voting of LPs.
Rather, the LPs are free to join and leave, while the investment thesis of each
fund is maintained by select GPs. The ATOM2.0 tokenomics model is akin to
taxing all investment funds and putting the proceeds into a giant super-fund
controlled by LPs. If this were to happen in the real world, the super-fund
would create such a large distortion of incentives as to destroy innovation in
general. It would turn into a game of media/mind/political control, and actual
innovators would fail to get the funding they need, and even if they did get
funding, the entire private sector would become swamped with the resulting dumb
money, making it harder for innovators to compete with incumbent politicians.
The world would not accept such a policy toward central planning, and we should
not accept it either, as it will lead to sure failure not only of the Cosmos
Hub, but of the entire ecosystem.

This Constitution defines Treasury DAOs that have local decision making
authority, where Treasury DAOs compete with each other.

### How to immunize against the mark of the beast

The powers of the world, as represented by the WEF, are intent on implementing
the mark of the beast. As per the Book of Revelation,

* The nations of the world were deceived by pharmakia/medicine (Revelation 18
  23)
* The mark (in original Koine greek, a needle prick) is required to buy or sell
  (Revelation 13:17)
* The mark gives you sores (Revelation 16:2) // NOTE: have no fear about it
  even if you got the shot.

How could it be that a two millennia old document can predict what is happening
today? Well, most people don't read the bible at all, and anyone can see the
light and turn into a white-hat. It appears to me that the white-hats have
shepherded the black-hats into following a script that ends up exposes them
when it is too late. And now the "true Christians" have indeed exposed the NWO
agenda, and this awareness is growing exponentially.

This control grid was leaked by whistleblower Senator Larry McDonald in the
70's, whose plane was soon after shot down.
[(ODY)](https://odysee.com/@Commentator:e4/Former_US_Congressman_Larry_McDonald:1)[(TWT)](https://twitter.com/Xx17965797N/status/1578662395358384128?s=20&t=MrwxzKymkKv6ehdfhvKlAA).
The "monolithic and ruthless conspiracy that relies on covert means" was leaked
before by JFK who was assassinated in 1963.
[(ODY)](https://odysee.com/@Real_Solutions:b/JFK's-Monolithic-Conspiracy-Revelation:7)[(YT)](https://www.youtube.com/watch?v=RhkjYJAHCjM).
Now we have experienced the NWO control grid by the WEF, and its young global
leaders, such as Fauci, Gates, Gavin Newsom, Trudeau; and even experiencing the
war between Zelenski and Putin. The WEF, whose leader Schwab boasted about
having infiltrated government cabinets, also wants us to "own nothing and be
happy".  If it isn't clear enough, their logo even includes a subtle 666.

The fact of the matter is, we probably do want some regulatory system to deal
with large scale theft of coins resulting from bugs or human error/malice. Even
if such a regulatory system is not imposed upon any zones, zones may want to
voluntarily adopt some kind of regulatory policy. And zones probably want to
enforce these policies across zones that choose to adopt the same policies.

It follows that, in a minimal system, zones should be allowed to choose their
own set of regulatory policies, and the Hub can help enforce these policies
when it comes to IBC transfers across zones, or from hub to zone. From the
perspective of the Hub, this is still a permissionless, voluntary system.

When it comes to transactions on the hub, and transactions from zones to the
hub, we should adopt the most minimal regulatory system. We could arguably do
nothing--until it is too late, and we learn our lesson that we need *something*
in the case of large scale theft from malice or bugs. The minimal nonzero
policy we can adopt, is to enable one or more bonded DAOs to designate one or
more addresses as being affected, with full justification, to temporarily
freeze those coins, where the coins will unfreeze automatically until the hub's
governance votes to act upon the frozen coins or to freeze the coins for longer
to determine the facts. Something along these lines is a minimal regulatory
system. Also, it is probably sensible for the hub to implement a kind of
delayed transfer system, so that accounts with large token amounts can be
protected by this regulatory system in the case of theft. Perhaps accounts can
opt out of this transfer-delay protection.

The above minimal regulatory system still begs for a full system of checks and
balances. That's another rabbit hole for another issue/document. The AtomOne
constitution hints at a system of checks and balances, but IMO it isn't
complete yet, at least not in its current written form.

The challenge is to (a) further refine and minimize the aforementioned hub
internal regulatory policy, and (b) to define the inter-zone permissionless
regulatory framework. With these implemented, the hub can ensure the rights to
property, protect property in the case of theft, and allow zones to
permissionlessly set their own policies. This is a critical
architecture/constitutional/regulatory problem we should be discussing today.
Until it is implemented, IMO crypto will not be ready for the general
population.

The problem is, getting this wrong is _significantly worse_ than not
implementing it. When done wrong, either it will become destructive (tokens
being frozen/stolen by the regulatory system), or it will become abusive (think
1984 global dictatorship as per the NWO agenda, followed by global human
depopulation). It follows that the AtomOne Hub should not implement a regulatory
framework such that it can allow the permission innovation and proofing of
regulatory frameworks over time. This best increases the chances of success in
designing the right regulatory framework.

A Rule should be added to the constitution to succinctly represent the above
paragraph.
