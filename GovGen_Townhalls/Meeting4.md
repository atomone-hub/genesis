### **GovGen Townhall Meeting 4 Notes**

### Date: February 20th, 2024, 10:00 AM PST

<br> 

- Link to the [full transcript here](https://docs.google.com/document/d/1qM0-xYqBg3KqIV0cQD353Axa31BBdeoR6cq1igsJB1g/edit?usp=sharing)

- Link to the [full recording here](https://drive.google.com/file/d/1jJOYRJyz3ebX4-LK06d8fUsHEsTIm9we/view?usp=sharing)

**Main action items**

- If you're interested in becoming a validator, make a PR on the [validator directory](https://github.com/atomone-hub/validator) following the process listed on there. Please submit your GenTx into the validators' repo following this readme [insert GenTx readme]

- Take a look at this code for [GovGen Genesis Creation](https://github.com/atomone-hub/govgen-genesis) and the [GovGen Chain Software](https://github.com/atomone-hub/govgen), submit an issue if you have something you want to bring up.

- [Join our Discord](https://discord.gg/atomone)

<br> 

**Brief Summary of Discussion**

- Validator Onboarding Process & Timeline
  - Tuesday, 20th Feb- publish the Genesis - link will be up in a link in the README of https://github.com/atomone-hub/govgen-genesis
  - Friday 23rd, 5PM PST- validators submit your Gentx cut off time
  - GovGen launch - Tuesday 27th, 6 am PST
- Technical Updates and Governance Discussions
  - Working on x/gov changes as discussed in last townhall ([GitHub EPIC](https://github.com/atomone-hub/govgen/issues/6))
    - Forked x/gov to implement discussed modifications
        - https://github.com/atomone-hub/govgen/pull/4    
        - https://github.com/atomone-hub/govgen/pull/9
    - Remove ability for validators to vote
        - https://github.com/atomone-hub/govgen/pull/10
    - Allow different voting periods for different types of proposals
      - https://github.com/atomone-hub/govgen/pull/16
  - Updated base genesis to reflect software changes
    - https://github.com/atomone-hub/govgen-genesis/pull/1
    - 365 days voting period for text, 14 days for parameter change and 28 days for software upgrade
  - Potential discussions:
    - [refund text proposal deposit before the end of the voting period](https://github.com/atomone-hub/govgen/issues/8)
    - [consider disallowing validator vote on text proposal only](https://github.com/atomone-hub/govgen/issues/15)
- Discussion on rationale for constitutional chain, included diving into the topic of Atom inflation in relation to proposal 848
    - A couple references to [$ATOM Must NOT be Money](https://forum.cosmos.network/t/proposal-set-max-inflation-at-10/11841/240?u=aib)
- Discussion on intricacies of calculating rewards from staking tokens and points made against treating them solely as income for tax purposes. 
- Discussion on majority of rewards being more akin to token splits 
- Discussion on current tax guidance and efforts to communicate with regulators about the unique nature of staking tokens. Emphasis on  the importance of resolving these issues for the security and functionality of blockchains, drawing parallels with corporate shares as a control mechanism.
- Discussion on validator rewards model where validators are paid equally to promote decentralization and efficiency. Jae also acknowledges the need to address flaws in the current incentive mechanism of the Cosmos Hub including the need to change the incentive model for validators to ensure their participation in all ICS chains. 
- The conversation delves into the political dynamics within the Cosmos ecosystem, with discussions on the dominance of certain parties and the need for more competition, with suggestions of creating a new political party to counterbalance existing power structures, i.e "Stride" and hence, "AtomOne.  
- The discussion also touches on the role of ICS in scaling consumer chains and the potential limitations of targeting specific customers like Stride.
- Discussion continues into the purpose of ICS and the need for scalability in blockchain ecosystems, highlighting the importance of usability and programming languages in smart contract platforms.
- The discussion concludes with plans for validator applications and the upcoming GovGen launch.


<br>
