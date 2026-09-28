# ADR 0024: Assess the LRU draft as a reusable generic container

- Status: proposed, pending independent review
- Date: 2026-09-28
- Scope: the existing bounded-lru-cache-001 draft in PR61

## Problem and decision

The integer draft supplies the complete Cache interface and method stubs. Its
runtime defects do not assess the declared interfaces.implicit-satisfaction
primary. Receiver semantics is not an honest replacement for supplied pointer
receivers. Keep the substantive SHAD recency/eviction/Range/Clear behavior and
revise this one draft into a reusable generic container:

```go
type Cache[K comparable, V any] struct { /* private state */ }
func New[K comparable, V any](capacity int) *Cache[K, V]
func (c *Cache[K, V]) Get(key K) (V, bool)
func (c *Cache[K, V]) Set(key K, value V)
func (c *Cache[K, V]) Range(f func(K, V) bool)
func (c *Cache[K, V]) Clear()
```

Use the existing generics.reusable-data-structures primary, implementation
activity and fixed-tests profile. No taxonomy identifier, schema, profile or
field meaning changes. Required prerequisites include generic constraints and
zero values. No theory expansion is needed.

The zero Cache behaves as capacity zero; nil receivers are outside the contract.
Capacity is nonnegative; calls are sequential; callbacks are non-nil and do not
call the cache. Do not copy Cache after use. Keys must compare equal to themselves
without panic. Values follow ordinary Go assignment semantics, including shared
references inside a value; no deep-copy guarantee or requirement is introduced.
No asymptotic complexity or allocation threshold is imposed.

## Observable assessment

Public tests must exercise nonempty generic values across update, lookup,
eviction and traversal, not merely instantiate types. Cover usable zero Cache;
exact typed zero on a miss; presence of stored zero/nil values; named/composite
and pointer-identity keys; noncomparable values such as slices and structs with
slice fields; independent instances; and the complete source-derived operation
sequence. Preserve useful actual/expected diagnostics and realistic compiling
mutants. Overrestricting V to comparable is a compile-time API violation, not a
compiling behavioral mutant or proof of semantic detection.

The source remains SHAD at its previously admitted commit and four artifacts.
Provenance/attribution distinguish the integer source contract from the locally
authored generic API, bounds and deterministic tests. Do not import additional
upstream material or label this adaptation original.

## Migration and rollback

This is a draft-only API revision. The draft has never been ready or routed; its
old state is retained at 75df9ec6d745ea669a21803e62194878e94345f4. Keep its stable
ID, recompute content digest and review evidence, and use task-promote for both
normal lifecycle transitions after independent content approval. Keep all 150
existing route addresses and every existing ready package unchanged. No learner
progress/workspace migration, compatibility exception or reset is introduced.
A manually extracted authoring draft is not a supported learner workspace.

The new package stays outside the mandatory route until placement is separately
reviewed. Reverting this content change restores the old draft and its explicit
assessment rejection; it never retroactively approves that draft. Release/export
still needs its normal commit binding and coverage gate. Preparation follows
ADR0022, independent review and exact-head GitHub validate before merge.
