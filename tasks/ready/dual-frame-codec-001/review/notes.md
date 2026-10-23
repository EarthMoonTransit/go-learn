# I11 dual-frame independent actual review

Original160-word root/sole-approved contract. Official task-init, original ADR0008 creation/digest with sources=[]; independent actual review enables disposable checks, not independent acceptance. Earlier text-level and binary-record-tests are unchanged scope preparation, not donor material. Learner implements all four methods; supplied Frame/API and encoding interface compatibility give no additional type/interface credit.

| Requirement | Public assertion | Required finite defects |
| --- | --- | --- |
| Literal binary version/tag/length/full payload | ExactFrameFormats vectors and exact bytes | binary-version-out,binary-tag-zero,binary-length-minus-one,binary-drop-tail |
| Strict binary framing | BinaryRejectsInvalidFrames named cases | binary-accept-version,binary-accept-trailing,binary-accept-truncated |
| Stable canonical lowercase text | ExactFrameFormats and TextRejectsInvalidFrames | text-uppercase-output,text-accept-version,text-accept-uppercase,text-ignore-length,text-ignore-tail,text-bad-nibble-zero |
| Max32 and nil oversize result | maximum-payload vectors, OversizeMarshalIsAtomic, oversize input cases | accept-oversize |
| Replace populated state, exact errors, rollback | ExactFrameFormats, all rejection snapshots | merge-old-payload,clear-receiver-on-error,fresh-error |
| Preserve input and own all output/decoded storage | FrameStorageOwnership, input/error snapshots | shared-output,binary-retain-input,mutate-input,marshal-reuses-payload |

Public vector/rejection helpers iterate cases without per-case PASS subtests; actual errors still name format/case and values. This feedback repair keeps late text/ownership assertions inside the existing report diagnostic budget. Earlier truncated reports are retained without useful-diagnostic credit.

Twenty-one required wrong implementations are finite and guard short input before indexing. Public wrong runs must fail with named ordinary assertions; panic/race/timeout/compile errors give no credit. Inputs<=512; arbitrary payload bytes and every tag valid; nil/empty payload equality is bytes-based. Literal expected bytes in both directions prevent paired round-trip defects hiding representation errors. Grader varies nontrivial tags/payload lengths, computes independent standard-hex expected vectors and checks every proper prefix of two frames. No hidden text normalization or unspecified capacity limit.

## Actual method and valid controls

Reference validates version/declared bound/full encoded length before allocating payload, validates text characters into local storage and commits receiver only after success. Marshal validates bound before constructing output and leaves receiver storage intact. Separate private nibble helpers are learner-authored. Correct encoding/hex plus append construction and explicit-copy/nil-empty alternatives pass public/full; helper names, capacities and empty storage identity are unrestricted.

Early bounded allocation before validating version/length is an intentional behavior-pass method-reject control. Outputs alone do not prove preallocation order; actual source review checks this requirement without an AST/allocation-count checker. Source review also verifies both real method bodies, full-input validation and ownership, not just provided signatures. Final sole source/method/evidence acceptance and root promotion/gates remain pending.

Independent actual acceptance: current canonical-v3 and28current copied fixtures accepted by /root/reviewer on2026-10-23. Superseded attempts retained without current credit. Public grouping preserves all original cases/assertions and useful labels. Final integration/delivery gates remain separate.
