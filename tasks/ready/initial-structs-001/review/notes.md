# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- NewMotor copies speed and drain into the corresponding fields.
- A new motor has Battery 100 and Distance zero.
- NewTrack stores the supplied distance.
- All inputs are nonnegative integers; zero values are valid inputs.

Public/reference assertions (actual/expected diagnostics):

- Case01: `want:=subject.Motor{Speed:5,Drain:2,Battery:100};if got:=subject.NewMotor(5,2);got!=want{t.Errorf("NewMotor got %+v want %+v",got,want)}`
- Case02: `want:=subject.Motor{Battery:100};if got:=subject.NewMotor(0,0);got!=want{t.Errorf("zero inputs got %+v want %+v",got,want)}`
- Case03: `if got:=subject.NewTrack(800);got.Distance!=800{t.Errorf("track got %+v want distance800",got)}`
- Case04: `if got:=subject.NewTrack(0);got.Distance!=0{t.Errorf("zero track got %+v",got)}`

Realistic required defect models:

- swapped-fields
- empty-battery
- distance-starts-speed
- track-zero
- zero-speed-default

Retained source behavior: need-for-speed selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| constructor fields/defaults | 01–02 | swapped-fields; empty-battery; distance-starts-speed; zero-speed-default |
| track distance including zero | 03–04 | track-zero |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public cases cover copied constructor fields, independent battery/distance defaults and zero-valued motor/track inputs.

Reviewed the prompt, API, actual checks, all 5 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 5 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
