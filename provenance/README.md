# Provenance и attribution

`task.yaml.source` хранит основной источник, а `provenance.yaml` — полную цепочку
происхождения: все source ids, paths, commits, hashes, редакторские решения и
attribution notices. Нельзя заменять commit веткой или тегом.

Минимальная форма `provenance.yaml`:

```yaml
version: 1
task_id: slices-append-aliasing-001
inputs:
  - source_id: go-practice-bank
    repository: https://example.invalid/go-practice-bank
    path: authoring/slices-append-aliasing-001
    commit: 0123456789abcdef0123456789abcdef01234567
    license: CC-BY-4.0
    content_hash: sha256:...
decisions:
  - kind: original
    note: Written from the taxonomy and language specification; no source text copied.
attribution: []
```

Реальный task package не может использовать заглушки из примера.
