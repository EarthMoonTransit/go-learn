# Taxonomy v3 audit against Exercism Go

- Source commit: `ae1477abfeeb9fbeaf971caee220e91f9e4df790`
- Taxonomy SHA-256: `8fc463c93634564e7972117ea7f7c1f90d4e5cb82659f2f631698ef68ba4ab64`

Taxonomy version: **3**. Explicit concept mappings: **38**. Accepted reviewed candidate evidence records: **18**. Candidate evidence is not ready-task coverage.

## Main findings

### Broad mechanisms resolved in taxonomy v3

Taxonomy v3 applies the reviewed decompositions for slice storage, rune and UTF-8 semantics, interfaces, errors, testing, goroutine lifecycle, channels, synchronization, and context. Floating-point semantics, string formatting, `fmt.Stringer`, and the standard-library `strings` package now have explicit assignable skills.

### Deliberate non-mappings

`basics` and `numbers` remain taxonomy gaps because each is an upstream umbrella rather than an atomic learner capability. `comments`, `randomness`, and `regular-expressions` remain scope-deferred.

## Audit trail for the applied decomposition

The detailed accepted operations are frozen by `docs/adr/0003-taxonomy-v3.md` and `curriculum/migrations/v2-to-v3.yaml`. The observations below explain the original pressure for the split.

| Area | v2 observation | v3 treatment | Candidate evidence |
|---|---|---|---|
| Slices | \`slices.model\`, \`capacity-bounds\`, \`append-aliasing\`, \`growth-allocation\` overlap across nil/empty state, slicing, allocation and aliasing. | Consider construction/nil/empty; len/cap/slicing; append/reallocation; aliasing/defensive-copy; copy/overlap; preallocation. | `anagram`, `animal-magic`, `bird-watcher`, `card-tricks`, `census`, `chessboard` (+7 more) |
| Strings / bytes / runes / UTF-8 | \`strings.construction-conversion\` and \`unicode.utf8-runes\` each cover several independently testable mechanics. | Split construction/allocation from string↔[]byte↔[]rune conversion; add strings package and formatting; separate rune model from UTF-8 encode/decode. | `acronym`, `affine-cipher`, `airport-robot`, `anagram`, `atbash-cipher`, `blackjack` (+47 more); `welcome-to-tech-palace`; `booking-up-for-beauty`, `census`, `clock`, `election-day`, `jedliks-toys`, `ledger` (+5 more); `logs-logs-logs`, `micro-blog`, `parsing-log-files`, `reverse-string`; `meteorology` |
| Methods / pointers | \`methods.receivers-method-sets\` combines declaration, receiver mutation, method sets and addressable-call convenience; \`pointers.receiver-choice\` overlaps it. | Split receiver semantics from method sets; merge receiver-choice into receiver semantics; require manual evidence for method values/expressions. | `airport-robot`, `census`, `high-scores`, `jedliks-toys`, `meteorology`, `parsing-log-files` (+2 more); `census`, `election-day`, `jedliks-toys` |
| Interfaces | Dynamic interface representation is coupled with typed-nil behavior; assertions are coupled with switches. | Split dynamic type/value from typed nil, and type assertions from type switches. | `airport-robot`, `meteorology`, `paasio`, `react`, `simple-cipher`, `sorting-room` (+1 more); `sorting-room` |
| Errors | Creation is coupled with propagation, wrapping with \`Is\`, and typed-error design with \`As\`. | Separate creation/sentinels, propagation/context, wrapping-chain construction, chain inspection, typed errors, and panic/recover boundaries. | `collatz-conjecture`, `error-handling`, `expenses`, `nth-prime`, `nucleotide-count`, `queen-attack` (+1 more) |
| Testing | Exercism has tests but does not label testing as a learning objective; current table/subtests, helper/cleanup, parallel/race and examples/golden pairs are broad. | Add basic test discovery/\`testing.T\`/failure diagnostics, then consider splitting the existing paired skills. | No direct Exercism curriculum evidence |
| Goroutines / channels / sync | Only low-confidence topic \`concurrency\` exists; it cannot select a specific concurrency skill. Several current skills combine ownership/results, leaks/panics, operations/close, select/nil and pipeline/fan/cancellation. | Split those mechanics; consider directional channels and data-race/happens-before skills through a separate ADR. | `bank-account`, `paasio`, `parallel-letter-frequency`, `robot-simulator` |
| Context | No Exercism concept or candidate provides context evidence. | Keep current decomposition unchanged from this audit; consider cancellation causes only through an independent curriculum decision. | No direct Exercism curriculum evidence |

## Mapping policy

1. Version mappings by source id and exact commit.
2. Sort labels and candidates lexicographically.
3. Count only `concepts`/`practices` as declared targets.
4. Keep prerequisites as dependency evidence and topics as low-confidence inventory.
5. Do not normalize hyphens, underscores, plurals, or use fuzzy/stemming matches.
6. Mark one-to-many labels ambiguous instead of assigning every possible skill.
7. Permit only reviewed overrides keyed by source commit, exercise kind, and slug.
8. Report WIP/deprecated exercises separately.
