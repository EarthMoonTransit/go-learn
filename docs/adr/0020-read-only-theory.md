# ADR 0020: Local read-only Theory

- Status: accepted
- Date: 2026-09-10
- Basis: foundation-next P0/P1; user-approved local explanation workflow

## Contract

Use the existing learner Service, CLI and local HTTP UI. Add optional
`curriculum/topics.yaml`, version 1, with a `topics` list. Each topic has a stable
`id`, existing route `section_id`, English `title`, relative `path` under
`theory/`, prerequisite topic IDs, existing `skills`, 1–3 ready v2 `task_ids`,
`go_version`, `reviewed_at` (ISO date), and 1–3 `official_links` containing
`url` and `description`. YAML unknown fields, duplicate IDs/links, missing
references and prerequisite cycles are errors. An absent file or empty list is
valid. This is an additive theory contract; task schemas and skill IDs do not change.

`learn topics` and `learn theory <topic-id>` read the same views as authenticated
GET API endpoints. Topic task links show existing lesson numbers and navigate to
TaskView. Only existing explicit start/resume/Check actions affect progress.
Show Theory links from tasks and after public feedback so the learner can return
to the explanation. There is no second scheduler, completion button, reading
tracking, implicit pause, progress write or migration.

Load only bounded regular UTF-8 files. Reject symlinks in every path component,
traversal, absolute paths, oversized metadata/pages, unknown task/skill/section
IDs and invalid URLs. Markdown uses text nodes and a narrow link subset: HTTPS
and local anchors only; reject raw HTML and unsafe link targets. Never fetch
remote content. Official metadata links are restricted to go.dev and pkg.go.dev;
external anchors open without opener/referrer or session credentials. Existing
loopback authentication, origin and CSP controls remain unchanged.

## Compatibility and migration

Route v2 and progress v3 retain identical interpretation and bytes. Theory
prerequisites organize reading only, never task availability. Empty theory must
not break ordinary lessons. Rollback removes optional theory files/UI support;
no learner-state downgrade is needed. Future metadata changes need a new version
and explicit migration; there is no migration in this batch.

## Acceptance

Synthetic fixtures verify CLI/API agreement, invalid files/URLs/IDs, prerequisite
cycles, safe link rendering, missing theory, and byte preservation of learner
solutions, progress and markers. P1 delivers tooling; P2 separately supplies six
reviewed English pages with disposable example checks. Existing tooling gates
remain mandatory. Theory examples do not disclose the linked exercise solution.
