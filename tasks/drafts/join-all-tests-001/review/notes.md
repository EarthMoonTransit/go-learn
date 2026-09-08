# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, public verification runs learner tests against the correct implementation and full verification challenges them with required defects. The table names the strong reference learner suite; the starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Every callback starts independently once; prompt return while blocked | TestContractCase01 and Case02 three independently owned slots and gates | skip-last, double-work, serial-work | callback counts / not independently started / blocked before callbacks finished |
| Completion waits for every callback, including the final unfinished one | Case02 releases callbacks in order 0, 2, 1 with intermediate Wait and open-channel assertions | close-before-wait, wait-first-only | completion before workers / after only first worker / before final worker |
| Completion closes with no sent values; empty calls independent and non-nil | awaitJoined and TestContractCase03 | completion-value, nil-empty | completion sent value / nil completion / shared completion |
| Race-safe test-owned state and gate cleanup | Per-slot callback writes; synchronized Wait reads; all gates deferred for release | Correct reference under required race profile | reference/race pass without wall-clock scheduling deadlines |

## Instructional and technical evidence

Primary skill: `sync.waitgroup-lifecycle`. Declared prerequisites resolve at the actual numbered route position; the final concurrency order was independently checked against preceding outcomes. New concepts: `sync.waitgroup-lifecycle`, `goroutines.state-ownership`. Short technical English prompt; minimal implementation stub or explicit debugging defect, or intentionally weak learner tests. No decorative narrative. Buffered/unbuffered closed input and non-nil callback preconditions are explicit where applicable. Select checks accept either ready branch, and synctest synchronization avoids wall-clock scheduling assumptions.

Required defects reviewed: `close-before-wait`, `skip-last`, `double-work`, `serial-work`, `nil-empty`, `completion-value`, `wait-first-only`. Author-provided `.tmp/checks/join-all-tests-review.yaml` records technical PASS, reference PASS, race PASS and 7/7 detected required defects. For mutation-tests the count describes the strong learner reference suite; starter-full PASS denotes the checker's insufficiency gate, not that the starter kills all defects.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
