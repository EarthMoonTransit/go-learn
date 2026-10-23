# Author evidence: bounded live buffer

Contract approved in .tmp/v1-applications-07/automatic-contract-review.json against proposal b07ecd5b4daa10a76d0d867b14a3d6f8b116df2f5eccf2f20fc42d5e389cd8bc. Actual independent source/evidence approval granted by /root/reviewer.

The learner owns a live receive-then-send worker and derives capacity: at full output, one further value can be held by that worker, so capacity+1 <= limit; the largest capacity is limit-1. A reservation-before-receive design may safely use a different capacity but is outside the explicitly declared method. This differs from the owned-values first's synchronous whole-batch snapshot. No extra ownership/launch/range application is claimed.

| Requirement | Public assertion | Required wrong specimen | Diagnostic |
| --- | --- | --- | --- |
| prompt output | windowCase handle before EOF | collect-before-return, nil-output | Window must return before EOF; output is nil |
| held+queued bound | cap(output), input consumed while consumer paused | forget-held-value | want limit-1; maximum |
| largest capacity | cap(output) | under-buffer | want limit-1 |
| exact streaming | first receive before EOF; final complete sequence | skip-zero, duplicate-value, reverse-pairs | stream actual/expected |
| closure | bounded drain after input close | never-close | output remains open after input exhaustion |
| independent channels | TestPrivateOutputs and held overlapping calls | reuse-output (active calls only) | calls share output |
| independent mutable state | TestOverlappingWindowCalls holds B while A reaches EOF/drains, then releases B | shared-current-input | overlapping A/B stream actual/expected |

Cleanup closes each test-owned input once, retrieves any late constructor handle, and performs 40 bounded receive/Wait steps. The maximum incorrect duplication emits 24 values. A missing final close has no remaining sender. No timeout/crash/race is detection evidence. No Sleep/Now/After, AST extension or source-execution is used.

The explicit held-value select alternative obeys receive-before-send without reservation and must pass public/full/race. The two-forwarder acknowledgement control preserves behavior/bounds and must pass public/full, but actual method review rejects its worker count. Tests do not prove architecture or capacity reasoning; actual source and learner explanation must be inspected. Canonical/reference/public-only controls and wrong diagnostics are saved as stream-* receipts.

BUFFER-INDEPENDENCE-01 adds nonempty held overlapping calls and a mutex-protected shared-current-input specimen. Both inputs are closed and both outputs drained even after assertions. Removing only TestOverlappingWindowCalls in a public-only copied fixture must allow that wrong reference to pass and the same mutant to survive; this records the regression's causal boundary.

BUFFER-RUNTIME-02 rejects the earlier permanently cached output specimen, which attempted to read a channel from a previous synctest bubble. Its original source/hash/report are retained as stream-rejected-reuse-output-v1.* and earn no detection credit. The repaired specimen shares output only while a worker is active and resets its mutex-protected cache before closing output.

## Recorded author checks

Canonical stream-buffer-canonical-final.yaml passed starter compilation/expected failure, reference public/full, two determinism runs, vet, race, and all required mutants. 16 copied fixtures passed their expected outcomes: 10 public-only ordinary wrong diagnostics, public-only canonical reference, correct alternative public/full, and method-wrong behavior control public/full. Every disposable fixture was removed. No automatic architecture credit is inferred from those controls. Exact commands, source/report hashes and diagnostics are retained in .tmp/v1-applications-07/stream-author-receipt.json, stream-source-manifest.json and stream-fixture-results.json. Independent actual review follows this author evidence.

The extra omission fixture removes only TestOverlappingWindowCalls, leaves all other public checks intact and removes grader checks. The shared-current-input wrong implementation then passes reference/public/race checks and survives mutation, causing the expected mutant-coverage failure. With the regression present it receives the ordinary overlapping-sequence diagnostic. This distinguishes shared-state coverage from the separate output-identity assertion. The rejected cross-bubble cached-output source/report/receipt remain stream-rejected-reuse-output-v1.* and earn no detection credit.

## Independent actual acceptance

Sole independent reviewer /root/reviewer approved actual source, executed evidence, provenance, public feedback and finite method on2026-10-22. Current decision is accepted for official promotion; final batch local gates, final diff review, exactHEADvalidate and merge remain separate.

Actual receive-then-send singleworker without output reservation; cap limit-1 follows queued plus held bound. Independent nonempty overlapping calls, ordinary public defects, correct state-machine alternative and behavior-passing two-worker method rejection are separately evidenced.
