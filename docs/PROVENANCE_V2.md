# Provenance v2

`provenance.yaml` with `schema_version: 2` distinguishes project identity from
audited external inputs.

## Project and creation record

`project` must exactly match `config/project-identity.yaml`. `content_digest` is
the SHA-256 framing digest over `TASK.md` and regular files under `starter/`,
`tests/`, `solution/`, `workspace/`, `verification/` and `hints/`. Lifecycle,
review, task metadata, attribution and provenance files are excluded so review
transitions do not rewrite the authored-content identity.

`creation` records author, date and a concise transformation boundary. Original
content has `kind: original` and an empty `sources` array. Its future commit is
not guessed; `release.manifest` reserves a deterministic path under
`provenance/releases/` for release-time binding.

## External source records

Each source must exactly match `sources/registry.yaml` by id, repository, commit,
license and usage mode. Each artifact has an exact source-relative path, SHA-256
and one usage classification.

Allowed roles are `concept-seed`, `api-shape`, `adapted-code`,
`adapted-test-structure`, `factual-reference` and `comparison-only`.

- `adapt` may use copied/adapted/structure roles when license obligations are
  retained and compatibility policy permits the license.
- `reference-only` is limited to factual-reference or comparison-only; artifacts
  may be factual, comparative or explicitly excluded.
- `concepts-only` is limited to concept, factual or comparison evidence and may
  not claim copied/adapted material.
- `blocked` cannot back a task.

`kind: adapted` requires at least one source. `kind: combined` requires at least
two. A per-task attribution file is mandatory. Generated `THIRD_PARTY_NOTICES`
will report upstream terms but never replace or broaden them.
