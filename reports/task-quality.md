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
| `access-logic-001` | `ready` | `abd8161b7d68b4b2b1443b0354ce30243787931015ffbf1d6f2e538add6d0258` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `7e3c9a2d826ae21161d5efc48abd3a707cebacefa281869967c38acad32b01dc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `cb992437408e6e04664cd920be9df00aabeea153fe444264e0772c8eac308a65` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `d2efa24a7eb50bc4f8e15d16a1e49bb5bcfc8739180a0c31409449c494d84abe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `2672cfdb714488aed2e5e7954a88f1acbf9411eea981441edb129508b7600a49` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `621d9af5e88c7686e52a0b0f507b4f7742a961b9951b3419cd694695b3f27c9b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `3f65759660ac6c9969ecac4375f83fcf7f53c32a97c991fa38fc941b85348b2e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `8b80f47bc08f4cf94074d38b99fd6163c71a61cb3dfd6daa285b6e3f8a26ac6d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `32d9fd33a772e4ab61ee39c70a6b1258d04d1b970a4b2fa696b19c7aaa3747e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `71f8d0819f0ce379134771db98b5f4416d58eb39e658606318e6c8c0d97a1326` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `4d9399c1cead615bf798e9615845cbaaa8f2696435f6f24df9186dfee708638c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `191abf75f911ac79af7e84999f55b45d36a03fb5e9486dbc964bd9bff6a9e183` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `55e5b5211dac17208332f5a3ea014b80e357938a1c81ec2f7677b4f59a41e0f3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `1417b3a2d60bcc8c9aa80a7033628bf9d9e2da252bf4be2cf3d666145c45d016` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `9ecf7b35244b783d61730ec6cd8550af21512b799e163360a17bd452b9423156` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `1647b0999a48ff74d2563339211509250243456c68a5cc58d2a23650a08ff018` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `cancelable-forward-001` | `ready` | `746fec096aaf8df44d4c628fab1c41c38dc8d057691ffed3161c5bc8b8ebd4ca` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `capture-panic-001` | `ready` | `d5b805035109193458c6e844ea81e98898b94f2b33c04f15ba6836382128f122` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `captured-method-values-001` | `ready` | `095bb17d19677a01ddc032d4541da732442c83f646519516a29e96aa4ef192be` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `category-closure-tests-001` | `ready` | `cb99c3350612556026f9c16b6502aa023f899ecd56ea5ba9b8d2984aee58ee3c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `e013366244f3afc1e12a1f9b641343a59403d6b0b2598d848e53305695ac30a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `389319b1927782cf2109e3088c640f1aa19c4cb95852e927bad346c8932af9ae` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `b40348c7099e86ceb009d83cccf0110aa1e54ce0ef612e8ae0d9deab4c0176ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `035209518fe53f2518b39f436ce203295ebc7920f3bb8124f05596970bfdaef6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `dd1693dafb4951dd433713faa04d785c0feee3fc13298dc7f3761346b3d39ab4` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `3e329e4d1dd6c778540d999293ddb902a92199375bc396f539459b03e7fb894c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `composite-key-table-001` | `ready` | `f1be9ab6b830385744f158c8e732c07ce6bf1695733e6c77d1544b592f569b42` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `ceae20379ac9f4d6c9a9198f9a0c09497fb4fbfb57712a4e31a89ab2481779d2` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `d7654be7cb6caef6af9c83718925bf1c4530e3a59702fd9924ab628d63209c3f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `3953829d7b05746b20e134e8cb98529df1c4d8cc18c41fef2bffe68ee5e983b6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `f3afec3c479145c61fbd5b7c1d761727e620f269b93a9394bd72021adf38b3c8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `54b41546aee9412f63cb5d0a46f16a6662c3a021b36608ef2eb06e3c8330145c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `f2598e4a7e921054744ad623842c38728a782e17adc234f246c03dbb2d232505` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `63556f4f000ec30cd83aa1203a2ef9db8c9824e06b603c9b0192ab95c30acd96` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `249525bd5b4a25ee6611e7fb5d9f94be6fc0dc3facf712d95b7f87b57da0e209` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `d1f713f982f00b2f7161d48122ed6c5481e8fa066e5c746e1217cefd29ee9911` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defer-trace-001` | `ready` | `95317289d48715a2433eedac9fb1ccb28c4035daa0cf4acb28518c5cf8b9fe5c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `deferred-release-001` | `ready` | `108f642d623d91b4d4c4b54804b52b05f3708c1314723eae073da69f52077fa6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `890f5d3d58579b9143440fa1438bd047d5d68a4cbbff9863eb2ed943c402cdc5` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `1bb8c30816f9129d22ba747c3a38106a20ed710384881d72a9bbbd532e795963` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `ce968cbab4ad6b50824b8d93b4e952c7eade1f7f4132b760f4673fd7046fb271` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `2ad9fe521d96cc5e7c55ee1858fdc89164b5ff419377dfda8478b8db6d39a492` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `5c7d42aebb567ce5917a32c9cfb5689a50b1d7c5dad48d14d5797af2ae95fca8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `9f6fdc570868a06cfe926a0f115cddbbcaf92be2ed1c7e75d8081399a3bcb8d4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `2f597f472e69335c12e0d6f1c4bc0fcc3726a4846487faf56e2f16df22da1763` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `f3e528263342e0b86e01d61f85e04da595bce3589888cd6119fe441b689c4b43` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `5874eaf7c8bac7c3e4b889acc29d644bff655b3f9733e117b3d57f1903608fe9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `228e7cb7bdb37771de3d618933ebce3f5cd358ddf47c5b7c6b9cb19854c7e05e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `ff0c6b9b0d6c397d889eabb3fa1e5d6ff679247b559238c0556c0d0bd8a93bae` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `55f9eee2b5c65b55815a44eec8d3b033113c396cbc9dc4ae9a060dfe878baf91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `c5ea3cbbae8cd99a41120c9165b2e0cdf6805fbdb26d342e511f50f944c4b303` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `08f36bf10c62b1d41563e6cf9c985ba3619b81cbe2166a45806b2d0f56408d78` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `4a9f416e4041e08b2f3687ca923023ba6bf61e1c416523b31616615566a29665` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `48d8d8f264a562592aaecd0bb134c989cf95ebfc3fc9804f879694729ea13698` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `6768ca93bc5849fc5c70372782519454cbee63e2db9028439bfcd6a213be2baf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `0737c90779605c79062f515e30ce2b6fdd3d97095eb0cd2a70c447d48d7559a9` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `4705b0f9c80763fedfe55c9d5059d090f0616c79aafc71eb0b5fdb236774dc95` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `b0923b298d2d8519dd793fd0249c155b1b518a27ece3e34833c151d26694aa6b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `6f9336470b4851017103b1d1a89c3599fc15ced36914c6afe9f5c3e3bcbf882e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `2bb51af6f4a7e480e34b373d7f8ea36645147392739d664b6a667287128a15e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `5a5008ba5bc229086802ad995328e64e32e641602e1abbb57ebe9604f44d9480` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `b8ab6e2824bc4f806f61bc8f8242ff9dcac764e59a82babebf89caa15b30c0c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `5e76090b4653605e6511d652c3b5c32d072b4b7fb0ec5fe0915f252aa367847f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `59d9811c43dbac87aa2388d8dd2dd4896f7657a330e09c7284f92acaecb608f9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `2f5d3558414fec0523c5c12bc0c55cea47b43254e5a13dfa98a6c2e1f5e07976` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `ready` | `77af78d796511b8d7cccba66b78ca2fb9533056270504cb0059a37c901ee0532` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `fractional-band-001` | `ready` | `475c62b49f55c430a81328e9e7114757af94b98e37612a902cd68f415022a640` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gated-worker-arguments-001` | `ready` | `cc57dc7502459fc5063970ec1d6bea4b81d4c25a7326b3dcc3218e4bede83d0f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `gcd-reduction-001` | `ready` | `9e33f978c7a044c70302c6be0759677fafe8ed0ae3d3fde5c58b9fce1f81eb5d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `0fe98a34286475539c06cdc80ca01e20c8255d1487acc654b11117f8862774eb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `1562e3a7b04be853a14c0a3a56e82e65064844b9d9f6aa3469a844af134894ff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `df1f6a9b2a13b2924898ec6016a55fb3aad13ff76e302e572275be31ef5f892a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `fa0ca9c224af1921e8e5b9662a0069d073e96deb2f046b9812d98648f3b0e9ca` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `414c5d8a07132baf19af5618ef164dbddf9beb343bc0cc1e74f44d0176cb8d21` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `ac8378c2506a0f771ce858bcfea41f302f526455a3f03b90e691b906d3dadcb8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `464f1c06e0f350a1e4e6b85cbd867e4dccf11a6de71d62ff5543892b8a882f24` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `4d2cced78fe0411aaedad0531d73f515f1a5067653b635590156253b2b93cbbb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `indexed-job-results-001` | `ready` | `9e8e8aa6706073337ed49a3cae521c6a37dba349a84042a1b104f99fb463d6b1` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `initial-structs-001` | `ready` | `21e1e9a17eace971d29cc59df3d95fd1396d1c188cc8b069a18b04a826462972` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `2407b5444e056ccc1a401ca59375307c07899ad680c31f9b7caef9bd8ee930a7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `8f6e0bba301b8aea294b32b74908b7bd3eb5dfb08cbff8a70ecb36614e498b5a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `6cd62eee295664e33ce8545129849cf0d395935cce80ed3fd15abbb9f424f7bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `c345b90c2f02f8f3dbe9a1fcc4e8ff5442f5d13ee83d414db01ade6596b3b2cf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `a2e305effdd38b57a972f9aa5529e46345256b5aae322dfc547e1411e4cb5e71` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `launch-provided-job-001` | `ready` | `88d5b631b306b27b8bad117d9d4fb3d9cfe42f81a5851abb4313069490a0a741` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `lazy-probe-001` | `ready` | `dba7c712f3abcb59d6fcf056ed0a24fdc54378ed55ebd16b8ca9a3a782295804` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `9f29426f85273270bcead6f14db5c7c7c11b6b4f327f30a9c39391c4508aef84` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `042ad5817f1cbc66854eb275ad3de5cbe5f50e171c6fdd14f25deb47cfdaedb8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `23455bdb6446ce2bcb133d26cf6148307751d89a954ecf1bd1c1fefcf56a204a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `3ff69e2dc36f89e96d726696f6f13419a4462d0b9f4fa7e5626a5e99f4c311f8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `cffb2e70bbf6db3ccb7c396cfab0472985b7ad11204c075b61c610f4914b43a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `mapped-records-001` | `ready` | `b745d45b1ed698ddbfd303e97432ce9e3848ab676d9fad1edd5bf77deaec89ae` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `merge-two-streams-001` | `ready` | `8196a80acb20814d42756820b500496872cc95ddf8688304696f1d8d9a948ed4` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `month-enumeration-001` | `ready` | `73301fef9cf47e002764efd917a3ce43f01d61e570d2490efa9fbecf057f1e8b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `e47d53f1b36684c6d55199c2b6ca60bd7dc614984a344221cc86dc1c5ad46234` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `511a59457547956c24e046769d3a258723cb9de2aba0a07154d262646628c7f0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `29fe43b2392960bbc5787f165644c731d19b5e116661115daf15ff8dc5b2381c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `15080f3a1e053cd41ac13e42b99a2f2b22f462bf1e459db7eb226f33903a0b8e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `12548e5527571c1f4e26a8b35da7725559cd1ea77ff648f9cce48772cda2b418` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `named-result-adjustment-001` | `ready` | `fd9a8e64f8fb852fedf5e799f2d2465b636a03ed64cdf1c6993eb9e26da582e0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `nested-slot-count-001` | `ready` | `89e5c19e519500b2deefb2616aebc96d49d77a19bb36fcbc95961bb19deeeaba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `583ccd71531a6d5b92a61d52e89ccebab7249559bc8b39e70e6ead920d853900` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `ccde6b81fdc96200d78e0bee592d098aa6b5cb936ef6bd7ebf1d0d3e4781752f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `456a1735c98b14fbe5f5dfe4546e073b1ccc41483d11e84d99a5521040b204fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `7130bd3d014bff064b56b4920ba883f49cb82736a407fe9b2b4e101b009e49cd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `fec28740b6d20f2cf0d0328123c072ba1b5422984c26a64df618baf824c42390` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `9b7e81fa0deb015ec4fed44ba4d8efb5b0cb024a59b205a3cca5e707a03a6c14` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `e8f21f86ad9b11457d94658c4d2b922b5ff25586159fd6f35c43408e48de0d1d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `76492e2d1f7afb107ab0c5ffe6562c5563a74729a804e164725233e50f87b6cb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `806365bab5b1b4da17396af8dcc60c27dad33efab4e3585529b9c5d88f5af77e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `8047e7a9703f73f8b09688337e5644305eedf4ba7d871346c22bd6a0a5a82a18` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `cbdc1e3cb44b16c21bc9478f317a182409a6768dd3afd3a577c0bc6a9b9ff290` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `cef349e66343da57426d99f2b0db09c0e7184e55bbf6900371bf87cafe2f25cd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `d68c59fb8cf2aa9241f5d5e851d3cb1bf955ceb714397296e4ab8e45a318d526` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `7231d7ceb8a74f292571f56fd1fc89209073c800d7f9d2e39b859b9daf1f79cf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `cf5f5ec6149bf38c2de9caaafc63efc5f198178aa1a3d75c8bd75d6d01de5055` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `2ce0a131de9bee2c7fba7e5b7d25890a58c251ff972bc391b2dcbd06f5bb0fa9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `0783445c28318a141ac5f8c5dd216858b62418ca7f97c3f0ff3d843f41ea5ee3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `90beab3466c629b61f6c904a2bb65892f85eb1b00c8127977038776986674ae9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `b5e235283ab43d5ab9f67ac26533b0b1a63efaa8997862455c3790e243c3f3e7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `f2e36e637c6959657b450f47c28fd8cb3d7dc3aa1602e895c12e8e81e37b1286` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `0d34f934b9c747205b0643de4066852dd5fee21f8d1dfa02e4b4e0a338917460` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `e87f0d8c75ed18c81c5d7ab557e2c328815d8ee9bc3993da92acca62121adaab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `1645c148da3303db10ba474f101149db486bd5f60bd150f15e8f7f0a9a111922` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `995e58116a6e2a3481ee62865b9b987dcbb093869a5dfec807e1d8c7527831b9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `364b435e620256eafc5af3b66568f11e766d599999e2aa728fe347e91501a94e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `088c5dedb39b9f0a0ecfe85b449b7688f76d4d55e7232419fd3287d95e10a4c1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `replace-pending-value-001` | `ready` | `11c86c01fb3fa0169d7feef7cf48dab294edc51f206b04cc56e6529039b4f3c3` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `5fad35769a87224f061a12690c14caf88f65f54f72c238db3cf91646788e41df` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `ready` | `0f0f17065b16ffde3bf6e5dbafbfc6cdfcbabd192963f280dca61b5d89743637` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `1df3837c28f382ccdef62ae326bf82e5a74b8ab927e8a77de8947ae84ae33ccd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `bb5be064a8d5129aa5bc259432e918aa180f2fce18353a2dce48748d26e329f7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `b21aaa173fd3f2f4d5bbc436393a2b14eb3d9807c1cd39ba9675687f41fa0386` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `1942a7e02865614aa00f9329da87098e4e753af1f92c1c15bc26669cb7d504f5` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `b46608d9f7d31c347a90a335f1738d6e23f2306255dd19a9f077e197eebe7f8c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `3da8b5a985aa502d8c742052876492a2b684732839a888e4833f184238c1c044` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `8cd674ac69be22aeff02270539db93e9ddf6465989cd9da132c1825d90e5ffa4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `3799f1d2e1014856aba4a00e5e5b79c340e74f8c7c2f92cd33eb013d6a0e7301` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `ac55ed3f85ac0ba5bdcacc50f2c1590a210913f231122a91d3a39f5361310006` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `e36524a4c8831b8e96205c33db0c85b2d702c50a1d163464c6e101ff759669b8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `b26bc4df14782d42a67c0c926bf6fd16fb7eb3fdff685d4535bcc3e852bc066a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `f526763f7aa1e4a5eb8d4661977f086523153dd5ebb39a2c00662932ac695ad3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `1f5002ccf9830cb223077e38dbccc249b68ecc61f8d434cf5737d0820c0b287d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `e9e89ff24e2da6c7ef6ae488ffde684a83d2d86dc526804092712f834661768d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `35d0ea1443c8fa8ce6d5d2ddfd777dd7b80dc5bc4781e185319e436b76ed53df` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `9d2162c61b7dbe98b15b706b9bb09f144c7444bbcdaba2551ffca502ecd105ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `fa2105574b317b2e43097e405c1561dde0e4c3427679b9f199517e82e3c3ad31` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `c74ff3d56a58d99e55f144110fa1acf71b544fc0f2233ff0261387ddb4339023` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `ac4cbc36e9382376c9336a34e34544dae7eee1556e62f46bf40bcb0a24bdefe5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `426f8b2548f1eeb870699337281c6762c22cd4155adaada098f40fa77a7b6142` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `1f0fc11cd54a68cf5a6fc1d3cf9bd1b4dce0400fda9c4b3c9da60a46c6c7328e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `bc81add5e4b213cc0888108daa28f94cb91cf21265af2a3083934ba0865eac21` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `accbb0cb5fb414b7a8cfe833c92699dabfa38cbc0577e0d69c4accb5e7c454cd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `aef268867e0217e72aab6c43731970f74a678d95d5a04014ba5a1959947dbfce` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `timeout-operation-001` | `ready` | `95561ce884e99df6afee13198c8acd407a1ccbb914fc0abcf83d1e07606ecb30` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `token-values-001` | `ready` | `8bcc682b3c9ecce91cd1bbf6dd230254d0263f963cc02900074dda76e9e6b22c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `67287e570623e20bf6171d5a5e5b3cdd2bb1fbd5e87cad0041d2013ae3a7474d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `b7c20628e8f953e1d3dfed765c9135c3b5d37f734f5a794827281f3071314883` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `a1503f9a6d10ce676bac39fd1c7f7c7a833b112b94ae8c9a0148adcfbeada254` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `d566e0c0f4f14be5be2fae3304c8a7e63a09dc3dcb3985560513c81456a15c85` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `6cc9c4ec5ab10c2240c851afb1a36307125df07e4b8ec7d0cfa6b856d2814bd7` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `5039be02651a308aa7f9e1c2133cbf5408ce97e209356bdb4789340d92a93ca3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `4fc2cbd8440f4a6d2c3e3a303d979d6ecd7d05eb6b39ff62cf4449e5b68c24a4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `9c954f1ff2c1e639ec05baf6c9513ada10ace1af5fab07b0f7a3cc87163b1b71` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `c80fd706ea5d0aa28364cc957375e1074d023b9d2ffca2202f712436bdbe8267` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `cf6e2ab1ce8d6b784384998b75ccbe1448a5248a73d637b545b1460a8e047ba3` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `c04d04931e98112a0e7e509004c191cd02df3ee78b5a3b174940eb0bd2396df1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `132459fa5bc25020c6e0760088dabd49af4840b72753a5aee86fe638b905a9b5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `2da6b1f209e3b7573c411f46e9801a566ac03ae08e8246915c3b411f79157333` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `d60c7fdf2e4d76df1ba39c34148a1604a8fb13946804e978a3c61fdd73b70a3f` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `d48d5cf7222a3b41359c4b31d1858e8f7482198990e9fde3a4fb472a39dc8bf9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `958a74deae2045ac36f2d0713a0f03fe1604942f903ed1d714fd0e0a2e161263` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `f05add16292f28c9e25bacef6502622d815e209984ead8e6592096daae86c4b1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `04a470631b421a2864a4091632e34ac711a05fca1ff63bc81d519d390301792e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `83395dc4e14f828d3f3080eae259c2251be4a574779fa7b2447c086ff3789222` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `b0613ef07d55fd654f58fcf6bdbaf0cc78a91fb439b02c0720d896d015b7a5e4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `0dd28d98af0959f2c6785eb4b0a28ce4aaa507ca4f5726f40ed8f59c69b94e2c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `1e0030aa381d2e04e3f8d4e8295bb428821eda7b8d83f678d00006d13fe8b919` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `cf9b305ca88d2bbe62b84ed7635583f3cb44b8df427f6dee0416ffe9a59ad5ef` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
