# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Add amount to the receiver Units, visible through every alias of the same pointer.
- Repeated calls accumulate from the current balance; separate Balance values remain independent.
- Receiver is non-nil, initial Units and amounts are nonnegative and sums fit int.
- Zero deposits leave Units unchanged.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Pointer receiver adds to current balance, visible through aliases | TestContractCase01/02 nonzero initial state and repeated updates | replace-balance, no-update, subtract, shared-total | deposit amount, balance and alias values |
| Zero deposit preserves balance | TestContractCase02 includes zero after positive update | zero-reset | deposit0 actual/expected state |
| Separate receivers remain independent | TestContractCase02 deposits into other and then checks original balance | retained-receiver | independent balances actual/expected |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- replace-balance
- no-update
- subtract
- zero-reset
- shared-total
- retained-receiver

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, selected operation in `pointers/v4/wallet.go`. All declared consulted artifacts and SHA256 values were verified against the pinned local source. MIT notices and attribution were preserved; independently authored boundaries and defects extend the selected observable behavior. No generated canonical or reference-only material was reused. Private reference code and mutants are excluded from learner payloads.

## Independent review

The retained-receiver defect behaves correctly for one receiver and requires an independent receiver check. The saved one-receiver weak submission remains insufficient through the disposable checker.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/deposit-state-tests-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.

The saved `review/weak-submission.txt` checks one nonzero receiver with repeated and zero deposits. The author-provided `.tmp/checks/deposit-state-tests-weak.yaml` reports starter-public pass and starter-full insufficiency-gate pass in a disposable copy; its six-mutant summary describes the strong reference suite, not the weak suite. This is technical evidence, not learner calibration.
