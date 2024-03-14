### **GovGen Townhall Meeting 7 Notes**

### Date: March 12th, 2024, 11:00 PM PST

<br> 

- Link to the [full transcript here](https://docs.google.com/document/d/11i9BY3Vo6GGLlv9DmM7T0LDPS8ufRchMiqmmF_Kyuco/edit?usp=sharing)
- Link to the [full recording here](https://drive.google.com/file/d/1w1SIscLxnN1FzOdclheFq0GL9ePXCb2k/view?usp=sharing)

**Main action items**

- Contribute to the [Constitution](https://github.com/atomone-hub/genesis/blob/a9b9d9d5a2440fb623d3bad3c672ae4754377b00/CONSTITUTION.md)
- Contribute to the [issue consolidating ideas for the Constitution](https://github.com/atomone-hub/genesis/issues/136)
- [Join our Discord](https://discord.gg/atomone)
- [Help Define ICS](https://github.com/atomone-hub/genesis/issues/66)
- Consider applying for [open roles](https://jobs.lever.co/allinbits)

<br> 

**Brief Summary of Discussion**

- Disclaimer to warn users to never enter mnemonics
- How to Sign offline [guide](https://github.com/tbruyelle/govgen-proposals/blob/tbruyelle/doc/submit-tx-securely/submit-tx-securely.md) 
     - Reach out to security@allinbits.com if you find any issues.
- Cosmostation, Keplr (Testnet) and Leap have integrated support for GovGen
- Mintscan explorer has added support for GovGen - https://www.mintscan.io/govgen/
- Distribution proposal draft discussion
     - Where the "blend" `B` is the sum of the relative percentages of the distribution of YES, NO, and NWV with their respective multipliers, but *not* taking into account any bonus. Given the prop848 tally results, B is ~~1.94~~ ~2.46. The multiplier `B` is computed as `B = Y + (N + V) * 4` where `Y` is the *relative* percentage of *YES* votes (i.e. the total $ATOMs that voted *YES* over the sum of total $ATOMs that voted either *YES*, *NO*, or *NWV*), and `N` and `V` are the *relative percentages* of respectively *NO* and *NWV* votes.
     - Bonus and malus are less than 5% (set to 3% currently)
     - ~~This makes the total supply of $ATONE about 4x times higher than the supply of $ATOM as of now. [**NOTE**: might have miscounted accounts that *did not vote* at all and had not staked, we are reviewing internally, final supply could be higher]~~
     - *ERRATA CORRIGE* on last point: given that the term `B` should be "neutral" with respect to how it alters the power distribution against the NO, NWV and YES relative percentages and multipliers, the expected resulting $ATONE supply should increase of a term close to `B` with respect to the $ATOM supply. And indeed, the ratio of (as of current calculations) increase of $ATONE with respect to $ATOM is of around ~2.36x (amounting to ~809M $ATONE)
       
     |                    |  DIDN'T VOTE  | YES | ABSTAIN | NO |    NWV    |
     |:------------------:|:-------------:|:---:|:-------:|:--:|:---------:|
     | Staking multiplier | B x malus |  1  |  B  |  4 | 4 x bonus |
     | Liquid multiplier  | B x malus |  -  |    -    |  - |     -     |

- Distribution prop draft: https://github.com/atomone-hub/govgen-proposals/pull/5
     -  Be sure to open the raw version or open the file in an editor so that you can see the comments that express some guidelines on the content of each section.
 
- Conversation on Defining ICS
- Conversation on what the relationship between Photon Atom vs Photon Atone might look like, to be discussed as a community over the coming weeks.
- Conversation on the "Future of blockchain as a 3 layer cake"

     - Layer 1 (the infrastructure that allows a blockchain to run) - VAAS, scaled security - any application or blockchain can be hosted by any validator set

     - Layer 2 (Application itself) - i.e, Ethereum, Hub Logic, Liquid Staking provider, smart contract platform, any VM

          - Most popular applications that get the most users or tx/s would be smart contract systems, i.e EVM, GnoVM, CosmWasm. Smart contract systems that are good will enable the most application to be built on them.

     - Layer 3 (Dapp Layer) - Dapps or smart contracts built on top of these applications.

- Detailed converastion on "How do we design a system that allows for any blockchain application to be run more or less permissionlessly?"
     - One idea is to turn a blockchain application into a linux container image, like a docker image. Support images that AtomOne can take from a shard. 

**Other suggested talking points that weren't covered due to time**
<br>
- Urgencies in constitution for minimal working draft to move forward such as,
  - How exactly to reform governance
  - Remove NWV or remove delegations altogether
  - Remove validator commission, include fixed rate
  - Validators have machines onsite, i.e no white labels, no AWS
  - Validators should declare conflicts of interest
  - Should not roll out fixes that validators don't understand
  - If anything to be added to constitution, should be unchallenged for a given term (3 months?)
  - Potentially funding multiple teams of leading developers to compete to build ICS
