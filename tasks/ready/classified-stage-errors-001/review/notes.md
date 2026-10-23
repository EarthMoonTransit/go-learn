# Author mapping and method boundary

Learner owns stage execution, exposed versus opaque cause construction, then errors.Is/errors.As on the returned error. ErrTemporary and nil-safe Detail are supplied preparation, with no extra sentinel/typed-design credit. The starter supplies empty private context types and zero Error/Unwrap stubs solely to preserve the existing method API fingerprint. Learner owns their representation and bodies as well as the loop, boundary choice and inspection; no cause is retained or exposed by the starter. Three proposed applications only: errors.wrapping, errors.is, errors.as; new_concepts is empty.

| Requirement | Public assertion | Wrong variants | Useful diagnostic |
| --- | --- | --- | --- |
| once/order/first failure and exact value | TestStageOrderAndValues | continue-after-error; repeat-stage; discard-failure-value | stage calls / failure value actual and expected |
| exact context and direct identity | TestContextTextAndIdentity | missing-context; trimmed-context; formatter-text; flatten-cause | quoted context text or direct cause |
| exposed/opaque returned boundary | TestOpaqueBoundary and temporary/detail tests | always-opaque; always-wrap; classify-raw-cause | direct cause, opaque exposure or returned classification |
| standard Is semantics | TestTemporaryClassification | identity-only; text-matching; first-branch-only; ignore-custom-is | Temporary actual/wanted |
| standard As first identity and boolean | TestDetailClassification | direct-detail-only; copy-detail; reverse-detail-order; ignore-custom-as; typed-nil-is-miss | detail pointer / HasDetail actual/wanted |
| empty/success no classification | TestStageOrderAndValues | success-flags | zero/success Outcome |

Grader extends four stage positions, context text and three-branch finite trees. All callbacks are synchronous, finite and non-panicking. Supplied Detail.Error handles nil; no wrong As target or nil dereference is used as detection. Tests never require byte-identical error concrete representation. Both direct-cause identity and exact Error text are intentional; formatter-text is not an alternative because Formatter can differ from Error.

Reference wraps only exposed failures, directly to the original cause, using Error text rather than Formatter. Opaque translation stores no unwrap cause. Classification happens after boundary choice, with standard Is and As. It preserves HasDetail independently from a possibly nil pointer. The correct alternative computes text lazily using distinct visible/opaque types and a helper; its opaque type may retain a private cause but cannot unwrap it. That private implementation choice is accepted.

The manual tree-traversal control handles direct/custom/one/many cases and may pass every finite assertion, but violates the explicit standard Is/As method. Its passing public/full receipts are not method acceptance. Actual independent source review is approved in review/independent-review.json; no new AST or source mechanism. Every required wrong binding must fail ordinary public assertions in a copied public-only fixture. Correct alternative and method control must pass public/full. All execution is existing task-check disposable, with no timeout/crash credit or real learner state.

## Independent actual acceptance

Reviewer /root/reviewer accepted this final source and finite evidence on2026-10-23; exact immutable decision is retained in review/independent-review.json. Behavior-passing method controls are explicitly rejected by the required checklist. Official lifecycle transitions are complete; integration and all mandatory batch local/CI gates remain separate.
