# Repair two package boundaries

Prerequisites: modules-build-001 and accepted consumer-text-contract-001. This manual unit assesses import cycles and internal-parent visibility separately.

1. Predict the consumer/producer import cycle. Reproduce its compiler diagnostic, including the alias-only wrong repair.
2. Remove the cycle by making the consumer own a minimal contract and removing the concrete back-dependencies. Preserve Render's exact bracketed text.
3. Independently predict and reproduce the outside caller's forbidden internal import. Explain why sharing a module does not grant access.
4. Author a narrow public domain facade and call it from cmd/outside; keep the formatter implementation in its original location.
5. Verify the legal inside-parent caller, internal formatter and repaired public caller. Preserve command outputs [alpha] and <beta>, each followed by newline.
6. Submit actual before/after files, predictions, commands, diagnostics and explanations under PROCEDURE.md.

Edit only the four listed implementation files. Keep module, tests and outputs unchanged. No application or automatic coverage credit is implied.
