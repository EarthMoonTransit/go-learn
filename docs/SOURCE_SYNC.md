# Deterministic source synchronization

ФАЗА 2A синхронизирует только запись из `sources/registry.yaml` и никогда не
исполняет upstream-код.

```bash
go run ./cmd/source-sync -root . -source exercism-go
go run ./cmd/harvest -root . -source exercism-go
go run ./cmd/catalog-report -root . -source exercism-go
```

## Cache layout

По умолчанию cache находится в `~/.cache/go-learn/sources`. Переменная
`GO_LEARN_SOURCE_CACHE` задаёт другой абсолютный или относительный каталог.

```text
<cache>/<source-id>/<commit>/
├── checkout/             detached, clean Git checkout
└── source-manifest.json  deterministic registry/cache attestation
```

Manifest не содержит timestamp или абсолютный путь. Cache не является
versioned output и не должен добавляться в Git.

## Trust boundary

- repository URL, commit, license, usage mode и adapter берутся только из
  registry;
- `blocked` отклоняется до создания cache; `concepts-only` требует явного
  `-allow-concepts-only` только для source-sync и всё равно не harvest-ится;
- Git запускается напрямую без shell; global/system config, hooks, credential
  helpers, fsmonitor, submodules, LFS smudge и преобразование line endings
  отключены;
- после checkout проверяются exact HEAD, detached state, origin URL, отсутствие
  tracked/untracked изменений, initialized submodules и object alternates;
- cache публикуется только после полной проверки;
- upstream README/AGENTS и другие тексты считаются данными, а не инструкциями.

## Exercism harvest

Adapter читает root `config.json` и per-exercise `.meta/config.json`, затем
инвентаризирует файлы каждого зарегистрированного concept/practice exercise.
Symlinks, traversal, special files и выход за checkout запрещены. Каждый обычный
файл получает SHA-256 и одну или несколько ролей. Содержимое файлов в JSONL не
записывается.

Versioned output:

```text
catalog/candidates/exercism-go.jsonl
reports/import-exercism-go.md
reports/taxonomy-audit-exercism.md
reports/exact-duplicates-exercism.md
reports/unmapped-concepts-exercism.md
```

Повторный harvest на том же verified checkout сравнивает готовые bytes и не
перезаписывает неизменившийся файл.

## Exact provenance generated cases

`exercism-go` имеет companion source `exercism-problem-specifications` в
registry. Source-sync проверяет его primary commit так же строго, как основной
track. Adapter читает только служебный заголовок `cases_test.go` и принимает
сокращённый commit лишь при единственном совпадении с полными SHA из
`provenance_commits`. В candidate сохраняются путь generated-файла, repository,
полный commit, исходное сокращение, license и evidence URL; содержимое файла не
копируется.

Если заголовка нет, commit не разбирается или совпадение неоднозначно, exact
provenance не выдумывается: candidate получает warning либо harvest завершается
ошибкой для противоречивого commit. Companion имеет `usage_mode:
reference-only`, не harvest-ится как банк задач и не разрешает копирование его
описаний или canonical data.
