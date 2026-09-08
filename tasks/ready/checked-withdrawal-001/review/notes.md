# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- For amount <= Units subtract amount and return nil. Withdrawing exactly the balance is allowed.
- If amount exceeds Units return the original non-nil ErrInsufficient sentinel and leave Units unchanged.
- The receiver is non-nil; initial Units and amount are nonnegative integers <=1000000. Zero withdrawals are valid.
- Update the shared receiver visible through aliases, preserve the sentinel, and use the current balance on repeated calls.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Subtract current balance, including exact and zero withdrawals | TestPublicCase01 repeated calls with alias | strict-boundary, no-subtract, add-instead | amount, actual balance/error, expected balance |
| Failed withdrawal preserves balance and returns original sentinel | TestPublicCase02/03; then a valid withdrawal after failure | subtract-on-failure, fresh-error, failure-nil | failed balance/error versus original sentinel |
| Preserve original sentinel during successful calls too | Each public/grader case captures original non-nil sentinel and checks after calls | success-reassigned | sentinel identity changed |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- strict-boundary
- subtract-on-failure
- fresh-error
- no-subtract
- add-instead
- failure-nil
- success-reassigned

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, operation in `pointers/v4/wallet.go`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

Repeated shared-state updates, failure preservation and success-time sentinel reassignment are separately discriminated.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 7 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/checked-withdrawal-001.yaml` reports technical pass and 7/7 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
