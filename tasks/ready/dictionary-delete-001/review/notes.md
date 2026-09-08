# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Delete a present key and return nil, including a key with an empty stored value.
- For an absent key return ErrAbsent and preserve the entire dictionary; nil maps contain no keys.
- Keys are exact strings and may be empty. Preserve all unrelated entries.
- The error must be the supplied DictionaryError constant, not a fresh generic error.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Delete only a present key, including empty stored value/key | TestPublicCase01/02 | empty-is-absent, clear-all, keep-zero | dictionary contents and error |
| Absent and nil preserve entire dictionary and return typed constant | TestPublicCase03 cloned nil/empty/case-sensitive map; repeated-delete grader | always-nil, fresh-error, missing-clears | actual dictionary/error versus unchanged,ErrAbsent |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- always-nil
- empty-is-absent
- clear-all
- keep-zero
- fresh-error
- missing-clears

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, selected operation in `maps/v7/dictionary.go`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Presence, typed absent error and preservation of unrelated entries distinguish Delete from earlier Search/Add/Update operations.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 6 required defects, hints, route prerequisites and provenance. `.tmp/checks/dictionary-delete-001.yaml` reports technical pass with 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
