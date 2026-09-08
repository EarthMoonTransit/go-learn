# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Write exactly "Hello, " followed by name, preserving empty text, Unicode and spaces; append no newline.
- Make one Write call on the supplied non-nil writer. Return its exact error if non-nil, even with a partial count.
- If Write reports fewer bytes than requested with nil error, return io.ErrShortWrite; on a full successful write return nil.
- Do not retry, close the writer or write elsewhere. Writer counts are between zero and the supplied byte length.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact greeting, empty/Unicode/spaces, no newline | TestPublicCase01/02 Buffer and custom writer | extra-newline, trim-name | quoted greeting/name and output |
| One call and original write error, including partial count | TestPublicCase03 | ignore-error, fresh-error, retry-write | failure error and call count |
| Short nil-error write returns io.ErrShortWrite | TestPublicCase04 and zero-count grader | ignore-short | short actual error and call count |
| Caller retains ownership on success, error and short write | TestWriterRemainsOpen covers all three outcomes with Close probe | close-writer, close-on-error, close-on-short | closes got N want 0 with writer count/error |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- ignore-error
- ignore-short
- fresh-error
- extra-newline
- trim-name
- retry-write
- close-writer
- close-on-error
- close-on-short

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, selected operation in `di/v2/di.go`. All declared consulted artifacts and SHA256 values were verified against the pinned local source. MIT notices and attribution were preserved; independently authored boundaries and defects extend the selected observable behavior. No generated canonical or reference-only material was reused. Private reference code and mutants are excluded from learner payloads.

## Independent review

The public Close probe now checks successful, failing and short writes with explicit close-count diagnostics; each outcome has a dedicated realistic defect.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 9 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/writer-greeting-001.yaml` reports technical pass and 9/9 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
