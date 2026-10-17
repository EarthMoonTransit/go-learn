# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v3`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 8 |
| `in-review` | 0 |
| `ready` | 158 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 39 |
| `implement` | 84 |
| `refactor` | 3 |
| `test-writing` | 41 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `5934feed02691f14f40488505bdb872e4f168a2fc20f173a68b846f70e90cc7c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `89205b9a692d88f13caecaaf1ea67a0511c4d830bfb1d0d08c098adf0cbfecb8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `535c5fa47b2b7eb18fdae8e21b8c3d8f972d704688bee0cfdc44ecfcffed6918` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `cbebb4925a8db9683bef79eb61cc0ae50ed254cc06e9cb9c58883ac8b41f3218` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `ef9efc7bf682090bcfb0253ef2698a2021c59f9f39131028aa2b021202d29b16` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `f643a047f72c7d00b95f04b8c65f4f41720c6bd2267c16fadb5a1dcd1e1a7dcd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `49c3131fb17b599a1d04b596e27c45ef2597b9e453c7b9d32404998b7b0fbe62` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `dccc302f7193d37fa7f1615f96f1e42ffca4e1be504671b12b7023d09a6e3e9f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `6125c2b4768672b6652693b18bba4865e8e927fc818961560ac4aa80c9b1f393` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `b5237d625a227f84f32c994ef3ebfc24dd64d6e60b5ea3a0f722d19ef3e37b61` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `04c644538ea0931ddba1be47fb79daca7fbe5e1283a33c07be1387a2ec3716e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `a2cad21a97ea4e6e214ac72a643f25a6c0794d7eb4fb057f2c2a3c62754b1ac0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `44be7add5819fb3921b6a157ffde2e9c980fa96396b9f98a9de4a4f28fff09ff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `08e2bae8bfa48cffd4d88b8b77915b078231b6c64a8ed46b1828f8472077c2ce` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `c0ae1ba9424439de5aad0e5818bff52247086b8bcefbaeda3949473a047a3f04` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `0dff7c90374e8ede696533219baa6cc88fd3c44ab182879be4c4d3510ce0d841` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `capture-panic-001` | `ready` | `2f4b3cdd6aa21079d4c89f5f989a4fa6885db911a5599c6519a2d0b9017d2abd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `category-closure-tests-001` | `ready` | `309fc62767f76da9d78d2e7dc7a8d55f2b5472f0489234008dd7828e1d34dada` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `c057bc7f492ac0cf88dd3cc6caae167c284ab41900d9d5c6c9ea67b1bd33fc47` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `19b035ac570aa5fac463570c7fbd6460418e4edbdcad7e4c64bf1b79fa114e13` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `a4af4b861ba571fc238793d564f11c0465f06ef135394aadb625f581f96672f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `cb09d1f6573708e3d0b8744d586e252597211e232701f13dd139cdc1f212b58e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `bc19bcddf9f311398272e53df9d68c8b82a4d107c9c71a5b38e174b520205e6e` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `bbff2ac58f1f7df09d94d35111e249c98c93af1d458d397e8e57f9f50f26c95f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `90146b816085d438ceb4b9e1a70b4905ce0cc2b04786c00bc24fd8dfa86717d0` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `f5baf75a92ee761c893871f1eb1dbdf6a84c71633895c17f0fb30c657783b1e6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `58f9968bbb2f378a9bb12558da4695894afac2251337635b9486dd4acf73d3f4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `b756a0831e18be8675341d6d1fb0891fb7adfb6e478d8f34c9095b612aedc135` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `d4339dcea77e5e28ffadbf0510ebbf50c0e919b54785f62c3927b1e75beb7115` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `8a1c900e94708eed28276d661ad04ee75113fe05b4626add75f4a63067551d04` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `c1ee1a0292cfccb4aff915cf0898c8b48980a8b5fed0a01a16352ef8719f3d15` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `0135f4e4c7d0d5849b70785508c1f2b50c472686e81816f562415887538372e3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `a08efbdae116df5cd5b6de34d192aba3d7d348675f6004f23e983be71de44d2d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `deferred-release-001` | `ready` | `be39863f4feed73d358f2337aa96a5feaab5d77119cda72f0f86a09fff7c8110` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `052a3371b02e90ef0608847af18fa39e5107d4632ae81ea70ca07e3fe7c5056f` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `e13c0403590fd6a25c4302eede4eb4a8e7dfbfcdd0a636452862e2fc770e286d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `3b6ab7f45969a817958c61ff6378172bf654e9caf5a831624777a6069bd539b7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `9d37503af6eab6bc9569d43cbf82571173036af5f33ff7c03b144b55eba8d78d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `b14918c1f8c0478aabe9fd3689147fcea1c66f6765163ba8b87bb45ff7863f22` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `679746604db1ccc2846bf2da87e5e3bc3d5fed18889c0cd42233c779b77fe78f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `3b628c44a24cce014cd8666fe953624834961d17d11bf7adbaf5fbaab797021f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `f20568b37c428c41c9f98912bb95e4756cc928c3f532cb9d58b1f8f09a4dd2b0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `211b655944456284f7064cdd8f79946a1acb622a0692ef29a3d5dc5c10a7f995` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `2919ed627974111cb953478f599fd91ffb67ed3ec32aaa500d2dfc2c9a9e3880` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `03e91445de6af31fa82a7f74325f1fc6f28deb173ffcd600b847eeed99137f49` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `b9a5920cbc2af2debad762d954d46e9c80a667468f851c00c02de72cb7e5b688` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `2437e75701b24687a2538a9b2db4a9f1aa4569643630a3f855d2d3772f65339e` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `66e6c1e2b1d5f54fe0020ee662edbfe32c2d0a21cd622837b6b518fe4f1603ba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `ae08dd34af9bafbff058724988a8a738aa155a77c45e7347316f337aad92595f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `44e5b61732ba308b72d2e65229986b0cc3844e4e37c80c6f9320d901d9fff1c4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `ad881a0d1803e025452dbca4b7ff9e5d5faa144c6a599eeb58802af281e14813` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `01c4c6103423b8a313bb99c1ccdd2dad2009ef7f4f4fc4c484aae2ffaf023a2c` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `3baec58ae59cd8cc78f977680cd77e38c173ebfa4d876239d16b973c2e2426d4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `202767850a657c57e4acf2468027f98cee6ffbb4e42e3b03bffbdc90fa0d468f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `7b3e058b16957bcea1bf0648746886b6ade838713d88c2d1a8a071ec809d1d91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `a5511e8f8e93a048377b809698d79508f03534380e510c185f3e90042447769a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `11d0ee3ccb56246c2e6fffb696e06f2f897fdddd5b42215ac9b5a9b8ac9aa265` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `f8e9070328d6e711049117b2d8fccc7d65239085723980dafae357a4e8b50d63` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `afeddbfcacdb42cfdad6b7a7c07235db938de95bcb92a510254f656c159ce3cb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `bbb3ceef82ef1501fe60bb191a740a9b66014874dd04b9efa7a1f26b4aedd905` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `0852d1d456b960b3a680d3619b26c7f627ce0005d22eacb4cac01cf2b2e4cf57` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `draft` | `41bf77f3b1b5ffff87accefbea0de01c989b3cd4aa5a4bb12acb9f03d854e984` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `fractional-band-001` | `ready` | `55b14caeefc3892aea67364114c72047eb2c9df6eb3a82d48d372601c3361084` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gcd-reduction-001` | `ready` | `42dfd86d5b755705c2e3bc3f0cf26b0d2ad5a9786a7bd1b8f17b3d021d8e00ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `bc3ef42e3968df7e303255e2ba5c404d5c948b8c668782bd1b6dca4b01709fd3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `93530d9a5c1141ec91b2876c7dcd1e1abcfd1aa9eda41acee2726bf3136386cd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `ec8658778f4a789dcca45ddeb31f449195512e44e87b6cd3aa15b436cd369fc8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `3c74c613507a506e31ac3505f23fa797c3d6d222cf1cb0a4aacdea80ed1cfceb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `5952ae146eee61c0556bd95a119ba61d3e62b314b8be978ec39f4cbb7b9e6d82` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `6883e7a4bea5d47560807c3fc45162724ff0ec752c47d62df0c5defde85f8933` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `b6cf94fa81b92390ed1ca1f5addff9ea880f4885bf20d42cfa88898689dfbec3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `55e3684e2d6688d81ed36c53ec18c5fa996d81f0ed74cd5bfca155543d3142c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `initial-structs-001` | `ready` | `ffbf04083ef6f4c76b1698b90d72db446036d0f4591331a196f4d75b25503761` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `31f364fac0daaadbfc856698ec571d79fa9fa191bab24215f9ad79da1b084f18` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `1e0b5ea5b1cc63462c9bf596c0c1a8a3d036c5738a3ccefa0357a90ce7ab0508` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `de343f85f123c5abdf24c665ec1342e62b81a94601b1c92556fde54f82034065` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `4d89138c22b2f4dbb042c3dbfd116463f22747f961b55ec2666d243241fe51a2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `d721760945132f6f88a7da89a6b7c1c90a0a467ad27dcd1f34514d665b5e9430` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lazy-probe-001` | `ready` | `c689bf5392707c7ced3150a9b8ee5eaf7ec7d42746bc279365065ef64927b60f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `9439d86c916aec9a443b9fecf32d6742cfe84b5595d304cb8fe5e22f45ac1994` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `47a7d5e18118c50646dc4700d9d2778c967f8d6534cc0904fce7865d0cc2b0ae` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `60a0a747cf4209dce4990c25cf2bd64ed7154d8080c02c58bdced738070abc22` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `2efb8ee4d1c8ab5da1ecf395c6451c9173e2dc6ff922ce7ae18ccf98530fad89` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `ded5cd52c875b69f1e4fa00acf8e2bab6dc9255f6068824f98ae09470c69d6d7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `month-enumeration-001` | `ready` | `90a21a1673ebbf36c2dbf1f26ea4537770269a42911b19383bc8e8fb3766184a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `f2c8802d82ae64cb44e2f531b3cf34d70dbf54b4b3bc913dd583ba59a18bef19` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `38717d46b4b84a030e9161cc0b8ecac08b7ebab12e3c1e44cc6d940ceb4e9187` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `a7df40028fe0f8674d9650e24d532842825b74012d5cf53c998bc4247926a91a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `49c93a0e4b13a00893df0bf1836cc83952ae1c4e8604d67933e5ecd4852baa53` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `763851af6e3eb7b3a2a5b75402e47f6b895eed6e3d469f51a2da7e141e859dd3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-slot-count-001` | `ready` | `e4d4984a54a94b2d981bb9fce8dd2ce05e689f3aa0a5aa9b6b9e7e2914ca4272` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `43b2ca26134bf4a0814861e9c7144c184672db304b7b4d51c7be7f4541c28300` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `0e265667a9b96c61513480f37a4668d08e0a0a8e5c8c05c3f5df74ff469e152f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `72c74b36cd9141c25cad33b7843e9b5464c9a8a4c366397665ab58a25a9c3da2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `d1faa8958123bf90b50fde038cd625595dc0e9591d460385a429dfa7f4cd9047` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `1acc67523b2dd3a07f736d99684c76e4b30743650785eb12e6450d6219104d89` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `8f4f11f8ef5320da05f87d4ba885b31e27313ca6e32f140d9c979d3af0ece014` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `82b5a9a2534a15ddcdb49dcf587daa7977e14b717054f83c3eb2a4914912aced` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `856c97f7836fdfe85760b2d9aff75980cf903563351336431f405a281d5dc4f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `0721d680431b7178cc5a5a36cc7585b8a3c5ce8d442c9946e978cef02d63d55a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `84e0e990990aaf7d04c3421e4c6a42306414cbc26e466ab75c4d2e106edf0cde` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `ed9b0b5608bb494f1d7584fa99482e786c7861d7d2d077c42a569a4c8baed2b8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `adefa616db67ea0c7189e6d3f0d708c552a79342592bd6b295a01d74d6f057c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `9ce2c031342167de0d167682b2eec0f32863a1991a0a3286a7888da286d110e7` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `112118ec76f903ed8208f4627ddaf132271a432b15fb74e7075264d566994273` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `b786ad51f2ca67b340a94d8abe9d2d893bd928ec29c6ff736cbb4391afdc20cc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `628c46b8844391a349c55791451d102d9df19ae3129a786776a492de64b5ae1d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `b27d6e15015bd5b08b9f2d7b1eb7ddea554945b85dba676e8e657e99d5e3b69b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `4c34671b7bfb4da8c41cafe017286ac6decd038bf11e95882b66d9f71863b346` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `28423bab4243713777f294a2a4a764da92e8ea8198a9d605797493713c3aa4bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `3d7505791cd077c2872a89b1c6948bfee8adb3588517f74cdc51ba3562e14849` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `5ab4631cd8f29ba5f588c53880f8f9808e1c8e2baebd3cee620adcb6aa7cfec6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `d5768c732998583124d6759378c6617a6bde536ccce83d57b70e580603122446` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `1c251cb5115c6bbadc01e0ad3cecbf68ee1f68036f1f67b34f0e5465f78abf18` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `241be1a45e3404f5f0b2c1f2afa405b3499f411e22f42390a52456199ce2103c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `f5f91e8d47bee783e8266e5de061baf0afcaa57c2452b2e95cd190f9eb7cb492` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `6418885b9bdb46424be97b4467aca83a9afb5dfd84e1ba721eb8f165a3e1b64e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `0a70b131061c0b54961eec1bbba18eb4083d08a5d0a65178d44ce99f771925ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `draft` | `15235dbefc0f98bdb85ec73105cf52cee01768e62391881d4b0b076b73270632` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `2110841a8110c0230910dfece941f5c0191b8b738d1863db92887ba87545d6ef` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `843ab0c9e588ed1d0bacd0383c52b64bbcd532bbaef40566ebd0d1d2e0fa2438` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `3e4e21917fc9fdc0f39ff55c49a9e6799f0a50b5f572e09d5b3aec81dd1ef06b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `ba1b4962f9d137f498e0e32b3b01a10ceb4de0f1b43a4f05e8c7656e37a90838` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `9effce66fe41a937d0b100fc0701c06ae1341b51385d8a754a3508f7aeef5f45` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `f9fd44d718685d2aef5718109caa9cab650ae99a828b840ab6b2b9b1042f13f0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `88f1e9738d58ad3c49470a24090fcb6a69ced7cd2059e6c9fc74d97e8aca9711` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `2686f87187fb04a974662f820d123f8d2ed4914286e47335c53776a355c46794` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `c1e194a3add4733aab2b7ab9a08a6cda8ffb282d3ca9c2d258e29c53b2c90eee` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `1b9c6d5d76f7f57a20a17bb96f4c3e3255003a797a217b2c3074a1cecbe4c0ff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `1daee6376c56302a7886c97a86d362b9c75b484004779305182822e9eb5a524e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `0e0c5d3d235515c16013338c559d72bd74600608c575e0449b2656bcde57a787` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `17e1ae82afd007a76af59a5793537fc838af790813054e2e7de71801506fe76d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `beed1699f547798b398b1526c9de0769ba3d08504212002eb820595e9e68af05` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `4b7debd9a217ce7b2b1da79ab8ee62258fdd0ab3bc4cc4c420775f80e65be970` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `cd95356c9e7cca63241a55f9909657047ec2d14b801df56362b731b13db7633f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `fd1f250f8e1eb9727486a164035c94f50ff2dc1ef58e8f96c59312f90d49f613` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `e1fc37ed4774c57fe79d45f4df64e5980947300008680a6c00bbf8da074ebb86` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `a4cee5169bbba19397fb53ce1e6f5d2738c489ca1c7da7c530853232bd60a89d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `b2110fb03a8eaf64296de918a1a54874f22737b333ff8dc97233f7f57a4d3ded` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `c6cd248ba2e84e80123c0420a22a951dd0a9c6ccb1992e5f3bc6459b21952291` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `1ab703f33dcb9956baa6cf879fa9b32409941647fc8f2b74563334c46cae8ac5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `0c178d5d722d188b4f86418ac81d6475885c736a9464898cc16707006811080c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `5e3822f0451c6e0c0988ae7e8288142934dcb259b077a1d0d82e9528b4aa303d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `token-values-001` | `ready` | `836bb939e2d111f4165be0434ebb3162989eaceb663c229b10c0b7578a3dc8a2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `da3a05a91e3b74a5edd12f0fed92bcb960012fee81f6c88200a7f0399445ed68` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `2d145973add50b5e270d0fc4186bb200505e44091ec0dab283d9a5f97b60128a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `f618882e79f210d37d67c88b52702b2fdbfc9ff868ab1fcab761bc86560fe8e8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `338416a27fe7d7bd115b75c7a876ad369e8a4aa493ffa9f8d2e5c8600101320f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `8090ef5ff91c715e1bdf269a0e0cbaf35474e3464f868b74b14545290e2ac463` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `7016e8a1afcdd8c117a707748f38e763d136c309d4aff11eb0a0e8664db42fdb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `2e1efdaa6f0e124521e71b30e05d3839863534629c50003d977610e28cd4d0fb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `d00f7fe5b6f080bf6245d0b7fb6d45ce833fd91eb8424e3eb385c78392efe319` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `875808495841abaa7f4dad368a2b53ab87315d223765fc776a82ca1570a07bc0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `6d14b2d8f099f8e2edc0748b41cf6ac7ab3f766132e1239e4fb1373f69a7abe4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `def4b11cd1d799b4ea4092b3459ebc13164596119cc1a65f3f8cabb85c4fa28f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `a8a349d33020f0cf9447ae2b3e0fd429588b6265825143127cb8efcb84b78ab5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `c3dfeacaa823866a1722f8844dd100f7c10876663cd3180bc01252b92048afe4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `2716a0b8a46a72edb08ce3ba6c2c12e22a719e0c942cc706c47cbfc3deb64700` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `8bd6b266ebbba96b0bb797a7b1fe875578179f402ed8dfd3a09ba96cde0b112c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `7426d19e149870441a4817bdb90f7951ace1f6cc9e0a2974ddfcfa9b2675c5cc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `07f26a9441ca135c802a1ad560868266f156e58ab741e25982bd1a5435954e3e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `6f9b7c51880685b5d3cedf30a7dc90da20c68d1756207e8a8cc54a71a661c367` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `4a17241dbb6a73ecddfbbdb996ca07d5a84ce36fcb59d53837ca2c66c218a6a3` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `df2e0e39a531a43b78b5cc2d56a409c5a6a4801d4c2c2b606c77bab688cd9550` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `9d69fdd547810239bfd90a9a0aa95e7fc720c8c72666e46cc04073eed8e262ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `30c37d74f835543b0f6963f6fbe41e9b5cd233b87053ed12a8eb06179f1628eb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `7a18d473341c491f8ea214b2ed3a17d358e94c7fcb36ec39578927554f880f83` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
