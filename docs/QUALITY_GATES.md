# Quality gates

Задача продвигается только вперёд по состояниям `draft → review → validated →
ready`. `quarantined` блокирует публикацию; `retired` сохраняет id и provenance.

## Gate 1 — metadata

- `task.yaml` проходит JSON Schema и semantic validator;
- все skill ids и source id существуют;
- source metadata совпадает с registry, commit является exact SHA;
- нет конфликтов primary/secondary/prerequisites/new concepts;
- id уникален и совпадает с именем task directory.

## Gate 2 — legal/provenance

- usage mode допускает transformation type;
- полная цепочка есть в `provenance.yaml`;
- notices и attribution сформированы;
- reference-only материал не воспроизведён;
- неоднозначность переводит задачу в quarantine, а не маскируется.

## Gate 3 — instructional design

- один наблюдаемый primary skill;
- prerequisites достаточны и не образуют скрытый скачок;
- условие без сюжета, с входами, выходами, ограничениями и примерами;
- difficulty/time согласуются с rubric;
- задача не является поверхностным дубликатом;
- встреча выполняет роль introduction, practice или transfer.

## Gate 4 — executable correctness

- starter компилируется в ожидаемом режиме;
- public и grader tests действительно запускаются;
- reference solution проходит оба набора;
- минимум одна реалистичная неверная реализация проваливает grader tests;
- tests детерминированны, изолированы от сети/часов/случайности либо контролируют
  эти зависимости;
- для concurrency запускается `go test -race`; для fuzz/benchmark есть отдельные
  ограниченные команды.

## Gate 5 — maintainability

- `go fmt`, `go vet` и repository validation зелёные;
- нет secrets, grader leakage и незафиксированных бинарных fixtures;
- task package самодостаточен;
- reviewer, toolchain version и команды проверки воспроизводимы в CI.

## Repository gates

ФАЗА 1:

```bash
go test ./...
go run ./cmd/validate -root . -phase1
```

Неполные authoring batches до достижения полного curriculum coverage:

```bash
go test ./...
go run ./cmd/validate -root . -phase1 -run-commands
```

ФАЗА 2+ и release:

```bash
go test ./...
go run ./cmd/validate -root . -coverage -run-commands
```

`-phase1` разрешает неполный `tasks/ready`, но не ослабляет schema, taxonomy,
registry или проверки уже существующих manifests. С `-run-commands` он также
исполняет validation commands каждой ready-задачи. `-coverage` требует все
minimum occurrences и диапазон 350–500 ready tasks фундаментального трека.
