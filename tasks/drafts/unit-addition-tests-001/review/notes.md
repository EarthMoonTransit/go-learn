# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- values is non-nil; units may be nil. Quantities are nonnegative and additions fit int.
- An absent unit returns false and leaves values unchanged.
- A present unit returns true and adds its quantity to the named entry, creating an absent entry.
- A present zero unit still succeeds and creates an absent key with zero.
- Preserve other entries and the entire units map; repeated calls accumulate.

Public/reference assertions (actual/expected diagnostics):

- Case01: `v:=map[string]int{};u:=map[string]int{"dozen":12};if !subject.AddUnit(v,u,"item","dozen")||v["item"]!=12{t.Errorf("add got %v want item:12 and true",v)}`
- Case02: `v:=map[string]int{"a":4,"b":3};u:=map[string]int{"triple":3};if !subject.AddUnit(v,u,"a","triple")||!subject.AddUnit(v,u,"a","triple")||v["a"]!=10||v["b"]!=3{t.Errorf("accumulation got %v want a:10 b:3",v)}`
- Case03: `v:=map[string]int{"a":4};if subject.AddUnit(v,map[string]int{},"a","missing")||len(v)!=1||v["a"]!=4{t.Errorf("unknown unit changed %v or returned true",v)}`
- Case04: `v:=map[string]int{};if subject.AddUnit(v,nil,"a","missing")||len(v)!=0{t.Errorf("nil units changed %v or returned true",v)}`
- Case05: `v:=map[string]int{};ok:=subject.AddUnit(v,map[string]int{"zero":0},"a","zero");n,present:=v["a"];if !ok||!present||n!=0{t.Errorf("zero unit got %v,%v want present a:0 and true",v,ok)}`
- Case06: `v:=map[string]int{};u:=map[string]int{"triple":3,"other":7};subject.AddUnit(v,u,"a","triple");if len(u)!=2||u["triple"]!=3||u["other"]!=7{t.Errorf("units changed: %v",u)}`

Realistic required defect models:

- overwrite-instead
- unknown-succeeds
- unknown-inserts
- zero-rejected
- zero-not-created
- clear-other-entries
- consume-unit

Retained source behavior: gross-store selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| new item/quantity | 01 | zero-not-created; overwrite-instead |
| accumulation/other items | 02 | overwrite-instead; clear-other-entries |
| missing/nil unit rejected without mutation | 03–04 | unknown-succeeds; unknown-inserts |
| present zero creates key | 05 | zero-rejected; zero-not-created |
| unit map preserved | 06 | consume-unit |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.
