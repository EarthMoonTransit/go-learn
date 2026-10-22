# Finite-v1 review

Contract approved by /root/reviewer before authoring. Actual content/method and canonical/public-only7/7 accepted by /root/reviewer2026-10-22; applicable race passes.

| Requirement | Public assertion | Defect/method |
| --- | --- | --- |
| prompt return and streaming before EOF | TestStreamingBeforeEOF prefix after each send | collect-until-eof; useful prefix diagnostic |
| transform then filter, zero, order, duplicates | TestStreamingBeforeEOF; TestZeroNegativeDeltaAndDuplicates | wrong-delta/filter-before-transform/drop-zero/reverse-order/deduplicate |
| empty/filter-only exhaustion and close | TestEmptyAndAllFiltered; final close assertion | never-close-output |
| two actual range stages and separate send/close ownership | actual reference checklist | fused stage remains behavioral control, not method acceptance |

Public-only specimens must detect every behavioral defect with useful assertions. Panic, timeout or compile failure does not establish detection. Method review examines actual source; behavior pass alone earns no method credit.

Independent application evidence: range-drain changes Collect-to-slice into two live output stages before EOF; ownership changes one ready-gated producer to distinct intermediate/final send-close paths. Actual reference ranges, completes both stages, closes only its owned channels. Fused correct control passes7/7/race but fails two-stage method.
