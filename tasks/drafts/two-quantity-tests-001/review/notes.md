# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Each exact solid entry contributes 50 to the integer result; each exact liquid entry contributes 0.2 to the floating result.
- Other strings contribute nothing; matching is case-sensitive.
- Return integer quantity first and floating quantity second; nil or empty input returns two zeros.
- Preserve the input. At most 10000 entries are supplied; allow absolute floating error 0.000000001.

Public/reference assertions (actual/expected diagnostics):

- Case01: `gotA,gotB:=subject.Quantities([]string{"solid","liquid","solid"});if gotA!=100||math.IsNaN(gotB)||math.Abs(gotB-0.2)>1e-9{t.Errorf("Quantities got (%d,%g) want (100,0.2)",gotA,gotB)}`
- Case02: `gotA,gotB:=subject.Quantities([]string{"liquid","other","liquid"});if gotA!=0||math.IsNaN(gotB)||math.Abs(gotB-0.4)>1e-9{t.Errorf("Quantities got (%d,%g) want (0,0.4)",gotA,gotB)}`
- Case03: `gotA,gotB:=subject.Quantities([]string{"solid","other","solid"});if gotA!=100||math.IsNaN(gotB)||math.Abs(gotB-0)>1e-9{t.Errorf("Quantities got (%d,%g) want (100,0)",gotA,gotB)}`
- Case04: `gotA,gotB:=subject.Quantities(nil);if gotA!=0||math.IsNaN(gotB)||math.Abs(gotB-0)>1e-9{t.Errorf("Quantities got (%d,%g) want (0,0)",gotA,gotB)}`
- Case05: `gotA,gotB:=subject.Quantities([]string{"SOLID","LIQUID",""});if gotA!=0||math.IsNaN(gotB)||math.Abs(gotB-0)>1e-9{t.Errorf("Quantities got (%d,%g) want (0,0)",gotA,gotB)}`
- Case06: `in:=[]string{"solid","liquid"};subject.Quantities(in);if !slices.Equal(in,[]string{"solid","liquid"}){t.Errorf("input changed: %q",in)}`

Realistic required defect models:

- one-of-each
- count-unknown
- integer-liquid
- case-normalized
- empty-default
- input-consumed

Retained source behavior: lasagna-master selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| both independent repeated quantities | 01–03 | one-of-each; integer-liquid |
| unknown/exact kinds | 02–03,05 | count-unknown; case-normalized |
| empty zeros | 04 | empty-default |
| preserve input | 06 | input-consumed |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.
