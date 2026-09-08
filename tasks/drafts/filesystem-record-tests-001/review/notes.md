# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Substantive adaptation from `learn-go-with-tests` at `4675d96ad50b815d3e9fdc2122ab3dbc14636252`. Reviewed the exact declared source files and checked every SHA256 against the pinned checkout: `reading-files/blogposts.go` (adapted), `LICENSE.md` (copied), `reading-files/post.go` (adapted), `reading-files/blogposts_test.go` (adapted). Registry permits adaptation; complete declared-license notice, author/contributor attribution and modification boundary are retained. No reference-only content or generated canonical cases used.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Lexically ordered root records, directories skipped, exact headers/tags/body | TestContractCase01–02 | reverse-files, unsplit-tags, trim-body | records/exact fields with opens/closes |
| Every opened file closed exactly once before return, allowing any closure order | TestContractCase01–04 compare closure multisets | no-close, close-before-read | opens/closes including duplicate counts |
| Directory/open/read/close errors preserve identity and discard partial collection | TestContractCase03 | ignore-close-error, partial-results, ignore-reader-error, fresh-open-error | named failure path with returned records/error |
| Malformed headers use original non-nil ErrFormat; read/format error wins over close error | TestContractCase04 and package-level originalFormatError | accept-malformed-separator, fresh-format-error | malformed; error priority |

## Instructional and technical evidence

Primary skill: `io.files-cleanup`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `io.files-cleanup`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `no-close`, `ignore-close-error`, `partial-results`, `reverse-files`, `ignore-reader-error`, `unsplit-tags`, `trim-body`, `fresh-open-error`, `close-before-read`, `accept-malformed-separator`, `fresh-format-error`. The author-provided `.tmp/checks/filesystem-record-tests-review.yaml` records technical PASS, reference PASS and 11/11 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
