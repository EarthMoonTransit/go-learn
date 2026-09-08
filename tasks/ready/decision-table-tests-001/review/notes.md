# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Substantive adaptation from `exercism-go` at `ae1477abfeeb9fbeaf971caee220e91f9e4df790`. Reviewed the exact declared source files and checked every SHA256 against the pinned checkout: `exercises/concept/blackjack/.docs/instructions.md` (adapted), `exercises/concept/blackjack/blackjack_test.go` (adapted), `exercises/concept/blackjack/blackjack.go` (adapted), `exercises/concept/blackjack/.meta/config.json` (factual-reference), `LICENSE` (copied). Registry permits adaptation; complete MIT notice, author/contributor attribution and modification boundary are retained. No reference-only content or generated canonical cases used.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Ordered 22/21 decisions and visible-token threshold | TestContractCase01–02 | pair-after-range, win-against-ten | pair; decision got/want |
| 17..20, 12..16 and below-12 boundaries | TestContractCase02 | stand-at-sixteen, exclude-twelve, visible-seven-low | decision input triple with got/want |
| Exact token weights, unknown tokens, case and whitespace | TestContractCase02–03 | normalize-token, unknown-is-ten | exact token quoted inputs |

## Instructional and technical evidence

Primary skill: `control-flow.if`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: none. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `pair-after-range`, `win-against-ten`, `stand-at-sixteen`, `exclude-twelve`, `visible-seven-low`, `normalize-token`, `unknown-is-ten`. The author-provided `.tmp/checks/decision-table-tests-001.yaml` records technical PASS, reference PASS and 7/7 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
