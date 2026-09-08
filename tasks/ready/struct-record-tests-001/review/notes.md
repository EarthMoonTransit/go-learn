# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return a non-nil pointer with the supplied Name, Age and Attributes contents.
- Empty name and zero age are preserved; all integer ages are accepted.
- A nil Attributes input remains nil; a non-nil empty map remains non-nil.
- Do not change the supplied map contents. Each call returns a distinct record; map storage may be shared or copied.

Public/reference assertions (actual/expected diagnostics):

- Case01: `want:=subject.Entry{Name:"record",Age:29,Attributes:map[string]string{"street":"Main St."}};got:=subject.NewEntry(want.Name,want.Age,want.Attributes);if !reflect.DeepEqual(got,&want){t.Errorf("record got %+v want %+v",got,want)}`
- Case02: `got:=subject.NewEntry("",0,nil);if got==nil||got.Name!=""||got.Age!=0||got.Attributes!=nil{t.Errorf("zero record got %+v",got)}`
- Case03: `got:=subject.NewEntry("n",-2,map[string]string{});if got==nil||got.Age!=-2||got.Attributes==nil||len(got.Attributes)!=0{t.Errorf("empty map/negative age got %+v",got)}`
- Case04: `a:=map[string]string{"k":"v"};subject.NewEntry("n",3,a);if len(a)!=1||a["k"]!="v"{t.Errorf("attributes changed: %v",a)}`
- Case05: `a:=subject.NewEntry("a",1,nil);b:=subject.NewEntry("b",2,nil);if a==nil||b==nil{t.Fatal("got nil record")};a.Name="changed";if b.Name!="b"{t.Errorf("independent record got %q want b",b.Name)}`

Realistic required defect models:

- nil-result
- name-lost
- age-lost
- attributes-lost
- nil-normalized
- empty-normalized
- clear-input
- shared-record
- default-name
- default-zero-age
- clamped-negative-age

Retained source behavior: census selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| nonnull fields/map | 01 | nil-result; name-lost; age-lost; attributes-lost |
| nil/empty maps distinct | 02–03 | nil-normalized; empty-normalized |
| empty name, zero and negative age preserved | 02–03 | default-name; default-zero-age; clamped-negative-age |
| input preserved | 04 | clear-input |
| independent records | 05 | shared-record |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.

Added distinct default-name, default-zero-age and clamped-negative-age defects. Case02 checks empty name/zero age and Case03 negative age. Saved weak fixture retains map/independence checks using only nonempty names and positive ages, so the specific omitted domain remains necessary.


## Independent review

The review gap is closed by separate default-name, default-zero-age and clamped-negative-age defects. The saved weak suite retains map and record-independence checks while omitting those input categories; its disposable report confirms insufficiency. The strong suite detects all eleven required defects.

Reviewed the prompt, API, actual checks, all 11 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 11 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
