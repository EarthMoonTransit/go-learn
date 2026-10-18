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
| `access-logic-001` | `ready` | `a8596dc170ef98a74ecfb9b20321aa0e6e9e2ed70e2b3f4678fe1fd5fe7b06ca` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `ff6a71ebc8fde75a0fa79cc2328269b0be407f8a5ffb662f042980400d7cfbed` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `860fd53daa945f3b59d75bcee613d4dcd148429a834be340c932b0547b4ddfe8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `7ec729706d8563e26c454abc148a179142b99e677db7fec3964f269c8349e933` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `bf1c57e62e0003a9762364410d190d6c28d00a067d79f3eaa2d68aa60cc559ac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `f400ec317fb8bf2dc67725621305e85e6a748ca2cc3b68a0c25faf891ed82c79` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `dbe03e8c45143fcf710db261b59010232aaeb1a39397bdc81c8513fdc92e59bd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `7407fe84ade5a78589d1d957d2282183321f5f2486f30a5d79678c39e5cbc351` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `34e32a6d0f84a1959ff1951f5ee37e2dcfbe6ffa00e41ad876dd8e92f17d6d34` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `c3a60704b8bebfdcde98e7027be3d62946672494db0819f3165c0760990d1fc1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `8589f1587f10556213a84620fde82774fb559011db549fa99b2c16b0793a93f3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `ffd3ab7febdd9aef5b8724aee302eb9cfb53569a308c23d347e956a2e869ef35` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `fa774145dd9dd53461af1afab859c76a899b4fbb977952fe4837b2842a7af658` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `b828ae87706f245a1af2ce6b3b35fb9100c4f542d963e4d91be9cf0974efcdb7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `5752fd92e22c9f6c990075381fff5a586c87c347587d1b3ca43909286060fcfd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `2ed0b7f6c422b18a8af484c7c6608165e1a00504cc78d0d6d9fa70a0f9f3db4a` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `cancelable-forward-001` | `ready` | `89640b916e36596e8b839153a65b98a47dedc9793375445468c62f7c3e73d039` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `capture-panic-001` | `ready` | `11d6f087512d091e8bde31e6ae880c39d101a8920104107ec8a01879295b0874` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `captured-method-values-001` | `ready` | `c292491e09989b4b1a08a47cbf114e86be5fd6074d0a423026dc1c7cca0a580d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `category-closure-tests-001` | `ready` | `57b087a6cf7d3ccfa00e11b70600f59a98f8badbd3a7c1d47020044702d37330` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `a5388ebd715f903a0e8c80a9292a82cf077a31dea71b9203386aa6bd6bbecbfe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `aa4a255976f53206655b4960c2f59e8e661e098bbb12a2aae67d421ef202041d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `c50b331263e2fcf1220f4beb7ad07d5b3602cdc7ee56136c63c829bab3e91e30` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `a5145f6c6981d7fb2cce8d9a4cecb29b3f083cfa3bca946590f7bd60dbc8b49e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `7628cf51b2ef007817e1c09b6557198aeba6ee2940f8a5e6c3e908f29cea893e` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `a54375049c68a545ccf42780dd822b1a7d96321c556ff4324fa4b9b026bc228b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `composite-key-table-001` | `ready` | `91aa4406a47ebf9e72734c3b7cc729ea3d8b27cb08721b941b2dbf1b3be94f8d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `b58e6d8ba12bbb3e4846ef69f7e4b62b25129f2d5f8adcf2adc1cd4c437d3bf9` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `07653acf9b6ea4c5e0f1bc047b8b83cffeefeb298edc85ab72f66d8d01c15b56` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `0d7c734e027d241e2b66c0f7b561b9b1cad6c1200006ef48ad61f4f0880decab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `b170bd5965110393e86abd5c6d48dceb51e04ffee0a285bd76e6f7f51a248f28` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `61ce38b6f03b36a6e49771102e287fc08f75ec8a24e087a2acaae06745d01bfc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `65fa170d93af6cc0a7fe393596075c37ddd9cc51214435482fc078fa9d599730` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `8d487b2c3b3e14854c432d73b33b4cae428f3d79709dcc96113bb814657f3d3f` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `deac529631d598362c4cbec24c379ae531f9519c16c2d530481b8b83fffeda0d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `ddaf32e98fa283189cccc37ed82907fdee8aff3fa9fbed9e3b7ea6f80e44d406` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defer-trace-001` | `ready` | `546c3dbdffa54727fb54699f04d1f27f84889bcd7b4f2ebdd66863b31472a8ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `deferred-release-001` | `ready` | `e22c1ea97c1bf239fab888c5c52a2948a466ad5674220d863f97d1a900d7f851` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `a15dc6c4aeda6f249627bfa26b7676a3b241d9ff5b6a9841b5230105f28327d8` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `56e9b77be78c38228349df0a6ae2df39ed490fab86ab3d17266fa0357faeec95` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `66b57ebf424e3a50a81ce1ffeb46a473c52c33a4faac4150f612d7a1d238e131` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `7498fd77ad9fd863cfabfb800ef572c635abd3907edecbd24edd3795cb6e3bbd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `7c373070922acd9a1e533c984c9e314097ccfd8377d9b113c666b36830eafd3c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `697907dbd35d146b00df4e9381e8e35e5aaade0fc0d907e8db870ea465acc4bf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `2dada55d9eae9090991bc2cfe2d2786d7dd3a19bd3e6dc25613fa3463f76c323` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `350e627767751c206501c2af20ccd1f8d2b064115e4174cb49c06991758e27b9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `f17ee882a2113e9e6b26aed3590780d4105c0a507b5c4b8f0fd8c5960250f770` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `4f44af7bfdd319f6fb85929b9208237c2c96ce5889686fda13717a1d9cbbefa2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `917ec2cb3af3184e5db31b4a0c1a7ee7d4412a8dcdc74463912b5962a5695c06` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `0cb0372eb13f240c9668ecc86272c1a216ce19dcfd1c5e6d1043ddcdb85f6a94` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `2b49e5b7c3426995dcb53f684e5ce2187ea5d246f8bc1741ffc136917c89aa1d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `0e601ed27ac6e2742e5c011cdd0d1e8b85efd024ecce2d53dbaef9939dec7519` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `0633d0a760dc5822e0f34f84c09897052c5a5aa1e3314bc4c2c3b83e5e58b18b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `e19e838878d9c2283cf11b33bdedc3274319894757be6be7071156855d82a249` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `6d8bdbbc8d63f0573c047daedb098d10f82ad0d5eba91a77faf6e77da162a6b9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `3d1a1d5359969ab25a958808b6092d962a8ae4d21d37d8cfc8d289e94ed89df0` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `69e195a959edf16f02ebaba6356895f464970ec83b7bf862420d0415e543eba1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `a38735d3ce61ba9f2aa68d66d0d6c6bc2fcf7ab59506ea73a9136049e71df88a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `4335c00b983e716718ec07b89521e4f76fa22fc102c4fbc8ceb01820ad0f54e7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `a8c3d4dd1bd6ac3c5a7c78685e6cb23bd8b4065667328ada555588dc49518ae7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `757bd437e3314d05609283b3d82641e0925bcf25c367b165139bfa00329e7572` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `c12d2f1dfccf9856833beb01991f6b37d40c47adec9b2d24b6ffdb3155fcda6c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `4293269cc2ca195f8c86f7624e3d139b07fed61b386025018b79aa3ff68b331e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `f2767f2fb0405d124eb83448cb2eaa3998945ea30873c0299202a9cca51c79f6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `7474bf0217a51d50e06a63407a68eb0deb3e2649d0c5b4f94fa61db2c489252b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `ready` | `d8868f7f99951514ad17bf801a1d10ecb7ea60e22a6563d780fd292dc2f0c3fa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `fractional-band-001` | `ready` | `553d6284186626048a26493fc254ab7797e8e655c528c36ca7bea3d19ea83e6a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gated-worker-arguments-001` | `ready` | `ff5bcc6346a71c784df99f1226ea0fcad87cd82dd0dd27932e16270ab55956b4` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `gcd-reduction-001` | `ready` | `cb42ae4de1e3a60ec9ddc0729c87b01554759b2651ae081f40d47027072bea75` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `65d1c003c50da53413ee01878b06ec436100cdda5fe7e509e131b25876d46554` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `a550381e20c2fe3f8c56847e7edf8b8af2112dc041edc1ddb05ce95de9f613c9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `82b8882b85dc86d104a2f1e6bea8c6844f283e0480860f826103b51cc641b111` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `1a7a90f0f9f80fc537aa7fff98f9f8c1205ba821aac38d64b7ea3bd143156b2f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `046068d686ae7a4402d82dfec3ce987042f21059a2ed4eaf16fee8099ec295e1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `96eb2e523460d74dc2b4dbfc43746f55e5cca94a6c4708c83efd510b45918596` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `539c920b8af5459c5ac0e46cf1491dd0b1365203d3a5adb77ee18eaa5250db9f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `0a8018cea003d2f1e85ad8c5b8810a54e87a3b20f635269485912b0c9cfcba44` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `indexed-job-results-001` | `ready` | `af68f9125146fdd8abc25a8d2031669b197b6fd34b6904c266efa1ce173465b7` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `initial-structs-001` | `ready` | `345e288002f38c0fcec4c1762b2740c699178d31bbe0fa69d19b92869c6aed60` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `8a4a5464b70b8e29ccc74af31d4dcfd23e5bb0f70581cbca500aa9dcd307d281` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `952fc257ab1ca271c3c749fd4b7b86e5107d7de3615f5e00d0b8b73cb36ef78f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `feb481da695d672be8a61c56376b5a0d7a4d81ca52ffe0d41c95134c71a8a8c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `e9860620d2ddf2d1287aa1cbeeb3479fffdca057368ca98001bde16108deacb6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `581cd5ca614111aca4513581fea3ba59f739221bb0bbc4070b43ad2af0bf4cc6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `launch-provided-job-001` | `ready` | `5ffa7f04c765ec13646fb3b07620b67678938b6d480b62b202596550c7703812` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `lazy-probe-001` | `ready` | `e0502eaac7b8d9020c3315ff026831c33539615565d05684ce2f619137bcbe57` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `6e7b4226a863f55cb7e85d8f45fdc05f322ec228dd7c91b6645133791e919215` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `9e9ba368ad09333328fc1af1d404d30317de12a14678fecf17c61580fe30bfef` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `cc99c579a08b6fa7a9d4f611a56f012d814f4a2a6bfd7ab4b2b30df9459c8967` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `32812969e9b524b6221db918bffc83e20bdd45d24696d9aef8125fb48b90c57f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `f570fef0c331c5d158ae0572eecfef535269a79ad9a9fe9c4d66a86919c36834` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `mapped-records-001` | `ready` | `09c952dae45aa92064eff959ad92067050b3ebc2813f42e9c97fc9fc543276dd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `merge-two-streams-001` | `ready` | `066091df58cbf0d538c1edc7c11f6f284f93a24a5af2532507cbec755c6282eb` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `month-enumeration-001` | `ready` | `7309326ee4c94776b834c13f55c74bf84565ff07659983e4f3b444d9b15e7a89` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `eacfc2e90dd9cc38a50c3be676e4cb33983886b04251fc45568aa186c459d299` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `bf0ce35a76618f4aafc620b1ebb60bf57d65703c63295530337122c837ef9e6a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `fcf215edac7bc97ec3aaea0220f51eb566fd263ad72bed8542dd56ecbe454af7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `e51fddd16459dcb8ef57d9765869558ec5356fce6fe62bd6631123df19154450` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `2c3a5b498db5a8eaddfcbc5163018fca58706cf5579523a65bf66b6aeb9200e5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `named-result-adjustment-001` | `ready` | `70ff3547e9998e14b2ec63565f9b315f9647cdeb95ba870018e6e3e84f938bf2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `nested-slot-count-001` | `ready` | `14ec551aca28d93a19d3788405780035dd645e5294ede8ed884e59ac2cd5fc4e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `cc1a50413503fd8216d2782283d2b4e731492ad7f45b641557bb9767576e920c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `038a2a2eff6518a3f24ac52f0c1c5b1e582e7694935b61e405ae382bc5a43cbc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `82754c5d100c62b07f135f96148d9a414a651dbb65d27034f70ea93c9a74ec73` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `3af4c9a9ad661f1859821964a3c737166d4d2d64ed286b7b948cd82714e911cb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `87ae774b41e138208ad2a960fd9fa0f9fec0abff12136237dad5fe91d9d99d72` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `d5a7d927ccbe096a4c1fb1ef417f3e5fd2fe6973b79435e315ccbb134a0a86bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `e5868e5692518b1a7086d6cb861699bf51cbc21508f6fe30ae192e3323739279` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `30f26a1dcab4fe03bf6a064b7de3f0c4e6e5cce32f2c592a78a867734490f986` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `020a974d4536df101577ceb64aac6ac75a2bd2dbac79eb18e73c8c9ff1326661` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `3fa02bbc9790688c18090e9d66d4eb1dce31cd759533c19485f0621264272a3d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `bf6555230af6560153f6f570f3ebeb79a6732acd0369a1c2ec8299d75a0859e6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `048b3827a035251a6023037c2592ba5b3b4ee96149b88b57f8774119864f7b36` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `c757b2134b7ce15368dd499e93bbf4de9a41f713dd10cd9ce38f80834cdd6a53` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `7de73784935bccecd1c44cd95bb55b3893c37a60605e906554020f387fc8815f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `7e351bce6854ea525151b1bc2d5f3847e4c433fb27618ef54e46043899036432` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `418b3b8739e0adf9933ca406af8b997c7b2f3b9d8a21d4f8d77994f766e3536a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `86ee05ac5ef7d0632d2b5567e5b9b0ebafadc8c84a0b5c4f8c797ce6a11689ff` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `a6e7fc3eaf0e235428e834a414e3ba8833e27ffecdd12d7787dc84376a66c5c5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `91d1c06ad0afa38566d58fae304b68579a7e96e51d0a2a697f21d2fdfa4dbf1c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `9708cea0ba80e578cc66313bbf0912c075d0e18dc28e8a35aecc8d2ebda9d39a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `b2d15a19c6b149986ea016386e700d46faffb06ce4a57f1779f93b3ff284fb91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `418c922facb14238b61db8364e101fc54cc27a89ecaf1a5178122ee666911b1b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `2cafd3838b029673eb3552702994a833f97a878954ceecf38f6b622e9c527abf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `15d73907131c9e7ac7689ee891a21e3b183d3b6c829ff63fdd39f76d71aa8762` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `f0429ae33bceefaa37a8b3a3b7966d6bc810faaf8fea65a2fec4a5b86d012c1c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `f29e6e3312226732ceec1c1102d8f767b0750c2d14b0feab54094d030ba8813d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `replace-pending-value-001` | `ready` | `d54f839c3b3f6521774279e870106485ae28a1dfc3264c63bf5e989495fd7b29` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `48ed3bcefae85a284ca1c97e648a3d668f5fff31413a9abcbd4cb3d83d004d4a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `ready` | `dd1ca2a8fb6af1c22e8be91953c78fbd8a9313855bb326a2ce836819c0f667e9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `b264edb570bd9425dade7433f9ed8be81efb2a20b88dc78719e17cec00cccd92` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `87f7d92589c8683bab9030fb5ea23db9494bd57f4e2c7d937464faa1256ab6a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `f112c72d367ee4c8856df172775bdd70c398e1de4e1731ab9e1cc0b9de4f57b2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `b10d8904fb993f2eeaa3b4ccd061032f7b8bee9e6350661ee0cb79c93ea12f7c` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `9ea4443e37c33cac28f46f90e583f393d6ded189e303e05bfd1b655728fe36a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `b7a946aee3c0e4897b9b417fabb9dc2026975661fc4928c95d5cdea50932f3f4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `c0f54955cda6fd81f8dab021c21178b48ee2c3244d6fb143f04a38f8e230a85f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `9a8f062c01192668680bd9faac30b3d49dea3f0e22603882c20361ec1bf687e3` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `72eb1dbefed1e5d960955ea035c9d8f2d7580a35987418363fbf8535625cb80f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `549814dfd3b86cd7cd8630c19c976afa4d69b260a77c236a67d31dc83dcc21e2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `dd9835e897ab260e2db7cbf4ab4206293989aabcf7525fca1d573e300060c3dd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `f0fd3a722aa50fad9f5f5f091267d9d2d8de83f4f2218a531fb0bcaae2e27174` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `32d011611578b0ca83fc3ded1a5c51d6d7a3b0bb9e40fb58fa67f30f49333c5f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `f4f0750ced49f1b131436df60a88433ea7cd6c29fa826ca289e55ef1db6bb083` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `357da4eb555c2ba36af877a5a7428d2d124267d85443458b83bde90a9fbc3643` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `d3baf3582406d1d82356550ad79bee75db1384bffa03c547e5c9f2f9befd2d5e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `630d56abbafe7e832f6cbc3f02fc48674a0524c08e8804dbcab42a3fd4bc5f41` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `dd1b09de038b78ca5f3011b6d302f3ea425f67377c537aa2e6101289efc97d35` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `32323ed681574088d0212316576e5d08b7e6cf8044cd416adc4fdf1905d5d9b7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `15d147a8dc8fb7e5e3f2e7dec86529f3680e79b860e98465e11065ee79d7dcc1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `be26c34dfe5b7ca98986cb42ae3ff2cc62e4a6408f06d8aae49aaa66be840895` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `2fab764beefca6c3c0cba66f61b83da2589524ada4ebef096b85e17d081eff5c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `6e4561a23808dd4fa1708df213a6f2cae81f31117774510c605a0a7a93eb1bc7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `a8e47f6edc98aa1bf42c5e0dbe95e648e529976562b4b15d35a95df60c766d00` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `timeout-operation-001` | `ready` | `a2bbca3361527dee75b4e44543d10fb2c3ddff9ac538822e0e46916ef0eabb28` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `token-values-001` | `ready` | `5f2deb0a34d50867b46ed926f36052a15772a8fcbd077184a363319a1a2cf2f1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `dc3b3c3783c5838f40678587f722b78d4c392cc0d4079e56e382133886d1e7f2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `99d276e7a7c59a845e50fff46cd1b18d59cc9796e5942137df1179a64edd55a7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `7255d22b762fb2cec88bf9ef81f92220158d01616ed4b5dc2b598924cfec34c2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `29d65a822c916e8e59406f88d82122bf0559f14f2544f89c257638ef81ef96d3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `8e869013d0ac831d2a7099b1ca81f75fb7c59b0d2ddded67534962f29ba0d88b` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `5f86a236375fb201744e4c57cf462def934e78ad9bc3c0650cffc6e5d80f75d2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `b3a4c82b05ab38a4605688deb8e82d316dba6b6bbbfeb239d3a0839f5d0226d7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `d645bac66096b1ed636b62e1787204deba133e7ba043d189717537e87d7d8c55` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `a1ec9071f0c649617901356b40bd1d066a498591bf3fc3ebafafa4b9bd400784` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `9b2368d20e1ee910a08f279f3e5fe4dc1d7a485586def69134bcf7765282de8d` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `0ec6e50e74698c36058bd567d96aeb716de3911e53e4cf31d85ddb2e9b492dd8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `7173fccc54555f387c70d87301a5786f6b5580cdb9e492247ae3c579ffb92e29` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `12a2faabc0910c7c166ee315fee7c7c09e1d4c847a2ad907edbd5ee71070b2bb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `a302fa95d0fe9be61b1fd43fedad37cb782d7dde3ac0f631823d3fa913923a11` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `1af717644d9ad46814d58e9d047370c4f971f47292525dd4a4d5c79e9971a303` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `31cc48ca3d64208bcef0d459fb9e21bac53ca8de41667ae2e3c2f9ca2ba5fe79` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `7614a4ba7f4e0f4394d68f16e6bbd49bc2a6628820a8b9f420382303e76e36db` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `3b8c262fd5a4c14bff7c92c3de3836683d544751d24d21492ca2c4344767b980` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `d31699fa1efd5de0ca20fceac842dfd301255da6c1d866b574ac5d37a259e2e1` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `72903c4478e4158d620d7c83e6eb9e89fa670826f8ad8e780ec024bb57337757` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `e86a1cde1f5c28a5288c0dfa0830ee9309bce1a82546c303ff1374b1c5b93a13` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `e996af8c1d86327d3a9bc354c05ab7432d25026b57b8dffb4c97c2a39a21763f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `ecf1e2afda96163dd61e48373619d5c61b6b52fef0c9d48c31597e941f60eb90` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
