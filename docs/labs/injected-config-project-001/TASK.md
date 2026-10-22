# Inject and verify a configuration dependency

Build module finite.local/configkit from the editable starters. Source is nonnil;
requests contain 0–4 uppercase/underscore keys of 1–12 bytes; values have at most
32 bytes.

- Author a minimal named consumer interface in snapshot. Collect accepts it; production lookup and unrelated adapter methods stay outside snapshot.
- Lookup every requested key exactly once, in order including duplicates. Observe current responses on every invocation; no cache or environment fallback.
- Return KEY="value" with Go %q quoting when present, otherwise KEY=<missing>. Empty requests return nil without lookup.
- Implement and delegate to snapshot/internal/row.Format. Wire the supplied process adapter in cmd/snapshot; preserve the independent client's Lookup-only provider.
- Author an external-package recording double and the two named tests covering injection, order/count, duplicates, missing versus empty, empty requests and changed responses.
- Use PROCEDURE.md to build/run, inspect imports and verify public/client/tests plus wrong examples. Explain module ownership, the narrow internal boundary and each interface method.
