# Content review

Independent application of `slices.full-slice-expressions`. Source/design approved by /root/m2_2_review. Distinct contract: bounded-slice-view-001 implements a view with zero spare capacity; this tests partial allowed reuse and its exact boundary using nonzero spare and hidden capacity beyond len.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| exact window and bounded capacity | TestSpareViewBoundaries/interior,zero-spare,hidden-capacity | unbounded-capacity, zero-spare, one-extra-capacity, spare-as-length, ignore-low |
| sharing and append reuse/detachment | TestSpareViewBoundaries | detached-view, unbounded-capacity, zero-spare |
| empty view with reserve | TestSpareViewBoundaries/empty-interior,empty-start,empty-end | empty-loses-reserve |
| preservation of selected and spare cells during call | TestSpareViewBoundaries | mutate-selected, mutate-spare |

Diagnostics report actual and expected values or a named state/alias violation. Compile errors and timeouts are not defect detection. All behaviorally equivalent implementations are accepted; no source-shape, allocation-timing or human-mastery claim.

Public checks establish learner-test compatibility with the correct implementation. Full mutation checks establish all required defect detection; the weak starter is intentionally insufficient.

Independent content review approved 2026-10-21 by m2_2_review at the provenance content digest. Targeted technical checks passed; 9/9 required defects detected. Strong learner tests detected all nine defects; weak starter remained insufficient. Public correctness and full mutation detection are separate evidence.
