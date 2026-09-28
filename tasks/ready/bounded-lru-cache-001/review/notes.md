# P5 / C17 generic LRU content review

One substantive SHAD adaptation with an algorithmic facet; intended transfer
role. ADR0024 (PR62) replaces the rejected integer draft's unassessed interface
primary with existing `generics.reusable-data-structures`. No schema, skill ID,
profile or mandatory route changes. The previous draft is preserved at 75df9ec;
it was never ready/routed. The generic decision and empty private-state scaffold
were independently accepted for authoring, not automatically for promotion.

The learner implements zero-state behavior, typed misses/presence, unrestricted
values, key identity and cross-operation invariants. Tests use nonempty slices
and structs containing slices, named/composite keys, and distinct pointer keys
with equal contents. Compile-time `V comparable` is an invalid API; it is not
counted among compiling behavioral mutants. No secondary map coverage is claimed:
a linear implementation is valid. This is distinct from generic-stack's LIFO
contract and dictionary/embedded-cache exercises: keyed bounded storage, recency,
eviction and ordered stoppable traversal interact across multiple K/V types.
Estimated 35–60 minutes remains uncalibrated. No learner evidence is claimed.

## Public requirement mapping

| Requirement | Public tests | Defects | Useful diagnostic |
| --- | --- | --- | --- |
| New independent/empty; usable zero Cache | EmptyZeroAndIndependent, ZeroCache, AssignmentAndIndependentGenericInstances | shared-instance, zero-value-cache | Range length/entries or Get actual/expected |
| Set assigns any V; zero capacity | GenericSlicesAndRecency, NamedCompositeKeysAndStructValues, EmptyZeroAndIndependent, AssignmentAndIndependentGenericInstances | noncomparable-value-rejected, named-key-rejected, zero-capacity, value-deep-copy | Get/Range values actual/expected |
| Stored zero/nil present; typed zero on miss | TypedMissesAndStoredNil, EmptyZeroAndIndependent, GenericSlicesAndRecency | stored-zero-missing, nil-value-rejected, missing-value | Get actual pair versus expected value/presence |
| Successful Get and Set refresh; misses inert | GetRefreshAndEviction, UpdateRefreshWithoutEviction, MissDoesNotChangeOrder, GenericSlicesAndRecency | get-no-refresh, update-no-refresh, miss-removes-oldest | ordered entries or missing lookup actual/expected |
| Oldest eviction; updates retain other entries | GetRefreshAndEviction, UpdateRefreshWithoutEviction, GenericSlicesAndRecency | evict-newest, update-duplicates | Get pair or Range length/entries actual/expected |
| Range once, oldest-first, early stop, no changes | RangeStopAndNoMutation, GenericSlicesAndRecency | range-reversed, range-ignores-stop | callback count/entries actual/expected |
| Clear repeatable, empty, capacity reusable | ClearAndReuse, GenericSlicesAndRecency, ZeroCache | clear-noop, clear-loses-capacity | post-Clear Get/Range actual/expected |
| Key equality, including zero/composite/pointer keys | ZeroKeyAndNegativeValue, NamedCompositeKeysAndStructValues, PointerKeyIdentity | zero-key-rejected, named-key-rejected, pointer-key-content | Get pair or Range keys actual/expected |

Names in the table have `TestPublic` prefix. Reference must pass both public and
full suites; every listed mutant must compile and fail public-only with a useful
diagnostic, independently of hidden coverage. The full suite adds deterministic
integer and composite-key/slice-value traces using an independent logical access
counter/map oracle. No randomness, sleeps, network, machine timing or private
state assertions. Assignment tests check retained slice aliasing and exact pointer
value identity without requiring deep copies. Cache independence does not imply
copying objects referenced by values.

## Provenance

PR60 admits SHAD commit 6718601587ec77cfaa632bbd29a8ca0716fe81ec and only
lrucache/{README.md,cache.go,lru.go,cache_test.go}, plus MIT notice preservation.
All five hashes were rechecked on 2026-09-28 against the pinned checkout and
match provenance.yaml. No upstream code was executed. The integer source's
recency, eviction, Range and Clear behavior is substantively adapted; the generic
API, usable zero Cache, explicit bounds and deterministic tests are local changes.
This is an adaptation, not original. Historical GitLab identity and copyright
notice remain in ATTRIBUTION.md. No reference-only or canonical Exercism material
was used. Source PR54/PR60 and decision PR62 must land with their gates before
this dependent content PR merges. Existing Theory needs no further development.

## Review and execution evidence

Independent review by `/root/review_intake` approved exact
8557cb332c619e0f970f4f85ca7f5fee083fe27c on 2026-09-28 with no blocking findings.
This supersedes the previous integer-draft assessment rejection.
Targeted checker and public-only probe execute authored code solely through the
checker-owned disposable runner. The probe uses a temporary Go test overlay in
internal/tasking; no production tooling file or canonical task execution is added.
Targeted technical checks pass with 20/20 mutants; the public-only reference
passes and all 20 mutants compile and fail useful public assertions. Logs:
.tmp/results/lru-generic-target.yaml and lru-public.log. Content digest:
f57610fc17ba67e99e5efc2a0923d199899d6855c16d9b7b4fdb7ee01e9cf9aa.
Normal draft → in-review → ready promotion is required before the final collection. Fixture
runs are technical evidence, not learner completions or mastery.
