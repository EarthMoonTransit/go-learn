# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- If Battery is at least Drain, subtract Drain and add Speed to Distance.
- Exactly enough battery permits one advance; insufficient battery leaves every field unchanged.
- Speed and Drain never change. Return the updated value without changing the caller value.
- All fields are nonnegative, Drain is positive, and arithmetic fits int.

Public/reference assertions (actual/expected diagnostics):

- Case01: `m:=subject.Motor{Speed:5,Drain:2,Battery:100};want:=subject.Motor{Speed:5,Drain:2,Battery:98,Distance:5};if got:=subject.Advance(m);got!=want{t.Errorf("advance got %+v want %+v",got,want)}`
- Case02: `m:=subject.Motor{Speed:18,Drain:10,Battery:10};want:=subject.Motor{Speed:18,Drain:10,Battery:0,Distance:18};if got:=subject.Advance(m);got!=want{t.Errorf("exact battery got %+v want %+v",got,want)}`
- Case03: `m:=subject.Motor{Speed:5,Drain:7,Battery:3,Distance:9};if got:=subject.Advance(m);got!=m{t.Errorf("insufficient got %+v want %+v",got,m)}`
- Case04: `m:=subject.Motor{Speed:5,Drain:3,Battery:100,Distance:1};before:=m;want:=m;want.Battery=97;want.Distance=6;if got:=subject.Advance(m);got!=want{t.Errorf("nonzero distance got %+v want %+v",got,want)};if m!=before{t.Errorf("caller value changed from %+v to %+v",before,m)}`
- Case05: `m:=subject.Motor{Speed:2,Drain:1,Battery:0,Distance:5};if got:=subject.Advance(m);got!=m{t.Errorf("empty battery got %+v want %+v",got,m)}`

Realistic required defect models:

- strict-battery
- always-advance
- reset-distance
- no-drain
- swap-specification
- insufficient-zero

Retained source behavior: need-for-speed selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| advance fields | 01,04 | no-drain; reset-distance; swap-specification |
| exact battery boundary | 02 | strict-battery |
| insufficient unchanged | 03,05 | always-advance; insufficient-zero |
| caller copy preserved | 04 | value parameter prevents caller mutation structurally |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Required defects distinguish exact battery sufficiency, insufficient-state preservation, accumulating distance, battery drain and unchanged specification fields. The value-parameter API structurally preserves the caller value.

Reviewed the prompt, API, actual checks, all 6 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
