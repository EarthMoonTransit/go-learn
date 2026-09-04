# Phase 3A.1b performance and acceptance

## Outcome

Полная проверка банка теперь выполняется один раз. Один `task-check -all`
создаёт validated CheckCollection; validator и `task-report -check` используют
её без повторного task execution. Persistent result cache, sharding и
concurrent task execution не добавлены.

## Наблюдаемое время

Wall-clock приведён только как диагностический сигнал; gate основан на
структурных counters.

| Локальный запуск | До | После |
| --- | ---: | ---: |
| `go test ./...` | примерно 96–100 s | 10.86 s |
| `go test -race ./...` | примерно 103–104 s | 13.82 s |
| `go test -tags=taskintegration ./...` | отдельной suite не было | 27.52 s |

Для baseline GitHub Actions пользователь наблюдал около шести минут только на
обычном `go test ./...`. Значения зависят от machine/cache и не используются как
жёсткий threshold.

## CI command graph

```text
config/toolchain.yaml
        |
        v
fast tests --> fast race --> vet --> tagged subprocess integration
                                              |
                                              v
                                     build three CLIs once
                                              |
                                              v
task-check -all --> temporary checks.yaml --> validate -checks
                                  `---------> task-report -check -checks
```

Production full-bank execution: 4 passes → 1. Подробная декомпозиция находится
в `reports/phase-3a1b-execution-counts.md`.

## CheckCollection и stale protection

- strict versioned schema, checker contract и toolchain digest;
- deterministic task order без timestamps;
- per-task digest учитывает state, paths, permissions, bytes, task contracts и
  toolchain policy;
- missing, unknown, duplicate, reordered или state-mismatched results
  отклоняются;
- изменённые package/lifecycle/schema/toolchain inputs отклоняются;
- aggregate report фиксирует source contract, toolchain digest и task digests.

Collection не заявляется attestation: CI доверяет только обязательному
предыдущему producer step в том же job.

## Acceptance

| Проверка | Результат |
| --- | --- |
| `go test ./...` | pass |
| `go test -race ./...` | pass |
| `go vet ./...` | pass |
| `go test -tags=taskintegration ./...` | pass |
| build `task-check`, `validate`, `task-report` | pass |
| single `task-check -all -output` | pass |
| validator reuse через `-checks` | pass |
| report reuse/check через `-checks` | pass |
| повторная collection byte comparison | pass |

## Protected artifacts

SHA-256 до и после совпадает. Directory digest — SHA-256 от сортированного
списка `SHA-256  relative/path` для каждого regular file.

| Artifact | До | После |
| --- | --- | --- |
| `tasks/ready/defined-types-001` | `a56d73571d38124735653800129e5029e03b5374d018b6a32f4c79a40f9f30e2` | same |
| `tasks/ready/variadic-sum-001` | `c5a429509d6b20f0d130d593d605a3947e7e10af3880f372a10905c0cf83c7b2` | same |
| `tasks/ready/zero-state-001` | `f21b87df35fa46deeb69b36027154197cc5e0529bce42b6858626ec815fbec6d` | same |
| `schemas/task.schema.json` | `987c1830cbdcb02702fd7e6ddb89362e76d479df14ade510cabcd20dc4dc30fb` | same |
| `curriculum/skills.yaml` | `4cc68407da591576cbce72117bc9a92993d491084ec736c26e5921323b269f16` | same |
| `catalog/reviews/exercism-go.yaml` | `05bf10d9bdbc282ca66c8b0c5bc9129deab1475be5a53bab3c303d51ec698f03` | same |
| `catalog/seed-plans/exercism-foundation-pilot.yaml` | `0cbbf900b5da2ba75460bd0f2d54c832f64e1debeaa2f0daadfe2025c134a2c5` | same |

## Спорные решения

1. Compile status определяется test-run event одного `go test -json`, а не
   отдельным compile-only process. Task contract требует реальные tests, поэтому
   отсутствие такого event считается compile failure.
2. Локальный `task-report` без `-checks` сохраняет удобный one-pass fallback;
   CI всегда передаёт collection.
3. Validator без `-run-commands`/`-checks` является static-only и сверяет
   committed aggregate; executable acceptance использует один из явных режимов.
4. File permissions входят в task digest, хотя Git надёжно переносит прежде
   всего executable bit. Это делает локальную подмену mode видимой.
5. Absolute output path разрешён для `RUNNER_TEMP`; tracked collection в
   repository не создаётся.
6. Shared cache допустим только внутри invocation для доверенных repository
   packages; он не объявлен sandbox boundary.

## Suggested commit

`perf(authoring): reuse single-pass task verification`
