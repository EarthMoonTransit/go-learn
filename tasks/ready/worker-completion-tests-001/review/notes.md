# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Asynchronous start and exactly one callback | TestWorker0–01 in synctest.Test with controlled release | inline-work, skip-work, double-work | Start blocked; callback calls/did not start |
| Non-nil channel closes only after callback, without sent values | awaitClosed and TestWorker1 pre-release check | nil-channel, never-close, close-before-work, send-before-close | nil completion; did not close; before release; sent a value |
| Overlapping calls own separate workers/completion channels | TestWorker2 holds both callbacks on a gate before comparison | shared-active-completion | overlapping callbacks; active calls share completion channel |
| Race-safe deterministic synchronization and released callback lifetime | synctest.Wait, deferred gate release, required race profile | Correct reference accepted under deterministic and race checks | test completion without wall-clock deadlines |

## Instructional and technical evidence

Primary skill: `goroutines.launch-lifecycle`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `goroutines.launch-lifecycle`, `channels.operations-blocking`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `inline-work`, `never-close`, `close-before-work`, `skip-work`, `double-work`, `nil-channel`, `shared-active-completion`, `send-before-close`. The author-provided `.tmp/checks/worker-completion-tests-review.yaml` records technical PASS, reference PASS and 8/8 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.


## P13 scaffold revision, 2026-10-17

Applied exactly the future starter/hints from merged compatibility PR #77,
ADR0028. Start, original creation record, empty external sources, eight mandatory
mutants, primary/secondary skills and introduction role remain. The named
blocked/released/completed example and idempotent release helper do not distribute
premature-completion or independent-call cases; those remain learner-owned.
All fixed exported files and editable paths are unchanged. Historical learner
workspaces retain their exact bytes/markers/paths/history; no reset is required.

The separate tooling integration proves public success of old/new starters,
private failure of their incomplete mutation coverage, named actual/want phase
feedback, all eight unchanged reference kills, race and controlled cleanup.
The new starter rejects inline-work through an assertion after safe release;
close-before-work and shared-active-completion remain intentionally uncovered.
Supplied synctest/helper internals are guidance, not independent harness mastery.
Fresh content/provenance review and final technical/batch gates follow.

Independent exact-content review approved by /root/clamp_e01_content_review on 2026-10-17. Actual targeted reference, 8/8 required mutants and race passed; accepted alternative solutions also passed where recorded. Final ready-state checks and batch acceptance follow this metadata freeze. Separate persisted public-feedback regressions remain an explicit rollout gate.

P13 final content review approved the exact ADR 0028 fixture digest. The eight-defect full matrix and preserved old/new workspaces remain protected by the separate tooling integration tests; the scaffold intentionally leaves premature completion and overlapping-call independence to the learner.
