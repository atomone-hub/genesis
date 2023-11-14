# Preamble

This document is being edited live here:
https://docs.google.com/document/d/1UlZJ3Ud4M0uUeBtnEHofEVgj07Sq7n56jAN_DIiNzLU/edit?usp=sharing

# Declaration of Genesis

There comes a time when there is too much disagreement among community members
about how to move forward, that it makes the most sense to have an alternative
fork running on standby to prepare for all contingencies, and also operating as
political party base.

This seems like a natural and necessary way to split the community to enable a
diverse ecosystem of specialization and coopetition and minimize friction.

TODO: add more in the spirit of the Declaration of Independence 1776.

For sake of bootstrapping we go straight into the plan, and show segments of
specifications to show the intended end goal.

# Terms

* Decentralists: github.com/decentralists is the name of the party.
* AtomOne: is the name of the hub of the party.

# Objectives

From https://x.com/jaekwon/status/1719437899429761420:

1. Define $ATOM: The staking token of the largest minimal ICS IBC Cosmos Hub
   that keeps 2/3 of ATOMs staked.
2. Hub minimalism, and most importantly without a complex VM implementation on
   the root shard, but rather forcing VMs to live on non-root shards. This is
   the only way to scale to billions of users.
3. Fix Validator incentives so that every validator is PAID to run ICS consumer
   chains and hub shards. Actually model the minimal economic model that
   describes physical reality sufficiently to be intuitive and adaptive to all
   scenarios. Every ICS pays supermajority to validators!
4. Fix governance. We already have proposals in the works on github
   decentralists DAO / governance. Contribute there so that Governance is fixed
   and we stop spending money like a flock of headless geese.
5. Fix "liquid staking". What we have isn't liquid staking, they are
   "partyhubs". What people want are liquid fungible tokens that aren't so
   inflationary, I get it. ICS scaling will fix the problem, but if you really
   want more, fix LS to delegate pro-rata, NOT have its own gov. Or, over time,
   possibly create a bicameral gov structure w/ exponentially inflating $ATOM
   stakers, and limit how $PHOTON reverts to $ATOM to prevent hostile takeover.
   VP(staked $ATOM) >= VP(all $PHOTON).
6. Adopt a Declaration of Independence and Constitution with cryptographic
   signatures.
7. Solve IBC1.5, or ICS1.5, where the validator sets are implicit, for fast
   inter-hub communication with implied IBC, WITHOUT sacrificing independent
   BFT consensus layers.
8. Create a permissioned and completely accountable and 100%
   predetermined-finite-time-delayed transparent security reporting system
   where ABSOLUTELY EVERYTHING within it eventually becomes public knowledge to
   help deal with zero day vulnerabilities and current attacks & fund it.
9. Create and support competing marketing, growth, infra, dapp subDAOs, and
   especially help them foster the best in class in Cosmos; from the user level
   down to the VM, every component should have a good selection of competition.
10. Create a team tasked with minimizing and simplifying code and reducing
    unnecessary dependencies, taking the best examples from various forks taken
    into consideration, so that all the best ideas from all forks can integrate
    into one where-ever possible. FINISH software.
11. Ossify the partyhub after it has become its own competing IBC/ICS hub.
    Allow others to likewise fork from you by enabling ICA partyhubs when there
    is disagreement. Multiply by meiosis and conquer the world.

## 1. Define $ATOM

## 2. Hub Minimalism

## 3. Validator Incentives

## 4. Governance

