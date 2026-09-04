# Raw candidate catalog

Здесь хранятся **только метаданные** ещё не принятых кандидатов, по одному
детерминированному JSONL-файлу на source id:

```text
catalog/candidates/<source-id>.jsonl
```

Candidate содержит paths, upstream metadata, file roles, SHA-256, license,
usage mode и warnings. Он не является задачей и не учитывается в curriculum
coverage. Снимки чужих репозиториев, полные условия и решения сюда не коммитятся. Переход в
`tasks/ready/` возможен только после license review, дедупликации, authoring и
прохождения quality gates.
