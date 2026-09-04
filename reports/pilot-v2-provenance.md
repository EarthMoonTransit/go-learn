# Foundation pilot v2: provenance

All task text, tests, implementations, mutants, hints, and review notes were
independently authored. No `problem-specifications` statement or canonical case
was used. Exercism Go contributed concept-direction metadata only, under the
pinned `exercism-go` registry record:

- repository: `https://github.com/exercism/go`
- commit: `ae1477abfeeb9fbeaf971caee220e91f9e4df790`
- license: `MIT`
- usage mode: `adapt`
- artifact usage in every adapted task: `concept-only`

## Adapted tasks

| Task | Pinned metadata artifact | SHA-256 |
| --- | --- | --- |
| `effective-rate-001` | `exercises/concept/cars-assemble/.meta/config.json` | `f26918431d75d89c1c2c0ec0d2b6e40b1492e608f50df89eeb7406f63eb8710a` |
| `detached-append-001` | `exercises/concept/card-tricks/.meta/config.json` | `36de25af5278bafa8add6eea7eebfe6a15a229e3f347351d11685d77e1a73dbb` |
| `utf8-prefix-001` | `exercises/practice/micro-blog/.meta/config.json` | `f16255121c09ac8d78c531e414dd701386cc105d58f335049602d455f2266eae` |
| `map-take-001` | `exercises/concept/gross-store/.meta/config.json` | `c6c37b8652475867b14518ed1981d6218d34bacd64a06b96c4b6925774efe249` |
| `counter-receiver-001` | `exercises/concept/jedliks-toys/.meta/config.json` | `2c4734f0e13f4673a656db93609095ee0f57cee34edc9627cb042acea6024f16` |
| `formatter-contract-001` | `exercises/concept/airport-robot/.meta/config.json` | `7268a45f4d7920ed81c8d34f74f4233c24801778915b13fa31cc406d0fef4b58` |
| `stream-normalize-001` | `exercises/practice/tournament/.meta/config.json` | `3d52a832037de5e80d189906b688cdaf2d9132daa161fb4d9cb05f725bc0ff35` |
| `clamp-tests-001` | `exercises/practice/counter/.meta/config.json` | `3912de6ab6afb5ec21ace1948e91b3b47df577f58583538007a1f48c1dd746e6` |
| `rune-offset-tests-001` | `exercises/concept/logs-logs-logs/.meta/config.json` | `09cada090c81b88da41d31f8dccc2c903372e86cc33843b893d25a163c100d59` |

All nine directions remain selected in
`catalog/seed-plans/exercism-foundation-pilot.yaml`. The requested
`micro-blog`/`reverse-string` choice uses `micro-blog`; the second approved
Unicode direction uses `logs-logs-logs`. No source direction was silently
substituted.

## Original tasks

| Task | Direction |
| --- | --- |
| `typed-nil-resource-001` | typed nil inside an interface |
| `error-chain-query-001` | wrapping plus `errors.Is` and `errors.As` |
| `resource-finalize-001` | defer, resource cleanup, and error propagation |

The three original packages use project identity `go-practice-bank`, project
repository `https://github.com/EarthMoonTransit/go-learn`, license
`Apache-2.0`, and an empty external source list as required by provenance v2.

Each task pins its authored content digest, attribution path, creation date, and
deferred release-manifest path. Release binding remains deferred because no
pilot task is ready or approved.

## Deviation

The requested Unicode fuzz-or-mutation direction is implemented as a mutation
test submission. `fuzz-seed-tests` exists in schema v2 but `task-check/v2`
currently rejects it as reserved and non-executable. This is a checker
capability deviation, not a seed or taxonomy substitution.
