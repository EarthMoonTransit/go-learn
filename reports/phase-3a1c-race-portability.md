# Phase 3A.1c race runner portability

`remote_ci: pending`

## Scope

Phase 3A.1c исправляет только portability внутреннего race runner. Учебные
задачи, task/provenance schema v2, taxonomy, candidates, seed plans, sources и
UI не изменялись.

## Reproduction

На текущей darwin/arm64 toolchain поведение `-race` с отключённым cgo не
воспроизводит Linux requirement. Поэтому исходный CI target был проверен
compile-only с одинаковыми GOOS/GOARCH и package:

```text
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -c -o /dev/null ./internal/tasking
exit 0

$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -race -c -o /dev/null ./internal/tasking
go: -race requires cgo
exit 2
```

Красный GitHub Actions run
[`33870475260`](https://github.com/EarthMoonTransit/go-learn/actions/runs/33870475260)
на baseline `a75ee43bae297a2a8e5aeee98e5603f00c126e88` подтвердил тот же defect в
реальном integration runner.

## Decision

- `RunGoOptions` несёт typed execution mode из runner в isolation layer;
- normal mode устанавливает `CGO_ENABLED=0`;
- race mode устанавливает `CGO_ENABLED=1`;
- legacy parser преобразует разрешённый `-race` в typed mode до исполнения;
- shell execution не используется, arbitrary command text mode не определяет;
- race support определяется явной GOOS/GOARCH matrix; linux/amd64 всегда
  считается supported;
- unsupported platform возвращает failing `infrastructure-unsupported`, а
  integration subtest может skip только после этой проверки.

## Pure-Go boundary

CGO=1 доступен только инструментальному race detector. Static package gate
отклоняет `import "C"`, cgo/compiler/linker directives, native source/object
files, external module dependencies и `toolchain`. Offline stdlib-only policy
и изоляция workspace сохранены.

## Failure diagnostics

Internal variant result различает assembly, compile, test, timeout,
race-unsupported и race/cgo-incompatible failures. Diagnostic извлекается из
Go JSON output, очищается от isolation/workspace/task paths и ограничивается
2048 Unicode code points. Starter и mutant засчитываются только при настоящем
test failure; infrastructure и compile failures не маскируются.

## Local acceptance

Проверено локально на darwin/arm64:

| Command | Result |
| --- | --- |
| `go test ./...` | pass |
| `go test -race ./...` | pass |
| `go vet ./...` | pass |
| `go test -tags=taskintegration ./...` | pass; supported race variant executed |
| build `task-check`, `validate`, `task-report` | pass |
| one `task-check -all -output` | pass |
| validator reuse через `-checks` | pass; no task execution |
| report check reuse через `-checks` | pass; no task execution |
| second `task-check` + byte comparison | pass |

Обе collection имеют SHA-256
`b30042bab3daca08962aee1447b221e0eac80733e264683b482938391de674da`.

## Execution counters

Production collection:

```yaml
sessions: 1
task_checks: 3
workspaces: 15
variant_runs: 0
legacy_commands: 15
child_go_runs: 15
```

Tagged integration session на поддерживаемой платформе вычисляет expectation из
шести dynamic checks × пяти variants, одного race variant и одной legacy
command: 31 variants + 1 legacy = 32 child Go commands/workspaces. Structural
invariants остаются строгими. Fast tests используют fake runner; validator и
report consumers не запускают task code повторно.

## Protected artifacts

Значения до и после совпадают; для task directories hash считается от
сортированного per-file `shasum -a 256` listing.

| Artifact | Before | After |
| --- | --- | --- |
| `tasks/ready/defined-types-001` | `a56d73571d38124735653800129e5029e03b5374d018b6a32f4c79a40f9f30e2` | same |
| `tasks/ready/variadic-sum-001` | `c5a429509d6b20f0d130d593d605a3947e7e10af3880f372a10905c0cf83c7b2` | same |
| `tasks/ready/zero-state-001` | `f21b87df35fa46deeb69b36027154197cc5e0529bce42b6858626ec815fbec6d` | same |
| `schemas/task.schema.json` | `987c1830cbdcb02702fd7e6ddb89362e76d479df14ade510cabcd20dc4dc30fb` | same |
| `curriculum/skills.yaml` | `4cc68407da591576cbce72117bc9a92993d491084ec736c26e5921323b269f16` | same |
| `catalog/reviews/exercism-go.yaml` | `05bf10d9bdbc282ca66c8b0c5bc9129deab1475be5a53bab3c303d51ec698f03` | same |
| `catalog/seed-plans/exercism-foundation-pilot.yaml` | `0cbbf900b5da2ba75460bd0f2d54c832f64e1debeaa2f0daadfe2025c134a2c5` | same |

До фактического post-push GitHub Actions run remote status остаётся `pending`.
