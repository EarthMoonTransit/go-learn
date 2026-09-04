# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v2`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 6 |
| `in-review` | 0 |
| `ready` | 3 |
| `rejected` | 0 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 3 |
| `implement` | 4 |
| `refactor` | 2 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **not-applicable** — fewer than 10 lifecycle tasks

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `counter-receiver-001` | `draft` | `4af83ae2b949960fc16dd9d4d9756f5aa5c8a3ebb644e4aabca68a3e0effa584` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `defined-types-001` | `ready` | `c08c4e3554e127e537523fb2c6e26120c86acd6f94fe1953ad5473355d63c2ea` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `detached-append-001` | `draft` | `34185a5bd02994ca3a8aebf2e15c86874837dd3af87ba509ca8beb58e9d64e00` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `35ec59c7ca9254e036f01fe31bc29e647aa833076f8cc371b222f024b34adf7c` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatter-contract-001` | `draft` | `e0ae2d891143a270f07831691b8bee324dc0cf44b92c046c2f6a6b2b7130d1b1` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `draft` | `ef3670f2152eea9da80f2e23353b64c0ef5f3c666ca7931debec92d50bee7514` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-prefix-001` | `draft` | `ec3413095ae741e274037c5f20c6c0cf78d3745ab393b49805815364de1af8b5` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `variadic-sum-001` | `ready` | `93c3d07556ec11789574e2536e1be4cf127b614f90c60b0187da9276baac0112` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `zero-state-001` | `ready` | `fd51358857af2a0711e30fc320be40c17841548c90107ae990ec7527e7fddda9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
