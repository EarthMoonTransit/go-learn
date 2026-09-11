# P5 / C17 content review — request changes: assessment

One substantive adaptation with an algorithmic facet; intended role transfer.
Independent review by `/root/review_intake` found a ready blocker: provisional
`interfaces.implicit-satisfaction` is not assessed because the compiling starter
already supplies every method. Mutants preserve satisfaction. Do not count this
label as coverage. `methods.receiver-semantics` is not a justified replacement:
pointer receivers are already supplied. No secondary map coverage is claimed
because a slice implementation is valid. The package remains draft, with no
promotion until a separately reviewed assessment decision resolves this mismatch.
Estimated 35–60 minutes is uncalibrated. The task stays outside the mandatory
150-step route; no old ID, number, progress, hint or timing data changes. The
interface-contracts Theory page is prepared separately in P2/PR58, whose merge
is a release dependency, not duplicated in this content batch.

## Public requirement mapping

| Requirement | Public assertion | Defect | Diagnostic |
| --- | --- | --- | --- |
| Independent empty usable instances | TestPublicEmptyZeroAndIndependent | shared-instance | Range actual/expected per instance |
| Arbitrary integers, including key zero and negative values | TestPublicZeroKeyAndNegativeValue | zero-key-rejected, negative-value-clamped | entries/Get actual and expected signed values |
| Store/update; zero capacity | TestPublicEmptyZeroAndIndependent, TestPublicUpdateRefreshWithoutEviction | zero-capacity, update-duplicates | Get or ordered entries actual/expected |
| Presence including zero, missing pair | TestPublicEmptyZeroAndIndependent | stored-zero-missing, missing-value | Get(key) actual pair; expected pair |
| Successful Get refresh | TestPublicGetRefreshAndEviction | get-no-refresh | ordered entries actual/expected |
| Set refresh without extra eviction | TestPublicUpdateRefreshWithoutEviction | update-no-refresh, update-duplicates | ordered entries actual/expected |
| Miss does not mutate | TestPublicMissDoesNotChangeOrder | miss-removes-oldest | ordered entries actual/expected |
| Least-recent eviction | TestPublicGetRefreshAndEviction | evict-newest | missing Get or ordered entries actual/expected |
| Range order and early stop; unchanged state | TestPublicRangeStopAndNoMutation | range-reversed, range-ignores-stop | callback entries/count actual/expected |
| Clear, repeat and reuse capacity | TestPublicClearAndReuse | clear-noop | ordered entries actual/expected |

The public assertions map every listed defect to an ordinary contract requirement.
Full reference/mutant runs do not by themselves prove public-only detection. The full
suite adds deterministic mixed traces with an independent map/logical-counter
oracle; no random seed, sleep, benchmark threshold or implementation-shape test.
The linear reference deliberately demonstrates that no O(1) implementation is
required. Grader uses shared public test helpers in the disposable combined
package. No learner solutions or real learner history are inspected.

## Provenance and distinction

P3/PR60 admitted exactly four SHAD artifacts and root MIT notice preservation;
all exact hashes are in provenance.yaml and independently verifiable in the
pinned source-sync checkout. No upstream source code was executed. The content
adapts the interface and behavioral scenarios substantively; it is not original.
Recency across Get/Set, eviction, ordered callback traversal and reusable Clear
are distinct from dictionary-update, generic-stack and embedded-cache practice.
No other source is included via transitive inference. Source PR54/PR60 and P2
must pass exact-head validate and merge before this dependent content PR merges.

## Review outcome

Source provenance, license, behavior and reference were independently accepted;
assessment was not. The public zero-key and negative-value gap was repaired with
explicit assertions and two realistic mutants. No lifecycle movement is requested.
The full technical collection records draft status and request-changes review;
its technical PASS does not authorize ready, release or assessed coverage.

Targeted `task-check` passes reference/public/full and detects 14/14 defects.
A temporary tooling overlay (no canonical task execution) used the existing
checker-owned RunVariant isolation runner with only tests/public. All 14 defects
compiled and failed with expected/actual diagnostics, and the reference passed;
canonical digest remained 251bae4444704eebe062afe05499d98771820645954faafe1f16d00b42fc8be6.
The local log is .tmp/results/lru-public.log. This probe does not supply missing
assessment evidence or bypass the request-changes decision.
