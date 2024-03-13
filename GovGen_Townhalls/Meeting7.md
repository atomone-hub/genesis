### **GovGen Townhall Meeting 7 Notes**

### Date: March 12th, 2024, 11:00 PM PST

<br> 

- Link to the [full transcript here] will go here.

- Link to the [full recording here] will go here.

**Main action items**

- Contribute to the [Constitution](https://github.com/atomone-hub/genesis/blob/a9b9d9d5a2440fb623d3bad3c672ae4754377b00/CONSTITUTION.md)
- Contribute to the [issue consolidating ideas for the Constitution](https://github.com/atomone-hub/genesis/issues/136)
- [Join our Discord](https://discord.gg/atomone)

<br> 

**Brief Summary of Discussion will go here**

- [Placeholder]
<br>

**Suggested Talking Points**

- Disclaimer to warn users to never enter mnemonics
- How to Sign offline [guide](https://github.com/tbruyelle/govgen-proposals/blob/tbruyelle/doc/submit-tx-securely/submit-tx-securely.md) 
     - Reach out to security@allinbits.com if you find any issues.
- Cosmostation, Keplr (Testnet) and Leap have integrated support for GovGen
- Mintscan explorer has added support for GovGen - https://www.mintscan.io/govgen/
- Distribution proposal draft discussion
     - Where the "blend" `B` is the sum of the relative percentages of the distribution of YES, NO, and NWV with their respective multipliers, but *not* taking into account any bonus. Given the prop848 tally results, B is ~1.9486. The multiplier `B` is computed as `B = Y + (N + V) * 4` where `Y` is the *relative* percentage of *YES* votes (i.e. the total $ATOMs that voted *YES* over the sum of total $ATOMs that voted either *YES*, *NO*, or *NWV*), and `N` and `V` are the *relative percentages* of respectively *NO* and *NWV* votes.
     - Bonus and malus are less than 5% (set to 3% currently)
     - This makes the total supply of $ATONE about 4x times higher than the supply of $ATOM as of now. [**NOTE**: might have miscounted accounts that *did not vote* at all and had not staked, we are reviewing internally, final supply could be higher]
       
     |                    |  DIDN'T VOTE  | YES | ABSTAIN | NO |    NWV    |
     |:------------------:|:-------------:|:---:|:-------:|:--:|:---------:|
     | Staking multiplier | B x malus |  1  |  B  |  4 | 4 x bonus |
     | Liquid multiplier  | B x malus |  -  |    -    |  - |     -     |

- Distribution prop draft: https://github.com/atomone-hub/govgen-proposals/pull/5
 
- Tech Details of ICS
- Liquid Staking / Photon Atom vs Photon Atone
- Urgencies in constitution for minimal working draft to move forward such as,
  - How exactly to reform governance
  - Remove NWV or remove delegations altogether
  - Remove validator commission, include fixed rate
  - Validator distribution
  - Validators have machines onsite, i.e no white labels, no AWS
  - Validators should declare conflicts of interest
  - Should not roll out fixes that validators don't understand
  - If anything to be added to constitution, should be unchallenged for a given term (3 months?)
  - Potentially funding multiple teams of leading developers to compete to build ICS
