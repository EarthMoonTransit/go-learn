# Обязательные правила для агентов

Эти правила действуют для всего репозитория и имеют приоритет над локальными
предпочтениями реализации.

1. Никогда не копировать и не перефразировать близко код, тесты, решения или
   учебный текст из источников с `usage_mode: reference-only`. Такие источники
   можно использовать только для проверки полноты тем и фактов.
2. Никогда не терять provenance. Любой кандидат и любая задача обязаны ссылаться
   на запись в `sources/registry.yaml`, точный commit и исходный path. Для
   оригинальной задачи используется внутренняя запись `go-practice-bank`; пока у
   репозитория нет canonical remote и первого commit, создавать такую запись и
   принимать original-задачи в `tasks/ready/` запрещено.
3. Не создавать сюжетные постановки. Условия должны описывать техническую
   ситуацию, входы, выходы, ограничения и наблюдаемое поведение без персонажей,
   компаний и декоративной истории.
4. Не добавлять задачу в `tasks/ready/` без работающих public и grader tests.
   Reference solution должна проходить оба набора; заведомо неверная реализация
   должна проваливать хотя бы grader tests.
5. Работать небольшими reviewable batches: не более 10 новых или существенно
   изменённых задач на один batch. Изменения taxonomy, adapters и tooling
   оформлять отдельно от content batches.
6. Не менять `schemas/task.schema.json`, идентификаторы skills или смысл полей
   без отдельного зафиксированного решения (ADR/issue) и миграционного плана.
7. Не помещать секретные grader tests и reference solutions в публичный export.
8. Перед завершением tooling batch запускать fast `go test ./...`,
   `go test -race ./...`, `go vet ./...`, затем один
   `go test -tags=taskintegration ./...`. Полный `task-check -all` выполняется
   один раз с `-output`; validator и `task-report -check` обязаны использовать
   тот же файл через `-checks`. Для release content batch validator использует
   `-coverage`; для infrastructure допустим `-phase1`.
9. Никогда не запускать authored task code или task tests непосредственно из
   source checkout. Использовать `task-check`; legacy validation commands также
   обязаны выполняться только в disposable copy.
10. Не перемещать task packages между lifecycle-каталогами вручную. Использовать
    `task-promote`, не обходить review metadata и не перезаписывать ready id.
11. Не коммитить и не пушить изменения напрямую в `master`. Каждый batch
    выполняется в отдельной ветке `codex/<topic>`, публикуется как PR и
    сливается только после успешного GitHub Actions check `validate`. Пока
    server-side branch protection недоступна, это обязательный process gate;
    локальный hook нельзя обходить через `--no-verify`.
