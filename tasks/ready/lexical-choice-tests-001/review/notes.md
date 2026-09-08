# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Choose the string that comes first under Go string comparison.
- Return selected: followed by one space and the chosen string.
- Swapping the arguments does not change the result; equal strings return that same string after the prefix.
- Empty strings are valid; preserve exact case and whitespace and compare the complete strings.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.ChooseLabel("Beta","Delta");got!="selected: Beta"{t.Errorf("result got %q want %q",got,"selected: Beta")}`
- Case02: `if got:=subject.ChooseLabel("Delta","Beta");got!="selected: Beta"{t.Errorf("result got %q want %q",got,"selected: Beta")}`
- Case03: `if got:=subject.ChooseLabel("abz","aba");got!="selected: aba"{t.Errorf("result got %q want %q",got,"selected: aba")}`
- Case04: `if got:=subject.ChooseLabel("z","A");got!="selected: A"{t.Errorf("result got %q want %q",got,"selected: A")}`
- Case05: `if got:=subject.ChooseLabel("x","x");got!="selected: x"{t.Errorf("result got %q want %q",got,"selected: x")}`
- Case06: `if got:=subject.ChooseLabel("x","");got!="selected: "{t.Errorf("result got %q want %q",got,"selected: ")}`
- Case07: `if got:=subject.ChooseLabel(" a","a");got!="selected:  a"{t.Errorf("result got %q want %q",got,"selected:  a")}`

Realistic required defect models:

- first-only
- largest
- first-byte-only
- case-insensitive
- empty-skipped
- equal-empty
- trimmed-output
- prefix-missing

Retained source behavior: vehicle-purchase selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| symmetric choice | 01–02 | first-only; largest |
| complete strings | 03 | first-byte-only |
| case-sensitive comparison | 04,08 | case-insensitive |
| equal inputs | 05 | equal-empty |
| empty input | 06 | empty-skipped |
| whitespace/prefix | 07 | trimmed-output; prefix-missing |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.


## Independent review

The first-byte-only defect now handles empty strings correctly, isolating complete-string comparison. The saved weak suite omitting the common-prefix case remains insufficient in the author-provided disposable report. Symmetry, equal/empty strings, exact case, whitespace and prefix behavior are covered.

Reviewed the actual task contract, starter, public/reference assertions, all 8 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 8 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
