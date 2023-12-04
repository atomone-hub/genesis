_NOTE: from https://github.com/decentralists/DAO/edit/main/governance/README.md.
Prob everything there should be refactored and moved into some file in this repo._

# Improving and fixing governance on Cosmos Hub and beyond

As of today, the Cosmos Hub has arguably achieved a lot of what was its [initial vision](https://v1.cosmos.network/resources/whitepaper) and [beyond](https://hub.cosmos.network/main/roadmap/): from inter-blockchain communication (IBC) and multi-token pegging, to having a working on-chain governance system, to interchain security (ICS) in its initial version as _Replicated Shared Security_.

However, as the Cosmos Hub matures and grows - alongside the whole Cosmos ecosystem - shortcomings are becoming more and more evident, as they impact the overall security and user experience of the network. _Governance_, in particular, is a key aspect of Cosmos which uniquely positions it in the blockchain industry. The ability to vote on-chain on various topics, from funds allocation to protocol-level changes and chain upgrades, is undoubtedly powerful. But the lack of a written constitution that defines a common vision and guiding principles, and can serve as an arbiter when there is any doubt, is surely a major missing piece. On top of that, there are more technical and practical changes that we deem essential for improving and fixing the on-chain governance system. Among these there are for instance implementing mechanisms for spam prevention and switching over delegation-less voting for most governance proposals.

This document aims to outline a roadmap that addresses the issues we have identified and sets the direction for future improvements in the governance of the Cosmos Hub and beyond. The key priorities in this roadmap include extending the voting period for late quorum achievement, implementing a proposal deposit auto-throttler, separating staking and governance, dynamically adjusting the quorum, and introducing 2/3 supermajority proposals.

## Priorities and Roadmap

As we strive to improve the governance of the Cosmos Hub, we have identified several key areas that require attention. These priorities have been chosen based on their potential to significantly enhance the effectiveness of the governance system and the overall user experience.

For each of these priorities, work is either already in progress or planned to start as we move along in the list. The roadmap for these improvements is presented in order of planned execution, although things may evolve and priorities may change as the ecosystem further develops and matures.

### 1. Extension of Voting Period for Late Quorum Achievement

In the current implementation of the `x/governance` module in the Cosmos SDK, a governance proposal that reaches quorum near the end of the voting period may not have sufficient time for discussion and deliberation among stakeholders.

The majority of community interest in a proposal arises when the vote quorum is reached. This is because a proposal's outcome only becomes valid after achieving quorum. Reaching quorum near the end of the proposal voting period could therefore be a symptom of low awareness about the proposal or could potentially be a form of governance manipulation.

To address this issue, we propose an extension mechanism for the voting period. By enforcing an extension when quorum is reached too close to the end of the voting period, we ensure that the community has enough time to understand all the proposal's implications and potentially react accordingly without the worry of an imminent end to the voting period.

The work-in-progress text of the related _Architecture Description Record_ (ADR) can be found [here](../adrs/voting-period-extension-late-quorum.md), while the code can be found [here](https://github.com/allinbits/cosmos-sdk/tree/giunatale/late-quorum-vote-extension). Both are still subject to changes and improvements.

### 2. Spam Prevention: Proposals Deposit Auto-Throttler

The Cosmos Hub governance - as well as other Cosmos chains - has been recently suffering from governance spam, with several proposals being submitted that are aiming at spreading misinformation and scams. While ultimately the impact for the chain is only an increase in the governance proposal index, the majority of these proposals contain harmful links that could potentially pose a risk to unsuspecting users.

Despite some mitigations have been designed, like the [initial deposit requirement for proposals](https://github.com/cosmos/cosmos-sdk/pull/12771) - waiting for the Hub to upgrade to v0.47 of the Cosmos SDK to be deployed - and a more minimal [temporary mitigation](https://github.com/cosmos/gaia/issues/2246) that was deployed in the _v9-Lambda_ upgrade of the Cosmos Hub, we believe a better system could be put in place. Moreover, the deployed temporary mitigation seems to be largely ineffective, as spam proposals are still being regularly submitted, albeit arguably in lower quantities.

Using a similar mechanism to the auto-adjusting inflation rate that targets a 2/3 bonding ratio, we can auto-adjust the deposit amount (between reasonable minimum and maximum bounds, just like the inflation rate) to target having on average *N* proposals active at any time, with *N* being a low number, e.g. 1 or 2.

More details on the idea and some initial considerations can be found on a Cosmos Hub Forum [post](https://forum.cosmos.network/t/governance-proposal-deposit-auto-throttler/10121).

### 3. Staking and Governance Separation: Introduction of Delegation-less Voting

The objective is to separate the Proof of Stake and governance layers, eliminating governance delegations and switching over *direct voting* for most cases. 

In the current system validators hold too much political weight. However, validators should not be politicians, generally speaking. They should be chosen because of their technical excellency, and their ability to efficiently serve the blockchain to maintain high levels of availability and security.

In practice, two kinds of governance proposals - from a high level perspective - would exist:

1. **Urgent infrastructure and/or validator-related proposals**: In this scenario, validators would retain their delegated governance voting power. These proposals would carry over similarly to what happens today with governance delegations mirroring staking delegations.
2. **Every other proposal**: Only direct voting would be allowed, with no delegations. However, this would require a mechanism for _quorum auto-adjustment_ that would account for the inevitable lower participation (in terms of voting power) with delegation-less voting.

This would also require a governance body that would be in charge of checking if proposals are correctly categorized, and possibly punishing proposers otherwise.

### 4. Quorum Threshold Auto-Adjustment

The idea would be to linearly decrease the quorum threshold when consecutive proposals fail because they do not meet quorum.
This is a mechanism that would allow us to find the proper quorum threshold in an environment where only direct voting is allowed, and as a consequence, lower participation (with respect to total voting power) is to be expected.

A specific governance proposal could then be used to set the quorum threshold to a higher value should the need arise.

This mechanism would also work in tandem with the spam prevention *deposit auto-throttler*, meaning that generally speaking no more than a few proposals - for instance, just 2 - would be active at any point in time.


### 5. Introduction of 2/3 Supermajority Proposals

We aim at introducing an option to allow proposals to pass only if a 2/3+ supermajority votes *Yes*, and consequently remove the *NoWithVeto* vote option for these proposals as it would be not needed.

The supermajority pass threshold might be only used on certain kinds of proposals - e.g. constitution changes - or enacted for all proposals. This is yet to be fully decided and is subject to further discussion.

## Conclusions

In conclusion, the roadmap outlined in this document presents an approach to improving and fixing governance on the Cosmos Hub and beyond. It is clear that while the Cosmos Hub has achieved significant milestones, there are areas that require attention and improvement as the ecosystem continues to mature and grow.

The enhancement of on-chain governance mechanisms and the active involvement of the community - as well as the establishment of a written constitution - are all critical steps towards creating a more robust, secure, and democratic governance system. These improvements will not only address the current shortcomings but also ensure that the Cosmos Hub is well-equipped to handle future challenges and opportunities.

This roadmap is a living document and will continue to evolve as the Cosmos Hub and the wider Cosmos ecosystem grow and develop. We welcome feedback and suggestions from the community as we work towards these goals.
