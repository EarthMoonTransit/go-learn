# Task specification

Статус документа: normative, ФАЗА 1. Машиночитаемый контракт находится в
`schemas/task.schema.json`; при расхождении schema имеет приоритет.

## Единица поставки

Каждая готовая задача — самодостаточный каталог `tasks/ready/<task-id>/`. Задача
проверяет ровно один primary skill; secondary skills допустимы только когда без
них нельзя получить реалистичное упражнение. Primary, secondary и prerequisites
не пересекаются; `new_concepts` может включать primary/secondary, но не
prerequisite.

Обязательные артефакты:

- `task.yaml` — индексируемые данные;
- `README.md` — техническое условие, сигнатуры, примеры и ограничения;
- `starter/` — минимальный компилируемый модуль или пакет учащегося;
- `tests/public/` — видимые примеры основного контракта;
- `tests/grader/` — edge cases и анти-hardcode проверки;
- `solution/reference/` — минимальное идиоматичное решение;
- `provenance.yaml` — полная цепочка источников и редакторских решений.

Grader tests и reference solution хранятся в репозитории для внутренней сборки,
но исключаются из публичного student export.

## Поля `task.yaml`

- `schema_version`: целая версия контракта; сейчас `1`.
- `id`: стабильный kebab-case идентификатор. После публикации не переиспользуется.
- `title`: короткое техническое название без сюжета и маркетинговых формулировок.
- `task_type`: одно из `implement`, `debug`, `predict`, `refactor`, `test-writing`,
  `performance`.
- `skills.primary`: ровно один id из `curriculum/skills.yaml`.
- `skills.secondary`: уникальные ids taxonomy; обычно 0–2.
- `prerequisites`: уже освоенные skills, необходимые до начала задачи.
- `new_concepts`: skills, которые задача впервые вводит явно. Primary skill обязан
  входить сюда для introduction-встречи и отсутствовать для practice/transfer.
- `difficulty`: `introductory`, `basic`, `intermediate` или `advanced`.
- `estimated_time_minutes.min/max`: реалистичный диапазон активной работы; min не
  больше max.
- `source`: primary provenance (`id`, repository, path, exact 40-char commit SHA,
  SPDX license). `source.id` обязан существовать в registry; repository и license
  совпадают с записью, а для адаптации совпадает и audited commit. Для собственной
  авторской работы используется внутренняя source entry `go-practice-bank` и
  commit, на котором задача впервые принята. Эту запись можно создать только
  после появления canonical remote и первого реального commit.
- `transformation_type`: `original`, `adapted`, `translated-adaptation`,
  `combined` или `metadata-only`.
- `validation_commands`: argv-free shell commands, выполняемые из task package.
  Команды должны быть детерминированны и пригодны для CI.
- `quality_status`: `draft`, `review`, `validated`, `ready`, `quarantined` или
  `retired`.

Дополнительные поля запрещены, чтобы опечатки не становились молчаливыми данными.
Обязательные provenance-детали, которым тесно в `task.yaml`, находятся в
`provenance.yaml`, но не заменяют основной `source`.

## Пример формы (не готовая задача)

```yaml
schema_version: 1
id: slices-append-aliasing-001
title: Control slice aliasing during append
task_type: debug
skills:
  primary: slices.append-aliasing
  secondary: [slices.capacity-bounds]
prerequisites: [slices.model]
new_concepts: [slices.append-aliasing]
difficulty: basic
estimated_time_minutes: {min: 20, max: 35}
source:
  id: go-practice-bank
  repository: https://example.invalid/go-practice-bank
  path: authoring/slices-append-aliasing-001
  commit: 0123456789abcdef0123456789abcdef01234567
  license: CC-BY-4.0
transformation_type: original
validation_commands:
  - go test ./...
quality_status: draft
```

Заглушки `example.invalid` и фиктивный SHA допустимы только в документации.

## Стабильность

Изменение смысла поля, enum или skill id требует отдельного ADR и миграции всех
затронутых manifests. Расширение schema не смешивается с content batch.
