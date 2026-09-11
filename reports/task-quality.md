# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v3`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 8 |
| `in-review` | 0 |
| `ready` | 153 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 38 |
| `implement` | 80 |
| `refactor` | 3 |
| `test-writing` | 41 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `10cb7baf0a4742d1767820356c0b1240983feeeec3afe0acd644459b119fc344` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `1756318f1755471efb1413a027684b5176b5c71aa5fb32476414e35a409bbfca` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `bb187728f7c3b93b2f382ebf0924efb019be0ca83f14046210fd2522e4a48484` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `fd1e63e7704120dc7e597b95153307b767bb68fc58ee3557d6d7b3e7dabbf628` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `c5bb95c3aa3b5893fb59dd174e099753ed806231594bd3579edc9216e64f9b46` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `5544c190e5c56baabfe5806fd8658ef88f7211fd348a91f09f6af72c9d4842b3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `ef314525f1a9f48e5894224bc997a223d6e0417d112a9bb03ee283f18a923637` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `d49066857c9e4736f219f19007024b4c306f14bb5d06df173118e614272e505b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `73badaf7719740ec9815436f32c546f9f5e1c7f8089b964c26f26c2957a3f926` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `5804e5915445ade7e27a7cc77f131fc3894e5ec97733235bcaa300daa739babe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `fc1e5192041ffeaddce9aced13b99d5625004ee306834c1374e605baf3608a4f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `f76152316453f466d5d22103d75a66453c2d1e8a51787700e4f0015e424dc46b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `96cbef54ad6a57e874babf497af4f7806241d156d489bc39f7554781bd7913a9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `68f9d25c93109b1f0bf2f182a12204a6fcebc75c2b6133eff9e60e77e271580a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `8190137f0df52485c2097a6972b06f2e7db13d278930aa06942c696bbd1c2a49` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `capture-panic-001` | `ready` | `cfd8d03403c22169ef0d5147b37fe9a10d103ad6f8d34e6a0628728dff330598` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `category-closure-tests-001` | `ready` | `c9a94c1f72836262eba21a20ea3f1b97753b238ecec7be35fb2cce86bf882fe2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `1fb686096928d0796ed71aac35978dce2dbf02a756add372a3f070701b5063eb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `35c0c0218878d5866ff0a967b43b0e042296c3e77c027e7c646c8b060eae9439` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `83d584f4c5699daa600d707cad75c11db49c5b8b331c9d43dfd0293db1ed16fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `1530e92d90f10f312018d4a1d5a4ea32f9ba464136d096bc82da38a8f4ce1144` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `742b9f4d614b2fb15d2c319bc2a891ac20171fb08fbbe50d687022dc7e1b7201` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `5facd08afe168aa8e12a17ab550c0a8cea60006d111af9559d5f0680a6401804` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `3b3cc4a4b0903a65005bc24d92a8f315029002564d90f7976aa87ce7ac0ae28a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `constant-kinds-001` | `ready` | `d493222e9d601d5554cfd7b78a98d8d1a63155182b35a2652b612480bdf05164` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `2b783e84a6628408aca4b512fff1af3cc32dbdcf3c028c0a28f6f9a07987b2e3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `c4c9331932819b1ab6e012a90e118240d6a49315c3a0eda4b593d16877ee5a0e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `5581da85a0cabce3f07a32af7b8e8e813dbd90903ddc49aed9d9b4d3f6dc2296` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `619ab5e1255b3093615fa01748b7d5990e50048feec737a61ab376aa5afb8561` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `0b0021440f805d09b294dfdc04ea6c7da31741c495b66b515fcdef3d4fd95a9f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `4e0e263998b5f53e5548283e81f02cc78bd75ac79accf28e47ef14bd065429bf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `deferred-release-001` | `ready` | `6a632270dde2ba90f3ea63f3d3a0b3819d40ed84fba6835d46ce4c0bb29c1d69` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `65c6b49b54ae70393f8fe2c6e0c7b392bfe6c5a2cd7de65adfae4d38ed2852fd` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `60b2fa695f3acda368541e26541eb18ec32bd918a61906f8eddfee9ea2dbde51` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `5110917a9f6fb621e65cf8d882192fb7120eea9e0bbf23b6c52edefd8f1e142b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `f0b842099e2d75678b3033d0745415a2b84130494a5f341bdd410959294db6f8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `48cd19064d974bd7bb13ccb50e3f87eb8a113524f18c086b319a6f349c508678` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `c552ab8165bb2e787fa718fe36c151e1dfb36d4c45c453cea2a039fff86ca512` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `da30dd28983254d2ba47a9efb1d27c8caa3ae00d21fac6f9e3c655cd2cb3c91f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `76bb7e17344a577092d25f5b5f5ffd067aa9b26fdeecdf1cd709055d056b2214` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `ea70c792546038c279866d33efa5a95c486f0d80b900529297ba99e1c71cd982` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `39e6aba6327853656fb54635987bbff8b6b490e364614dca887c5cab2bd8d4ba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `91d910b9ce6a286ef94e6e9af7ed959f57462a0e35bdbc9f0902368eca6f007d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `f449b098caf13e6af84ef86d8bb5b5245c24ac67e9306b9eed226b12ddf56631` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `d1f4f2006392369e6e1fcd491f532a7d7a4d3aca8ff8aee7cc5230298eae8966` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `c0f1ca841b9ef8218ec28abe689135e63a1a4f34b4c12617004f8a402da0d786` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `f84a3fb4c1fb506eb95fdf0bf40fe0097c9bca42e6775ea807744ed8987a07a7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `f71dafddfeb6652454a88975ac048f9b03227728dd3786f5c3c4baa22ec53845` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `8e71b07ca8ed07754a34e743db49e2090d04c44303472e1136e8cceaa2a54ebb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `31996d32a5e4f8c65b16379976100f68694c3f4b8c101907f3e8fdbb5b74c3c6` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `24f67b11584da56b106ea8808eb6b222238d8d832ac3439dd209ae572b56067b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `d4430a2f1575c184e67496f666d6d92e7445f5a3ae50aca975f6fc5005bcd1ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `18b63fada6318588327bcc0b976e2a904d52a0bde35ea2ef6464539ef2061033` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `1871ecdd2679320fc50e43ccd9031d0228d590acc71ba6dcc5fe817a8bbd27ed` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `b92f6ceb2eb78e695c813fc37d6e6352ac77ee0277615744c80e74e37760efa0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `04ad7347aaf6970cf85951879d5d0d55d1f5b5d8b9ae9c5be5f2972d9872b75d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `6a1992ce4734ee19827ca78e8f642869e9927f6dd21c551f42207aaecd43c76c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `7b907fec682edac371507c4f833db980940a6919df619008baa2b7f14df0d104` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `9dcfd855f2451a40b5be5df90b333caa95df7283e9314a61d5043a64b57141ef` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `draft` | `23fb0795d48a895d24f8ab4439992039dd37570044f7bb79b21b2c1fa18c68ca` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `fractional-band-001` | `ready` | `83de2a48b7486a1c1b45aa9cc5e8e8b7edde295e7c14a660afe7686f4922f13b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-find-tests-001` | `ready` | `c283101e76678f2dee0fdbf2442a068389c0344fa2135f51d2acf9eeb7f05c7e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `4bb7320a902e518c5df9895814141d22477b417b5c4454e90e545ce79d8b2838` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `fac8a22ce9b90083206af1ee0d8ae6e67bce5dd5a516c5fc8820fd388492a6a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `1851bda589d5e4081fcb161cec44e9a0fbe06e9f8e34a15499dc0ca174279729` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `e0a93b13ae9069454c986cfddb6ddc0f52f8a468f71e5aaef20c510780bf9a20` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `01007a02db35aa3c43a65e8a338b68cd88137cc24e25ecd8a77af7966584f250` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `fc4e2e8bf089bacd3114471a5ebe5f6a42d2b68560bf040e6fa444bffea30805` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `8149e4c82aed3f38d6b617741760c67304348dec428785930d04f72b866a0149` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `initial-structs-001` | `ready` | `4c01e5bfd423119bb5c75844457659738c62ba71c0b78334134bab56028c52ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `38bb1194cf4d9452f8b43b64c112c12a5a93597d2cb09e114bee5d5a8e438ddd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `af63d2eac50c75c47ef048ff897ad600e79d12a1afd5eb2730a7235747857bd7` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `37653699f3387d7f1fe9a026cd916d1506dc4082d7adcb7a03636669bb05d564` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `b186e8ef1506001899a530fc67ac802a37d1863a78114836001a5702b65218c9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `10cd5dc3128b5d14957a3b71bbcf2c466fdc912beb9325961fc2341d33b42937` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `5431bf3ea98759cce1bff1d70fb556b82320cd18c6a52daa5d30479a3d57a12f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `8c691614d2dbcc20f522c694ac4478449151948f1bc063a9bd0ef3d44a42fba0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `ff1589c2e2f79147009de2156f59276a95bddac67f410c9f7ec2daa0a85abac0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `d95c6f9a7b759f332dc713693de1579364a27369965ec3056b1c123b12f98452` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `a7fe14ec6fcbd6158f3566156bf22e343492d3db445eb6d8fd0862ba1f8da386` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `month-enumeration-001` | `ready` | `3a524c09622ed51a3cbf283eeb7afdeaeb665ecc71551fbfe3528308dda1123a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `35657ba38ab82e626c812c35ed2f018f7b972609d622b549e2450bc2e53b0fdb` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `adc4b89bb70fe2d8305994093a4d9643a9e106d5a3f6a5a02e538915cb2d2ce0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `145e84df66d3de4ae4e033ebe2fe5a89a56d9888c5e726dfd0b095be3f314100` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `6494d2cc8e8ee1d074b58fdb321799d4f7770d0dab4a5074eb01369836f380dc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `49ef5d71c11153a019c679fff51feb62b3a6e8feea0bab3631dd8e7ee63661c5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-slot-count-001` | `ready` | `719884b9bef62836679b325ac6ce963603733055e41e66f046310971f7a17b81` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `5145eb8c8f1e07ae4e52c1525904784a8c209c0b561ba6e785bf5fac1cc10653` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `5947584932c01da5989fa1cf51e9c8ed9ab3a014b84b5de44f554a5524509224` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `8815f9ae0358da0598e83fc991eaaac3d96948701959202a9930ba1d6ee26c39` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `5b9c1ca9f30250954c1f222fb9d67bd3824b700862da445f1b7a8896855912d4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `0764c57bf15244528f49210289a86a493ec5a6075ffd36e8ec638a94ecc369bc` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `84290730a845cb6d4e30849b66521bea16785d5e5b119d0e30c16461f3873f0d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `f2b8ac157703e67f942f479ca68fbc8100cb26a1322504823d3e7e033ddada3f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `a1edbc5b224525f3058a605803258db16393fb5a5b2e530353c612a982b7a467` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `f7c17053d1637d7aa427f17cc784c8fb3374d678be3b2c2fe276acac11812379` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `7255d89bcfb80448e40c6493799b76ef0a1312fb788dcfbc188ead315cbefabd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `f5d47843a1b95774fbcab66c88dea0f1e841e7dc1e3722db025c1e022a79b96b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `2840a24bda791bf26064bc5cb6806890cbe7fe01b0ccc30898732557c4faaeec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `693521070a6b57ad649cf1ef21603fe40d47f86bebb30a1459d8637597eb95a2` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quantity-presence-001` | `ready` | `4d5d57729c287109710c2165cc19ebedb3e8cc746b9623bdb72a65fe8c680fea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `b2029e5184d7fae1289064ee139dac7272430f04d251cfd3a183f7017e0718a0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `79aa23fcce82bd47a76e95a71d871e114e5709d817433a16596353f9671335b4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `c0c89b0a7d3303cd379fcca4706c0efa27fe48ed7793526fce88f14606434898` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `27b942cd8691a82ad3f6d890329ab09eb7b92cd3c2213fec9a894eb0d3baae31` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `8d8a9d34d3c0ba9b811fca9e0c522e7a193400512d46b56d648d2f19f7e421cc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `57f31479b5ac024b56843117ec0253111b6a008d3554fd88ebd681a8b518260d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `74ba6daa1d60207847767c2a6c02810e0a139ad1b59b5b98756644947561ce9a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `55fe6db077380142f80ee0d806f924e075260856a1ce3b752ccfe770919aa68c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `f0ea4481828a7c558a30b67a8420667f4c86e6e54cb92e7f93b99d70f2da0051` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `3104208910af2f8f1aae1f3a51bec22bf84fab70c067f9b9ff6d63c6c6f55503` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `f7c54578798b2afa01b47b0e191f474761c3f132e56854cebac9316974a46166` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `19c8a41f97ff2c153e59015dba0d164a804cbe1fab961ce77bbee00c6e6c7586` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `draft` | `00d5ebd5b98c3fc040a5e6533c2b35d7433f94f02cab12e42de96875cd6f902b` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `4166c0efc1a5ec54359f181ad22236bc849a809b94a7f33f5e7764fa06629919` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `9e742282f8774467d87638b81fc5fc8450a24179aee55d6d64fd21a2d2d5e32f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `f727aacd42a6737b5734e8566073f876d84ca710a52f7698ada89d1854151668` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `e87df66e529c41d4f31661ec812f636fecee758fbd7a97ca31db6202917e3a99` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `49b17b483b6a3dbdc52b5bcb7b2ea55727161f4cba03a5be085290511ef9c17c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `4ce31716ebc46a32601438a462886a12c1476b889a2597471784fd506b582c6a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `5bf0e9383fa601580c4ec7502b8b8e639c4d6112dda656cabd75c4ffa199618a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `7f11ce8ff084f90e21a96866b4572e458e066cf3c161e2027df891596fc8ba1f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `c0c02a6ccd419ebc1b1eda4baf28416118ee09223d3b298234ea433867b22e88` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `a6cf12feec97115b27c39c03a9a785c1cf461a75fe8feff17c228c66c0138272` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `a9859b0aa03f9c956fc2253b3a2223ec67f571a72297493ce52c1976e5592dc0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `14e1bc6e9b39640a58e3bf5c4d1ea71224a3039c75ca841ad317e9d1bb27430e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `375552b485f99004dc2aa5f144d3aaa514a339b82d3ff952dec755aedb11b8bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `216811f5a662c27a1e44a13ca618937f1f41d89ab784f5acb98b6d2374d98b57` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `ccdcae2ff334a01a3c1441cf3fc02b1aab2e9f7b9577ac571ce7971759c9ae98` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `604c35b23bb2a01d4f747892923ee7aa84642860129828e937d47f2976a02742` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `c916d863c94ffcd0b55f6e11b9c5b1ae39189d45c492d24b98a64f2799eb337a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `194f02f4395deeaccaf82441112f9778c5bb23399fbdbd7793f48d1e60d4a4f7` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `584aa0b1884d2c15d6a9f80809794f8c2023d8f31d0a77d8f805ee3bf0fc46cb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `bf5dcd3acfde1cac07c94a763ea8e2bc448cb1629afd594aefe13ee209216376` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `4df0ba6f604230d49f7bed9319d50eb1812119bca77a9a38b2c4f2e19664f451` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `9f1920b84b57f3ea9f1612de8a698d9d9c729e6d2502c82d5ef03f1f2878a141` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `94f0e9610633f05eae731e360e2e9f816388f4d861d5b814fe1ba51fda213755` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `646facab0f5da49a235f3f491c51b4a495806fc49ee41db48eea18c12bf9ac56` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `token-values-001` | `ready` | `668c6557aa06f51aba4d8cfa4e707af32ca3f6fa89e8b3ede1db734056d47c05` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `193bbafdc0435e30177eb654f723d561fad5b020798742bb55f54c22a2d919ee` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `51f74d6848bdb1195a6009adce2e9673c445595e4ae40b1476ee815c579974e9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `20ddc459ee57ec2394726af35052409eb9ad9eacd1fe02ca6609d9b391ce9ee4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `34335c79d1ac0b6434b11c8369b1c2c439f7e53db81dd3021bfcfb2c6ff37140` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `547d61fe75c9d068a683bf2de1838e889cc85523302b0c1789c54f26e3ece0ef` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `3b9a34754c9f20a91066df59ebabef73b836f5ac589a9755fc38f2f1093f7dbc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `69fc4e876d44d869a233f0b675494801644125f59cddf6bed273a50df1498dbf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `49d7291e1acefe11d7fd0a00f9bafcbb6474d216f69c74c1427d67ea67379bcb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `c8baf7dfa6e3785bd4331ba1d6c2939622c79fc6ba3048d7f0a56076ea6ac7ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `47e7481ba2f6b47b3ba4bc0205961323f977b031ea62eb8a815338ed797a950e` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `b03a491f4755697fcd9ffdd459bab1797172b69ae382a67eff07b78d9bfed86a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `015f2efa0adc3afb1c3a61d6ad6f1e3af2f2288c40cedfc3f7220f1171168c25` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `396c41a4d183097260ff2a77d2d31876602364569571898c4368a99e43dea2f8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `49b75d72cd4e9a6e3bf48ca1eb079d6dda1354fac4a003eca0e50a4fc5573dd6` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `ffef027d34c6b3d63b2d821993619d398343bd6daecef0967aaab94a0300fc25` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `f8514e633602761f3e92a9b0678e9350dd67327dbe720abd63d63a0f4d4b3189` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `1ff1be74827eeb3d04d680af9fb5e1ddc12ff8ef5cb3ab910cf7c4530aaca9ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `b94a7e5bb0ddd932c0f56fc09ed0aa196847600ae21071ef15025a6aac236b36` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `90c3512c326073a95a5e53a6961c5ae35b74b1ccfe9e49ae67dbf1ca6e1dde8a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `5959044e32c75578ec0a205cd3cecf6091d0020af86c177b4390203f14ccd902` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `82f3d6ae92605310592ee6fa262510cb9de8a21a74257b46ed4c28f771b10535` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `38a7164c685b2995f78274886caacd44a37dfd82095f5fbc74e486898ad10198` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `aca636efd49a5988ed9a42685baa98557b744deed9983621ae05628b9ef0ce8a` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
