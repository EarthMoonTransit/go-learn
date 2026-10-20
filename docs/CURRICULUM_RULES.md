# Curriculum rules

[ADR 0034](adr/0034-finite-v1-course.md) defines a separate v1 route-readiness milestone with two reviewed encounters per foundation skill. The deeper bank minima, modes and types below remain unchanged and are not claimed by v1. Its limited hybrid-recursion rule supersedes the machine-proof prerequisite below; numeric credit and task lifecycle rules remain unchanged.

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
- lifecycle states `draft`, `in-review`, `rejected`, `retired` не входят в
  coverage; legacy schema value `validated` также не учитывается без `ready`.

`minimum_occurrences.primary` и `.total` — нижние границы, не квоты. Taxonomy v4
сохраняет три обоснованных уровня: 2/4 для узких механик, 3/5 как базовый
контракт и 4/6 для навыков с высоким риском ошибок владения, lifetime, encoding
или concurrency. Для 145 assignable skills это 418 primary и 708 total встреч.
Все 142 прежних active-контракта и 35 deprecated stubs сохранены; три новых
контракта добавлены через ADR 0031 и `curriculum/migrations/v3-to-v4.yaml`.
Историческая v2→v3 миграция продолжает проверяться против immutable v3 snapshot.

Encounter tier не задаёт scope. По ADR 0031 foundation включает 141 active skill,
specialist — четыре; полный `-coverage` сохраняет требования всех 145 и target
350–500 ready tasks. Раздельная числовая scope-проверка — следующая tooling
партия. Нынешний executable coverage gate проверяет primary/total и размер банка;
фактические modes, distinct types и outcomes требуют отдельного semantic review.
Ни количество встреч в банке, ни успешная числовая проверка не доказывают learner
completion или mastery. Recursion authoring дополнительно заблокирован до
проверенного assessment-механизма из ADR 0031.

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

## Первая волна после milestone 150

Активный план — [foundation-next](plans/foundation-next.md). Read-only подтемы
дополняют разделы 01–09 и ссылаются на stable task IDs; они не назначают
новый порядок, номера или prerequisites scheduler. Теория не является встречей,
attempt, hint, assistance, completion или mastery (ADR 0020).

Алгоритмичность допустима при явных prerequisites и наблюдаемом контракте.
Это дополнительный facet, а не новый activity или skill; отказы пересматриваются
с сохранением исторического решения и license gate (ADR 0021).

Новые LGWT-задачи и теория не выбираются. Исходные упражнения, наши адаптации,
originals, concept seeds и theory pages считаются отдельно. Отсутствие
пользовательских решений не блокирует подготовку и не доказывает mastery.
Существующая механика двух успехов сохраняется до отдельного retention ADR;
она не доказывает долгосрочное удержание навыка.
