# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v2`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 9 |
| `in-review` | 0 |
| `ready` | 5 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 4 |
| `implement` | 6 |
| `refactor` | 3 |
| `test-writing` | 2 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `clamp-tests-001` | `draft` | `3f7f29c5635d0a5581c33e08eaa0a296604c71855da8455acf13849ae3852277` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `counter-receiver-001` | `draft` | `7e46814fd59cc686a0bdd4d7dad718e79f30d039e51ea146af53dce9d4db8f04` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `defined-types-001` | `ready` | `c08c4e3554e127e537523fb2c6e26120c86acd6f94fe1953ad5473355d63c2ea` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `detached-append-001` | `ready` | `92eadf85530879648f7368ec23871fee045d6fef1ddb897fdd101a5d1be2ddd6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `acb4dd90a9ca9321bb3f84e4be6738806c4af67f64b511e2b2b3f8c2002f87b4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `17abcfb02a2d62a89bcc904a563f69a99c8a4103762989f0ef9cc0a625b1d03d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `formatter-contract-001` | `draft` | `a84c2b7ac9c60653164aa69605a346e7656285f14faf5a04932e526123ebb401` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `4579c0556922eb92d723a9b2e1be09176b301323a6741325d7029c9bb20be26e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `resource-finalize-001` | `draft` | `4c4ca3c70afe266879ed5346879e5669f6dece91335f69682f99e7c9729605c5` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `beb0a532940414f1e3e9346f1c96e85698b711592381e636a88aeec0ee3dc15d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `stream-normalize-001` | `draft` | `d829725c995bc5e200fb41b268c96c02df51cd4bda6a22778435191123eec267` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `107bc09aeffede5853ab4fd8306a3ab185e559fda9ad5025daceeeeb1fc90985` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `utf8-prefix-001` | `draft` | `5aed6135299bc4b811ab292a5fad6581d922249e5cbc85263a619f78bf447e14` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `variadic-sum-001` | `ready` | `93c3d07556ec11789574e2536e1be4cf127b614f90c60b0187da9276baac0112` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `zero-state-001` | `ready` | `fd51358857af2a0711e30fc320be40c17841548c90107ae990ec7527e7fddda9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
