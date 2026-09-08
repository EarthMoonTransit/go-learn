# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return the exact prefix INPUT: followed by one space and the uppercased input.
- Use Go Unicode uppercase behavior for every letter.
- Preserve punctuation, digits and whitespace from the input.
- Input is valid UTF-8; empty input returns only the prefix and its space.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.UpperLabel("Judy");got!="INPUT: JUDY"{t.Errorf("result got %q want %q",got,"INPUT: JUDY")}`
- Case02: `if got:=subject.UpperLabel("lars");got!="INPUT: LARS"{t.Errorf("result got %q want %q",got,"INPUT: LARS")}`
- Case03: `if got:=subject.UpperLabel("Peter-James");got!="INPUT: PETER-JAMES"{t.Errorf("result got %q want %q",got,"INPUT: PETER-JAMES")}`
- Case04: `if got:=subject.UpperLabel("MJ");got!="INPUT: MJ"{t.Errorf("result got %q want %q",got,"INPUT: MJ")}`
- Case05: `if got:=subject.UpperLabel("éж9");got!="INPUT: \u00c9\u04169"{t.Errorf("result got %q want %q",got,"INPUT: \u00c9\u04169")}`
- Case06: `if got:=subject.UpperLabel(" a ");got!="INPUT:  A "{t.Errorf("result got %q want %q",got,"INPUT:  A ")}`
- Case07: `if got:=subject.UpperLabel("");got!="INPUT: "{t.Errorf("result got %q want %q",got,"INPUT: ")}`

Realistic required defect models:

- no-uppercase
- lowercase
- ascii-only
- trimmed-input
- prefix-missing
- empty-special

Retained source behavior: welcome-to-tech-palace selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| uppercase all letters | 01–04 | no-uppercase; lowercase |
| Unicode conversion | 05 | ascii-only |
| whitespace/punctuation | 03,06 | trimmed-input |
| prefix/empty | 07 | prefix-missing; empty-special |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.


## Independent review

Public checks exercise the exact prefix, Unicode uppercase across Latin and Cyrillic letters, punctuation/digits, retained whitespace and empty input.

Reviewed the actual task contract, starter, public/reference assertions, all 6 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
