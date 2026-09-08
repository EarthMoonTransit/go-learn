# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Double returns value plus itself, preserving its type.
- The private numeric constraint must admit int, int64 and float64 and named types with those underlying types.
- Do not admit strings or other numeric kinds. Use a union of approximated underlying types.
- Inputs are finite and doubling does not overflow. Both negative and zero values are valid.

Public/reference assertions (actual/expected diagnostics):

- Case01: `type units int;type total int64;type ratio float64;if got:=subject.Double(units(-3));got!=units(-6){t.Errorf("named int got%d",got)};if got:=subject.Double(total(5));got!=total(10){t.Errorf("named int64 got%d",got)};if got:=subject.Double(ratio(.25));got!=ratio(.5){t.Errorf("named float got%g",got)}`
- Case02: `if subject.Double(0)!=0||subject.Double[int64](-8)!=-16||subject.Double(1.25)!=2.5{t.Error("numeric doubling incorrect")}`

Realistic required defect models:

- exact-types-only
- missing-int64
- identity
- absolute-result
- integer-narrowing

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
