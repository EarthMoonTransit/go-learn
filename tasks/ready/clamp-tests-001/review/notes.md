# Review notes

## Prior rejection retained

The 2026-09-04 review returned this package to draft because mutation checks
cannot prove table-driven structure, named cases or diagnostic quality. The
lifecycle history is unchanged. Earlier acceptance reports are historical.

## Substantive redesign, 2026-09-05

This task now assesses expressions.comparison through verification practice.
The learner compares expected values for inclusive interval partitions. Test
writing is the submission activity; no testing, table or diagnostic skill is
credited. A discoverable test and assertion pattern is supplied. Table layout,
test names beyond normal discovery and failure-message style are not graded.

Six required behavior models distinguish missing lower/upper limiting, wrong
results exactly at each endpoint, an altered interior result and a zero-valued
singleton result. The exact-boundary models change output; merely replacing
< with <= while returning the same endpoint would be equivalent and is avoided.
The reference includes each partition; the starter passes the correct function
and leaves defects alive. Every correct test organization is accepted.

## Source boundary

The existing exact Exercism counter .meta/config.json record is preserved and
its SHA-256 was checked against the cached pinned artifact. Only that metadata
and root LICENSE were read during this redesign. The metadata supplies a
test-suite direction; it contains no interval API or cases. Attribution now
names its authors, contributors, path, commit and MIT notice. No upstream
exercise statement, canonical cases, implementation, tests or solution entered
the authoring inputs. Independent final review approved on 2026-09-05; all six required defects are detected in disposable verification.
