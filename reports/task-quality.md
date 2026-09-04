# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v2`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 8 |
| `in-review` | 0 |
| `ready` | 20 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 10 |
| `implement` | 14 |
| `refactor` | 3 |
| `test-writing` | 2 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `0b961a7e937b235bee75bce33fe24c9ddc4e587c797e690e0b5cae91fc3e4ff8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `append-value-001` | `ready` | `ccbd6f1ce1257c2b408709f348109b4c236bdae7ee43796d8a55776cd08aba64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `0c6f87243593f195b85f00f77b208ebafdf7aae3e11972968a55824ecf4035e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `bounded-value-001` | `ready` | `4697ac79d9cab7e4c75cf01963a8bbf2287b63d343ad85dd24af59746e3e439f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `clamp-tests-001` | `ready` | `808fb55df8762ddf87884cbcc44bed1decbb552b58ec7772b7be1b31a1e09786` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `7e46814fd59cc686a0bdd4d7dad718e79f30d039e51ea146af53dce9d4db8f04` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `defined-types-001` | `ready` | `c08c4e3554e127e537523fb2c6e26120c86acd6f94fe1953ad5473355d63c2ea` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `detached-append-001` | `ready` | `92eadf85530879648f7368ec23871fee045d6fef1ddb897fdd101a5d1be2ddd6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `division-results-001` | `ready` | `24683f4e4c8d6a22bc310c8553a21bbea69b30d1b256983985e410ae39879100` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `acb4dd90a9ca9321bb3f84e4be6738806c4af67f64b511e2b2b3f8c2002f87b4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `c4b683f1485ecd0bf5024b711e69dcaca3e78e8660f5d2fb4305b56af04bd622` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-query-001` | `draft` | `17abcfb02a2d62a89bcc904a563f69a99c8a4103762989f0ef9cc0a625b1d03d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `exclusive-loop-001` | `ready` | `4c5af5fc8b5db6bdfbda29eb2859f0d4667da5b4d3940ed4caefb7f7568f9d78` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `formatter-contract-001` | `draft` | `a84c2b7ac9c60653164aa69605a346e7656285f14faf5a04932e526123ebb401` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `literal-bindings-001` | `ready` | `00e44d1098a4969fb6474a1b39292eff1bd4c2aa280ace26c90c5fb82afd86a7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-presence-001` | `ready` | `d41f9e745fbef836b03160342b005b73c0a6edb0bb256952eebedd9aa7167692` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `4579c0556922eb92d723a9b2e1be09176b301323a6741325d7029c9bb20be26e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-block-001` | `ready` | `669644e5d4f83e8c35ce64edb97c84b5d4520524c3d62e06ae8613c8b9db0a56` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `range-offset-001` | `ready` | `2106d091cd92eacd035ae44eaa34dc2714767b249dcc341b597f8809e85685c8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `resource-finalize-001` | `draft` | `4c4ca3c70afe266879ed5346879e5669f6dece91335f69682f99e7c9729605c5` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `beb0a532940414f1e3e9346f1c96e85698b711592381e636a88aeec0ee3dc15d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `slice-window-001` | `ready` | `a66220fae3390e098a975b7880e59a5df05007ee7f562aad1685ec92cd95bdc7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `d829725c995bc5e200fb41b268c96c02df51cd4bda6a22778435191123eec267` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `107bc09aeffede5853ab4fd8306a3ab185e559fda9ad5025daceeeeb1fc90985` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `utf8-prefix-001` | `draft` | `5aed6135299bc4b811ab292a5fad6581d922249e5cbc85263a619f78bf447e14` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `variadic-sum-001` | `ready` | `93c3d07556ec11789574e2536e1be4cf127b614f90c60b0187da9276baac0112` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `wide-sum-001` | `ready` | `9ee2fe93f5c9e10dfb0fd01bdddce3440945a264cc6ea894d7269d6a8becc197` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `zero-results-001` | `ready` | `9a8e5a91db311fabe0ed3b6063d5b58d5e10fe39936c7fff5547fc55bd5b801a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `fd51358857af2a0711e30fc320be40c17841548c90107ae990ec7527e7fddda9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
