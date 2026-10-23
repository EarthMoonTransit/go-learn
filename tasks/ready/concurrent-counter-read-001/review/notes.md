# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Substantive adaptation of the permitted Learn Go with Tests `sync/v2/sync.go` and `sync/v2/sync_test.go`, pinned at `4675d96ad50b815d3e9fdc2122ab3dbc14636252`. Both exact source hashes and `LICENSE.md` hash verified against the local pinned checkout; MIT notice including literal copyright placeholders retained. The counter operations and three-increment/1000-worker test substance are retained. Concurrent reads explicitly extend the source after-Wait read semantics. No reference-only or generated canonical material reused.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Stable exact reads without mutation | TestPublicCase01/02 and grader sequential fresh reads | zero-read, read-increments, read-clears | stable read / final count / fresh read |
| Readers synchronize with the supplied writer | TestPublicCase02 overlaps 1000 reads and increments under required race profile | unlocked-read, unlock-before-read | DATA RACE |
| Independent new/zero counters | TestPublicCase01/03 | Reference acceptance probes | counter state is shared / zero read |

## Instructional and technical evidence

Primary skill: `sync.data-races-happens-before`. Prerequisites resolve at the actual numbered route position. Short technical English prompt; minimal implementation stub or explicit debugging defect. The two counter lessons distinguish a shared read-modify-write critical section from overlapping reader/writer synchronization; no extra scheduling order or absolute-time requirement is imposed.

Required defects: `unlocked-read`, `unlock-before-read`, `zero-read`, `read-increments`, `read-clears`. Author-provided `.tmp/checks/concurrent-counter-read.yaml` records technical PASS, reference PASS, race PASS and 5/5 required defects detected. No authored code or task tests were run from the checkout by this reviewer; no real learner state was accessed. Full collection, ordered gates, PR validation and lifecycle movement remain separate. Automated fixtures do not establish human calibration or learner mastery.

## I12 CI regression — bounded common-start workload (2026-10-23)

PR117 CI39933918517 at ba919937 retained an `unlock-before-read` survivor.
That failed collection is preserved in `.tmp/v1-applications-12/ci-attempt1-failure.txt`;
it earns no detection credit for that mutant. The historical review above remains
unchanged. Its old one-reader/one-writer workload did not require a common start;
this follow-up does not claim that every possible schedule must expose a race.

The approved content repair preserves TASK, API, supplied Inc, starter, reference,
all five required mutants, the three original public test bodies and grader.
`TestConcurrentReadCommonStart` adds four writers and four readers with 1000 calls
each. All eight acknowledge readiness before one start-channel close. Each worker
yields with Gosched only between calls; reader failures occupy disjoint slots,
read after WaitGroup join. Two final reads must equal 4000. The serial test raises
GOMAXPROCS to two only when necessary and restores it with Cleanup. It adds no
sleep, wall-clock deadline, private-field hook or required interleaving.

Current canonical task-check passes reference, race, vet and all 5/5 mutants.
The closed sixteen-scope matrix used only official task-check disposable
execution: one canonical scope plus fifteen copied fixtures, each removed.
Public-only reference and an explicit-lock/local-read/unlock correct alternative
(public/full) pass. The five unchanged public wrong variants are detected: three
ordinary value/state assertions and two real DATA RACE diagnostics, kept distinct.

One focused clean-reference fixture and three predeclared fresh focused fixtures
per race mutant contain only the new public test and an empty grader. All six
wrong fixtures report DATA RACE with Value/Inc accesses and goroutine creation
stacks naming TestConcurrentReadCommonStart. Thus this evidence is attributable
to the new workload, not the preserved Case02. These are six actual finite
observations, not a probability estimate or proof about all schedules. No retry
until success, timeout, panic or compilation failure is counted as detection.

Exact argv, statuses, source hashes, outputs and cleanup are retained under
`.tmp/v1-applications-12/counter-*`; `counter-author-receipt.json` indexes the
closed matrix. Adaptation registry/commit/artifact/license pins are unchanged;
only the project content digest is refreshed for the added public test. Sole
actual review is a separate ci-regression-review.json receipt; full-bank collection,
final consistency pins, mandatory gates and exact-head CI remain root's work.
