# Phase 3A.1b execution counts

Структурные метрики получены из call graph baseline `9636410` и из
instrumented runner после изменения. В `child Go commands` считаются только
nested вызовы Go toolchain из task verification; outer CI-команды `go test`,
`go vet` и `go build` не включены.

## До и после

| Метрика полного workflow | До | После |
| --- | ---: | ---: |
| Full-bank execution passes | 4 | 1 |
| Execution-capable `CheckAll` invocations | 17 | 1 |
| Bank task package checks | 12 | 3 |
| Все реальные task package checks, включая synthetic tests | 72 | 10 |
| Execution sessions / isolated cache roots | 60 | 2 |
| Task variant runs | 168 | 31 |
| Child Go commands в обычном `go test ./...` | 189 | 0 |
| Child Go commands в outer race suite | 189 | 0 |
| Child Go commands в tagged integration suite | 0 | 32 |
| Child Go commands в production CLI graph | 60 | 15 |
| Child Go commands всего | 438 | 47 |

`CheckAll` в fake-runner unit tests не считается execution-capable invocation:
он проверяет orchestration и transformations, но не создаёт процесс и cache.

## Раскладка baseline

- обычные tests: 173 tasking + 15 validation + 1 isolation = 189 child Go
  commands;
- outer race повторял те же 189 nested commands;
- validator: один `BuildAggregateReport → CheckAll` (15 legacy commands) и
  второй прямой legacy pass (15);
- `task-check -all`: ещё 15 legacy commands;
- `task-report -check → BuildAggregateReport → CheckAll`: ещё 15;
- production CI поэтому проверял каждую из трёх ready packages четыре раза.

Семнадцать execution-capable `CheckAll` складывались из семи в обычных tests,
семи в outer race и трёх CLI invocations. Отдельные legacy helpers объясняют,
почему full-bank passes было четыре, хотя production `CheckAll` — три.

## Новая раскладка

- fast unit и fast race: 0 nested Go commands и 0 real execution sessions;
- tagged integration: одна session, 31 variants и одна legacy command = 32;
- production `task-check -all`: одна session, три task checks, 15 отдельных
  legacy workspaces/commands;
- validator с `-checks`: 0 task checks и 0 child commands;
- `task-report -checks`: 0 task checks и 0 child commands.

Внутри любого непустого CheckCollection enforced invariants:
`sessions == 1`, `task_checks == len(tasks)`,
`child_go_runs == variant_runs + legacy_commands` и
`workspaces == child_go_runs`.