See
[github.com/decentralists/DAO](https://github.com/decentralists/DAO/tree/main/governance)

## 5. "Liquid Staking"

## 6. Declaration of Independence & Constitution

See https://github.com/tendermint/atom_one

## 7. IBC1.5

## 8. Transparent Security System

## 9. Fund SubDAOs

See https://gitub.com/gnolang/gno

## 10. Engineering Task Force

See https://github.com/gnolang/gno/tree/master/tm2 for Tendermint2

## 11. Enable Meiosis

See https://github.com/gnolang/gno/pull/1224 for prototype WIP 
 

It should be some distribution of the Cosmos Hub, with those who voted against
the spirit of this project removed.  Additionally, the Interchain Foundation
playing a key role in the evolution of the hub, should also be removed.  In
return, 10% of the ATOMs can go to DAOs to be managed on chain with specific
objective key results mandated.

# Plan

The AtomOne hub exists as a separate minimalist fork of Gaia. Both are separate
and distinct from gno.land, though gno.land and the GnoVM (as well as other
VMs) will play significant roles in completing the hub and maintaining its
upkeep.

The main goal is to fix what must be fixed in governance and the need for an
explicit constitution, before launching the full IBC and ICS functionality of
the chain.

First we describe the tokenomics of the atomone hub, followed by the main
milestones, with an emphasis on completion and even phase-out.

## Genesis Distribution

The Genesis Distribution is largely an opionated fork of the cosmoshub4 $ATOM
(judged by alignment based on voting activity).

The Interchain Foundation will excluded from this distribution, so as to create
a separation of concerns, and instead 10% of the final total amount will be
allocated toward contributors and onchain DAOs.

Of the 10% premine, 
 - 1% to general pre-launch contributors and early adopters.
 - 1% to reserved for IBC contributions (and all that it entails) and early
   adopters.
 - 1% to reserved for ICS1.5 contributors (and all that it entails thereafter)
   and early adopters.
 - 7% reserved for gov distribution to subDAOs for remainder of plan and
   constitution (but nothing more).

In addition to these premines, the earned tax revenue (rewards) and inflation
will be split the community pool will be allocated to the following.
 - 80% of the inflation+rewards going to the stakers who select validators.
 - 10% of the inflation+rewards going to validators equally.
 -  5% of the inflation+rewards going to pool/treasury (no governance).
 -  2% of the inflation+rewards going to open source blockchain explorer hosting.
 -  2% of the inflation+rewards going to securing open source wallet systems (w/ airgap).
 -  1% of the inflation+rewards going to public relations and growth.

 The last three following the pool/treasury will intially go to multisigs set
 in consensus params of the chain, until they get set as URIs pointing at
 blockchain based DAOs hosted on ICS1.5.

## Tokenomics

We opt to replace the market-based "commission" system with a flat
distribution to all validators, to incentivize the creation of peer validators
(who should all plan to become datacenter providers). 

The maximum bounds on the auto-adjust inflation parameter will be set at 20%.

The inflation will target 2/3 of $ATOM1 to be bonded. 

### ICS Fee Distribution

Every ICS zone should be paid for somehow. AtomOne owned ICS shards should be
paid for from the treasury of AtomOne.  Other ICS "consumer chains" can be paid
for by the the chain itself, and in emegencies anyone can step in and pay on
the zone's behalf. 

In short, every ICS zone should be profitable to every validator.

The DISTRIBUTION_FRACTION parameter is the fraction (between 0 and 1) of ICS
shard and consumer chain payments that are shared among the validators equally.
This is initially set to 0.8, giving the majority to the validators, and only
20% as royalty to be paid to ATOM1 stakers, with the COMMUNITY_TAX taking its
portion.

### Staking

XXX how does this work now? Document changes due to equality of validators.

Suppose that 1/3 of the $ATOM1 stakers are slashed due to a complex double
spend attack. Assuming that we want to allow the recompensation of victims upon
double spend attacks (within the bounds specified clearly in the constitution)
only from the recently slashed $ATOM1s, some nonzero portion of the slashed
stake must be burned to prevent using the double spend attack as a fast way to
unbond.

If no victims need to be made whole, then it could be appropriate to burn the
slashed $ATOM1s of the perpetrators. The end result is that the remaining
stakers own the network, and in a steady state this would result in the price
of $ATOM1s increasing due to the reduced supply, assuming that the convidence
and usage of atomone hasn't changed; though in perfect theory it should take a
bit of a hit, at least in proportion to the destruction of the reputation of
those validators.

If victims are to be made whole with slashed $ATOM1s, this may require the
selling of $ATOM1s into the market, or result in it, therefore the price of
$ATOM1s will be pushed lower, and the composition of the $ATOM1 holders mutated
according to market conditions.

### $phATOM the Deflationary Derivative of $ATOM1

Auto-staking (staking across all validators proportionally to existing voting
power) doesn't solve the "inflation problem", but it does give a way for people
who don't care about staking decisions to make better-than-random staking
decisions.

TODO show simplest example that demonstrates slashing.

Auto-staking if done by an external IBC zone, or an individual staker manually,
would like any other staking earn the pro-rata revenue and pay the various
taxes. So auto-staking per se does not make for a deflationary holding, but it
comes with the benefit of automatic diversification. Since it comes with some
benefits, it is better for the atomone hub to provide a standard minimal
correct implementation under its control, such that it can also regulate its
usage. More on why follows.

In order to incentivize the usage of $phATOM, the atomone hub offers a trade
that makes $phATOM deflationary: non-atom rewards nor taxes are applied to
auto-staked ATOM bonded $phATOM holders, and with the right conversion equation
(which adjusts for ATOM inflation) we can construct a perfectly fixed $phATOM
supply (say of 1 billion $phATOMs) no matter how many $ATOM1s bond to $phATOMs.

Should this "more monetary" construction of the $phATOM incentivize a large
liquid pool, it becomes more susceptible to hostile takeovers, simply because
there are more liquid $ATOM1 staking tokens available in comparison to the
total bonded voting power. Therefore for a more secure atomone hub we also
limit the conversion back from $phATOM to $ATOM1 so as to make hostile 
takeovers more expensive. (NOTE HOW is an active area of research).

In the case of validator & delegator $ATOM1 slashing, $phATOM holders will of
course also get slashed, but the ratio of $phATOM-bonded $ATOM1s and all other
(atomone hub staked or unstaked) $ATOM1s remains the same. The conversion
factor from $phATOM to $ATOM1 will change because of slashing, but the
conversion factor from $ATOM1 to $phATOM will not before and after slashing,
thereby making the total possible supply of $phATOM lower than before (more
deflationary) and over time making the cost of conversion to $phATOM more
expensive in comparison to the inverse, allowing the exchange rate between the
two tokens to naturally float between two reasonable bounds.

NOTE: This uses the market imperfection of the $ATOM1 and $phATOM tokens to
create a larger gap in the conversion price, thereby making the tokens more
independent of each other. $phATOM holders might be happy that their token has
become more deflationary (total supply reduction), and while they can only get
the post slash amount of $ATOM1s, the value of those $ATOM1s might be preserved
or catch up soon after new validators start operations. The alternative where
the total amount of $phATOM remains invariatn in comparison appears strictly
worse for $phATOM holders.

XXX conversion limitations based on periodic auctions or similar curve.
XXX comparison to liquid staking
XXX more about use as monetary token

## AtomOne Governance

Ultimately this hub is owned by the $ATOM1 holders. 

We will prioritize all of these items:
[github.com/decentralists/DAO](https://github.com/decentralists/DAO/tree/main/governance)

## Milestones

There are largely four phases to this plan.

### AtomOne Phase 1: Pre-IBC

1. Define Constitution
2. Governance Fixes
3. Launch Governance-Only Chain
4. Implement IBC

### AtomOne Phase 2: Post-IBC

1. $PHOTON with Auto-Staking
2. Fix Validator Incentives
3. Implement ICS1.5
4. Prototypes with SubDAOs (including GNO)

### AtomOne Phase 3: ICS1.5 scaling

1. Migrate $PHOTON to ICS
2. Promote Smart Contract Use Cases
3. Develop Scalable Validator Infrastructure
4. Develop Recovery Preocedures

### AtomOne Phase 4: Maintenance

1. Create OnChain Education Curriculum
2. Promote Good Forks and Projects
3. Promote Other Common Goods
4. Finalize the Software

Finalization should not be seen as a thing to avoid, but rather a necessity for
preserving immutability and thus providing real security benefits.

Everyone who wants something different is given a way to create their own
variation to compete and cooperate with the atomone hub. We should all be
familiar with this concept, as it is how atomone itself was born--by exodus
from Gaia.

It is possible that what we arrive at is not sufficient in the long run, and
that is still OK; the ultimate goal is to be a standard reference, in the very
least in relation to an improved fork; a reference that will last a thousand
years or more.

In short, the goal is nothing more than to create timeless code, even knowing
that in the end even atomone will be phased out, but never forgotten; the
template will have split into a million different forks and conquered the
world.

AtomOne Eschatology will be well documented and planned, for a time when nobody
was around for these early beginnings.

# FAQ

## AtomOne vs Gaia

AtomOne is just built different.

## AtomOne vs Gno.land

Gno.land will be a hub for GnoVM based smart contracts. It may benefit from
ICS1.5 in the future, but we will first offer GnoVM scaling on the AtomOne ICS
economic zone. Gno.land will not connect to Gaia except indirectly through
atomone, or its minimal successor.

## AtomOne & Tendermint2

NewTendermint, LLC will lead Tendermint2 development, and AIB, Inc will
contribute by building its own engineering team within AIB, Inc that will
transition into long term co-maintainers of atomone, and co-members of
gno.land. 

AtomOne is a collaborative project that includes many members including AIB Inc
and NewTendermint, LLC. Any rights to the brand of AtomOne that NewTendermint
might have (or not have) is donated to the sovereignty of the atomone hub, or
managed by AIB Inc on behalf of the hub.
