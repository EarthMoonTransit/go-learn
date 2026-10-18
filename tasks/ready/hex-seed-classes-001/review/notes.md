# Independent content review

This two-task P20 batch is incremental preparation under ADR 0022, not a release,
export or foundation acceptance. Concrete content and supplemental technical evidence were independently approved for normal branch-local lifecycle. Final repository and release gates remain distinct.

Primary outcome: `fuzzing.targets-seeds`.
Only seed selection is novel: the supplied callback already checks decoded value and status. No wire oracle implementation is supplied in this lesson.
- All test files are editable; seed focus is instructional, with no partial-file
restriction. No randomized search, invalid-image decoder, reproduction, ownership,
capacity or implementation credit. Nil and empty are equal byte values.

## Requirement and public feedback mapping

| Requirement | Public assertion / automated control | Full adequacy model | Diagnostic |
|---|---|---|---|
| A valid native target and registration | starter/reference target; wrong-target, no-registration | Eligibility precedes credit | finite grammar or original type-check message |
| Correct value and true status after Encode | supplied reference callback; public against canonical, loop and empty | empty-status, zero-byte, high-bit, letter-case, byte-order, run-length | Decode status or decoded bytes mismatch |
| Canonical two lowercase digits per byte | Documented codec contract; this lesson supplies round-trip property only | Unilateral letter-case model; no independent-oracle skill claim | Decode status mismatch |
| Equal nil/empty values; no representation restriction | nil-only-wire and nonnil-only-wire deliberately fail a correct implementation | Not a private defect kill | invalid nil-only / nonnil-only wire assumption |
| Learner-owned finite seeds / useful invariants | alternative corpus; one omission control per semantic seed class | Only actual submitted F.Add values execute | full adequacy reports surviving models privately |

## Source and rights

The exact positive source decision is
`catalog/intake/2026-10-18/p20-hex-source-review.json`. Provenance records
`golang-go`, commit `2c31595783c74564c1299a5401dc6e606f379cc7`,
`src/encoding/hex/hex_test.go`, SHA-256
`b59de7b29f9b3d0bd75d661666d57fed17d87574b79f5a299a92f9c0f88123bb`,
role concept-seed, usage concept-only. ATTRIBUTION retains LICENSE/PATENTS lineage.
This is not an upstream native-fuzz donor or substantive adaptation. No upstream
vectors/code/text and no reference-only material were copied. Local project
creation/digest is distinct from the source commit; release binding stays deferred.

## Concrete matrix to review and run

`control-matrix.json` and `controls/*.go.txt` retain actual regression submissions.
The canonical implementation and two alternatives vary algorithms and empty slice
representations. Every model has the same synchronous exported API and never
mutates its arguments or invokes another public codec API internally.
The six unilateral predicates are independent on the reference corpus: empty; any zero byte; any high-bit byte; any hexadecimal letter; non-palindromic byte order; any adjacent equal-byte run. No singleton position or hidden run-length threshold is required. Omission controls must show the corresponding model survives, while an independently chosen adequate corpus must pass.

Authored task code is never executed from this checkout. All executions below use disposable runners. Existing catalog addresses and learner state are preserved.

## Fairness correction from independent prereview

Initial singleton-only and length-eight model triggers were too narrow for the
stated input variation. Models now apply zero/high-bit/letter behavior in any
position and collapse any adjacent equal bytes without a hidden length threshold.
`combined-classes` contains the zero/high-bit/letter classes together in a multibyte
input and a two-byte run. It must still detect all six models; no exact instructor
shape or literal is required. Disjoint reference seeds retain omission evidence.

## Independent review and actual execution

Reviewed by `/root/activation_final_review` on 2026-10-18. The concrete review
verified prerequisite separation, prompt/API alignment, exact concept-only source
lineage, clean-room authorship, license notices and absence of solution leakage.
The complete independent receipt is retained in the goal-chain evidence as
`p20-content-review-epoch-6.json`; its scope is branch-local content/lifecycle,
conditional on actual learner boundary checks and all final acceptance/CI gates.

The two-task supplemental matrix completed 130 cases twice: 260 matched rows,
252 real native runs and eight source-admission rejections before native execution.
All canonical repeat hashes agree. `control-results.json` retains this task's
actual source and implementation hashes, expected/observed detection and repeats.
This supplement uses no race/vet; each separate targeted task-check used the
declared race/vet, passed all technical gates and detected every required model.
The original draft reports honestly retained failing review-evidence; their
technical PASS was not treated as review approval.

Normal `task-promote` transitions and actual learner public/full controls completed successfully, as recorded below.
This approval does not assert final foundation coverage, empirical calibration,
a public release/export or successful remote CI. Final exact-head review remains
required after all package, report and repository acceptance evidence exists.

Independent content receipt SHA-256: `6bc26d4bad57f8dcdfca962ba6e703f6df600066d28190811bf407e3f2a08d28`.

Both normal task-promote transitions completed successfully with future-package
technical and ready gates; no manual moves or ready overwrite were used.

## Actual learner public/full regression evidence

The retained automatic drivers `.tmp/p20-fuzz-content/run-public-controls.go` and
`run-boundary-supplement.go` replay the saved submissions through CheckSession
and explicit temporary learner homes. They never import or run authored task
packages from the checkout. Their exact SHA-256 bindings and this task's actual
results are retained in `public-results.json`; regression source data remains in
`controls/*.go.txt`. No real learner home was accessed.

The main run matched all 24 expected probes across both tasks, including useful
public diagnostics for invalid nil/empty assumptions and target registration,
accepted independent adequate submissions, and privately rejected weak tests.
All full results have empty diagnostics and coarse summaries; private model
identities, paths and source are absent. The separate supplement proved both
actual starter submissions fail full verification, canonical readonly contract
bytes match, write bits are absent, a deliberate content change is rejected
with the specific read-only error, and restored public checks pass. Both homes
were removed and all ready/route inputs stayed byte-identical during execution.

Independent learner-boundary review by `/root/activation_final_review` found
no blockers; receipt `p20-learner-boundary-review-epoch-6.json`, SHA-256
`97c1fc12c9e3c2f540d11948712c6b496d54524f2f565149079d2e649eaef22f`. Comparison of task quality status back to the
prereview value was read-only in-memory normalization, never a backward
lifecycle move. All 169 previous route and public catalog records are preserved.

Preparation only: final full-bank collection, shared current reports, ordered
repository gates, final exact-head review and green exact-head CI remain required
before merge. These probes do not establish foundation coverage or learner mastery.
