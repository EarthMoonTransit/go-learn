# Task-check collection

`task-check -all` создаёт одноразовый, детерминированный снимок результатов
полной проверки банка:

```bash
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
```

Формат фиксирует `schema_version: 1`, `contract_version: task-check/v4`, digest
toolchain policy, `verifier_source_digest`, actual Go runtime (`version`, `goos`,
`goarch`), упорядоченные task results и структурные execution metrics.
Контракт задан `schemas/task-check-collection.schema.json`.

## Task input digest

Digest включает:

- checker contract version и lifecycle state;
- embedded production sources tasking, isolation и finite native kernel;
- относительный path, permissions и содержимое каждого файла task package;
- task v1/v2, lifecycle/review, provenance v1/v2, quality/collection schemas;
- taxonomy/toolchain schemas, `curriculum/skills.yaml`,
  `sources/registry.yaml`, project identity, outbound-license policy, `LICENSE`,
  `THIRD_PARTY_NOTICES` и `config/toolchain.yaml`.

Для fuzz-профиля дополнительно включён canonical `docs/FUZZ_SOURCE.md`.
TaskContentDigest старых задач остаётся прежним; копия grammar contract новой
задачи входит в её обычный content digest под `verification/`.

Generated reports и сама collection не включаются. Любое изменение package,
state или учитываемого contract input делает collection stale.

## Проверка при чтении

`validate -checks` и `task-report -checks` до использования проверяют:

- strict YAML и JSON Schema;
- checker contract, toolchain и verifier source digests;
- фактический runtime через trusted `go env` в disposable environment;
- полный текущий task inventory;
- отсутствие missing, unknown и duplicate IDs;
- deterministic ID order и совпадение lifecycle state;
- task input digests и identity вложенного quality report;
- согласованность числа sessions, workspaces, variants, legacy и child Go runs,
  runtime queries и отдельных fuzz/helper/native counters.

`CheckAll` запрашивает runtime один раз на сессию и подтверждает неизменность
Go executable перед завершением. Каждый validation pass читателя проверяет
runtime заново без запуска task code или native helper. Runtime metadata не
означает, что alternate SDK был реально протестирован. `variant_runs` сохраняет
смысл старых профилей; native invocations учитываются в `fuzz_runs`,
`fuzz_go_runs`, `fuzz_workspaces`, `fuzz_helper_runs`, `fuzz_native_runs`.

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

ADR 0019 сохраняет race-required public и defect checks. ADR 0032 добавляет
native eligibility и v4 freshness. Старые v3 collections нужно пересоздать
текущим checker; менять только version label запрещено. Validator, task-report
и coverage-report используют один новый файл через `-checks`.
