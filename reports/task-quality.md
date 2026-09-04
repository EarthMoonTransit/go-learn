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
| `defined-types-001` | `ready` | `a724a129ca2818ca4de85f0475f9ace4094549edb55de76b9e217d9f6339838b` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-sum-001` | `ready` | `5e506467f4c99b466747db36ecc736bd286b39bca362ba7d64b53f91a6493792` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `zero-state-001` | `ready` | `768cdcc42099fcc4ee8ec854335ee28126dc473334e9cb1be5fec8c3d4e6c67d` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
