# ADR 0023: Preserve workspaces through the bounded-view nil correction

- Status: accepted
- Date: 2026-09-11

bounded-slice-view-001 permits nil input with valid zero bounds. Its hidden check
requires length/capacity zero but the public suite omitted that input. Add the
same observable requirement publicly, with useful actual/expected diagnostics.
A non-nil empty zero-capacity result remains valid; no representation requirement
is introduced. TASK, starter, reference, hints and exported workspace files stay
unchanged. P4T tooling and P4C content remain separate PRs.

Use ADR0014's exact directional compatibility mechanism:

- Stored digest: 0e8e493addf0119352a834748e66ea26da49ef61b44db5a4e39f2bbf1212f963
- Corrected digest: 960ba2896fb8d14d5911f0e3e0564aea05f5b22b129c5cb5b5c54c56df1ddbf8

The corrected digest adds only the public nil assertion and one private mutant.
The manifest declares that mutant; neither edit changes exported payloads.
No progress/marker/solution file is migrated or reset. Unknown pairs and reverse
compatibility remain rejected. Rolling back content preserves old markers;
new markers need the corrected content retained.

The disposable Service.Check regression first supports the projected correction
before the content PR. P4C's persistent review note then disables that projection:
the regression executes canonical public tests, so removing the nil assertion
must fail for unexpected public success. It accepts reference and a non-nil
empty zero-capacity alternative, and checks byte-preserving old-workspace views.
Ordinary Check/Complete keep normal attempt/time/completion accounting.
