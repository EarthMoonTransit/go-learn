# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v2`

Toolchain policy digest: `51c9dce19825ed8c5cceccba7b81b0993be4f9213caa1571dd0a51590385dc8e`

## Lifecycle

| State | Count |
| --- | ---: |
| `draft` | 8 |
| `in-review` | 0 |
| `ready` | 95 |
| `rejected` | 1 |
| `retired` | 0 |

## Task types

| Type | Count |
| --- | ---: |
| `debug` | 26 |
| `implement` | 51 |
| `refactor` | 3 |
| `test-writing` | 24 |

## Aggregate gates

- `task-state-policy`: **pass**
- `task-type-diversity`: **pass**

## Packages

| ID | State | Input digest | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |
| --- | --- | --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |
| `access-logic-001` | `ready` | `0b961a7e937b235bee75bce33fe24c9ddc4e587c797e690e0b5cae91fc3e4ff8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `allocated-counter-001` | `ready` | `955fc0cbec1f989effda7b4be131c264fb63f6e8e3d4e0c31bb9c88a333d0337` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `5c8954e58b37e08b56cc25332ed401d68c974e1fd39f6226162b20e5c1fdf142` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `ccbd6f1ce1257c2b408709f348109b4c236bdae7ee43796d8a55776cd08aba64` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `0c6f87243593f195b85f00f77b208ebafdf7aae3e11972968a55824ecf4035e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `boolean-policies-001` | `ready` | `05ff50cf291d707d44e08e32f3133e4fa0afcb7b53a59993bc07fef10206df58` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `e5a93f051bb5c982f0f0b8c7e8de7ac9d8a764aaffbb976b4c92d14380adcf3a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `bounded-value-001` | `ready` | `4697ac79d9cab7e4c75cf01963a8bbf2287b63d343ad85dd24af59746e3e439f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `byte-window-tests-001` | `ready` | `cb56fb71992ba01bd1fa4c4475f50be14aba88206243ecb626b1959ec651d781` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `category-closure-tests-001` | `ready` | `42dce59a44d405558ec1a5203f32390f8e8102942303e51e4ac94459840ea907` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `0416647d14e7af80c9d577587b6085094835075957062196ec4fe6961f9b950e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `e5f96e86f399cd5ddfabd706deac00b7f78f52c060a6c199b718a833af2f24e5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `808fb55df8762ddf87884cbcc44bed1decbb552b58ec7772b7be1b31a1e09786` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `column-occupancy-001` | `ready` | `c1cf3c5852d3da6cbbe53e438b51f3d628ed473eaabb2e2d74d80447f0da9b9f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `copy-final-element-001` | `ready` | `a58c5a934af4eb3fa8c3a846015a1ba8bea4b780c60e9756b9e47e09d0e08bc7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `1fe04c3f971c236b14c7685cdca5467c100deac9b34200c31b73588f61153421` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `7e46814fd59cc686a0bdd4d7dad718e79f30d039e51ea146af53dce9d4db8f04` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decode-prefix-001` | `ready` | `34f1cba662ee214301819057cde04f5ecaf5985bef675f2b9482814e9c6ffb60` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `deferred-release-001` | `ready` | `e13f31002cec9ba166428ac3d2c3e9a4eeee45dbcc0e49756329ec2e8814277f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `c08c4e3554e127e537523fb2c6e26120c86acd6f94fe1953ad5473355d63c2ea` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `4c447426395cc022827c4fa98dea614adaa39deda31b8e3ce9d430444644bf99` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `92eadf85530879648f7368ec23871fee045d6fef1ddb897fdd101a5d1be2ddd6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `159486244e250f12406f7021928d301d0c115d63c4c3c8c02d42383c5e651f5a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `81db80be200552d784e2f133beb797b215730bce17c7beabbd8882452287c05b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-update-001` | `ready` | `2979afe17bebb154861673c785e239e5628966e3922c6c4006268a675b246a9a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `24683f4e4c8d6a22bc310c8553a21bbea69b30d1b256983985e410ae39879100` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `dynamic-description-001` | `ready` | `8ace2c913cf5f57666649485c9b4e55b13182b613fb2a3b682cc006ddde640fd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `c1c77c1feba92dbea2931bd61c9b08c154450448011fb3f90d9cbe75eb95e8ce` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `acb4dd90a9ca9321bb3f84e4be6738806c4af67f64b511e2b2b3f8c2002f87b4` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encode-scalar-001` | `ready` | `a4d5e38f080ef75bfe3a1adc38f3e5e55cc337d8db024c9fbfe70b48e59f9e16` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `c4b683f1485ecd0bf5024b711e69dcaca3e78e8660f5d2fb4305b56af04bd622` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-query-001` | `draft` | `17abcfb02a2d62a89bcc904a563f69a99c8a4103762989f0ef9cc0a625b1d03d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `bfcefe6a9c21b8e5dc070919cede874cb400979a0b8549765ac489abdfff0c75` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `a3359bed1419ea04b84da94f5f9551accceb3c2d07cfbe43b174f50a93b5baeb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `4c5af5fc8b5db6bdfbda29eb2859f0d4667da5b4d3940ed4caefb7f7568f9d78` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `formatter-contract-001` | `draft` | `a84c2b7ac9c60653164aa69605a346e7656285f14faf5a04932e526123ebb401` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `fractional-band-001` | `ready` | `5d0d5c45365575aba7c517fefbd43c3b2e5623e23b2384a3560a1bddb877c945` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `b24cf5ca3772f3761e59999bb07615e853d34eb77f1da2304256dcb1afc0e897` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `68bf72b475661590b87c5bd54af74776dcedeaa9a292416db709751e6355ba9e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `8f56d9d6a05441d2269ff60ca5a047626aec26c8675082c593e28f62a8f1ec0c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `dc5af7e1abaf50eeb9a9ab0679cc58ff0c432bbd0d1d280b044f92eca7d1aeff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `initial-structs-001` | `ready` | `11322919cd17b3965c8c75c5e1735b0dddb4b6f27ab9852b586a6ee1e6c14691` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-values-001` | `ready` | `3efd7118c041b9829db73bc7f2d174f895bf73a51a2a1c02754274ea22450b2a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `kind-condition-001` | `ready` | `043f490975556b92b54d720835435493191603d018b851f27e7487b36bfce15d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `460c9c09ff13e3bfefc22efd847695f0f6bf99422c1dee48ea751e7c75e39b90` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `eff0308ce3593e20b7717686c286c362aea26914225522675b58cbcd96e84dc4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-presence-001` | `ready` | `d41f9e745fbef836b03160342b005b73c0a6edb0bb256952eebedd9aa7167692` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `4579c0556922eb92d723a9b2e1be09176b301323a6741325d7029c9bb20be26e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-block-001` | `ready` | `669644e5d4f83e8c35ce64edb97c84b5d4520524c3d62e06ae8613c8b9db0a56` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `nested-slot-count-001` | `ready` | `19a1c69ba690389975e10c947f6a84a0f01f737aaf193ecf556d67307dddb548` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `9a4273a5292a6d2e462287c842a4f9bb7902910da29647c8525baea6219759d2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `ac7fe21b43877e5272a0639d42ee6452184c17496bbcf5da76c009e70246531a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `4d4cc426b5eaa35415ee77f0d2ba939789541bc1b3def39b3c3e39e6a066e5db` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `a4519851c58c2c02d038a5a5fdd310bf73b27dc0b6b938f38707a8a8c3aa8bd2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `5d3b6dc09fcb582e9ada3ea13e54f6f4a1311beb48895c8c4ef6c119713f8fe3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `predicate-filter-001` | `ready` | `c34046110ebf32ee80ca18dd64c7747dfce2585ef70c573bc5507da4fecaa0a3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `1de66ec9d8ceb57391270eeec383d86999bf4cffde87263b7379872363c6dc29` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `fc278a6780cd85934f24146e4c8ca96a0e54875477f8fca6fca77f4a53d97c3b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `range-offset-001` | `ready` | `2106d091cd92eacd035ae44eaa34dc2714767b249dcc341b597f8809e85685c8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `af20194d7b5cea894cfcc50427b10e5811330b74e67a8612e465bc3fc1b34c4a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `d0f48bdfbd5ee37c5e3c1a8c9d6888a2f9953fc7e351b306c6b9db41135a1abb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `6cba991b8762d3a4d056a8bd47fdefd2702ee56f083676c439b1b376d3dc1bc2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `86929b8deaa310b9af811a0741f9d820db80dd606d9554e56b65dbdfe335e6d5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `ba449e6abe513d23cf555d5220a5951cf78e0a35ab7c11b4651dc374558178c6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `dca40081548b6668e4db0c3cf0867d23bf25671749ba6457d3a9e74faea9d0a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `8ad14e169f907c5851cb52669ab503b752eacdc570a6ec207460a5bfab17eeca` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `draft` | `4c4ca3c70afe266879ed5346879e5669f6dece91335f69682f99e7c9729605c5` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `4647fe6615c1e623a3635a576f759fe4f16fe8313afd9fc7410244bb398974df` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `45d1a0c239f8486d735ff53ba8ebeed2a5cf67e9cac5cabb8e48d95211619118` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `2cebad45e3e029c9be9588bf2d25619ecdd33833d0515c24be5353179745af1e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `beb0a532940414f1e3e9346f1c96e85698b711592381e636a88aeec0ee3dc15d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `b63a7739eb3c951e6ddf30185301b0d5bafdeea9edf9c6ca8a7bb2df906c3e4b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `4bbeece258dfc2c640a32acd000e34c31af060565b4cb20c8acb4f9352608f9c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `e25c441af4c22ff23586cd8f76b10a849103511c2b894d0450606dcdb58ee226` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `205b907bb0f236660e1afd84d5738525f4d75db4bfa2075cf55a096b358851e0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `eabcb5ca5c97a146a3d9ec6f91935ff926b539c031218c5acba0d1c3d041e3fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `540f4f37510d69bac06c10e48d0a89cbf950a80d636ef0589c2c194c4018de96` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `f905cc81be8f30ec62345d8ba058a9392315828c41cc21b822848f074590b25b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `slice-window-001` | `ready` | `a66220fae3390e098a975b7880e59a5df05007ee7f562aad1685ec92cd95bdc7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `447294f381f893e8d2516dd56b10b4f5beb52bef746c77effce88ae131b0e12f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `d829725c995bc5e200fb41b268c96c02df51cd4bda6a22778435191123eec267` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-record-tests-001` | `ready` | `37a332d190b3c0821900d76f03997d4528507eefaad418a27fa09f7683c56110` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `d1f68ac26844e550bf68942817281ebb59ee8810d75b60fc4b0454d77e5cebf6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `1a9e9eef7391e613a6ce3e3d490e68d21d63a74c299c633dbb42d5499fca9dbb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `6c21f009d7d76ebe385d98de380eb972dba63c4177fb5969251e946fa4fa9d60` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `token-values-001` | `ready` | `1aaeed7527bb5c4a5d5e2475c187b75d9bed17ef844d4f5bfce8f440161de3f3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `a4e25c070059ec340a01435a9c11935843f29719939645ded8a9d455072a9a2e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `6e8ecc31a0bdc2ad27ee95ac1d2a13446966e643f98d5c6e0e72912a319d8691` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `typed-line-001` | `ready` | `e8f1902d7cb82f9a2b377b4bba5f5a45a2d554f04d1e0befe998ceb5ca03b78f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `107bc09aeffede5853ab4fd8306a3ab185e559fda9ad5025daceeeeb1fc90985` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `bf128b650d9c8f43c77742f2b439af4afe774e63b7ce0bac32a69febf9a40817` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `021919f52f0371f783c154ec557ba0b9290c5ca9f3c0eeae7547bfddcb4a1ba8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `bc0c79923105beaffcd7f779037447c0100fd29ffc146f00a29eac3d2ac23388` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `5aed6135299bc4b811ab292a5fad6581d922249e5cbc85263a619f78bf447e14` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `e7777788647719dda216d9d6179603191f487ed6c9a3f543b3d17ed03b117d51` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `a3b58ec8a4619190599baa337ef111e8672a7632c2fea02ed9555d046cd73aeb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-sum-001` | `ready` | `93c3d07556ec11789574e2536e1be4cf127b614f90c60b0187da9276baac0112` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `28060ef6f63986a0fd930c989dd21d2f089894a79a87f58fa1a634e5fc6ab3d0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `9c6eb03ec167c74b6f2f536d6f4edfe9ad5be87a74a9e6df31f00715eedbecfb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `9ee2fe93f5c9e10dfb0fd01bdddce3440945a264cc6ea894d7269d6a8becc197` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `37cd75086385e4c0aeb5837203a58a1181540ccddcabe3141014ee7b801834c4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `f363bee5ef47023e27f67f08fb6b9bc61f58921651d4a0b26ab81deb7f8570de` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `9a8e5a91db311fabe0ed3b6063d5b58d5e10fe39936c7fff5547fc55bd5b801a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `fd51358857af2a0711e30fc320be40c17841548c90107ae990ec7527e7fddda9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
