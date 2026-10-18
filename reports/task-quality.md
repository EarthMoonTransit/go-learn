# Task quality report

Generated deterministically by `task-report`; it contains no timestamps.

Source check contract: `task-check/v4`

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
| `access-logic-001` | `ready` | `672084b6d9a5d4a53e58de9b5956492eb385840dbc1f7884e98a290a014373bd` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `afternoon-window-001` | `ready` | `2ebe804f0c63842752795b2d3148386f968d2f08290d8dea6ec598bfbdb37d32` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `allocated-counter-001` | `ready` | `4802ec80e2d434a59eee2db71962297a1d447c100ac4e782e1d10bc69b3c6490` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `alternate-correction-001` | `ready` | `41a95b0f7001ddeacbc8d0de553801a911d53a750c6d9aa740c5f8aa6d229dd0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `append-value-001` | `ready` | `728c343af2af3c2b9f0cd1d2a4d0e2998b14ffc96925908a8498f8153217d537` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-001` | `ready` | `1028276920788307601ac31958cc89956688e2ac0b4564360f018804257cf94f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `array-copy-address-tests-001` | `ready` | `bfb9ed5514cbe646641727e8d95d22af3ae6cd4a813193724b2755f6003b5831` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `binary-record-tests-001` | `ready` | `1126ccf47687e89d311b8a8ed81eee7e2310c4fb4036fe031650b0ec2c75e319` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `boolean-policies-001` | `ready` | `7ad7299917977b83604ff164dc5ed9d814395cbcb383f1a1f2230e7d2ff22dc0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-copy-001` | `ready` | `c86bb15bf7cd03779236f52f8862fe12b2e30d53aad2d79c72544f236dfd8cf9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `bounded-lru-cache-001` | `ready` | `7f947edded693ee2e343a7889243b2d4c9ec7b742a59f7fc28f8e7f62ee4202c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 20/20 | Phase 3A |
| `bounded-slice-view-001` | `ready` | `5ba9465880f4b7aae43bda44b09be13d96303a042c80f7980d09c712c84fcb69` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `bounded-value-001` | `ready` | `e03b7d836558c358c719ac73587807193137cc19f7323a26b087cfaa60b23628` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `buffered-flush-001` | `ready` | `2ea44756a952344e887c9c25ea8bc613d755e602df9b4fdda3edd77d64300ca3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `byte-window-tests-001` | `ready` | `2b31edbc2f88db309fb1751ef9289275c50efc714927a845464482004498f652` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cancel-worker-001` | `ready` | `93af2f676dad244622c7400bc83cd16a6f9c254eeb30458a95154750aa24eae5` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `cancelable-forward-001` | `ready` | `609022751aaa799728d8451d26d0c67bf3d7ebc7c21e6d3081d76ae874d6e54c` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 9/9 | Phase 3A |
| `capture-panic-001` | `ready` | `439b8fcccb52c49ebf334dce1aa4f59dae0dc99476a3acbd141e74a75451f4d9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `captured-method-values-001` | `ready` | `1705e321db1d5644ed2e3f2cab62dca18d380ad89e26d2ce1d1d64abcb653059` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `category-closure-tests-001` | `ready` | `8514b90c541b240f92b86f2d8c45919126da7a837ead41b508b4b61f00e5fecf` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `category-total-001` | `ready` | `36b4d33a3e34c6590836607dfa56e05130942e85de8d6967e028893b3c4531f5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `checked-withdrawal-001` | `ready` | `7e7afccffc8443b6729e7e7482605e96d910750669a2b43978446654efb46c7d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `clamp-tests-001` | `ready` | `c3ec0a861d52d3a1e9b2b54ec50af86d657ec6d9033c8249a0cee2a805626a9e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `cleanup-order-001` | `ready` | `6156b8e39bee289b61301190c8174cd05ef1a485a3f320f46c0dd5cd3b257f5e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `closed-receive-001` | `ready` | `dba7cec705090e35ac76cb893216380d5f83ef4487782a685c51980d72160555` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `column-occupancy-001` | `ready` | `fe1eff820e4c1dab5da6ba5c4e94b757ca531d24304f9828b6a4d4f3cebeb95c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `composite-key-table-001` | `ready` | `a92c85a21572e716923ea7bab7d7c76fe4ed0e2e52a673b94f6b4053cd2d8bd4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `concurrent-counter-read-001` | `ready` | `9896e1b181f7148e7335b4f669e5895860ce9cec71a42528071900618ebd684f` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `conditional-binding-001` | `ready` | `6ff06b137f71f34a921cd67ba7fb96d305bd9d51098eb38739304d8f86424b28` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `constant-kinds-001` | `ready` | `a14310cd4ff70808dab5f32abe4beb4e2016e20f101c089c0e91f31a3c7254dc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `constant-value-tests-001` | `ready` | `f5467aadd5c8087d527ab7b28775d4e1b9a051154e79a687b5d360a1dcaaf818` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `copy-final-element-001` | `ready` | `0313c9dfdd884aabe134576ac5db08accfaae6c3d8d796675276e765add26314` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `counter-read-tests-001` | `ready` | `cd228e8ee912fb0f6bb4d7df08b10969d08877f6447efd55d639dcdc311e9dfb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `counter-receiver-001` | `draft` | `1db797d7ad51197056943b81717369ee4b990d1877f5da387ec19a9c046bd0a8` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `decision-table-tests-001` | `ready` | `903a6eeebbc57164a0ec41647f4d4341ab2518340ff6ff9b2bf3a695b72f14a1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `decode-prefix-001` | `ready` | `f86cd158fec3ed05f98490f2243cf11846bd1b5af39faf45009995e12953de5b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defer-trace-001` | `ready` | `dd4daa5f5873d4e80a0a9e9b43ff7584d415d1fb2fe0df804d9a85d1b10ea182` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `deferred-release-001` | `ready` | `8074789f85bba9ba65ec17c834c0f4c731efff24ad3bf18d0d60904987a8649c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `defined-types-001` | `ready` | `6883213c43653b261ee3614f48ee9a4ed43222263d711d4f52a0b99fa8873131` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `deposit-state-tests-001` | `ready` | `670c5a5d80de589016d05465117f84cd8d4aa481eaa1c706937b69ca93399052` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `detached-append-001` | `ready` | `c42bf24a34e47bcb588bc0542fafb33c0698273e7cac794f3483945f73ec7e5c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `detached-prefix-001` | `ready` | `daea074810fc48603c53dfe8103798f632703ec330aafddda762ba04c0ba2f91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `diagnostic-helper-001` | `ready` | `c10b2c6cbefd8bd72c18a01650bc948eddbcff21c4f206adbf200dc2969b3e84` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `dictionary-add-tests-001` | `ready` | `cd8c7b97bbe604c6380b4376c148330b77a4d3bd0e0f669f2c532120c04739e9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `dictionary-delete-001` | `ready` | `6e2e1cad724368e10786906938cf58792659bb2339806b61a99c27edf0dc29d8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dictionary-update-001` | `ready` | `86960201e50ff0405ad4dc4775aa5bc29cdd1cb6a532f96b34d8c34153cd8c6a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `division-results-001` | `ready` | `99dc41197b526d1abe661350e7af549eb13f1d74d2baeacbe94b6b73669403b9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `duration-rounding-001` | `ready` | `2e244dd2d79a5fb43f8ac1b02cfdf493d6a829368827c94d63256046840f2cea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `dynamic-description-001` | `ready` | `e13318e7f1677c285773c047f0936c08715500db3fb63d2ed90d5d3b0f36ff91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `edit-rune-copy-001` | `ready` | `22f042ccf8803de9ab47523d9b7be29ff4adfd052bd849a026289766caba2870` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `effective-rate-001` | `draft` | `89f3d83bddc12b51a25251fbe1ba512e1d28f55d2baf719aed1cc958f6c81dda` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `embedded-cache-tests-001` | `ready` | `3d78258ab8ea19434205a27ddaebf41884b495c9a84de8eb084a5fb59d0113ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `encode-scalar-001` | `ready` | `8be3fd115f9bbffce34686000a0bfecb1951c95a2e3a6d0630bd1ec20adb29fc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `encoded-width-001` | `ready` | `d6a090ed96657b83d840f06e72e2e06209653405ac1a6440179a2fdb69bc6608` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-chain-match-001` | `ready` | `18a0b83f15cb0077cc56bd5c881070ce1b7f3c8df2ebbaa9db0c0a661a4f0856` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `error-chain-query-001` | `draft` | `5937d9a11aac43c271581137f65aa11bb4aab7fa8dd9e663842b69406c134517` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `error-context-wrap-001` | `ready` | `e0e5280c1f9ecf9786554b6759cb1f0a32c1a25eac5132c518a134dfd3e37b0a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `error-detail-tests-001` | `ready` | `b9d94f3033171c690ec80c710cf7c1125bdd69261f318b169b2ca3f5948e5d91` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `exact-dynamic-number-001` | `ready` | `e3b542b45c3b3149b469568bde587d72098d71e7666fa51dab4e1d3efb77a22d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exact-prefix-001` | `ready` | `ee0265f4afec860da9d71dabf21641bdb8e2bffdd430503d5888ee5d958fd222` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `exclusive-loop-001` | `ready` | `0e8e7cf882ea05f5b89741d17ffd9b15fde44d13851b223b706a71fd9201fbe6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `external-window-tests-001` | `ready` | `08ace20726aecfc016c39c72da4e384b90b55efb1999f7b77b37df3e620888aa` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `filesystem-record-tests-001` | `ready` | `3e3450b27c7bf6a3a76de9be8d8c96617f9034ca7ce0a8baca2c9985bb5edbe6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `floating-categories-001` | `ready` | `46f0aea9234b71c77047ff823a8be45dd0a7a0ad2329ddb35869d0212a527829` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `formatted-assignment-001` | `ready` | `3295994731c414641fd3612faf63e37ce8e07aa5af6525369a98673058b15fa8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `formatter-contract-001` | `ready` | `8e7f67f7dbdfda55d269544e48836ce1b691eb4a0477355cb3996f6bfafe6f49` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `fractional-band-001` | `ready` | `af7df8e29948f5861e9360449efc7f872efb91601382715c60c49701d52a13a6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `gated-worker-arguments-001` | `ready` | `7d383bea6920adc03a2a513c0f1e68b503b9532a56615c31d1c54859f6bcef68` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 15/15 | Phase 3A |
| `gcd-reduction-001` | `ready` | `70866ac67fe065eefc71846bb4076d8701b717256a89974d5c364a65c7dbb694` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `generic-find-tests-001` | `ready` | `31feab102c51326ae33f046b0baec8452904428248bc0c6f9b0a7be11116bac4` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `generic-reduce-001` | `ready` | `792823ebd2b29125660f5b20e20787a9c60040097b4bea70b8489ae745f76a33` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `generic-stack-001` | `ready` | `b5a434d578e9cec3c9ea2dea4bd74ab95d665244a314172cdd36f58195afdcb1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `greeter-dispatch-tests-001` | `ready` | `2ae82f253474e73811d7a2865455b4e6e159a83131a37edf5962f91803087992` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `grouped-unit-cost-001` | `ready` | `5f331b9ec0d56c8d544d538c180f8b96b51920118d44c506e57f3d56e932906d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `guarded-item-001` | `ready` | `2543a8a4bf3063fb1d30fd0208bcc23d4d721031d13ff52892cf302497d32db7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-policies-001` | `ready` | `5b6efd283506b71ecd4efb218e8aab275ae28f6649133aac9df701e947269b78` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `independent-scaled-copy-001` | `ready` | `eb83840bf31fc116e0f3dad98c5b8d3d6277ef8a7476f37c51efe6c0a49fe79a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `indexed-job-results-001` | `ready` | `712d84f38d3577ea7a950674a7aa7eeecee7c487f9257098ecf6a356308afd53` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `initial-structs-001` | `ready` | `6cc90cf60d9eea2b461750c96be04a803c08a77bd4ef462b16dd6e2b1da1d1f5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `initialization-values-001` | `ready` | `0cf6e3cd9f9c565281cc8b28cd86ff11ab3005dea3a078cdc0344cdd8a8b8091` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `join-all-tests-001` | `ready` | `9f281f187c01a5c1adcbf4c6dc3245aed687117af6d2f5bc4b245ea504fbbaaa` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `join-values-001` | `ready` | `47f331174637adcb94e638291de86c5ca5cbd8ec2f7c16c3c492880eb9d9d6c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `json-field-shape-001` | `ready` | `a55594d6acf31243d400ed86c7e2bfc6732217deb59294c98af0659e1cec9d02` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `kind-condition-001` | `ready` | `00a0fa7a47575764043dcbed221db70a424e6cb349e149c2d93cc535e5989b21` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `launch-provided-job-001` | `ready` | `297176267070bad64c6cddd577851f36cbd45d7f76d86ac81711de4c3a824681` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `lazy-probe-001` | `ready` | `987e9ae74f6c2013d2b5a53e52c163bbf62bfea84a33d81c6338046362caa975` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `lexical-choice-tests-001` | `ready` | `af1cbcc4bfb136e28adef85dc2840f3d091d0245e912020609f966f10225658a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `literal-bindings-001` | `ready` | `40cdcc31cb528e673c07f5d59892ec9c4b528f759a7bbe64c94b8258759039a8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-entry-tests-001` | `ready` | `706ecbbc89d2ca9c59b05f40a06d18e559c4f889b95e7e202ae12d43613406ee` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `map-presence-001` | `ready` | `6c83fdb390b19ffee834faa312dab04025103c7c877cc64bb87832ccbc506535` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `map-take-001` | `ready` | `f65996a0d04fad4d14bf5caf6291475499584db48b8c2269d658cbf4174c8240` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `mapped-records-001` | `ready` | `f9d29ae3695b4360c6b2d86eb8221869ca6e2882f6c3244873dae2702806a7f5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 16/16 | Phase 3A |
| `merge-two-streams-001` | `ready` | `223c2b610419f83ecf92dcf96566fac2e37e6571697835c6f1a2969274b06c9c` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 7/7 | Phase 3A |
| `month-enumeration-001` | `ready` | `6e962beef6ca4ec5f7edddbb7da2c996abd5a35edc2aecbee9b6c317887d9909` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `mutex-counter-001` | `ready` | `b7ba65f69503842edd8aa5ae157c82aa59fd2cc4449bec289095844d3c9174c8` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `named-anonymous-struct-001` | `ready` | `e008544dbd77a25396a42cad2b4d42d828ad8e48129775d28ec4ad9b8258d1e5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-block-001` | `ready` | `a4d0f45e4b794afcbb48d97005558b11ec6521d6bcd6ecaa85d16e8977cea550` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `named-cases-001` | `ready` | `fc94531e329c95296867e44a9aaf02f2af71d0e547c7f45205f177bb8c5d18f6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `named-numeric-constraint-001` | `ready` | `1450de136f480f9dd6ab168996ff903ee04b77734de1908d281dd37edc5c8e22` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `named-result-adjustment-001` | `ready` | `8a7f45a219837f4c753724ea275a11e56d6d916840001f4ce49988c9f2f628e6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `nested-slot-count-001` | `ready` | `c2c75e615ce989808d035946493d340a4daea15066ef80f59ecda3dd1136ea09` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `nested-stringers-001` | `ready` | `92e448a2be8f9875d4b4205d12b18d9549792c96344750dea39be373b6aa526e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `number-box-tests-001` | `ready` | `2443ae63daabb02e2a67413173bde36385144c94ba6e9ebf1a88c54488559735` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `ordered-error-propagation-001` | `ready` | `94bdf260cd474d54ab88a4266d2216e47b0cf85e8cb4a97b69fcaaec74e5427e` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `overlapping-shift-tests-001` | `ready` | `220641931ca2ddfb3a72abf1a8bfaff9ca03a3da1269196c2e719166306273ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `owned-values-001` | `ready` | `1a2b27cea6268d8d5b80f2a8e1bd93e067c55985e15006497d086955f4a66eae` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `package-inventory-001` | `ready` | `3f438d009e719ccad45914f6eabd59baa1419e54da3d2840ac46e45157fe9036` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `parse-timestamp-001` | `ready` | `331621b356a370a7fa95f28f001d977b1e1a0879dcbc6d17d6e0bc7e484b96da` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `percentage-output-001` | `ready` | `59396c14df3c94b60bac35feb34ad8550cb48410185b620ac2a076348b942647` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `period-predicate-001` | `ready` | `f99c7b917394f0b4c35b924dec25370e57db3debcb04bcaedec50aa5ceae1659` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `pointer-chain-001` | `ready` | `4130bae940896121742e539dd9704f3e5bc6873e49bf049fda75fe3897170a21` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `predicate-filter-001` | `ready` | `2bb16c6be151059957cc704e48954de0ca95d906e5c25a132307c75551ecd7a6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `prepend-values-tests-001` | `ready` | `3132525689c6060676a060bfe465cf7915457906693842184f1b07a32b7059ae` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `propagate-context-001` | `ready` | `1a79e076f7a0c29bb50b4362460f6e59b1cafd2e0d6494b057f9ed787ac933e8` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 5/5 | Phase 3A |
| `quadrant-switch-001` | `ready` | `0b1f435a8e4e714be86c12b5e9806e9653925877b796d14d746909d848d2f771` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-presence-001` | `ready` | `2862e52d21b8a7c55e278f188e9ded6d1ef1cc520ee8422dc2dc05a4c83ac0f0` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `quantity-removal-001` | `ready` | `f44607f97294a9c10f2166824b04d1c42f6e7c18bc0039c50fedfdfca50117a2` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `range-offset-001` | `ready` | `a80dc5794366f34ef94775c08dc5bc65da07502ca5573fc69f3c548af9097090` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rank-occupancy-tests-001` | `ready` | `c4aa2c820635de0f2c4589e35e92ec44e3841f39d6d620ab786f962f9e259de3` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `receiver-reset-001` | `ready` | `2c14320171846b43f09408d772f6e72451595a77d54e18c51e73dd2eeb32203a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `receiver-speakers-001` | `ready` | `265442613524a5bb4eb6deba52b8f491af2e352d9c669f71975b6edd487dbaf1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `record-result-001` | `ready` | `e4c7be7391893f352c73621c10da7469df5f5ebd7212e6a907cf0563c59628b1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rectangle-perimeter-001` | `ready` | `7d5093f33ca6bfe29bf63b8acf237adf5dd111427a6c8e7ec203d5c89f7686af` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remaining-constant-001` | `ready` | `822af7fada7bdab4247ee534034f4e84bdf1a9e65d3079146aa69dc676ae9d5d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `remove-item-tests-001` | `ready` | `204515afffb5840342692dc5ab7a295dad126b72d6307a8bb993bef83bb8cddb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `repair-utf8-001` | `ready` | `61f0bdeabc19d4aac9353fb9f091383da1a6b13bbb0552a076ccd1a9d04bbaac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `replace-byte-001` | `ready` | `da0d22d6d152c3d95818431f9fb4cb317ee2bb5cd8a43a8d86e56f95ac13f2e8` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `replace-pending-value-001` | `ready` | `da9081825e23fc48fc96d57288bbcc1ae8823660d296f4cd05d8bf8f4c1a8e8b` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `reserved-copy-001` | `ready` | `688a13e4dc69f0179e903ad6224dac8a94ebcadde85c4fd27415d562e52a1b0f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `resource-finalize-001` | `ready` | `d4bb3cde09096fea9ec1a2285b0da5eed342631bc4457ecedc18efd9d6b2b5ab` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `rune-boundary-tests-001` | `ready` | `66831076960ac6f263c717704ce3938892d37aa9c3debf498d71a614be10dc27` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `rune-limit-001` | `ready` | `f24b09322d8a1e759152c24132de744492df50fa046cb87ad746c9c975c71e2b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-marker-001` | `ready` | `aa7a1054bf433584e8fd919317a8f117602ff6ab37cf1f998a9a5e8fe66a380d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-offset-tests-001` | `draft` | `ac874ea5b6cfd23425d17dbd4471e5fdb35f2ac0fa00406b29d4aaf8c35a3296` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `rune-offsets-001` | `ready` | `9eb44a9c65479587ffc5e7dda308ecb68c86df82969bbdc950e0df45d274dcdb` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `rune-replace-tests-001` | `ready` | `3e565d4607d9b41b54c1157b17a5ce59e225d3d3ff38baa14fae592b6c57acf6` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `scan-lines-tests-001` | `ready` | `2deb8c4e1f85be775587f4e6c761792edef38483a310ddca1057c27334bcde10` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `select-ready-case-001` | `ready` | `aa1f1795ea810e8f33293dc97c0bdc1e2cd021b05d96f662ce2eef353de20f49` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 6/6 | Phase 3A |
| `sentinel-lookup-001` | `ready` | `7da29505d3aaffbabc6d2c53b00e7be77372654728057236ac2f9c0b070e5dcc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `set-or-append-001` | `ready` | `ff6a5d3ac5142af550418e0de218c7f274edef007a690e0fd47a956fe3bdc62b` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `seven-value-window-001` | `ready` | `98d6603df85c88a0a64572e38a4ef01a092f1e5c71e1cab0153f6c53715bc116` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `shape-contract-tests-001` | `ready` | `8badbf1214b6bb6e9160f5866e5ac525b08fe7bb42882551324fe3c260716f82` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `shared-counter-update-001` | `ready` | `3ef4be2df1eff19f0991270a65814c4bcb064b41907d52004ed5c3d57fa51f3a` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `single-json-value-001` | `ready` | `ef01ceddc357019d8a5eac205e3a73c3ac9f83ce4164c3a4b303c5408df69160` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-sequence-tests-001` | `ready` | `bb3b4671f0cb5f83309c021e2515b57650d20cdf51ac204d8f38e2999a423a38` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `slice-window-001` | `ready` | `c68f0147cda29c6179535f009bdf3cbcff0526dfffc35123ee74c0dac69bab46` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stateful-counter-tests-001` | `ready` | `aa0a708e7265397f4e4999b9379c17754c7f5077bcfb8a761b2d793271c882c7` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `stream-normalize-001` | `draft` | `d860149401279666054a3ab37a3c408b523d205d7417aefe5d53bf39276664b0` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `struct-copy-tests-001` | `ready` | `3c34a03f4e798df6cdeb8278ab4518291a51b1d86ef87606b9fe9db4f0985bdc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `struct-record-tests-001` | `ready` | `a8d93c02c7f3b7e4e17137428979d627c1dd3b9ae3f29cdbd7d9e64116d0489f` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `struct-transition-tests-001` | `ready` | `fab4337564787a0419727fa66eee114ae73e0f0ab5cdac3e41e06ffc5f7c2481` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `structured-error-tests-001` | `ready` | `fe499c301ec4cf01a4cb1116568b8fc58972e9e393679c405877d804f682ccdc` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `text-border-001` | `ready` | `b4f28598183117812e4da736d451daf35e6625fcc270fe1188dd0b335ce28d13` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `text-level-001` | `ready` | `a000a4b164b086c0ea121bca0c75c39ad00e2733370502913f3e6442668bfeac` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `timeout-operation-001` | `ready` | `54545521f02da6fd195951d7ad84a683379ebdef3cae95c419dacb59e538af63` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 13/13 | Phase 3A |
| `token-values-001` | `ready` | `4aa5603a9bc29e0315c9aaa95715771e3101ffb7e02f651f197d3eccffe79700` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `total-counts-tests-001` | `ready` | `c8bd4976274531193da685d564c546b2c8ef24c8e90e0df6932848376c355e04` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `two-quantity-tests-001` | `ready` | `baec5bde5f850d683e38a61b409c7edaa84a498929db7c9570a6fded0e86f2ea` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `type-switch-default-tests-001` | `ready` | `9804a58194c92aa1fdf429f132ce95849bb0f14a20cfb7259501b3b162862b44` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-line-001` | `ready` | `7db7407a10ca2e69df296e49c971dc979a708b924ade5ec64b2753a6d32843b1` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-nil-resource-001` | `rejected` | `0d47ed2069a1275853ac85e8911f0616f3ce293b2cd4a309c35ee925b3649ce9` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | Phase 3A |
| `typed-nil-result-001` | `ready` | `c0e4b6dfabb26865d40d38885d4e4a5c4facd174d6301cfa04bf40f49fd5a032` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `typed-value-observations-001` | `ready` | `66012a0901432ce1d70edd37598b53112d4bb59f63357ea52759ae71d6cb74de` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `unit-addition-tests-001` | `ready` | `1479e5fe4645bff273c002014c3423fe2b324900395ab4105f45ad10bc3b4a86` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `uppercase-label-001` | `ready` | `d2de48d327e063ff6a80602571dbe44693689d7c5250f8fcf6a965def26ac9fe` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `utf8-prefix-001` | `draft` | `7e161a31172398ceed3ba23d9364e127540f747e0afaffea888ccb32ce9d3617` | `pass` | `fail` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `utf8-validity-tests-001` | `ready` | `fc0c7ff066ec0ccf424812fdcc6a5567638d4ac5e178e17e253ad7a22b061580` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 11/11 | Phase 3A |
| `value-average-001` | `ready` | `bd215aa2a9810a2081f4e25efbdbf92984dfc08f7d65f363d3469b71b0bb36d5` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 8/8 | Phase 3A |
| `variadic-aliasing-001` | `ready` | `e306a087365c5214a14760fd10a6d10e03039b230f63fc7102cb0dab99be4441` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `variadic-sum-001` | `ready` | `f7aef546b1a4e005e24842b2689e43af1dc213e0aa7321922bf71175fa0fc4cf` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
| `variadic-tail-sums-001` | `ready` | `0f056b32d553a71eab12985d535da33755adac9d33961b5c10a90fd233ad7204` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 7/7 | Phase 3A |
| `whole-minute-output-001` | `ready` | `f7cc26245550a65ddc197694229ee8e1fc46e0d3581001f03f09a1e6d4e9b432` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `wide-sum-001` | `ready` | `ef0509ff72ae9dc8730914e9a20ebbac11784c8dfb1f1b96fd47fe462ea51e2d` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 4/4 | Phase 3A |
| `widened-rate-001` | `ready` | `49182a9c0dfabf6d5a9690109afb5012897ea67dfa6aa62166ffc3a1fe05350c` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `worker-completion-tests-001` | `ready` | `af666324ac62584aca76eae08847a560d586e9a7acf2e2675cac6a8d020c2a74` | `pass` | `pass` | true | `pass` | `pass` | `pass` | 8/8 | Phase 3A |
| `writer-double-001` | `ready` | `f0d693b7d2833aa5dde07c246b2d9d34fd3dea49b43d476d56afe71ab554e4ba` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 6/6 | Phase 3A |
| `writer-greeting-001` | `ready` | `bed1eabdca7a468183c5a8c30c7ba23508ae2646f75355ca56bb126022b077ec` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 9/9 | Phase 3A |
| `zero-results-001` | `ready` | `346cd9149bc56d8c5d7834d51dc74eb8933126bfdf573a01695659310a456af9` | `pass` | `pass` | true | `pass` | `pass` | `not-applicable` | 5/5 | Phase 3A |
| `zero-state-001` | `ready` | `1a89134b20adb56d7db4f4495d629f8017ab738fd2916dd450f876af8c06ca06` | `pass` | `not-applicable` | false | `not-applicable` | `not-applicable` | `not-applicable` | 0/0 | legacy v1 |
