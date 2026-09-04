# Task-check collection

`task-check -all` создаёт одноразовый, детерминированный снимок результатов
полной проверки банка:

```bash
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
```

Формат фиксирует `schema_version: 1`, `contract_version: task-check/v1`, digest
toolchain policy, упорядоченные task results и структурные execution metrics.
Контракт задан `schemas/task-check-collection.schema.json`.

## Task input digest

Digest включает:

- checker contract version и lifecycle state;
- относительный path, permissions и содержимое каждого файла task package;
- task lifecycle/review/provenance/quality/collection schemas;
- taxonomy/toolchain schemas, `curriculum/skills.yaml`,
  `sources/registry.yaml` и `config/toolchain.yaml`.

Generated reports и сама collection не включаются. Любое изменение package,
state или учитываемого contract input делает collection stale.

## Проверка при чтении

`validate -checks` и `task-report -checks` до использования проверяют:

- strict YAML и JSON Schema;
- checker contract и toolchain digest;
- полный текущий task inventory;
- отсутствие missing, unknown и duplicate IDs;
- deterministic ID order и совпадение lifecycle state;
- task input digests и identity вложенного quality report;
- согласованность числа sessions, workspaces, variants, legacy и child Go runs.

Collection не является криптографическим attestation. Надёжный CI создаёт её и
потребляет в последовательных обязательных шагах одного job.

## Совместимость CLI

- `-checks` можно сочетать с `-phase1` или `-coverage`;
- `-checks` и `-run-commands` взаимоисключающие;
- без `-checks` `task-report` сохраняет локальный fallback с одним новым full
  check;
- без `-checks` и без `-run-commands` validator выполняет только static checks
  и проверяет committed aggregate artifacts;
- с `-run-commands` validator выполняет ровно один `CheckAll` и строит report из
  полученной in-memory collection;
- при `-checks` validator и report builder не запускают task code.

## CI usage

```bash
.tmp/bin/task-check -root . -all -output .tmp/results/task-checks.yaml
.tmp/bin/validate -root . -phase1 -checks .tmp/results/task-checks.yaml
.tmp/bin/task-report -root . -check -checks .tmp/results/task-checks.yaml
```

Файл collection остаётся временным CI artifact и не коммитится. Повторный
`task-check` на неизменных inputs обязан дать те же байты.
