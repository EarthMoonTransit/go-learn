# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Replace the final destination element with the final source element.
- Preserve every earlier destination element and the destination length.
- Changes are visible through existing destination views; preserve every source element.
- Both slices are nonempty and their backing storage does not overlap; lengths may differ.

Public/reference assertions (actual/expected diagnostics):

- Case01: `src:=[]string{"a","b","c","last"};dst:=[]string{"a","x","b","z","?"};subject.CopyLast(src,dst);want:=[]string{"a","x","b","z","last"};if !slices.Equal(dst,want)||!slices.Equal(src,[]string{"a","b","c","last"}){t.Errorf("source/destination got %q/%q want unchanged source/%q",src,dst,want)}`
- Case02: `src:=[]string{""};dst:=[]string{"?"};view:=dst[:];subject.CopyLast(src,dst);if view[0]!=""{t.Errorf("aliased final value got %q want empty",view[0])}`

Realistic required defect models:

- append-instead
- source-first
- destination-first
- copy-destination
- swap-last

Retained source behavior: lasagna-master selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| final source to final destination, lengths/order | 01 | append-instead; source-first; destination-first |
| source preserved | 01 | swap-last |
| shared destination update and empty value | 02 | copy-destination |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.


## Independent review

Public cases distinguish source and destination endpoints with unequal lengths, preservation of earlier destination/source elements and mutation visible through destination views. Non-overlapping storage is explicit.

Reviewed the actual task contract, starter, public/reference assertions, all 5 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 5 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
