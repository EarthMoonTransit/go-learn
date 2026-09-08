# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Set Name to empty, Age to zero and Attributes to nil.
- Changes must be visible through every pointer alias to the same record.
- The previously referenced Attributes map itself remains unchanged.
- The receiver is non-nil; clearing an already-zero record is valid and repeatable.

Public/reference assertions (actual/expected diagnostics):

- Case01: `e:=subject.Entry{Name:"record",Age:29,Attributes:map[string]string{"street":"Main St."}};alias:=&e;alias.Clear();if e.Name!=""||e.Age!=0||e.Attributes!=nil{t.Errorf("reset got %+v want zero record",e)}`
- Case02: `m:=map[string]string{"k":"v"};e:=subject.Entry{Attributes:m};e.Clear();if len(m)!=1||m["k"]!="v"{t.Errorf("old map changed: %v",m)}`
- Case03: `e:=subject.Entry{};e.Clear();e.Clear();if e.Name!=""||e.Age!=0||e.Attributes!=nil{t.Errorf("repeat reset got %+v",e)}`

Realistic required defect models:

- rebind-receiver
- name-retained
- age-retained
- empty-map
- destroy-old-map

Retained source behavior: census selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| all fields reset through aliases | 01 | rebind-receiver; name-retained; age-retained; empty-map |
| old map preserved | 02 | destroy-old-map |
| repeatable zero reset | 03 | empty-map |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public checks cover all zeroed fields through a pointer receiver, preservation of previously referenced map storage and repeatable clearing.

Reviewed the prompt, API, actual checks, all 5 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 5 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
