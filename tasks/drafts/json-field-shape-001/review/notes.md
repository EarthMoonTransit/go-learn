# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Keep supplied public tags and private secret; constructor stores every value | TestFieldContract; TestPublicCase01 | discard-secret, wrong-count, trim-name | private secret; JSON tags; item fields |
| Always include name, omit only zero count, exclude secret | TestPublicCase01 decoded object comparison | default-names, count-always, omit-empty-name, hide-count | JSON got/want |
| Preserve negative count and exact Unicode/text while allowing key order/escaping alternatives | TestPublicCase01 semantic JSON comparison | hide-count, trim-name | JSON object and unchanged item |

## Instructional and technical evidence

Primary skill: `encoding.json-shape-tags`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `encoding.json-shape-tags`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `default-names`, `count-always`, `omit-empty-name`, `hide-count`, `wrong-count`, `trim-name`, `discard-secret`. The author-provided `.tmp/checks/json-field-shape-001.yaml` records technical PASS, reference PASS and 7/7 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
