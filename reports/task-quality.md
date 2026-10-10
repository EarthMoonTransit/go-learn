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
| `access-logic-001` | `ready` | `6f54582e2e457d990dfb7bbf53577a9b9e95fd19efffeef3ac3713e0feb83022` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `4cc4e5974063454d0a57047eef052decd72bf7207fb354ca663c8f12d8a8053c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `1e6031388b28427eef7b8d27e4d74c6497ec9c7bca302b26dbc8d2a135d6117e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `70a78bde28e93edd4692da3a64f92237de366bff93175432af69c666ec9c6f86` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `0bf46a1569b6d23223db15ab16f49ae0bc3a6eaa5d438c5bacecf48417ed04a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `d4ecd4275163083e6c1301003cef9572e63c3c702426017e3576201ef16f6c50` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `98978147fa80ddd9f0247edf7518c505659746d83294faf3e201170d07f51dc4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `4699a43ed7457b5717fb92362d8a91cb84581593567a0d008cb83ceb569bfbd4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `d4499444471c2a10aa5242c83510cbaa168c6e04348ae991462bc2e8f1dc33f5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `fdaa80edf1a81420cbb216e5ac1e3fd98f82eeb4177e4a74522deb0495bdf0f6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `5a991112bd65062e6911368f7a5789e159530724d4b20deba0a80cd7043c2e3c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `bounded-value-001` | `ready` | `98c9e2b7d1f0531f99eb2c6e1ea413ff630e7335c077dd8c6a76d3ef09566345` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `115784ff7a2d93200e999d10108b6d0f56cca72b135da671a0fe3492cbdcfa64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `4c305177626d19a0479a0f1848a190d8aa530e52a605ca1a9fabfb80ab75a429` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `db3baa921a946f89a5050ab74cbb1c5f5f28e36a80711e4c31732bb4bef096b6` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `capture-panic-001` | `ready` | `37577b6717e2c9e0f9421110338f6965a4ccbfe7ebd3916739615760db96ff66` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `category-closure-tests-001` | `ready` | `a6179475a4427232be9ca28891c756ba72f745e7edc608ec7d67f968e2af2351` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `3bd9f9c9ac5989f9bebbc9d72cac9524d7aca2c17e27cf9602abdad95f7eae74` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `3b8697724ee0c197a8ab3cd78513cdb8e0f6e1c429f0b420cebe1f5b854d1e74` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `b0e7f5a37360cf4516a02effa1113d18546de4519433330aa4dd07c555a6d9c3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `f5db970d5a4264636156969e22fa6e2e1f0f25af627a67a452eeb7dc4d2a2722` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `bf45fc3fc6e89f874554ce56d2e002ec91e5c7d52de278b0d9ba0747303930ff` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `e4ad82053a791f99fbcb89482eb2a236e2061503c72ed5870e8a04e7a406263e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `13b72df6c3e14692b587ca37c52fb661d02d3201c23b918b0a0ed2b0b8bb0529` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `constant-kinds-001` | `ready` | `0870b6efeba87062020ad7e065779d89b5bc86e305c3807633a9e89ffc4d5b7b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `bb9f0fb79b2fbadc0c521e38cfd4b48140d8fb7d538024abb5330c3b1ace3480` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `8c6769d442fb0dbabf5dd56036bfce3983a3d33a237f6bf65903d58c8dd80b18` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `8396c03a69e8c780b43e329f7b78a6081fe1761c8a4f931bc160383524890fe6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `7f67b32cf403f23da17b694cae012e95fee9377c131637a380db547b938176d2` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `438fb2d3c30b94764bcbf5ab5e9761409ef80b6a028a7630681edf7d52c41cb6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `69aa43f7824a8f929b538a7715b4fc055e7f93d07829c82e42916437d96f5114` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `deferred-release-001` | `ready` | `deea605ccfeb7db6dfea7d0f1bd61a2afc35a86d65640b3eb4d4088beee18d92` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `7023b51e3eaacb9400eca582a51fccdafd61b37b775fe1ca62e72df46ce4e8ea` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `081c9568098f5424c1e0bb59f9f548db18e73144789cfdfd06332e2ca27d77d0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `bbec195047ba2a1e0fde2c6f6f7a3cae9218dc69123675f6422b02cd3d0822ad` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `fc0fd7a0119632f6146bbe46f23803e4059a6e87abeab433456fbe0e05ef5584` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `7df809a5f847763d441a93cdd6cb4ba8a230055ca4e63d2938bd7607c48dc381` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `9c6a12a9649dc6729ecd0381f03aae1add913a9b45324d9197af0f121fcba598` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `c6f1e723fdf49879e9560c6cd8ac2faef98a4d3f7fdc1d7dbc4466097db4a784` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `24c1128ee209567652861338ae1e37b08264ed2a468763af36388ede6a30255c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `452097373aec0b6cae8982bdb193cd319f0f0fb848976b38fc4b319d72c5e218` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `30d5e2442bcd367290234acf549a25c37afe7bcd87aad81bf2475b7a024486fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `0c445566ead2fe68e4d6cb44e4974a9b25a3d9efbd289e09b2ac8a0fb5c317cc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `5f0760102ef6ab92eff8df8a1b6586ce4d06a19cd24649f84f0f7a44138f0a03` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `2a0be76810a609762be745348fe6b2edb85c5084ae34a2f7e555ebcf504556fe` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `392cd076898b01e9f1d168c222b750e3178637b245bb4079a9901c4f97b9f5b7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `468f4fe5ecc4ae804c14bb3357e0a850f842b7376aea8e9e59dbdd3af33dd2b1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `6ef86b0cf5125630d247ad91ee2a1d3dc82003de95deb8348245890453b46cf6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `bd44e1ae7e6dad0419807e533812ea53a44d6660d1c854185379ed50c416cf06` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `4115d059553d789ba74fd441c134f98aef883b2299e4b192d7aec04431ed5bfa` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `60c2df87fe1caf46c903b0140ac62a7abcb2ce7d2b4da637f6ba9cbfc816a3cc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `b7ad153a603170cfb82e957bec6c0ae2cf0e8557a57361e0824e2739e16ca911` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `3248d48878e030ed90f24a7248c03d9780c22c797b3ae1483b57aab558b783f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `5c4c71e6e67e3aaa416628c2218e5b93ff25fbde2d9e338936b57b24c79a0114` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `080b72ca650915b45794e2bf8815401e8a0c5301cbf3d2c60b9d0c7df41e4f1c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `3311cddc26122c8e8fc4d877c4e3423b3a81c58764f29f2d33871f7165f5be7b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `a8be28b0d49cdfc30af3369a1098d93534c4002835a7d4233e52660b0fb3d312` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `96a6b25490ed6a9bc0025a5bfd62b531d2541e4fe764adeca376c43ead8bd8b5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `86c956ebe4f93cf62eb1f54c626cff618c935c528d540b925655b4ace39d34c2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `draft` | `6c55fa374b805fcd61f39317917fa52bf4f98d4a74d50121c23f3a8ead5cb380` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `fractional-band-001` | `ready` | `4e9ec2ed564b27086c9f19902ea6741893132b78ea28c22ff4cc4c51ae8430e4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-find-tests-001` | `ready` | `f67ea692855f3daee60af863e893d950a39875daaefeb28389fbfbce9e21b22c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `7ea0ba3a6f5d72e93ed2b6e89644c006a1db16c6af6da07bd0f00396dc81f1f7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `31d3c5d8578d6e33678315a144aa165e14388fabc53b9cc0daee227e067e93e7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `63b67fca1a088e12c863c5169fb8cdaaba682cb49b30cc354dbe35cea919d141` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `d7a202d125aeaedec1bb605dc7536354f2ec55f11a0b2bcaeb02c69676772bce` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `65175767eda383d64587819130e71173aabaff5a980abb74a576b5ec355414d6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `c3f32ade2c4c26b9b6c801f50ab8929c40dc433b175c09ca03a73c5a0c448465` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `16eaa81048cda7e7003679da99f4c66f4bdde2dddb1f5199936133db251c8fac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `initial-structs-001` | `ready` | `9a164058bf4cfdce36cc330bf0e6552cd83c41796774cd2bff21beaf587da1b2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `49abec560d2566aa9a62260e2bd02165df12f9aae220a6d70d23eb9a39ad0d20` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `a9206e9ec6f4fe15214cdceacafca595747ab6600f351038dc3f288c295d1a34` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `08d9f05fb4bb5247e79e57b822c577e1e27f03fa00ae6cd81c1732fbc7a5e126` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `ac69b107286e320806154f3b531edbc5b662fda874882058fb30bc0b0aa83ce2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `8bd85c475ec2c362d9db747a464f900b5627975eb3b1757b5c89170e77ac7df0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `89d45dad276a9c4785f48a69d47b34cb0c53dd7f603c095084d5274e260e23dc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `b14e721defa63c7bb493bcf064f9f018f5a776c23bc13d06639111d3b1d02f01` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `391a3f20af8bf84d6bb86d4d51dcd5a07238af4c7e75dbd526c8ff2a08d61fff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `b0140f0eb703739470d6b7fdd32e051cc8c244b8e1986778852386a3ea0b8e90` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `5e968e1978aed2cb2bcc04457cb68dba47eac78f04e5da9154038526bece0f58` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `month-enumeration-001` | `ready` | `35514cebbc39763debcae7e27db0b5c04629d777ef40041718d12064b477fbc8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `a7aea25127c59fcc7b1e346f065b88724985b2da8f244f2c040160f4823a88aa` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `8001fe3d3c01723f6635aa19e940604631ec1c4d4c71b49e879dd5066afe2230` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `87bb92f6324eb41ff7d8045a0bca2c571c490585a59e98e84d08efe1949bcccc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `740c4c64a6efcfc382b6d4e36ebddab3c9e69db02cf29efb8752a7d87735998c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `aae975ee6ee7e646bcc47c2c78287007074086929185b3626c70e45ba8ebd5ba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-slot-count-001` | `ready` | `26f1eedc09fcbfd9a937fc328d4f4626b31152e533d07b42d08577e3c1ac5fcb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `3ef34f924ccc89497cb91fa52afb5feab6cbd26906fca4e2a271c9b5237502e3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `ecbceba0dda11e054c8041af02b97a9a0eaf7f045bfbb22c98eb1c654d6779d9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `a34d43f56fb5c5e9cbf2c3549e6e7f1e0e1c21708a07582c75995c127f75e273` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `a7ecdc480b2af45ec9e8dc1d4cd8157979cad7582fa920577b6db4fa698f32f2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `f7f94907e86bb36adf7ccd5fcf20acfbabb51d84ff025b94159827fbacdef4c9` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `ec59f376f52bfe23ac437f0d16c1eb1dcdb57873bc65a9ed526959a0f74bc1be` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `742703f4902ec4257f3bec40457e2f61c2720f6df08484ea9f2722a43766ecbd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `d4d2a53ad5355d0e2aae97135701d24c8e6db46c5653ed9fd0ee0cfc2a286cb4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `df952918494f384120209c030d02bbde10f53c95d72bb5ab35a0890bfef82910` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `e17305f8139c331fec0ea13fbeb147c92a1beaea707c069c58cc480a2dd19314` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `a231d887399dcf8d3822f1980012fc443846ccb1054297e0246b4548455db523` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `47c6832a7894da4701d869ae754ce09c8c29b1067e969d215062d4b3124220cd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `316bc081fc4553cdd1190dcab40b4fa7aebed938a31f21cbf70756387477b294` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quantity-presence-001` | `ready` | `e20979e98af2772880c02a6db9b275e4fd7a04a5e5665c7340d1c2a73618e81f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `28e9caf5d81a79501306daebada91a6b30809a839b4bdcc009992cb18ab96189` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `0edb2b4544eacad4a68260ad0d793c24bb21edec77272076d557a7b335e11478` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `b849943e96c537fa731ca62edc81963649947a889c63b5c502e8ca5454272df1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `736fe6260af70d2197da2649dbcf55235a7d621660525cdbf22ad31782442201` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `c32a8d8454ea188b6dd6e3af41b7f1070f088b25dd647b5fdad39a259b3f7a64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `29ba15a44b3b73b983ab167c87283d477fbb303230a36075303d940094edd2e4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `263c37e6cff64f5de5a4ad258a7597678fa546266125c1acd3c8e2b4fce66cf2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `eb7d5f4fc2a43b291cad8816c4c1f04bf9e685f2e1d56dff89c29d841c92cd68` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `9f16cc7cfec7db852bace6579d5749f41f43ab61ee158a5f2e77f4646f0bd679` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `011a59f5d330ee67874f151132f5b99ea57975508eac88880032e8e99a0f178e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `8b7a38eff3605aa30c21b14c0792319303ebdae8f77be9608fe4372791114e4e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `b4a4798a134ae81ed2cab0466d7d206e266131d24034a5a6a5e038cb435df973` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `draft` | `c2a1ee114fe0ad4f26f36140fa3f868195624e248fb4fc4b143a359c65d21362` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `028bd9cac22778cafbf706abd8219eeb70f151db75791d971e990e861ae83252` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `b7fe3f97d8d6e5c18d5d61e50efbceddff98d96b26e9fcb77ceb97d65d762dc1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `706e68f5a6ae445ae02c02f85575483fc5103a4187f1c5219c703152895a4d6c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `0fb660adaab61aca26d7f664f5350b3d0c3a5d8d20b5841cc0437ecf2ee9b3cb` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `349864265b8cbbbdac3bd47ac855b76a30305e0e79882cde8012593f0002b58d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `1ff98bc2dc1c1f40b3d1a74ee3bc6daf1fbb444351c1c24765227c158bf80345` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `8cecaa295051c396c620af80b26d7bba2aa9fd68a2ac5b4f07c2249a78db78e5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `55ee74477f25fd4b190a726db351475fb4707ed9ec07a09ad08782fed991949d` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `ebe31de9145b1f5dcc896c1442f41a7c88acd1ad484cdebfee1bc20b47aa93da` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `a531751b4241edb71afe5934c9002f2759cdc4d09836aca222d2b437b08cf674` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `f49b4ed42a7f28e7574282725433b43584bc551cb0ef7ce37277ea3db2682e84` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `077b29baaa9462ab70709d230c614033c276fc9d1af3766015015c376fb4c183` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `eb86c532915ad3fd3ea99ea64387f7ac70d83278988f6867768d67dec13eed2b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `a8481d2f7970bb47beffff3979c51f565b87452bb3ed2be4bbf12cb7fec9be80` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `2fd3b7b2e1b7ff4b464d152f8098e61c523d4d630d275478d414b49a115b2a12` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `cab5940d4c51804492e408e5933df03da3bbc5fb08c93ddfd1c6cec8d45968e1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `a83562c5e0f14a7a50b729750cd1f1aa8ea25586e6c1807c6a8fefa5946e9b7e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `9bc22bae810035bdcfe1183fee6c1dfc88322281f414e07fe7aaf295b999d6f0` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `9447912e5d4d56e8556b9ec2f644006f8c6d12fa5ef14dc109bc752e2c21bd64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `a657538d80c15d287faac8ecfab1ec06572256216e5bb5e2f7038381e603af3a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `40be02dbfdeda225dc474d13e30dfe7561d2e0151ae6abf1ac6007d6c0e277a1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `964cfd6e0f47b626cab0c60d3bed7b77790ff75d5b1d52939766f61a7509fb75` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `f543f6cf5b2e9f8637060d88caca91916c21e710bc48a6c4b6a336957359f181` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `b78dcdc6be8106143e53a8f6d7bbdaa65adb1048db10f1135aa9baf1557018e8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `token-values-001` | `ready` | `fa048a3bcfb02f9a051f143c651269b6f48da7e63ae79c40485e712f2ca22080` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `9e85f85bd9784e1ee89a77a6027064a9d701645dc0b013c9bf4c63bddddc768a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `50180675fcda0cbb80ea32cfadb8ff3977f78086a02c6ce07c46709bba1d8bb3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `b24acfe1259361febde5086df2e6dc30cbee693da6a8b643eab5f02621068122` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `37bfe8d24fdd52ccd44b99daa0d33ee421a5e7d44aa50db6b850a5a95a38f317` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `c5ccfc7182a9b04e09ab744b84c65544fc197df67c77ada98fdc010efed7b9f8` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `1a2eca52212f461834791a0677b02c6bcf1adf2a69e13b94aa5f7952b6aa571b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `2c7ce1bc02c705a3b5472eaf343f8e576ffd5a1fab771f5d284ae699e47318c2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `b47cc993a1cc21c6de186fbabf86a6b1741c071ab070891790dafa6adfededd0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `a581dc0ba234a2954a411c78352db3fd999c226709715c5b03fa9ed97f66ee42` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `bbaeb528d6ec9d69ae841a44d3a4ec9f18908477c938bd31fe410aa433190b0f` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `422d9675c670a20b7ed586515855e2f232070405ed8d61692ceaa83ce2969f87` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `77d074d5af9d5473b9f28d369a127c99b3e35d6aef457e65b7a85112045d1ac7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `01c6fb6840e599b227bea613035f4ad3e9559c657fef4d0130717788740c0b98` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `39fa0fffeee6b4a8de659ffd17c72c7532bd21b65cfa77989e0263c174b7062f` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `8a2a2dbf7a391af20074a6517f01ec008ae95078b92b06505d0f4f802e079806` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `fc9c25cfdbfb974bcf9a5397e85694ee211fde44e396740fe6b93c557c135484` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `cff1f6e35bb0da3cdec3e837750ab7fa02832648c46f53a913dde32c470e4fe7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `7aa31b26925044e861fe62f3c0baf63874fa4ee9a9652f19a9141dace749d126` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `ae557ffb6acddb6a9ef1e2c7f85b8d86febf63afdb3d893d4f564c69454b26ac` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `9dfc669d1d71fa83cd5c2955943ef4fe236fff4db54617ce3d455186cdbb23ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `daf3e41cd984301dfd6b48934094d74f264f1f0b8385403a24eec843d553f478` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `08702a12077b96e3617d284b31e10e0b5df92344578fc7c0b09cd5e1587d7c1a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `5d6e0175015e49fd3f941571b949afadb164a72e0cdf53213b76dc58adffed81` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
