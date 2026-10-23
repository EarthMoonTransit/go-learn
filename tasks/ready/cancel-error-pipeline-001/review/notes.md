# Author evidence: cancellation across two stages

Contract approved in automatic-contract-review.json against proposal b07ecd5b4daa10a76d0d867b14a3d6f8b116df2f5eccf2f20fc42d5e389cd8bc. Actual independent content/method/evidence approved by /root/reviewer. Explicit context.derivation prerequisite documents owned WithCancel; it was already available transitively through context.cancellation and adds no credit.

The learner writes feeder and mapper workers connected by an internal channel, derives/owns child cancellation, cancels the sibling on the first step error, preserves the successful prefix, and joins before reporting terminal. The previous Forward first has one worker and no callback failure/sibling/joined terminal. The previous two-stage first requires full EOF/drain and excludes cancellation. No additional context/pipeline/ownership/WaitGroup/leak application is claimed.

| Requirement | Public observation | Required specimens | Useful diagnostics |
| --- | --- | --- | --- |
| prompt non-nil channels | startPipe before caller closes input | nil-handles, collect-before-handles | nil pipeline handle; pipeline must return before EOF |
| declared capacities | startPipe cap checks | buffer-output, unbuffer-terminal | output/terminal capacity actual/expected |
| exact stream and partial prefix | pipelineCase before EOF and exact final sequence | skip-zero, duplicate-result, emit-failing-value | mapped prefix actual/expected |
| first error stops processing | callback count, exact sentinel | continue-after-error, replace-error, suppress-error | no calls after first error; want exact sentinel |
| one closed terminal | bounded result collection | two-terminal-results, open-terminal | terminal count; terminal remains open |
| upstream cancellation | first-error result while input stays open | error-waits-for-eof | terminal still pending after step failure with input open |
| idle input cancellation | TestCancelInputWait | cancel-waits-for-eof | terminal still pending after cancellation while input open |
| paused output cancellation | cancel without ready receiver; Wait before observation | uncancelable-output-send | output remains open after paused-send cancellation |
| wait for active callback | TestJoinHeldStep | early-terminal | terminal arrived before active step completed |
| pre-canceled guard | TestAlreadyCanceled | pre-canceled-read | pre-canceled call consumed input/called step |
| close final output | bounded observation | open-output | output remains open after termination |

Every test closes its caller-owned input once, owns a parent cancel, releases held callbacks independently, retrieves any late returned handles, and drains up to 80 steps (at most 24 duplicate values/two terminal errors). The uncancelable output-send specimen has one externally releasable send; EOF-waiting specimens drain input without sending internally after cancellation. Their wrong policy is diagnosed before those rescue actions. Raw inaccessible internal send omissions and double-close panics are not required specimens or credited runtime detection.

Normal tests require EOF/drain; cancellation/error tests supply an explicit termination cause. There is no promised completion for an indefinitely open input, an absent output consumer, and no cancellation/error. Already-winning sends are permitted; the paused-send test deliberately supplies no receiver until after cancellation and synctest.Wait. Callback completion is independently held for the join assertion. Correct alternatives may choose different co-ready causes/read-ahead; no scheduler priority or exact prefetch is required.

Actual review must inspect each feeder input/internal send and mapper internal receive/output send for child cancellation, WithCancel(parent), mapper-only reason writes with Wait-before-read, close ownership, cancellation release, and join-before-terminal. Tests alone do not prove architecture or general leak absence. The completion-channel coordinator alternative must pass public/full/race; a fused worker must pass behavior but fails the two-worker method. These controls are retained under stream-* and are not credited as detected mutants. No AST/runtime extension is proposed.

## Recorded author checks

Canonical stream-cancel-canonical-final.yaml passed starter compilation/expected failure, reference public/full, two determinism runs, vet, race, and all required mutants. 23 copied fixtures passed their expected outcomes: 18 public-only ordinary wrong diagnostics, public-only canonical reference, correct alternative public/full, and method-wrong behavior control public/full. Every disposable fixture was removed. No automatic architecture credit is inferred from those controls. Exact commands, source/report hashes and diagnostics are retained in .tmp/v1-applications-07/stream-author-receipt.json, stream-source-manifest.json and stream-fixture-results.json. Independent actual review follows this author evidence.

A preliminary unbuffer-terminal specimen was rejected: its pre-canceled synchronous send could deadlock before returning handles. It was repaired to publish asynchronously, preserving the wrong zero capacity while making cleanup finite. The old source (matched to its original hash), failed report and receipt are retained as stream-rejected-unbuffer-terminal-v1.*; they earn no detection credit. Canonical and all 23 cancellation fixtures were rerun after this sole specimen change. Reference, TASK and public/grader tests were unchanged.

## Independent actual acceptance

Sole independent reviewer /root/reviewer approved actual source, executed evidence, provenance, public feedback and finite method on2026-10-22. Current decision is accepted for official promotion; final batch local gates, final diff review, exactHEADvalidate and merge remain separate.

Actual feeder and mapper own derived cancellation, every blocking operation observes it, and terminal reporting follows both joined stages and output closure. Public finite rescue paths, exact prefix/error behavior and correct two-done alternative pass; fused behavior-passing control is method-rejected. No inaccessible-deadlock detection credit.
