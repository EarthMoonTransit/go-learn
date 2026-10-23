# Author mapping and method boundary

This is an independent application of JSON shape/tags and struct field visibility. The model starter has an empty private recordState and zero method stubs; all public/private fields and tags are learner-authored. The shared type Record recordState preserves existing exported API fingerprint without supplying representation. Reference subject uses actual model.Record.Label and Quantity selectors and the Secret accessor, then New reconstructs private state. Tests outside model use stable API and guarded reflection so starter and visibility defects compile. All assertions live in external subject_test, and verifier selectors use ./subject; model remains a real imported implementation package. This keeps deterministic test-event ordering within one test package, without changing the cross-package contract. Tags alone are not the cross-package credit.

| Requirement | Public assertion | Wrong variants | Useful feedback |
| --- | --- | --- | --- |
| authored fields/export/tags | TestRecordDeclaration | missing-tags; private-label; exported-secret | field missing/export/type/JSON name or private secret missing |
| exact construction/accessors and independent values | TestRecordConstructionAndAccessors | discard-secret | record values actual/wanted |
| encoding shape and zero omission | TestEncodeDeclaredShape | missing-tags; omit-empty-label; retain-zero-quantity; drop-negative-quantity; leak-secret | JSON keys, label, qty presence/value |
| absent/null versus explicit zero | TestOptionalPatchValues | missing-zero-presence; null-clears; zero-ignored | record result actual/wanted |
| unknown/private keys and secret preservation | TestOptionalPatchValues | reject-unknown; overwrite-secret; reset-secret | Apply error or exact private value |
| atomic error result and base preservation | TestPatchFailureIsAtomic | ignore-decode-error; apply-partial-on-error | expected decoding error or unchanged record result |

Grader extends only valid text/number combinations and encoding after patch. JSON output is compared as parsed keys/values, never byte spelling/order. Reflection checks explicit field names/tags/visibility; it does not prescribe the omission option string. The alternate plain-int omitzero declaration and anonymous optional struct must pass public/full. Public field guards avoid panic when starter/private-label lacks Label; exported-secret still hides its value with json:"-", so visibility is independently checked rather than inferred solely from wire leakage. Private-label's consumer uses Values/New to keep it compilable and yield ordinary shape failure.

The learner owns optional typed fields, json.Unmarshal on that struct, error-before-application, exact cross-package value update and json.Marshal on the model representation. Unknown fields are ignored; absent and null both preserve base, while explicit zero/empty values apply. Private secret is never accepted from input or emitted. All state consists of strings/ints; no concurrency or mutable aliases, and no source checker extension.

The map-based decode/encode control retains correct tagged model and cross-package access and is expected to pass this finite public/full matrix, but violates the explicit struct-consumer method. Its behavior success is not method acceptance. Actual independent source review establishes learner ownership and consumer method; see review/independent-review.json. Every required wrong overlay must compile and fail useful ordinary public assertions with grader blanked. Correct and method controls are documented honestly. All authored execution uses existing task-check disposable copies; no timeout/crash credit or real learner state.

## Independent actual acceptance

Reviewer /root/reviewer accepted this final source and finite evidence on2026-10-23; exact immutable decision is retained in review/independent-review.json. Behavior-passing method controls are explicitly rejected by the required checklist. Official lifecycle transitions are complete; integration and all mandatory batch local/CI gates remain separate.
