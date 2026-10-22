# Finite-v1 review

Contract approved by /root/reviewer before authoring. Actual content/method and canonical/public-only8/8 accepted by /root/reviewer2026-10-22; applicable race passes.

| Requirement | Public assertion | Defect/method |
| --- | --- | --- |
| absent vs present empty and ancestor shadowing | TestAbsentEmptyAndShadow | drop-empty/keep-ancestor/always-present |
| literal id and collision protection | TestLiteralMetadata; TestUnrelatedAndWrongDynamicValues | wrong-id/string-key/trim-id |
| safe assertion for actual wrong dynamic type | wrongValue parent returns int23 | wrong-type-present; assertion must return absent |
| parent control and unrelated values | TestParentCancellationAndDeadline; unrelated-value assertion | drop-parent |
| private key and request scope | actual source and explanation | supplied Context API earns no authored-action credit |

Public-only specimens must detect every behavioral defect with useful assertions. Panic, timeout or compile failure does not establish detection. Method review examines actual source; behavior pass alone earns no method credit.

Independent application evidence: type assertion changes direct dynamic-number classification into a safe Context.Value accessor preserving absence versus present empty and rejecting actual int23 dynamic values. Private key and metadata scope reviewed from actual source.

The exact arbitrary-ID public case is protected by trim-id. Trimming whitespace must fail TestLiteralMetadata; grader repeated-shadow does not cover this case. A disposable omission of only that public function must leave exactly trim-id alive while the reference passes. This retained requirement→public assertion→mutant relationship needs no new tool.

Actual retained-regression results independently accepted2026-10-22: canonical/public-only8/8; omission reference passes and exactly trim-id survives7/8. No task inputs change after this accepted freeze.
