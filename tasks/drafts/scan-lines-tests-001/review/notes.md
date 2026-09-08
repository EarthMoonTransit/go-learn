# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Line endings, final unterminated line, spaces, Unicode and interior empties | TestContractCase01–02 | drop-unterminated, drop-last, trim-lines, skip-empty | final line; lines got/want |
| Retain collected lines, original reader error and caller ownership | TestContractCase03 | ignore-error, fresh-reader-error, close-reader | reader error/success with lines/error/closes |
| Accept 2048-byte lines and reject 10000-byte unbroken token | TestContractCase02 and Case04 | tiny-buffer, unbounded-token | lines; oversize token wantempty,error |

## Instructional and technical evidence

Primary skill: `io.scanning`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `io.scanning`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `ignore-error`, `trim-lines`, `skip-empty`, `tiny-buffer`, `unbounded-token`, `close-reader`, `drop-last`, `fresh-reader-error`, `drop-unterminated`. The author-provided `.tmp/checks/scan-lines-tests-001.yaml` records technical PASS, reference PASS and 9/9 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.

The saved `review/weak-final-line.txt` preserves the strong suite while changing otherwise accepted final lines to newline-terminated inputs. The author executed it only as a starter in a disposable task-check fixture; `.tmp/checks/scan-lines-weak.yaml` records reference PASS and the expected starter-full insufficiency gate. The final TeeReader defect preserves original scanner/reader errors and collected lines, isolating loss of an unterminated final line. Its necessity is independent of reader-error coverage. This is automated weakness evidence, not human calibration.
