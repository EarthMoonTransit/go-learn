# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- rate is an integer from zero through 1000000; percent is finite and between zero and 100.
- Return rate multiplied by percent divided by 100 as a float64.
- Preserve fractional percentages and fractional output; do not truncate to an integer.
- Zero rate or zero percent gives zero; 100 percent gives the original rate. Absolute error up to 1e-9 is accepted.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.SuccessfulPerHour(0,100);math.IsNaN(got)||math.Abs(got-(0))>1e-9 {t.Errorf("SuccessfulPerHour: got %.12g want %.12g",got,float64(0))}`
- Case02: `if got:=subject.SuccessfulPerHour(221,100);math.IsNaN(got)||math.Abs(got-(221))>1e-9 {t.Errorf("SuccessfulPerHour: got %.12g want %.12g",got,float64(221))}`
- Case03: `if got:=subject.SuccessfulPerHour(426,80);math.IsNaN(got)||math.Abs(got-(340.8))>1e-9 {t.Errorf("SuccessfulPerHour: got %.12g want %.12g",got,float64(340.8))}`
- Case04: `if got:=subject.SuccessfulPerHour(6824,20.5);math.IsNaN(got)||math.Abs(got-(1398.92))>1e-9 {t.Errorf("SuccessfulPerHour: got %.12g want %.12g",got,float64(1398.92))}`
- Case05: `if got:=subject.SuccessfulPerHour(8000,0);math.IsNaN(got)||math.Abs(got-(0))>1e-9 {t.Errorf("SuccessfulPerHour: got %.12g want %.12g",got,float64(0))}`

Realistic required defect models:

- percent-not-normalized
- percent-truncated
- output-truncated
- zero-percent-unlimited
- integer-normalization

Retained source behavior: cars-assemble selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| zero and100percent | 01–02,05 | zero-percent-unlimited; percent-not-normalized |
| percentage normalization | 03 | percent-not-normalized; integer-normalization |
| fractional percent/output | 04 | percent-truncated; output-truncated |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.


## Independent review

Fractional percentages and fractional output remain observable; public numeric diagnostics and explicit NaN rejection protect the stated tolerance.

Reviewed the English prompt, prerequisite placement, public/reference assertions, all 5 required defects, stable API and learner payload boundaries. The adaptation preserves a substantive permitted source operation rather than only its topic; exact pinned artifact hashes, authors/contributors and MIT notice contents match the local Exercism checkout. No reference-only or generated source artifact is claimed as adapted lineage.

The author-provided targeted task-check report is technically passing with all 5 required defects detected. The reviewer inspected files and reports without executing authored code. This approves content for the lifecycle process; full batch gates remain the author's responsibility. Automated checks are not learner calibration.
