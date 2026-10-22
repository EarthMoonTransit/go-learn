# I6 author precheck

The sole reviewer approved the bounded contract before authoring. Actual independent method/content/execution review approved by /root/reviewer2026-10-22. Final batch/dependency/local/CI gates still precede delivery credit.

TASK preserves all approved 142-word contract clauses, formatted into the existing mandatory Context/API/Requirements/Constraints sections. The checker counts 160 words including its syntax tokens; there are six requirements. The input assumption moved unchanged into Constraints. The public matrix covers every basic behavior; grader extends values, split positions and accepted boundaries only.

| Requirement | Public assertion | Wrong variant | Useful diagnostic |
| --- | --- | --- | --- |
| zero count, no closing | TestPublicZeroCount; all checkPublicSource calls | reads-on-zero, closes-input | Read calls / Close calls actual versus zero |
| partial reads and zero-byte progress | TestPublicBoundedPrefix/chunk-0,1,2 | assumes-full-read | accepted records or unexpected error versus expected complete values |
| three-byte big-endian values and order | TestPublicBoundedPrefix | little-endian, wide-frame-read | record values / requests exceeding current record |
| bounded prefix and no subsequent probe | TestPublicBoundedPrefix; bad-version-with-tail | probes-after-count, wide-frame-read | consumed bytes / current-record request bound |
| never publish partial or malformed records | one-byte-eof; incomplete-invalid-version; bad-version-with-tail/eof | publishes-partial, accepts-version | accepted prefix and exact selected cause |
| retain accepted prefix on failure | boundary-eof; partial-next-eof/error | drops-prefix | accepted records and wanted length |
| positive bytes processed before error | complete-final-eof/error; complete-intermediate-error | discards-bytes-with-error | accepted prefix / actual nil / exact transport cause |
| early and final exact EOF differ | boundary-eof; complete-final-eof | early-eof-success, final-eof-failure | missing failure versus required UnexpectedEOF / actual nil |
| non-EOF error wins even at count | complete-final-error | drops-error-at-count | missing RecordFailure with original transport cause |
| non-EOF beats same-read bad format | bad-version-reader-error | format-before-reader-error | Unwrap exact cause / errors.Is |
| stop after first read error | complete-final-error; partial-next-error | reads-after-error | reads after terminal error actual versus zero |
| exact transport cause, including wrapped EOF | wrapped-eof-is-reader-error | unwraps-reader-cause, no-unwrap | Unwrap exact cause / errors.Is |
| typed recovery and index may equal count | complete-final-error; TestPublicRetainedFailures | unstructured-error, fixed-index | failure dynamic type / RecordIndex actual versus prefix boundary |
| exact error context | every failure row | wrong-error-text | Error actual versus expected complete string |
| independent retained errors | TestPublicRetainedFailures final re-observation | shared-error-state | first boundary/cause/text after second call |
| actual nil on success | TestPublicZeroCount; TestPublicBoundedPrefix; complete-final-eof | typed-nil-success | dynamic type on success, no method invoked on typed nil |

There are 22 named, finite, ordinary wrong variants. The larger explicit set keeps unrelated terminal/type/ownership defects separate; no spin, panic, compile-error or timeout mutant is credited. Every variant must fail the public-only copied fixture with an assertion. Reader helpers preserve unread data and observe request bounds without panicking. They do not impose exact call counts or read sizes, except the count-zero no-read requirement. Typed-nil errors are inspected safely before methods.

## Actual method boundary

Starter provides Record/ErrVersion/RecordFailure and an empty private recordFailure with stub pointer-method signatures. It supplies no index/cause fields, constructor, parsing decisions or working error methods. Reference adds useful at/cause state, independently constructs each error at a failure boundary, preserves exact cause and retained error lifetime, and returns actual nil on success. The index explains the accepted prefix; cause enables classification. Incidental wire storage is not exposed. A separate correct control uses nested private recovery fields and one-byte reads; public/full must accept it.

Supplied interface, method signatures and pointer receiver selection earn no authorship credit. Method acceptance must inspect learner-owned storage/construction/behavior; tests alone establish behavior, not this design choice. The approved limited repair is errors.typed, with reader-side io.reader-writer-contracts and bounded-prefix encoding.streaming-strictness applications. Actual JSON-shape prerequisite is included. No binary-marshaling, errors.As/Is, test-double, copy-adapter, scanner, buffer, resource-close or pipe application is claimed from supplied APIs/harness.

## Provenance and execution boundary

Original ADR 0008 creation/digest and sources[]; no external task text/code/tests copied, no reference-only use. Existing structured-error-tests, bounded-copy, single-json-value and binary-record-tests are unchanged preparation. Contract/source provenance comparison is retained in the I6 planning proposal. Only task-check may execute task code, in its disposable workspaces. Canonical/public-only/correct-control receipts are author evidence; promotion and independent review belong to root/reviewer.

Independent actual method acceptance: Actual private index/cause storage, independent construction, useful recovery methods and parser decisions accepted.22ordinary public diagnostics and nested-state/one-byte-read correct alternate accepted. Component IO plus typed-error repair, reader and bounded-prefix streaming applications accepted; supplied receiver/interface/As/Is/doubles are not authorship credit.
