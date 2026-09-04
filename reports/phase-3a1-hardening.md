# Phase 3A.1 hardening report

Baseline: `6943879e13c8c0415b5d39e5d4c0e55443ea03ef`  
Date: 2026-09-04

## Scope

Изменены только execution model, lifecycle validation, quality reporting,
promotion, toolchain policy, CI и связанная документация. Новые учебные задачи,
UI, sources, taxonomy changes, candidate reviews и schema/provenance v2 не
создавались. Synthetic artifacts находятся только в `internal/testdata`.

## Закрытые обходы

- Legacy `validation_commands` больше не передаются `sh -c`: parser разрешает
  только `go test`/`go vet`, allowlisted flags и безопасные `./` package paths.
- Process environment создаётся из allowlist, а не из `os.Environ`.
- Reference/mutant `*_test.go` запрещены. Checker сам собирает `subject` и
  подключает canonical public/grader tests.
- Все lifecycle directories проверяются строго; malformed directory и duplicate
  ID не могут исчезнуть из отчёта из-за фильтра.
- Aggregate report содержит результаты исполнения, review, mutants и race, а
  не только inventory. `-check` не пишет и отклоняет drift.
- Promotion валидирует будущий state, использует per-ID lock, не допускает
  backdated transition и откатывает обычную metadata failure byte-for-byte.

## Execution trust boundary

Source checkout никогда не является command working directory. Каждый запуск
получает новый temporary workspace, отдельные home/cache/temp paths, offline
module resolution и прямой Go binary активного runtime. Symlinks и нерегулярные
файлы не копируются. Это защищает canonical checkout от обычных task checks, но
не является OS/container sandbox. Authoring code считается trusted repository
input; произвольный learner code потребует отдельной sandbox phase.

## State policy

| State | Required result |
| --- | --- |
| `draft` | technical pass; approve необязателен |
| `in-review` | technical pass; request-changes допустим |
| `ready` | technical pass и approve review |
| `rejected` | metadata, lifecycle, provenance audit |
| `retired` | metadata, lifecycle, provenance audit |
| legacy `ready` | explicit compatibility checks; promotion ineligible |

## Promotion guarantees

До source writes проверяются mutated schemas, lifecycle invariants и полная
копия future target state. Destination не перезаписывается. Два metadata files
пишутся atomic replacement, затем directory перемещается final rename. При
обычной возвращаемой ошибке восстанавливаются исходные bytes. Это не полная
транзакция: process crash между metadata writes и rename может потребовать VCS
recovery и ручного удаления stale lock после проверки.

## Toolchain и CI

`config/toolchain.yaml` фиксирует Go language 1.27, CI `1.27.x`, module prefix,
запрет task `toolchain` directive и три Go 1.24 legacy exceptions. Checker и
template/init читают policy. Уже используемые root dependencies находятся в
`vendor/`; CI запускает Go-команды с `-mod=vendor`, `GOPROXY=off`, `GOSUMDB=off`.

## Acceptance

| Command | Result |
| --- | --- |
| `go test ./...` | pass |
| `go test -race ./...` | pass |
| `go vet ./...` | pass |
| `go run ./cmd/validate -root . -phase1 -run-commands` | pass; 3/3 ready |
| `go run ./cmd/task-check -root . -all` | pass; 3 legacy reports |
| `go run ./cmd/task-report -root . -check` | pass; reports current |

## Legacy integrity

Directory digest is SHA-256 of the sorted per-file `shasum -a 256` listing.

| Directory | Before | After |
| --- | --- | --- |
| `defined-types-001` | `a56d73571d38124735653800129e5029e03b5374d018b6a32f4c79a40f9f30e2` | `a56d73571d38124735653800129e5029e03b5374d018b6a32f4c79a40f9f30e2` |
| `variadic-sum-001` | `c5a429509d6b20f0d130d593d605a3947e7e10af3880f372a10905c0cf83c7b2` | `c5a429509d6b20f0d130d593d605a3947e7e10af3880f372a10905c0cf83c7b2` |
| `zero-state-001` | `f21b87df35fa46deeb69b36027154197cc5e0529bce42b6858626ec815fbec6d` | `f21b87df35fa46deeb69b36027154197cc5e0529bce42b6858626ec815fbec6d` |

`schemas/task.schema.json` также не изменён:  
`987c1830cbdcb02702fd7e6ddb89362e76d479df14ade510cabcd20dc4dc30fb`.
