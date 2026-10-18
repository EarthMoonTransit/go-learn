# ADR 0005: hardening исполнения, gates и promotion

- Статус: accepted
- Дата: 2026-09-04
- Решение: ФАЗА 3A.1

## Контекст

Контракт ADR 0004 позволял task-authored `*_test.go` связывать reference и
mutants с suites. Пустой reference binding мог создать ложный pass, а
безусловный `t.Fatal` в mutant binding — ложную «детекцию». Legacy manifests
исполнялись через `sh -c`, процесс наследовал почти всё окружение, inventory
мог молча пропустить неверно названный каталог, а promotion проверял исходное,
но не подготовленное целевое состояние.

## Решение: checker-owned assembly

Новый package использует единственный исполняемый контракт:

- `starter/`, `solution/reference/` и каждый mutant содержат только обычные
  `*.go` с package `subject`; `*_test.go` там запрещены;
- `tests/public/` и `tests/grader/` содержат только `*_test.go` с package
  `subject_test` и обращаются к `task.local/<id>/subject`;
- checker создаёт новый temporary module, сам копирует выбранную реализацию в
  `subject/` и затем копирует туда canonical tests;
- starter проверяется с public и с public+grader; reference — дважды с
  public+grader; каждый mutant — с grader; concurrency reference — с `-race`.

Task package больше не выбирает, какие tests запускаются против reference или
mutant. Пустой public test обнаруживается тем, что starter неожиданно проходит.
Пустой grader обнаруживается выжившим реалистичным mutant.

## Execution boundary

Никакая команда не запускается в source checkout. Исполняется только Go binary
текущего runtime по абсолютному пути; произвольного process API у task checker
нет. Среда строится из allowlist: отдельные `HOME`, `TMPDIR`, `GOCACHE`,
`GOMODCACHE`, плюс `GOWORK=off`, `GOPROXY=off`, `GOTOOLCHAIN=local`,
`GOFLAGS=-mod=readonly`, `GOENV=off` и фиксированный `PATH`.
Уточнение реализации 2026-10-18: до первого Go command создаётся telemetry
mode file `off` в isolated user config; Windows `APPDATA` также указывает внутрь
disposable home. `GOTELEMETRY` — read-only значение `go env`, а не override
режима ([Go telemetry configuration](https://go.dev/doc/telemetry#configuration)).
Типизированный execution mode задаёт `CGO_ENABLED=0` обычным runs и
`CGO_ENABLED=1` только race runs; runner не выводит mode из текста команды.
Не связанные переменные, включая secrets, не наследуются.

Разрешение cgo для самого race detector не меняет pure-Go task contract.
Checker статически отклоняет `import "C"`, cgo compiler/linker directives,
C/C++/Objective-C и другие native source/object files. Task `go.mod` по-прежнему
не допускает `toolchain`, external dependencies и прочие dependency directives.

Legacy `validation_commands` проходят лексический parser до первого запуска.
Разрешены только `go test` и `go vet`, безопасные test flags и явные `./...`
package paths. Pipes, redirections, traversal, `$()`/environment expansion,
shell metacharacters и другие программы отклоняются. Затем каждая команда
запускается напрямую в отдельной копии package. Это hygiene boundary для
доверенного repository content, а не OS sandbox против hostile Go code.

Race support проверяется явно для текущей GOOS/GOARCH. Неподдерживаемая
платформа получает `infrastructure-unsupported`, а не pass; на поддерживаемой
платформе отсутствие compiler/cgo или любой race failure остаётся failing gate.
Variant failure различает assembly, compile, test, timeout и race/cgo; bounded
diagnostic очищается от absolute temporary paths до попадания в отчёт.

## Toolchain policy

`config/toolchain.yaml` — единый machine-readable источник версии языка,
CI-линии, module prefix и legacy exceptions. Новые task modules используют Go
1.27, только stdlib, без `toolchain`, `go.work` и dependency directives. Три
неизменяемые legacy packages явно закреплены на Go 1.24. Template, init,
checker, validator и CI согласованы с policy; schema task/provenance не менялась.

## State-aware gates и отчёт

`draft` и `in-review` обязаны проходить technical gates; review approval пока
не является их общим pass criterion. `ready` требует technical pass и review
approval. `rejected`/`retired` проверяют core lifecycle/provenance audit без
исполнения. Legacy ready имеет явный compatibility path и не считается
promotion-eligible.

`task-check -all` сначала проверяет все пять каталогов: malformed entries и
duplicate ID между любыми states являются ошибкой даже при state filter.
`task-report` запускает реальные state-aware checks, а `-check` только сравнивает
детерминированный результат с committed artifacts и никогда не пишет файлы.

## Promotion и crash semantics

`task-promote` берёт per-ID lock, запрещает дату раньше последнего transition,
готовит обе metadata в памяти, валидирует schemas/lifecycle и проверяет полную
копию будущего target state. Только после этого два файла заменяются atomic
writes, а directory перемещается одним rename без overwrite. При любой обычной
возвращаемой ошибке исходные metadata восстанавливаются byte-for-byte.

Это не полная filesystem transaction. Crash процесса между metadata writes и
directory rename может оставить несовпадающие metadata или stale lock.
`task-check -all` обнаружит mismatch; восстановление выполняется из VCS, а stale
lock удаляется вручную только после проверки отсутствия живого promotion.

## Последствия

Три legacy ready directories остаются byte-identical. Checker больше не доверяет
authored bindings, CI проверяет race/vet/all-task/report drift, а student workflow
работает с scratch-copy и не изменяет canonical repository package.
