Header and two-index slicing are both explicitly assessed. Mutants cover capacity-versus-length validation, full-slice capacity trimming, allocation instead of a view, and invalid empty-window policy. Valid inputs include spare capacity and nil. Grader uses fresh arrays per case and checks both directions of sharing.

Independent review approved on 2026-09-05 after repairing invalid-path preservation coverage and the non-compiling fallback defect. Reference and required defects passed disposable checks. Human calibration remains pending. No upstream task material was used.
