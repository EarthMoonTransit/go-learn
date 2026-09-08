# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- A nil pointer returns zero without panicking.
- A non-nil pointer returns the exact integer it points to, including zero and negative values.
- Read the current value on every call, including after the caller changes it.
- Leave the pointed-to integer unchanged.

Public/reference assertions (actual/expected diagnostics):

- Case01: `n:=2;if got:=subject.CounterValue(&n);got!=2{t.Errorf("value got %d want 2",got)}`
- Case02: `if got := subject.CounterValue(nil); got != 0 { t.Errorf("CounterValue: got %v, want %v", got, 0) }`
- Case03: `n:=0;if got:=subject.CounterValue(&n);got!=0{t.Errorf("zero got %d want 0",got)}`
- Case04: `n:=-7;if got:=subject.CounterValue(&n);got!=-7{t.Errorf("negative got %d want -7",got)}`
- Case05: `n:=3;subject.CounterValue(&n);n=8;if got:=subject.CounterValue(&n);got!=8{t.Errorf("current value got %d want 8",got)}`
- Case06: `n:=5;subject.CounterValue(&n);if n!=5{t.Errorf("input got %d want 5",n)}`

Realistic required defect models:

- nil-one
- always-zero
- absolute-value
- zero-one
- stale-value
- destructive-read

Retained source behavior: election-day selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| stored value | 01,03–04 | always-zero; zero-one; absolute-value |
| nil zero | 02 | nil-one |
| current read | 05 | stale-value |
| storage preserved | 06 | destructive-read |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.

The stale-value defect now caches by pointer identity; only a changed value at the same pointer exposes it. Saved weak-submission.txt omits that specific reread case, while retaining nil/zero/negative/preservation cases.


## Independent review

The stale-value defect now caches by pointer identity, so first reads of unrelated pointers cannot establish changed-pointee behavior. The strong suite checks the same pointer before and after a change. The saved weak suite omitting that reread remains insufficient according to the author-provided disposable task-check report.

Reviewed the prompt, API, actual assertions, all 6 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 6 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
