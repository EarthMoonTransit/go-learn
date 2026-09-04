# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v1`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 0 |
| `in-review` | 0 |
| `ready` | 3 |
| `rejected` | 0 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 1 |
| `implement` | 1 |
| `refactor` | 1 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **not-applicable** — fewer than 10 lifecycle tasks

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `defined-types-001` | `ready` | `1ca92b08192f60d8be360297356576a5e86be1766b126c95701b1c18e6b4a4be` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-sum-001` | `ready` | `34df0af6760295a955b081b76d364deddf71124051d77bd33344b5e54a010af9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `zero-state-001` | `ready` | `3118f1b85d3e6a9892ba6be4c5a1fbbcf7e1ad6f24f0de943614fcaa812316aa` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
