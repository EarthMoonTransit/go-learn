# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v3`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 6 |
| `in-review` | 0 |
| `ready` | 172 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 42 |
| `implement` | 93 |
| `refactor` | 3 |
| `test-writing` | 41 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `0bcf9e2bc7b8431d6de3ef428214cfba5b364c20b59532710df4f7f97193da26` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `d0a0af9477954d32765f366107f26679cc75cbfa6c6bf7d37331b06b6ebd5c0c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `a9e76426380a3be11f873b7b1d54f9a2f8b1dec5d045abad9cc08f1a3c7c43f3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `f0c2f3f5a46211fda0b3b0eaba457d731f23091ebedc1808f568cba728b9cc99` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `c30f90fc8347f4ebe95bb8927cf7a3438749d1fad8bf2c6a1d94d2e736734173` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `b9c49f4713fc6de1b2395bd5b98d64d17decf90307d2d446d807ffb1de4b2142` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `8bff057c6d456019b0278aa380eb94e67c61b8a2a51a2cc1e88ba12b7bbac0fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `36b6954e86bce1427594341b3b3a8142ef10d6318f8f844faca70225dc561c7d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `d175a3e036e6d02c0c55b1c90b652050f5ec36fe0b082105a549fd002b011bf7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `cf07504f15abf2f8ef95c10e7e6a787dd3cdd9565ddcb4413d4a6a8c5b32d260` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `689867c4e27aa826de87372c427dbe19f9e8a73a31efe53818866ee877dc4d42` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `bfd558f1497d8df6d97c57f68bc2854673ff33ca6425672cd13c590b19fd7b5e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `a629a78ffaf8a7b79dfbf5a0e369ea1826e850a87b629b99e799cbb036fe3563` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `5d95bc8eecdcaa6775799ca4fec16966027a4bedd43a3ac180fd56a574be4d4b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `77112e85d161c1d21ee1aaaeeb1079d14bd4d1c5ce61eacb23a886c0e71c1cb3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `793bdda8fc564d2b96270444f061a0d225cb68bfb188c3f80e8bbddae53ba5fe` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `cancelable-forward-001` | `ready` | `08fd80eb823ca226bc02a5f51384cdcb0740c606303ea5907e413b0fc98e6f02` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `capture-panic-001` | `ready` | `508cafeb127c2e64161d5d45e78bde66360afa205a5f00a65928f643a5aa754f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `captured-method-values-001` | `ready` | `74e85b6175ea0a3738760663b57bc63db0f46ef3e1496e8ab4073105db809f71` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `category-closure-tests-001` | `ready` | `afcfff0a778133f047254497b43d2d7af0924aa833315c2279ea4775e6394e3a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `bf83e4669722c33f7389b2e0359a0a86a1e8e464db3f40fde02e56c3c3aa725a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `a4a8a652f4bb6e9099f180328d59e88c8472f2b658a6d9c827afb30e63c2b90f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `91ba6826d7cf5642a5b082f3bf36e0b92ff1ee64fc877a6845cd8b0d19b46190` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `b3b1fb5efba2c1fd3599b637232425dfa7530373e5eb1e2619e21b1cd40fd6ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `8c6c7566f0fb46926ef911896b285ec6ebd8d2a64d00e89d0e2eda2bffffdcba` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `e2cb8264264d125e577b4bbcad36518cb6f5d1c94bb538d76b0b785f7034cd99` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `composite-key-table-001` | `ready` | `5445de4d4d352033e832fd222c02dcb4a911e6868409a5bb8da1037764996ffb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `16711a6d365542030cf33416a5f9e5e67c18c77bf2773a777fda3011d21faa3b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `c06e52c3f1ab33b9bacd5bef5652cc200080a6fd8076a2f82b2218332ac4ed2a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `fcc7c1eb2effa74ff50d83e174a75e47a64be32118debe34452670edeae9813e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `d49e5ecbfa7e7e9c2bf94d800e714d7668da42b87b202ea538613d2a00d76254` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `96918902537ce4247e0b87f4c137362096646198f2dbe99e2d7ab335bb61ec41` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `872e40d343875f8db642839e83631f1c68f21fde214810157096feb1eb4672a6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `4f219c63e2cf604996e30e6ad49dd58db8a75a7d9709fd35234b6381a2f218cc` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `23d2748495e2ce13715d49fd3313ef987ce0b861fba1b5e66610685b63c0c852` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `4d064b80737a2bc505c630d2222df8b17e740c5a4a088c798f34b094a7ac1a1a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defer-trace-001` | `ready` | `c0570c6332ef45546f2aae75410cf0f6dedc505336dfc03d14ae4c9e31360293` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `deferred-release-001` | `ready` | `eb6dfcdf93912e5222e6e70618385ad050f9720987178d74adb372f4e619822b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `275a8bbac1e52cb0c164be8f9cd072a44c55283a57a42b91cd010c4998ffc662` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `d61608e131de0f579c49d3cb83a28e716b68993e1218563077e58aa2d26d2f85` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `8913cddb6bb7d66e1bc83ddddde432a504ea041001d680121a7435793d28ffb8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `2a5183a4682fbb561894db1885196f14dbad2c9c527016d003fd60c8d2819578` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `81593c25990dba012db878e3bfb2092dde2dd132769a6f29867508cfcc41f3a0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `6996fe4804d5bdcbfcd6eca09d3c60382312f31b27e499c160ccd73d9ea64d40` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `90609a83b1d0d433bd49a4f324b8ad7d7b6e329a4cc62d5c879d01fdba20ccdf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `6fe73d409684591e79aad8b7ceb652431365d71ca2a2e321d2a328e240b800f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `e137a700eebc7771212501f01bd4a05835c8706c2752209a8599915ad2740177` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `2f1be89e118811fe06131ff2404c36a37dd266676e57b3666d9492f93e2d4f90` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `7adecfdb2245d682d715fbabc15a12d58ff5021f760d4916e9607da1693132d8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `a4ca8b3b780446aa96939d2380ca9309c7d42e3a172574cb7019ba4943f6ecce` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `d42bca84abe55b4fa99c45e31d1dd6b147540326caa6f4e37f78a077feb42476` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `5cb94f4b8521d03edfd5e180e40e5956dc43f6a9ccb6d616df88b08e1b4060c5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `cd2707dcf8552b47aa6044ce9174be6f516702a43501f4ddbfeedce7e95c9a30` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `884f110141156dff94494f143c3dd95242f8690f637a28e895cc632d22e8970e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `d5aa087eedaceb4e1d252647fcc172645c0fafb04e5e59071e2e2626cee464ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `8ade28cfd6ff0662bfc2f3c25c9f69675464f9ade055a2651d64d32b074e6e42` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `54a3309488666f3c1b33bfa15b28bd0ff93a9fabf5ee8e294c0dbcec76fb8ace` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `05150ba4a2ffa90a1b8181db944375174da6be253c7fdf4732ce1026a4fcc389` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `62100c2cd843b1d4e7df7d94038284a733e272e9cf5ea4a6bb387900784646f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `6227eb07d19adacda7f75102d081ed6aec29ac226047e84ddd7a4d6111c8b80c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `eb6c940adfcdbabe5fd52c09ec42a6d9ec7a214342c11ddbd40f7d7697b02298` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `71b7c19c13d0198ea0930af93c11e85f25c6517ec30b6d05d0059079dcab99ac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `924e55461fbd4c4b205d2282b1b4f4d96d39b132ef7e996c6a08b45721bedee7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `0c3af1c9f170aa51417dda58265fc8efe550f4feaaad41d1df09aaadb291278b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `0944bd358bef3aa3e3687213d4dd99a9ee08499985ae891e3e5a38286df3fb0e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `ready` | `296815ec004b5756e2e5e1dffc4851efcc2532b17be6c8ef09a5af2525bc5ba3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `fractional-band-001` | `ready` | `63d71bfa6747f0e20e10d18ee6f8516ce9f4a8580fe5f80f33e0391b9b791ac3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gated-worker-arguments-001` | `ready` | `ad943eb979ffa80758a21dfd56dd3ea5127451200101a8e66eb79b024057d298` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `gcd-reduction-001` | `ready` | `a1642c6f13db0ff3a9dc138a884a498ca24c96bf2795de829683215d227fecda` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `deefcc19acb3052194e00a9b1c7a096c4706f79947b1e9b01eda2fa95cc1e794` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `04bd5a45202d4f3d49f986b5043f5ee70e73aef7d0be2fd381b5f2ec402eb911` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `c9719f099ef82aa02117c4d066c79fc7df4561f47a8ea1fc8ed97f53285bc2c5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `72508b98ac29cd4dfc8fb03c596d8690333e4d80741c1876953c1bc00cee4148` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `3f30b464d3c0e1c8733cacd267a2138b05b0bfb29401fd4e6d936fb527cff818` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `7f1f41024eb7d83ebfa2f517f4ae74e201197b4a916029828d464fa0e8771d89` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `844ffa80b98a253f69e22d939807fbd7425cde45fd3c8e79d6b61f0afd0ce30a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `81f004d8e865b14cfe6b1940b733b3abbfc9dd6ba3c72302a332e84bf465e88c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `indexed-job-results-001` | `ready` | `cb90eea9bcfa51126fbd42a2f53c9dccd15f4c5b39707b5ab303eed33bdf0aa8` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `initial-structs-001` | `ready` | `7e028a4a3509f3f77ef7e2c28b9c81b705bb883ace987aadac12ddc751c929bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `c6c9af532a4101918e30edd8172732537240816bd8bccd7f263abcc12cc59449` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `c6a36cfd703c900ac52bd80015c803447c5f651955efaa642b63630c9ffb9292` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `f7eeafe32c8681ae2e2f30576868e748e8a9d831e2f52b6e88ad27f109858042` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `d18e0acaf9ce77aba3ba5258fad769352df08e8e94d74bc2573f4f56fae466bc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `47a10e80ff7205102477f85dd37b55343681c6620f18efce53f58143f6035616` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `launch-provided-job-001` | `ready` | `cb1ca1a60d4f2e84926d82cfcab2350995888f07583527f61963108864ddf09f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `lazy-probe-001` | `ready` | `1a69218f42c29df18befd36a13c4f02d7c59759c9562f50d9b662d4a9cc88b97` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `07bf57e471a1d60bac02f17d80af0f161ac904940e0bb9eaba833cd33575896e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `431ed1b1a5adf8cb0df03ad0d749b62b8956deeea5e3cc424b56fb959ef79553` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `fce4bec49cf27f12dd6a906e1ba472bc03ffb6a6bd914394d97dd69074570e37` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `1176159b03d075b29248e549fb4fb491226cb9b3052af7f5f34b1f2cd36ff315` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `64f402f223a3dc715b7d42af2c5cc7aee306806f4560984a76cfb4de6985291f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `mapped-records-001` | `ready` | `44db24936623ce081cdbb0b965f7cf473b089eafc5a3b129052215f58f32b5b2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `merge-two-streams-001` | `ready` | `106b45098c5bda401d81d91aedb0acb98d11e8d09294f0d6466ee6d35f91d43a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `month-enumeration-001` | `ready` | `5ea7611478cd6369acd616016cc8506fee0489fcb16977eec69a13700fab0197` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `052d66d430b11f613668026bf90cf46f48acd2283cadcddb789cde067ab15306` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `f6ca9c17790da980eff2392c06ea1089c8e57e8915f1dfd9b5742e6d10acc008` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `f2a303103fb226cd118c0b9c59725b339e4633d745e45fc2c4bbfff2b5a10a64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `564d0bc1eab0d4b91006b48047b255d42a80faaf6f9d1a66ff5f28a14a3a2e95` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `0f104de1bc8bb9bb0d87f651c50225a37f1f430757599829ae8f5263775baf7b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `named-result-adjustment-001` | `ready` | `f20783268d34ef1808a1e5915d3af11d08fef75413a123fa099f04b5eacf808c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `nested-slot-count-001` | `ready` | `6fa53d08ef758ff8923c1da97e5f4fcca2f4ea2d3c68cbd247174f6117b32a3e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `3fd119d0a11b7c21fbb60ed3a479cd3b46660949eeea1efdf419eb0bb8666c0f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `3d80aeb390ca29f3fd3c507c0eec60faaf8df26c7275ba80af5948416c31a83e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `a2999fbb3f19dc53d6a5ac0ed74a7edf5a7c08e6020900c8553039305a67aa80` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `6b7bb4af01e97e7f5d572a42e75be57b2f69e94165fff60db43005b5b50901fc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `cd720fa913bd1c0e2feebc0c8c64e95c94da71ae8fc3aabf064252609e03ec15` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `98ba1b8a54e49ec914de8ebdc3af9cea0c3896555efb2f562f209e110e8272d2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `69c4055519745be50296b80dfb29b820e2e56771f99e37615b0be3b95c46419d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `c2da767e646fd5dc80075e74a26d9907e391bc0ee3612af79b8594802d8d0899` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `9eb4da42ec5f3ff55a94ff0bfc95e27a38997b7b64776961ccaca1af11a473b0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `00fa4176f6ede222c53ab64a376e6a94ad2dd2647951bc1904d5628785d962f4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `20e96b085f4391ba99d8b7edcc1e5ab4c997c4ba6424865af36a8c4e07238d89` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `78076f6ff46162573f55d7a311e2ad50f6e82abb3f96df295258eacda171ffd6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `66901ea6b9deadca79358add0e4385717c333cb45ad8598b59721365df973a25` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `323f97b5d9dcb87dfd41a0157d3de92c5587bb8267b0a442444824e511a90ee4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `0a0dc6cf49e733d7f6c85f78d14a99899b96fafce2fd925ff538590d17c9d3ac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `ba634fe7b13eee470e1a78755f05957cfd1d00922e76924c43ce958727db940e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `24e53b1c68995cd8cb74a16cb4f299681a13c29ef4f3bbf0337655060e494ddc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `9d19da09d75b948452ee78586a329d24d6fcc9c8c1c60afab670004761754877` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `1ff8229a309b9932bef39800efa0ff7dfe8e16df64d0290229b034d09109ac77` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `8f33473bb5b85484dc7588aea92f249b571b12a50fc3f0b26b219c7708392eb4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `ede4df2e0b42523910437a88506bed5793457eb3fc855c6355fc1d97bf8e0981` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `2a201ab315858e4d8654e40036e34af1b6ddda101bb8fc26652b243c47803231` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `c728cf259b520c77c0ba7cb19dfe5a3f606324575ac3957fefc7f20c304a3720` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `889ab9d97911772d880bc9655d8ace1488677257c7e1ed9323cf30f1bbb9a295` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `4a597c7f9f8f619cc1bd4a996927a42e984fc10bbd08fc6129358a388e28e5ed` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `0243fd6bc5b9649d087a89b728651f96f815f7f1a5b9447821a7a20641c4999e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `replace-pending-value-001` | `ready` | `1020b6b753b2276820467d9928b72037c910085b3312426a0674e60022602fdd` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `1cf257e801170409b6050e36178a217f750607bacd00ea0ddbce47f977a4d9eb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `ready` | `02b11c232b0b2eaaa59fec979040a7d51c19346ef7b10d6bb85a3b85e1c9d08f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `6843862ad97deeeb14c7d780df2fed7d2a9bb0693e5d910bec5a59efe3d86d89` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `1f6983b986c1ca6d8c66e005dd06d9f850370f8646bc5ff5b8ba861369032c9c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `2049e74db104df1c725c1e32fa946817b98fd38d6e4a30e8794d5217163f76bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `2028018a5a71ed2f7565cb400ed60591dc47767f236b821a611551e6cbb09806` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `8d8545efbbc7b7769d667ea238082fde3c36dc40279cc06c492f89311c890716` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `b4a58026d29d3091e92663f4dc04d3bfa40cbaaa1fb6a519c10e8a6b4f7ae88f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `f1db0da4f8c8e0d92c2f4694496dd9507e1fad47ca3cba87fa521b0235cda0ba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `f3478e97f832ed7f76e12df64420596cf65ea44f53bcd909ee1df854ebeb0856` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `4a2fed55dd3a2db590d5907359adc63b05353cf84e56289a86c44570eaee2c20` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `c56a892f1ffb3bb858fb0b0ead53914f091a57110672be6a9b7820cdce13f01d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `72272e2aeb1af3ff7a78340c93992eda0ca8b97c19b03eb0657eb22d3617b10d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `f1b218a2e336de1f523f9212ce42cb13709befd8ae9b502f8743837a532171cb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `5c202be8e27bd1e02c0eeb7fad9066532c4a002fb250984044e604e8b3f41f61` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `3092d73a1b8ec6db28a95e95501b2ecd90d800d6b847cd283f2d1c9d1d5983f2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `dfa02ab1fdbe8b106ff6ff14cf5b9971d50f1115892a7c887dfc33cf5df837b3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `11949667a729f1625dc0b3e6bc6859b4157f68558391ae9ba799a4db6b27dfbd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `094f1c2fb254c9f69dc5166c1f1446ed467ea3505c8476aacf8a052c3f2676a9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `de2fba9a64d4225a6efa18809a79ba50c3c242ff38587594e3a0ce5ca3ea873b` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `dc10d8f9ffe142b329814a2cc134a5a7198ad0a9188625ab8bcec961dbcf9df9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `4b2cdbb09992363dbd0d0bc5f5f7d9f1175f80fbbea1aa4b59e93dcb1ba6f09b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `0fb1bcf26ae8f4411328f7e7f311f67da58ec5b2a04e43444d827e0a75cb3bed` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `0130e1423d311eb5229ad592199d9a76d1672304ad5a3802991bcf1e4640b714` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `860e7f88d226674697e9e48af0580112ff52e2398fe8fb9d1faf99fd266d1314` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `44ba9b5987fe451ad1fa413bcc104a114c367a0adc6fb43beab6824b4b58a614` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `timeout-operation-001` | `ready` | `3dbf7f42be2271f971c0737a681ddff4be8dd6b506653150be2091873d9086d3` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `token-values-001` | `ready` | `4e179140ce2eaaab86a1de7b6e71513cb475e7a182fe8b816c21b2ff4080772d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `5b6dc321510fa50c2ab6103342872a996fc5a5aa0a62c88f0b82b310f23fbdba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `87d3a4c8909fa3f00a8005f69aeed69318063ed4407ebd74cf1254c26b685207` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `9658dc37a7afb82bc09b5d7b69eb9a2df4f63cebc5d2295031a9de0c3068986e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `1ec2b2520002970e12ee602e55c5f05192a04f65ad1ee7c13964fc70233c3268` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `6a52f5649a94d913a7b81ce2b1af7679fb61acdd3637d43757387c5468d42b7a` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `a2a3d73f2759fa455e5376fc2451207153323e316a67010807a774470c030a40` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `19913fdf7c865260c8483fd704196946bee9c11560edfdf4ef7a708924e3c125` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `7170645ea36cbe8e3215b075842724b79a209ef3a47743bcfa27b7d7a6fa3fa6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `e8952b2aa3d70d2f290b7283bb34ae0fa0d6afea72532db491a7d0ecafb1696f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `2694eea20e413251358361fab45c8ccc16ac7123db06f2d31ef36051bcef9fa4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `1b9c5fa4ee43ada3fe9496accaf548ecaa749b13e8bc0d2187df4c2660e73fa6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `9a7136665a3c3c8fad1c12a15d2fa816bd4c4a1e092406219d7598e0492f011c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `62fef5d730fc593ef855215d78078ed0f8a95b23d60e876b2f8f2c0aedc5c114` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `f86650a824b83ffc5a91679f3991ee5ca7ed5c3339fc2cf4a93b4698b653bf39` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `bceb2af61b09374af77fc90bd0f8480cf3787ab633392eb24759b63aac3575e8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `0b02162dc147690ea1f0dd15fab15a704ca809925ed1989a15c187bf5dd64720` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `ba47b22233cd88ac6ff1ac6a7299342b9b851ce8cd22e6310238e659ccf58af2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `77def9c4ef194caeded845e39f6c1d7875a8da538fdb0aba7bb04fa99cb7ede8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `32fe89aeaec0951fc19c8e6a8364fe772cf00e7beaee411577691ed1c12888e6` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `8a4d226c83168190b7003eec56e3833e71df9ddb946c00d36085f36ab6ba7730` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `20fbdd4121648b429cb6526012d320f4cc7e5add2b8c1ba555e6051982d65a5c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `dbb007b0bf9af756a190fd2e9db14e0146411a62240ff3012facc46a6e8484fe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `06913fc665474f8b5ffa9fd371d656f201b2e9b11d5d9135cf3c3e66697d4b9e` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
