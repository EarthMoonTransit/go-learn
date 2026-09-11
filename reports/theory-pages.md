# P2 / C15 — six local English explanations

Six independently authored pages provide the initial reading for A1 workspace,
A5 testing, B2 slice bounds, B3 ownership, F1 interfaces and L1 cancellation.
Topics attach to existing sections and stable task IDs. F1/L1 stay in their
existing later course sections; reading does not reorder or complete lessons.

Each page has a goal, prerequisites, explanation, one small example, mistakes,
existing exercise links and 1–2 official reading links. Body content is local;
no remote assets or code run on page load. Official links were checked on
2026-09-11; package URLs are floating, while examples were checked with the
pinned local Go 1.27.0 executable. Source websites are linked for facts/further
reading, not imported instructional text. No LGWT material is used.

All six examples compiled and passed in separate temporary Go modules with
network module access disabled. A first harness invocation used the system Go
1.26 binary and refused the module version before compiling; rerunning with the
resolved pinned Go 1.27 executable passed all six. `learn topics` shows all six
ready explanations. These are theory pages, not six ready tasks or completions.

B2 demonstrates a two-index slice, not the bounded-view implementation. L1 shows
synchronous cancellation acknowledgement, not StartCancelable or worker lifecycle
solution code. A5 uses strings.Cut, unrelated to the linked clamp task's cases.

P1 tooling is the prerequisite. Full local acceptance and independent review are
required before publication is called delivery. GitHub Actions currently refuses
new jobs due to account billing/spending settings; no merge bypass is authorized.
Learner history remains unread; examples and test fixtures are not human evidence.
