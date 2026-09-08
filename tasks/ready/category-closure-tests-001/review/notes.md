# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return a non-nil predicate that accepts exactly matching Category strings.
- Matching is case-sensitive and whitespace-sensitive; empty categories are valid.
- Day and Amount do not affect the result.
- Predicates from separate calls retain independent categories and can be called repeatedly in any order.

Public/reference assertions (actual/expected diagnostics):

- Case01: `f:=subject.InCategory("a");if f==nil{t.Fatal("nil predicate")};if !f(subject.Record{Category:"a"})||f(subject.Record{Category:"b"}){t.Error("predicate did not match exact category")}`
- Case02: `f:=subject.InCategory("a");if f==nil{t.Fatal("nil predicate")};for _,c:=range []string{"A","a "," a"}{if f(subject.Record{Category:c}){t.Errorf("unexpected category %q",c)}}`
- Case03: `f:=subject.InCategory("");if f==nil||!f(subject.Record{}){t.Error("empty category rejected")}`
- Case04: `a:=subject.InCategory("a");b:=subject.InCategory("b");if a==nil||b==nil{t.Fatal("nil predicates")};for i:=0;i<2;i++{if !a(subject.Record{Category:"a"})||a(subject.Record{Category:"b"})||!b(subject.Record{Category:"b"})||b(subject.Record{Category:"a"}){t.Error("predicates share captured category")}}`
- Case05: `f:=subject.InCategory("a");if f==nil||!f(subject.Record{Day:-3,Amount:-2,Category:"a"}){t.Error("unrelated fields affected match")}`

Realistic required defect models:

- always-true
- case-fold
- trim-space
- empty-rejected
- shared-category
- positive-amount
- default-day-only
- default-amount-only

Retained source behavior: expenses selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| match and nonmatch | 01 | always-true |
| exact case/whitespace | 02 | case-fold; trim-space |
| empty category | 03 | empty-rejected |
| independent repeatable captures | 04 | shared-category |
| unrelated fields irrelevant | 05 | positive-amount; default-day-only; default-amount-only |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.

Required default-day-only and default-amount-only defects independently enforce Case05 varying the irrelevant fields; merely testing default struct fields is insufficient.


## Independent review

The review gap is closed by default-day-only and default-amount-only defects; each requires varying its irrelevant field beyond starter defaults. Exact case/whitespace, empty categories and independent repeated closure calls are also covered.

Reviewed the prompt, API, actual checks, all 8 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 8 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
