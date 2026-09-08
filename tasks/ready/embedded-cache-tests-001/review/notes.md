# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- NewCache uses the supplied non-nil Feed; Lookup is non-nil. Promoted Count returns the current Feed.Total.
- First Fetch calls Feed.Lookup through Feed.Fetch, increments Feed.Calls, and caches a successful Document under the exact requested key.
- Repeated successful keys return the cached document without another lookup. Different keys and cache instances are independent.
- Errors return a zero Document and the original error; failures are not cached. Successful empty documents are valid.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Promoted Count tracks supplied Feed.Total | TestContractCase02 changes Total after construction | fixed-count | promoted/live count actual/expected |
| Cache successful documents by exact requested key; independent instances | TestContractCase02 uses different document key and repeated a/A across instances | uncached, cache-by-document-key, shared-cache | key/document mismatch and lookup call count |
| Return zero document/original error and retry failures; cache empty success | TestContractCase03 transitions two errors to empty success | error-cached, partial-error-document, empty-not-cached, fresh-error | failure result/original identity and retry/empty call count |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- uncached
- cache-by-document-key
- fixed-count
- error-cached
- partial-error-document
- empty-not-cached
- shared-cache
- fresh-error

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/embedding/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

Promotion/override, exact-key caching and independent cache state are observable. The fresh-error defect independently enforces original error identity while preserving zero results and retries. This introduction does not claim ambiguity-resolution mastery.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 8 required defects, hints, route prerequisites and provenance. `.tmp/checks/embedded-cache-tests-001.yaml` reports technical pass and 8/8 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.

The saved `review/weak-error-identity.txt` removes only exact sentinel equality from the strong suite. Its author-provided `.tmp/checks/embedded-cache-weak.yaml` reports reference compatibility and the starter-full insufficiency gate passing through a disposable copy. The report's 8/8 detected summary concerns the strong reference, not this weak fixture.
