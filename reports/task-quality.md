# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v4`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 6 |
| `in-review` | 0 |
| `ready` | 229 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 51 |
| `implement` | 131 |
| `refactor` | 3 |
| `test-writing` | 51 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `bdfb23b33098effe6da2e5a5b7119c434d7b9ef7a07c17890adedb97dc7d34a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `fc34da604f294a416f45811630b0c18243a55667909e7ad3fee9c30511b88a9a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `1507e77b4a737984d61adef39faf799c81cfb0c22a889a1794e3df63090172e1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `1edf16d4857192ab01f156e4e1adefad271638f9a6e1c560438ff8770da2a0dd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `ba720ffec31b0f0da4032efdc4415647ade78da2b5257f8a7a03eda80b3f4fb0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `fe99ff79db04df403af89a2bc83cc5629a0955ab0aa9e65a5ec0441a1a9a3911` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `8b1fc2d308c5a7444924a07f9240ccf5043e8cf1e56a75f42cc9ea9b12c08c07` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `array-copy-stages-001` | `ready` | `4fef2bee403b4bb9ae3bfd95abba3cd8498c48d96a63088775dad99b07ae7f95` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `array-window-update-001` | `ready` | `770e30425fd5cf869270dc72dfb1f5a07d206cdc51591ec67c75225d43f29ce2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `atomic-count-tests-001` | `ready` | `eb39c3ff885c8e817eb71ae43e1ba9e18a257d82ad6fcd5fb2a5083849135798` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 13/13 | Phase 3A |
| `binary-record-tests-001` | `ready` | `2167596e0547a7c3fd55c892cbbc2b416aa828c797038e57110ac955c8342872` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `e2bbe29e07543bf1fa727149d3246d4d7d4d6c0b9f5963eed7e79a3dd7d2845e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bound-stringer-pair-001` | `ready` | `b9cb355bb2d166d45cd42dfab8e40c8b03d4cdc29ac14849a3a429529555a9e0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-copy-001` | `ready` | `9696ce824e0d628d19e4228027d11014e77cfdf8ba1e0785b447757aa41f738c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-job-pool-001` | `ready` | `0cfe6eea18d7e49c824a5b55a46865383813398d69b869deaf127ababab25998` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `8edd337d00cbd40abaeff6f37c5feb273f393775650c2a7db9ea9571218212a7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-record-read-001` | `ready` | `33dd4e1981a059845c03223358451f3905ab4546c8da330a8c6cb5021de09c21` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 22/22 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `b3a40f5a835ebc09e68b20f3f9d326257cdd5808cbc15ccd54d0b72f4482ab43` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `ca561b5655d0a065a8c9bbd6e0b79893076bc09403685581ed7743b6e012e8d2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `branch-paths-001` | `ready` | `c03a6c23cf71815414a3ebc9567f0215de2d595d3bda6afdfa81df64386a16ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `buffer-window-stream-001` | `ready` | `208ec9922c2d0c7fb7c291863e615b17a2db52803c63751f22b5681d2e566f7e` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 10/10 | Phase 3A |
| `buffered-flush-001` | `ready` | `7376fcaeaa25c38786be3291e8d46ef579855d0d725aa73e99736d8a356d9fad` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `buffered-word-stream-001` | `ready` | `8f2015963ca5d7f657395a3c872d422c986efe6bab594bd4d5a092bbfbd0cd7b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 18/18 | Phase 3A |
| `byte-window-tests-001` | `ready` | `aecb2b3def2d5501d993ce2b4f903f1adca109fca3bf9f779277986f0ff23f75` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-error-pipeline-001` | `ready` | `b4a7958276ce2f89616c8946fa8c491b0019d9e552096ace46429fc6e794e97c` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 18/18 | Phase 3A |
| `cancel-worker-001` | `ready` | `ef5cd4dad8eb4c8caa0a9151edd94c56e1d6c4e47c5f3bf74284243d64a6731d` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `cancelable-forward-001` | `ready` | `b1faa2490d037ae3f57d006db2270d959d9a094ece0b73cd6b11e82d898c3f5d` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `capture-panic-001` | `ready` | `7beea33c79a446d178419d6112f2fef01940a803025791198a2085ab7333b7a1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `captured-method-values-001` | `ready` | `d6743870954590bfe6eb3413f64c0de7fbba7fd08b78b732437e6e98cefa2565` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `category-closure-tests-001` | `ready` | `4b305afe17b0ab60f1d1e3aa1aab9b652563d13ec1dbc0e6e25dc5b48573e435` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `b468c25b7e37c5b5ebbacd11b152ca8dfa0c86192da075db0ed3f7a729c1beb7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `e5e0804bd5255ee422a1a8a0c0e51a1bc85e8c977b6cc1ff7a826a8f7f31f217` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `child-operation-scope-001` | `ready` | `a37213996c40ec0bd34eaf367bc3880a30c9f406a740461cc6d1379cb99af926` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 18/18 | Phase 3A |
| `clamp-tests-001` | `ready` | `e80f407f938832a6efcec84508fd9f117be44a0cdc627896525e2670775e0cdb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `classified-stage-errors-001` | `ready` | `3959eab8b7815a00179151a183ee9c9bffc7c6a09af452f14b1c3ff293f021d0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `cleanup-order-001` | `ready` | `80317901c7313b474fd24209bbb5365e91e65b9705b4d454d0a7d52a721a2f82` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `cd11197eeb405494556d75b56ba6362b5159dc542f561060a336a4cc5cf65195` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `37248c884e5f6272ed0b644637d59b2c3063b0677ec4e01b9ff8c8939a2b2ffe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `composite-key-table-001` | `ready` | `32b9836ba5c728a475c750ac4cac0c04be43ebc72e37a4850c5b8e92731d079e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `bc057d88939fb3f1e79f873dd74b9bebdfd77ebee4af0a81d253f637f0998ab0` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `3060e60707797dbe91fd47817df0fbf0946c9790d5a3055084f520a9b20aad61` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `3f314b87ff53acacd0b3bb48a342da4b415c7e3e7cdff91bcd2f948246518f96` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `28fa42949db8be630c68baed9be53153ef504f400566fc7ab157de2e403ff5f0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `consume-slice-prefix-001` | `ready` | `1c7e683be2d2189a50d82a04b1d662ce7961af54ab4ab6402131dd9b993c0dfe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 10/10 | Phase 3A |
| `context-business-boundary-001` | `ready` | `580aae08b09392d5f4f5045efe43f63d4173f25bd8812d4f21bb6c6eb126fe5b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `copied-journal-handles-001` | `ready` | `aa60eea64202177b093ad4887a6345ef847e9d6bd36df4b4d4f16414b10a3571` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `copy-final-element-001` | `ready` | `abed54a288dad593fbadedef334b9291c341937c8fb8460da07d52a14c68065c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `ba9d8a400cd0b7c32b937d03798a30a7644b375656d9300caab944874bfdf108` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `8979949ce98c00fef08cdc9f4525aa905d95d0ded5db6a95b17fbcda81705935` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `deadline-capped-batch-001` | `ready` | `0f11e6c78fe46c56baa44dba3742cb9429accc15a762fa66d3ea599514af9092` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 22/22 | Phase 3A |
| `decision-table-tests-001` | `ready` | `e3380f105e50b7c0fdeb9b2dbc55d5120600ca2bac3f482a3b4fa9488806bb2e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `16aab50851a48d052e8955e5a9bed8507b6ed43633ac64784edb2388607050a4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defer-trace-001` | `ready` | `61619cab4b03dfcbec228188b8e0e2a0995f4b405c51c39dd80aae76c8e1b8ad` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `deferred-release-001` | `ready` | `f07fdb3ddc11cc33d9b4580afb15fb08c5c6c526f388f2daf5d0abe31d7d100f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `55b2783cd04e088b1701b6f61852aea1c05db2e9cb49b91feb53dd340805aec1` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `47bec0b345f8f4b2724078545a4dd0563f551c14ac64734a58a02123815457c8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `c0a2d1dee1b0d9db280b6ecd3d3cd081457cb8b88cdd1d6a8b551b58de6291a6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `bc3060ee0be68ce79867fd11e86270f28799f6f537d737996b153bb5ae9e8d11` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `77bb74b0e3e3106ac5d68ca5b5e4059adcf315b43c02445039bdbe0b6a93c6c4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `b1b853ee2fb571e49f73b89253559ca8171fa812fc98b508f4fb72c529a834f7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `9af110220ca175bbf6704146d5928e26379a7493addb398adf6db42bd7a452b2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `43d595126a7fb999bb67d3aac29ec540555323076993e3e7908aa140ab71bf3b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `2566d2e4f1754d3829f4a960e911148a2bf6b08499ee45ea35cb33448df75b5c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `drain-values-001` | `ready` | `18e37dfbfdc96c455b4f09c0750beb26f826e2b0789e4bb17937494f94b222f0` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `dual-frame-codec-001` | `ready` | `3e7f7e6d4712ce594ae1fc7f208da1842f680d0e5313c88e0ea023c6f907779d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 21/21 | Phase 3A |
| `duration-rounding-001` | `ready` | `666d6f7eb3e42f66ff91b58e4fc3ed006ce2fc397d5e0dc453b5139da92b66ad` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `2d89d4c422f5cf8c8ca1eb4dbee85bb9661ae152255a36c955aebc980b2c58a0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `addfa6aee9474d7ec7b65d9c81cf8aae7c98438a0f6238ab71df72a3f0ed3183` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `642be720f59dcf5e54b0db05193d43cba6b4f9f8b4e894f5458882dd9fd781fe` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `9dd364e373d448e48ceaf3428456071a8a93c88f24779f44d73a22594f12dfc6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `embedded-field-update-001` | `ready` | `aa2c89406645b689eb7251de8ab56c9bfc638a9693519d3dc5f54ec04b094feb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `e7c0106eb488c275c589285741b5ef1f6749683cbd1ab2829cd914db6c988b53` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `a121c9b56f2b2007b2a28f383f911b8949df1630caf6782cc1866fc25535e62b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `86f0f25bd7546990cde780005fb8e60bde6b1d2fe11d33eb96e685fcd8b9bebc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `83768573b7f21e559b6f95ae2b5189e8ef9374b42af0d71cba4e185adab60998` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `99c976058bf471d4161d59f5b1337fdf703b7728ac9ad426587e7c59478d842b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `f2a3caf4c10d9a6c5f4c5c99d111ee1c773921f42e8a7872ae58945939ec5a94` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `69bf0bd7958a31acb8a4481402588bc89ccdcaa0c01bfed908e4a2e650a85c7a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `a2586951de7f8b99bf91916349bee5bf7eb6150b2d3c70ba66edb99ad9d959dd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `a17bc8f6a43b7da877d706e6fb2231769233160a6201fab6346f97c0f1c10e92` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `c0c0913cda81b51d3ca4c3395e09ff4913a2b751a76fa3785f986399c05a2c14` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `a9aee50f14ec8108dfd076bd04969a60f4c81969e1dac98fe1cb85a22989e5f7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `first-pair-positions-001` | `ready` | `0140ea8fed650c1dff647a6b3b55ea61c3689974254c626e7077f8df2c9d9101` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `fixed-width-shifts-001` | `ready` | `33135585e845af12fd8cf8112a2a75a09cdc6286eca5dec950f7a1589f11a663` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `floating-categories-001` | `ready` | `813971e105b80967eed278948f98e8b950334774938576759008cef17d11b053` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `5eb932d033cf7e0180ca42794e2a64ea681f18b44903b1f12896c219d046a3c8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `ready` | `ed66e9321665e21e4d6c2b429a08c5918e4ca4ffff707ce43c20ed09330886f3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `fractional-band-001` | `ready` | `0dde7ee0591355306575bf70d45fecf8d5bd03cca269dce7de0f64e9ecbe8349` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gated-worker-arguments-001` | `ready` | `aeef062948ec37a7913e11dd1f68a9d95ebc34752e4a956c97ee703bf7c8c424` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `gcd-reduction-001` | `ready` | `803289f9e875aeebd9a9d748eacaf5e1c5ea93abf33ebf080360a4f49b8a8905` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `ba675db33f66a5a91b6a28d46804db2eddaebf4c71924dbcd788e43adf21dda9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `d7e6e75bab4e2d4c421e01c72c5766eb4dc9e8611b7631f5368526dd3bfec8e1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `d6c12661f1db57457c0b584a82f0317cc73004f2d830fcd71f9117838fb38818` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `9633491b928cbb87df7bfa33e13cbd65e2a06109071c620818e150fd7487e6be` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `d58ac77c6c046b6e1ce57f15bde50f04555469e171ed6da87255876d23c79fe8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `0567c7f0b63c4b1e9c0c7ecd8669ec14b921624646e3189677cc9de6473d2d15` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `guarded-stage-values-001` | `ready` | `6f5471daaec9b55da3760c112fd70a6f5c1b9309a2fb10e787a9b5ea1bd68550` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `hex-seed-classes-001` | `ready` | `cbb9e8f33948f0786b9082cfe7e54e049af942d497c782c2e319b899cae5ff8b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `hex-wire-invariants-001` | `ready` | `f6caeebe715bc724e5e0c2345115a5b7aaf49545de97d9bdde242a3ed1dc3c95` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `idle-stream-001` | `ready` | `d27612122040c54276ada81618364e3897359ca07742de2848eaee0f227ba67f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `independent-policies-001` | `ready` | `799b3c020ff6e014f19a38cea181329abe063110dacff910ba9b23ca6dcecfe9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `36c9f18f0e73b55e03fbafbf2f594778fea68fff07225b9dec7bc1f3b052e09f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `indexed-job-results-001` | `ready` | `e951557789f0923c8926a436b02987b10b78055abc682790bea1bd2451412f8a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `initial-structs-001` | `ready` | `aeb48770efdb7193dcb3a44606f139c5065fff215b6f295444d67748f04bdb62` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `a36570e370abf51710b5858024022a0e46416652d3b389241974b50178736896` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `interval-table-tests-001` | `ready` | `7d8e8694cc42f20115c11de98cd1ae8799ea04c4e2586f5bd75770c7d27cc211` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `iota-layout-tests-001` | `ready` | `237e023ab8b45a24aae6462459ec6849bf27be5d5230adfe942c036e1eb1ff2b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `join-all-tests-001` | `ready` | `c47810775cdecf958f0cdf28337cb126c1dd0d7c6e1800f9485c4d158dc88bfc` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `c81835d0c0154449feb9d2af3363f16e9ebddea14c3cafaac26229204570adf4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `adb50271b8214efbe7b0effdc8930b1bbe451e5fbd40e474ecf95bd82434aa92` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `json-record-patch-001` | `ready` | `927730df2ee1acfed3539eecc178615d4789136d6eef076382042a1f9b92b8fd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `key-summary-examples-001` | `ready` | `347ecc0252790b7057e96edf5e0f711cd75251f8157cb6a51599102599ba4f39` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `keyed-once-registry-001` | `ready` | `8419ab8ab65f0763d778cb884f837b4337060b19a677041c9647bbbcb12f977d` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `kind-condition-001` | `ready` | `aeb350713036b9a12848b92c0f1ca9da04efdaaef3e08bcef848d1cd55099063` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `launch-provided-job-001` | `ready` | `fd5d2aec13d0025167bed584a5a8ca88988b0f328b3862cb51968f605a7976a0` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `lazy-probe-001` | `ready` | `b4a9d8d15f804afdf616238bd4bd57e312b1af80a8ffe37668e0c79d01f535db` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `2cbe94a16e75422aca7030cfb61e9ac5e0f0a2c81fca8c55a03f2f1d6ced7ca8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `limited-pipe-relay-001` | `ready` | `1327f800417aff7501267831530caeff29552a04fb2f064d7dc371c4616c9a1d` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `literal-bindings-001` | `ready` | `c72c96b1703a7c7739372fd35f63859c52c66d9e3bad44ca16a628436a3487f2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `e0f65f2868f861786a401cb057219aa972e17821f895cd4d6a21f9e0f41fceed` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `e8844fc4404bc87af33edfc9f16fbdde5780e744fdfb792df9dcc28549c65728` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `48a5cac02099d3e5f20e566c19ab1e4bcf14ae2a9cbc33f11b993efb0e3c10a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `mapped-records-001` | `ready` | `647d8b8a739841ff81bea349055f437652df7035ecf206ed65bb54b5c7d75ff5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `mask-state-tests-001` | `ready` | `f137802217036ba5cb67996891e3e0209c9a4dde3dd46d5b420c0155702f5180` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 13/13 | Phase 3A |
| `merge-two-streams-001` | `ready` | `df5aee0a06b90cea32604789db831a1c52347d43a8b0cc877d8704e808ace31b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `month-enumeration-001` | `ready` | `bd9e80a69a6f20b0c5be1944446391f766f317e5edd983200fdd7ebe327726fe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `move-overlap-range-001` | `ready` | `aaa395101eba62b57f6ae38a771ccf1aabbf29ab0c4853d9bfc56f97d59af2eb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `mutex-counter-001` | `ready` | `18f8c5e6b5d2e588d73205275b4268de54a204ec70f2be86dbe43221f7e710ff` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `80026a41a8cf7049d6d726a6adca777163791f9b3644f4ba416a7b711fcf7504` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `09f16ac82d665cdc6e492fbaee685d69ac777b35c4bd9acb9fc30773a5d97233` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `1511bc6e37d12a986720ae744d4fedf27bf90c132cb17d25b94d564703382e47` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `fbb1dd251316579c0187b11fbbd82d74720253ed9f39eccfd5534961c23bd0b7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `named-result-adjustment-001` | `ready` | `7b32eec41212bbbf8dfae5ec924cc86b02cdbdd6ee102d5bf79f6cf84ce2cc04` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `named-step-conversion-001` | `ready` | `20c43898260ad138cec112ddc7b37452d4a4653e75d79b35e2102be68b59a1b1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `nested-slot-count-001` | `ready` | `88f0a28ec1d21a0536e349f8418c3bbc2d6a45468840a1f009f9df484c9fbe85` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `7519698cbf9fa84b4a07c1703b280c2c9c2bdd5b5ff4e3be6ede6dc23e803506` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `nil-preserving-filter-001` | `ready` | `5c93a0ad99b4814e97aabc8660ffe22907d98a0b6c1b1a9c21deb024bef81197` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `number-box-tests-001` | `ready` | `318bfc498e072fd384425b7af9bc95cb54fafe75616eb9a75321bccd3049ffc9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `once-result-001` | `ready` | `e10c7b12c5902376c18abda6792972edb53888e6ace157d9062d7208e5121156` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `7fad9bb59fe9ae8333b38d796440c7c35b18146081cf3e777425370351576590` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `ea96ca44951eb82e4270c0d2a62edfb5665d0866088e94e14a30aa39a987ff8c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-batch-channel-001` | `ready` | `fdaf577db6cdb7385e8dbea077eaec671c52752d675898651c52fd54777d0203` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 11/11 | Phase 3A |
| `owned-values-001` | `ready` | `dde8a2b9c33b3b54da8a5bf3bd0de610fb8c24f8778a74c7fc90a7bc027d32d9` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `paced-forward-001` | `ready` | `dd2c9c820d92b95bc632ada6b95e2bf6cc51a947702c88028057986b85ecb4c2` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `package-inventory-001` | `ready` | `3a08c4d8087f93af6eecb4a726755f044467efae477d60f379192a6579960196` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `f12f8ca31f37f0f27223a2e77f52f5bfe56da43958bd0164944678134995a676` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `ba4b6910e82063bede60d63df2bdb35ff57b8885986cfd71f8f629df3d86e8f7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `b5b6d32f6d2ffa3ab1d14b86d328b3c96d8c2a76a6100e8009fa0b53c1387e4a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `periodic-budget-001` | `ready` | `091622129a2991fc1e9016da74de95b815281c5068dafcc7871e3d7cc4eb05e8` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 14/14 | Phase 3A |
| `pipe-exchange-001` | `ready` | `203674488170379038119f1c96cf368c1344b3890cf85fc75211f69a544cc7f2` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `pointer-chain-001` | `ready` | `f8d604729660ffec304c7d1467ca5cd7b2ef24baf0e39a78dc8648523ac8cb54` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `3c0637a477cd7cb7ced3b0203edea0107c352186f459cfa2652f7ddcda09609f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `eaf47f37858604025b774388f9140766b6e37b9f16d358c558af531294cb07b4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `presence-concat-tests-001` | `ready` | `516c6711a9e9a34918b4827fa42b1effb80f25b6176cea9409df45d1d06183e9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `promoted-field-selection-001` | `ready` | `c21d91c7ed3d228e84fea204feb54b9903ce07d0dbd2da5fe5762e60e2f75b07` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `a9d43fe4e34c1a1c4e22f243965d2e52de827b360e96167531819e18be3eb449` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `1e2e9affe62871d0a0411102bfe2e39bc7688ba990345411fcdc9a8cf4b91e88` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `ad9ee30ebb14f7929163249103fde128418636738038abb7673e87ffa89cb76c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `667d6689ae86a24907faab2154e3ed8d514edbe0a02b6820d22a179d16b5fae4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `17d00dc21fa9af132e07f4c61728594f2b4033facd00e9ca6af9c730c64252ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `89f90f00fb6301ad48ac13e44dcd37ca4df92812d7b6a136295247b42aa8647b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `5d79013be08bd9051db28580b5f5925f7b8026dc066a9259f743555a2915df5e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `2bf8567a350b06be1c749ebd151bef3b1564783e4e5c31efb3188fe925665914` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `65dc76260b9d5b7444b113486d39cf17056f51a98f81d0f24dda7ebd5181b9a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `recover-worker-001` | `ready` | `3915516cd7ca8e5fe7c871054f92dc6cb097822171e6c20d4bc2937ac58d5aaa` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 12/12 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `7e67c75f6c13a5bf01ccfe63d9cf21991a51fd76aabd4ff004ac2fe7ece96210` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `recursive-total-001` | `ready` | `8ff48ba041554963def433993b8bc3e775ee5102efb28a470cbf2cd24b2cce25` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `remaining-constant-001` | `ready` | `99d25aca0d03057759e126f6790e7d83926b0191d208c75c7f7270d3303f4878` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `258b97df650a35f36569d1c2b80d29c46d346eef6168b8c9266a4180340f5f54` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `f924e8ccfc322be6d157c5b7614af10c1e94184609dfed734ef15f555cda205d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `fb51efd4c03ec61e7a981a437c626707ca664b3f66e7aee51608c6e6b42e9458` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `replace-mask-field-001` | `ready` | `1011070b9e375542bbf7b8e3b65de3562c9cc94146045456e63ca78aa3c32a59` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `replace-pending-value-001` | `ready` | `30c4a93dd07c17513586e31ffd1ed538888b72fef2a885737f9671a3746246d1` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `replace-rune-range-001` | `ready` | `40eba1eb8309faf7852f169a748e345b903a210a1384f2a1386b549f7c242e74` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 10/10 | Phase 3A |
| `reserved-copy-001` | `ready` | `68cb121962bfef0de82b36c38b631153d80d9eb99b67f2d6b884eaa5832b4958` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `reset-selected-state-001` | `ready` | `9cebf3d2bdf009977d2f23ecae23c6438c9a99c19f22655098806d2584a0e13a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `resource-finalize-001` | `ready` | `7b1e12833b1f234d2c34773b38d47ab9da19303c31b0e276e6e05e5eb659ba7e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `f7042dff44d622315742c8f0fc46e2451d835b7b6ae1fb45d7649b5611ae147f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `9cc5680beca06bea9a40872725e333d120bd9736930cde3725669083f74ef6c6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `ee16f955178098ca68b6b96ae0bbc4acd06793ba3df90f1c79806f9936996fa9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `140278506e1a69bb46f9b56252e826b85b37d4b1d2e52f20502f7f08d4578911` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `51393e5cac3c801a1eba1d8380e84990263a736026ac5b2272e7f0a399dac7ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `17bbd6056bf577b073f7974ed150bf161a5163ac9c25b285df2fe7459a6ada15` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `e69dc15cf9d3a66c6e8241cf7df73d22cb9a45b28328fde502c7292182768ebc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `scatter-snapshot-001` | `ready` | `3d04a73d85b8d8f78c92f08985638acced858539cef6df0bef93f7802e8d347e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `select-ready-case-001` | `ready` | `ddd95e0003112ceef12da8b03ed26731b63b120ab9c47813bfb8acca7dfb82c4` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `selected-window-copies-001` | `ready` | `641a0a7b35ee477659ede6ba16de5a46e88c28f34a8a31ca28650335bca09545` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `6922c0c757db8d8284d327f889b69976c3c4ccaf7b56e2d9ebfefa7c4db8ca09` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `a3408f828c8c217047dd2d434946928a7e9c8877fc17ad1c8cd0d2ccc6ff105e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `3df86924740ffb825e5e7b5804ac93b4d8ad4057edf849820110261a50126fd7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `615ab9aef1991a2a2db089411bef4c3bb1ca941646b41699f1a309d8af8356df` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `23d39bb426b8a4b5edf3b3097ea0d70c3ff4f4b9632a44096c2daa1b1c15ad5f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shift-boundaries-tests-001` | `ready` | `1b6fb7a98f2c49b2dcae93948e870d8f87acedca75234b2b21b8c9b8f10b5f6f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 10/10 | Phase 3A |
| `single-json-value-001` | `ready` | `ae6a3052ed66e9998862cb326b4c76c5e66a866cb9ddf0afe80226c54f09fff6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `89003d1c4ee56655ebdd061f97466c9f4ec807976bb0091e4f1c6ce6373ddbd2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `78c428c69bd8c847834c8da015eed5d263fe1801d4f2111de7a877744ec3fbea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `snapshot-transfers-001` | `ready` | `575de4d55bb469eedad06641042dec373b647ace49567a7869e1c2e10be9084b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 19/19 | Phase 3A |
| `spare-view-tests-001` | `ready` | `ccc2894b816af87d10b17c8f50317f94b11a8fb1430e13a3958e7d2235d5a3ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `599814193278f7244a28288f447b6b780c061a6a8a597c8ba659e39b2412b002` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-job-outcomes-001` | `ready` | `798db29f65594d49eb5e9124c1324aa369988612a2fb3bdaebe3eba464992e7a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 17/17 | Phase 3A |
| `stream-normalize-001` | `draft` | `72ffd557dd017b747af2c575cc6086b7e92a7dac66c0a9528c5958b115e605c4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-comparability-001` | `ready` | `0c1202c23ca241cc77602a0c7cd944ae3118fe054e4101d548cd0cd35bcb9881` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 13/13 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `853777690b1f3c8bcaca2845a600fd2cd685a9b98b77f59055fb3ecc81f21be4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `dd88400d7e86262bbdc2c89496e3c24f2031f597fe553c12df6aa542c94fd313` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `1c76207ed80d08e11bf012031a46c598950352572fec690fc47a758adde9999e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `9532266e3eb1e3c1434a6e1ff272474201f2d0c92884b739159d1467ce45f356` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `795feac3dd98ba663cb03199e96c59743f5a70978182f0abba28684103c09065` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `00edf839d80e5abbba1e72ed083717a8efc8c1d324bd9738b4af3d1aa7986029` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `timeout-operation-001` | `ready` | `315b015d79303a874db29e3b65dca50cc727a85dedb1b0604b2b90061c0ae79b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `token-values-001` | `ready` | `86cb8b87edecbb50a2fe8c45908ed90f75e35311dd908bc89593914cc607e3e5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `4da2ea8e0e0a7afe0ac9bb9f5761d80e60c1778617cf088eba8e50290297a9d1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `trace-context-001` | `ready` | `17cf2559366b847eb45d8903ab8bdbf4c67e48802d5eb8c30bd743bb4cc7d410` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `transfer-shared-values-001` | `ready` | `7f66e94d7dc1a89a81e2e541a820f2420c04de78f89e2f0f740532263deb39d0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `bbbace23b18e3926a4281165af99a7fd529f857685bea7a900981874100b3663` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `two-stage-stream-001` | `ready` | `96e968307bc8baee34c66180083c8359197f314e5798896973d1c589d4bd07dc` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `2a9cc26834b0d60c3462d5bf008cced52a69c9855917e88ce40d203959cfd1c1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `8901bf2c40974a9bffa0603e942b045be4f7add95fe766fc8beeae940dda6261` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `ed2499f418474a615d3e20627e50e3f16ce3707605af4de8e26877f8610d8110` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `2d6080370954d9d349c0489a4ed6660608eb67973fb06889ab4442b9704a231c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `179d56c5105b3f51704320b54110df7d45099c07c1a9a01700e49b26024477af` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `af37aea3ce281c7529cbc5d68669bf6718b60e38130cdd658917863ad5cdd84a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `e1756de7b3d342f68288959eb753eacfe05bc91ca6ee93a1eaea34da8774693f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `098a8555831ac2b803cbe5a6e460a4875da75db35728240b371e665f7b3b31b4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `255742d027b52223fe9530e9b38fef251565532d337de74490342fe896a73961` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `92e7c45fc4ffc810282aa5373816980137419a19af99f7d9bc3a24f316b93df7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `d3e34e8b21f2254c4d9898f7c57d04e63846ffb39542b946676960fb8f1ab587` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `e5e2cdbc214d9091661015284f45730dc686e6beca0c44de866b115b50a2f1bc` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `e91716f5238017876a040e04d3d84fcc829b95147c5fb73cd64748b5387082aa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `9db129d11488c92e077145f9877845efdc121a34e30fba7b2c7a457ee9c892df` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `8e72c30bb9e51ff0321fdeeecd2eb1ab95f6dfe264424f869c6c307f93685c9f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `4a20e7a0f1a0c460d903a3a0a68d67ad4e28fcc543b1e96d86c81b0e9eb0c47e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `65b8cb0b4d1e3965b901d0811901cb40994027278e22f8295a70d24d24e08182` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `d24e5534c3e1f30023bbad482b9c088cf2c04a1a5cca76f18a84104b3179d838` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `f8c0d7e98fcb32ff2bae1462204462acd9d74737cd137b83e495465bb7b975e9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `4adf6a872298501b1f4508e2c05e04453e8a1fc5c7d2b988496b354f768864b6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `85f6cda3f3f9155c7a4112122ee96a7393951c7832a6d460f05ddb8fe5d68bc0` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
