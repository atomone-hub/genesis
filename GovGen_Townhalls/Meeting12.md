### **GovGen Townhall Meeting 12 Notes**

### Date: Date: April 23 2024, 8am PDT
<br> 

- Link to the [full transcript](https://docs.google.com/document/d/1lLrhe0pNKTsXMAtkJgOjelKSgwvCtCwLMimteEnDh7I/edit?usp=sharing)
- Link to the [full recording](https://drive.google.com/file/d/1OgyF0myC90_Yu7hpiogLQcHR_4waOdQG/view?usp=sharing)
- Link to the [full message log](https://drive.google.com/file/d/1iGB39795fad1aVIgr1C1oPmy2k25ggcF/view?usp=sharing)


**Main action items**

- Contribute to the [Constitution](https://github.com/atomone-hub/genesis/blob/a9b9d9d5a2440fb623d3bad3c672ae4754377b00/CONSTITUTION.md)
- Contribute to the [issue consolidating ideas for the Constitution](https://github.com/atomone-hub/genesis/issues/136)
- [Join our Discord](https://discord.gg/atomone)
- [Help Define ICS](https://github.com/atomone-hub/genesis/issues/66)
- [Participate in the Atom One Design Competition](https://github.com/atomone-hub/assets/tree/main/AtomOneDesignCompetition) 
- Consider applying for [open roles](https://jobs.lever.co/allinbits)
- First Constitution Working Group call is next Tuesday at 9am.
  -   Calls will alternate starting this week

<BR>

**Brief Summary of Discussion**

- The call starts with an introduction bringing attention to the main action items including the Atom One Design competition and a call for contribution to the Constitution.
- There was a discussion about efforts to prevent dependency injections in binary distributions, particularly focusing on reproducible builds. The conversation highlighted the reproducibility of builds in Go 1.20 and emphasized the importance of ensuring consistent builds for user confidence.
- The concept of autostaking, where users stake their tokens to receive rewards, was explained. Two staking methods were detailed: regular staking and autostaking, and discussions revolved around how rewards are distributed among validators based on uptime rather than direct delegation.
- There was also conversation on ideas to ensure equal rewards for validators performing the same work, regardless of the number of delegations they receive.
- Discussions also covered maintaining fairness in rewards distribution and preventing economic attacks on the network.
- The technical implementation of a pool-based delegation system was outlined, where tokens are pooled and distributed among validators based on uptime. This approach aims to simplify delegation for users and promote network security however there were some concerns covered, like how do you define the validator set to begin with?
- Various ideas for selecting validators were discussed, including validators deciding on new additions or community-driven proposals to expand the validator set based on network needs and performance.
- The conversation continued to cover various perspectives on incentivizing governance participation, with participants discussing the importance of validation and delegation mechanisms in ensuring network success and security.
- One participant highlighted the significance of governance participation and suggested considering the the idea of incentivizing governance participation in validating the network's success.
- Another expressed skepticism towards incentivizing governance participation, arguing that not incentivizing participation might lead to more meaningful engagement and better outcomes. They emphasized the need to preserve proof of stake mechanisms while encouraging intelligent delegation.
- A third participant proposed a strategy of distributing tokens equally among good validators to ensure network security, with penalties for validators who do not adhere to the rules.
- The discussion then shifted with a conclusion wrapping up the call focused on updates of the various ongoing projects and initiatives, including the development of the GovGen dapp, a bug bounty programs, and the logo design competition.

- #### Engineering Update
  - internally reviewing the available initial [constitution.md](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md)
  - code disclaimers and reproducible build for GovGen
    - govgen uses go1.20 which provides [reproducible builds](https://go.dev/blog/rebuild) when cgo is disabled
    - need to align goreleaser and Makefile so they produce the same binary signature
    
| Go version | Debian (glibc) | Alpine (musl) | Arch (glibc) | Reproducible build |
|------------|----------------|---------------|--------------|--------------------|
| 1.20.14    | c1d25ff        | c1d25ff       |              | ✅              |
| 1.21.9     | 6bbfd3c        | 6bbfd3c       |              | ✅              |
| 1.22.2     | 27affed        | 27affed       | 27affed      | ✅           |

  - analysis of PHOTON design
  <image src="https://atomone-hub.github.io/govbox/photon-design-draft.png" width="80%">

- #### GovGen dApp Update
  - Working on going public by finalizing License
  - Working on Main Network launch for dApp
  - Platform has 100% coverage for translations
  - Ramping up additional features to launch for the dApp after Main Network
  - Building a minimalistic Calisto inspired Indexer with a very small footprint

- #### Bug Bounty
  - Launching in coming days
  - Rewards 200-10k USD per finding
  - Scope covers Govgen and the governance dApp
  - Follow on Twitter to receive a link when it launches. Rewards go to the first finder of each vulnerability.


- #### Logo Competition Update
  -   All In Bits is hosting a design competition for the AtomOne project's logo. Participants have three weeks to submit their designs, followed by a week of voting. The competition starts on April 22nd, 2024, and ends on May 31st, 2024. Last day to submit is 13th of May.  The voting period is from April 22nd to May 20th, 2024.
  -   Participants must create unique logos that reflect their creativity and innovation. They must submit full-size and short logos, each in three colors (including white and black). Participants can submit up to two entries.
  -   The total prize pool is $28,000, with four grand prizes of $7,000 each. Winners must waive their rights to the intellectual property of their designs in favor of All In Bits, allowing them to use the logos for the AtomOne project.
  -   Participants must have a GitHub account to participate and submit their entries via the dedicated GitHub repository. They must also meet eligibility requirements and agree to transfer their submission's rights to All In Bits if selected as a winner.
  -   The judging process includes community voting on GitHub and evaluation by the All In Bits team. The winning logos may be used for the AtomOne project's future brand identity.
  -   This logo will be used interim until the proposed AtomOne project runs it's own vote
  -   Overall, it's a global design competition offering participants the chance to showcase their talent and potentially win cash prizes while contributing to the AtomOne project's branding. [Here's a link](https://github.com/atomone-hub/assets/tree/main/AtomOneDesignCompetition)


- #### Working Group Form Update
  - Closed on April 20th
  - 29 respondents
  - **Time Slot**
    - The most mentioned time slots (in PDT) were: 9am, 8am, 7am, 6am, 10am, 6pm
    - The 6am slot is crucial for folks in China and South Korea
    - Of those that mentioned dates, Tuesday and Wednesdays were predominant
    - Given that information I would suggest we continue with the Tuesday slot, alternating between 6am, 9am, 6pm time slots. Maybe an experimental 2am timeslot lead by someone in Europe.
  - **Blockers from participating**
    - Many people requested the meeting time a week or two in advance to ensure they have the time slot available, I think we can improve on this.
    - People mentioned time constraints as the main blocker
  - **What would make a good constitution?**
    - **Transparency and Decentralization**: Preventing abuse of power and ensuring fairness for all contributors.
    - **Flexibility and Clarity**: Balancing flexibility to adapt to changing circumstances with clarity on core principles, avoiding overly bureaucratic language, and maintaining a clear direction for the project.
    - **Community Engagement and Support**: Supporting community initiatives, funding projects aligned with core values, and fostering inclusivity within the AtomOne ecosystem.
    - **Ethical Framework and Dispute Resolution**: Incorporating an ethical framework to guide behavior and establishing clear procedures for resolving disputes, promoting respect, equity, and fairness.
    - **Minimalism and Accessibility**: Advocating for a minimalist approach to the constitution, using simple language for accessibility, and ensuring broad understanding and participation within the community.
  - **What specific topics would you like to contribute to?**
    - **Governance and Decentralization**: Several folks expressed a desire to contribute to aspects related to governance, decentralization, and the role of validators within the ecosystem.
    - **Economics and Economic Incentives**: Many indicated interest in contributing ideas for economic incentives to increase voting participation and discussions on voting mechanisms.
    - **Community Development and Innovation** **Support**: There was significant interest in contributing to crafting guidelines and support systems for innovation within the AtomOne ecosystem, including mentorship programs and collaborations with external partners.
    - **Philosophical and Fundamental Principles**: Some respondents expressed a desire to contribute to defining clear fundamental principles that would guide the constitution and the project as a whole.
    - **Proofreading and Legal Advisory**: A notable number of people expressed interest in proofreading the entire constitution, providing input on specific topics, and offering legal advice regarding the content, structure, and design of the AtomOne Constitution.

- Define a process for deciding what time the working group calls are held
