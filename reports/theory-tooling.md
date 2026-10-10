# P1 / T14 — local read-only Theory tooling

The learner service reads optional curriculum/topics.yaml and bounded local
Markdown pages. `learn topics`, `learn theory <topic-id>`, authenticated GET
endpoints and UI navigate the same stable topic/task IDs. Missing pages are
planned. Existing lessons retain their route numbers and explicit actions.

Topics group under existing route sections. A task's Theory button remains
available after public Check; a page links to the task view, existing GoLand
launcher and Check controls. No second scheduler, reading completion, hints,
assistance, time tracking or progress migration is introduced.

ADR0020 is clarified to keep skills in task manifests and permit missing pages,
as required by foundation-next. New topic schema is separate from task schemas.
Repository validation checks structure and semantic references. Existing
loopback auth/origin/CSP remains in force; HTTPS links use no-referrer/noopener,
anchors scroll without replacing the session fragment, and remote assets are
never fetched automatically. Unsafe files, raw HTML and unsafe links are rejected.

Synthetic tests cover CLI/API parity, stable lesson number, exact learner-file
preservation, unsafe/unknown metadata, size, UTF-8, symlinks, traversal, missing
pages, auth and origin. Targeted Go tests and the dependency-free Node renderer
checks passed. Mandatory full acceptance and exact-head CI remain PR gates.

No theory pages or ready tasks are supplied by this tooling batch. P2 supplies
six separately reviewed explanations; human completion evidence is unknown.
