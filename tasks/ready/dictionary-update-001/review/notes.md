# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- For an existing key replace its value exactly and return nil; a stored empty value counts as existing.
- For an absent key return ErrAbsent, preserving the entire dictionary; nil maps contain no keys.
- Empty keys and new empty values are valid. Preserve unrelated entries and do not normalize key spelling.
- The failure has dynamic type DictionaryError and Error text "cannot update absent key"; a newly allocated generic error is insufficient.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Replace only an existing entry, preserving unrelated entries | TestPublicCase01 | ignore-value, clear-others | map and error actual result |
| Absent key returns typed ErrAbsent and preserves entire map | TestPublicCase02 clones nil/empty/nonempty maps and compares them | insert-missing, generic-error, clear-on-missing | dynamic error type/text and absent changed map |
| Empty stored values/keys and new empty values are valid | TestPublicCase03 | empty-missing, empty-key-rejected | existing/new empty map and error result |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- insert-missing
- empty-missing
- generic-error
- ignore-value
- clear-others
- empty-key-rejected
- clear-on-missing

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, operation in `maps/v7/dictionary.go`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

Complete cloned-map comparison exposes clearing the map on a missing key, including preservation of empty/nil states.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 7 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/dictionary-update-001.yaml` reports technical pass and 7/7 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
