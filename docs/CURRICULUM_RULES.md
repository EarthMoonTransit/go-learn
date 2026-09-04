# Curriculum rules

## Объём и границы

Фундаментальный трек содержит 350–500 задач. Его границы заданы доменами в
`curriculum/skills.yaml`: от toolchain и типов до context включительно. HTTP,
SQL, observability, deployment и production architecture относятся к следующим
трекам и не увеличивают coverage фундамента.

Taxonomy — граф навыков, а не порядок глав учебника. Ребро `prerequisites`
означает, что prerequisite должен встречаться и пройти проверку раньше dependent
skill. Циклы запрещены.

## Что считается встречей

Встреча учитывается только для `ready`-задачи с прошедшими quality gates:

- `primary` — навык является главным объектом оценки;
- `total` — навык указан как primary или secondary;
- одна задача даёт не более одной встречи одному skill;
- prerequisite не считается встречей;
- quarantined/retired/draft/review/validated без `ready` не входят в coverage.

`minimum_occurrences.primary` и `.total` — нижние границы, не квоты. Taxonomy v3
использует три обоснованных уровня: 2/4 для specialist-механик, 3/5 как базовый
контракт и 4/6 для навыков с высоким риском ошибок владения, lifetime, encoding
или concurrency. Для 142 assignable skills это 409 primary и 693 total встреч.
Решение и полная миграция зафиксированы в
`docs/adr/0003-taxonomy-v3.md` и `curriculum/migrations/v2-to-v3.yaml`.

## Spacing и прогрессия

Для каждого skill встречи распределяются как минимум по трём различимым ролям:

1. introduction — изолированная механика с явным новым понятием;
2. practice — вариация с edge cases и меньшим scaffold;
3. transfer — применение вместе с ранее изученными навыками в новом API-контексте.

Если skill имеет четвёртую primary-встречу, она должна быть либо
debugging/test-writing, либо интеграционной. Две задачи с одинаковой сигнатурой,
fixtures и логикой решения не считаются разными встречами.

## Типы упражнений

Разрешены `implement`, `debug`, `predict`, `refactor`, `test-writing`,
`performance`. На каждый домен:

- не более 50% primary-встреч одного типа;
- минимум одна задача `debug`, `predict` или `test-writing`;
- performance используется только с доказуемым benchmark/complexity контрактом;
- формулировка техническая, без сюжетного обрамления;
- quiz без исполняемого проверяемого артефакта не считается задачей банка.

## Difficulty и время

- `introductory`: одна новая идея, 10–25 минут;
- `basic`: один primary skill и ограниченные edge cases, 15–40 минут;
- `intermediate`: сочетание 2–3 навыков, 30–75 минут;
- `advanced`: неоднозначные trade-offs, concurrency или performance, 45–120 минут.

Оценка времени проверяется на пилотной группе и хранится диапазоном. Difficulty не
выводится автоматически из номера темы.

## Mastery

Mastery по skill достигнуто, когда учащийся независимо проходит критерии из
taxonomy в двух разнесённых transfer/practice задачах, не только повторяет
introduction. Coverage — необходимое, но недостаточное условие mastery.

## Управление изменениями

Новый skill требует уникального стабильного id, обоснования атомарности,
prerequisites, misconceptions, exercise types, minimum occurrences и наблюдаемых
mastery criteria. Удаление/слияние skill требует миграции task manifests и
coverage baseline.

Assignable skill всегда имеет `status: active`, `assignable: true` и полный
skill contract. Deprecated stub имеет нулевые minima, не участвует в графе и не
может использоваться в новых задачах. Исключение для старой задачи допустимо
только через точную migration-запись с task id, полем и SHA-256 манифеста;
глобальные aliases запрещены.
