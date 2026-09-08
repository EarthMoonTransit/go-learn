# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Add delta to the integer referenced by counter.
- The caller and every alias to that storage observe the updated value.
- A zero delta leaves the integer unchanged; negative deltas subtract.
- counter is non-nil and the sum fits int; repeated calls accumulate.

Public/reference assertions (actual/expected diagnostics):

- Case01: `n:=2;subject.AddCounter(&n,2);if n!=4{t.Errorf("stored value got %d want 4",n)}`
- Case02: `n:=5;p:=&n;q:=p;subject.AddCounter(p,7);if *q!=12{t.Errorf("alias got %d want 12",*q)}`
- Case03: `n:=4;subject.AddCounter(&n,0);if n!=4{t.Errorf("zero delta got %d want 4",n)}`
- Case04: `n:=4;subject.AddCounter(&n,-6);if n!=-2{t.Errorf("negative delta got %d want -2",n)}`
- Case05: `n:=1;subject.AddCounter(&n,2);subject.AddCounter(&n,3);if n!=6{t.Errorf("repeated update got %d want 6",n)}`

Realistic required defect models:

- local-copy
- replace-instead
- absolute-delta
- zero-resets
- rebind-pointer

Retained source behavior: election-day selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| addition visible to aliases | 01–02,05 | local-copy; rebind-pointer; replace-instead |
| zero delta | 03 | zero-resets |
| negative delta | 04 | absolute-delta |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public assertions observe caller and alias storage, negative and zero deltas and accumulation across repeated calls.

Reviewed the prompt, API, actual checks, all 5 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 5 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
