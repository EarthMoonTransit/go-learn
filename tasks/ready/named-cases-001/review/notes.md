# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Real named child tests receive distinct *testing.T identities | TestPublicCase01–02 inspect child names and parent identity | parent-callback, constant-name | child first/second/single name |
| Serial table order, once-only callbacks and unchanged input names | TestPublicCase01 | skip-last, reverse-order, double-check | serial cases calls/names |
| Empty and single tables; independent later invocations | TestPublicCase02 and grader fresh call | skip-last, double-check | single/fresh callback count |

## Instructional and technical evidence

Primary skill: `testing.subtests`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `testing.subtests`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `parent-callback`, `constant-name`, `skip-last`, `reverse-order`, `double-check`. The author-provided `.tmp/checks/named-cases-001.yaml` records technical PASS, reference PASS and 5/5 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.

## M1 prerequisite review, 2026-10-17

The public Case API contains a function-valued Check field, and RunNamed accepts
a slice of Case structures. Add functions.function-values, structs.literals-zero
and slices.header-model to prerequisites. Existing introductions precede this
lesson in both address and proposed orders. The testing.subtests assessment,
serial execution requirements and public feedback remain unchanged.

Metadata-only preparation: TASK, starter, public/grader/reference bodies,
mutants, hints, authored digest, source chain and license notices are unchanged.
Reuse the previously admitted content and source review; no new upstream artifact
audit or learner calibration is claimed. Existing requirement/public-feedback
mappings remain applicable. Full technical acceptance and the metadata review
are separate gates.

Independent metadata review by `/root/clamp_e01_content_review` approved this
M1 correction on 2026-10-17 with no blockers. It verified API requirements, earlier
introductions in both orders, unchanged authored digests/source chains and the
limited metadata scope. It did not run suites or claim a new upstream audit.
