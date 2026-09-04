# Task specification

Статус документа: normative, ФАЗА 3A. Индексируемые данные определяет
`schemas/task.schema.json`; lifecycle, review, provenance и quality report имеют
отдельные versioned schemas. Совместимость schema v1 описана в ADR 0004.

## Единица поставки

Каждая новая задача — самодостаточный каталог `tasks/<state>/<task-id>/`. Она
проверяет ровно один primary skill; secondary skills допустимы только когда без
них нельзя получить реалистичное упражнение. Primary, secondary и prerequisites
не пересекаются. `new_concepts` не пересекаются с prerequisites и для нового
Phase 3A package содержат не более двух skills.

Обязательные артефакты:

- `task.yaml` — индексируемые данные schema v1;
- `lifecycle.yaml` — текущее состояние и непрерывная история переходов;
- `TASK.md` — техническое условие, API, требования и ограничения;
- `go.mod` — отдельный stdlib-only module `task.local/<task-id>`;
- `starter/` — минимальная компилируемая реализация учащегося;
- `tests/public/` — видимая suite и `*_test.go` binding к starter;
- `tests/grader/` — скрытая suite и `*_test.go` binding к starter;
- `solution/reference/` — реализация и binding к public/grader suites;
- `solution/mutants/<mutant-id>/` — различимые неверные реализации и bindings;
- `review/review.yaml`, `review/notes.md` — reviewer decision и evidence;
- `hints/HINTS.md` — прогрессивные подсказки без готового решения;
- `provenance.yaml` — полная цепочка источников и редакторских решений;
- `ATTRIBUTION.md` — файл, указанный provenance.

Grader tests, reference, mutants и внутренние review artifacts исключаются из
public student export.

## Контракт `TASK.md`

Обычное условие содержит не более 160 слов. Обязательны секции `Context`, `API`,
`Requirements`, `Constraints`: context занимает не более двух предложений; API
задан явно в fenced Go block; requirements содержат 4–7 проверяемых пунктов;
constraints содержат хотя бы один явный пункт. Условие не содержит сюжета,
объяснения решения или скрытых требований. Соответствие текста тестам,
story-free формулировка и отсутствие solution leakage подтверждаются reviewer.

## Исполняемый контракт

Implementation и bindings могут быть разными Go packages внутри одного модуля.
Public/grader bindings импортируют starter. Reference и каждый mutant имеют
собственный `*_test.go`, запускающий обе suites против выбранной реализации.
Starter компилируется, но падает хотя бы на обязательной public и full проверке;
reference проходит две повторные full проверки; каждый mutant компилируется и
обнаруживается grader. Для concurrency reference дополнительно проходит race.

`go.work`, symlinks, нерегулярные файлы и внешние module dependencies запрещены.
Authored code выполняется только в disposable copy согласно ADR 0004.

## Поля `task.yaml`

- `schema_version`: версия индексируемого контракта; сейчас `1`.
- `id`: стабильный kebab-case id с трёхзначным суффиксом; не переиспользуется.
- `title`: короткое техническое название без сюжета.
- `task_type`: `implement`, `debug`, `predict`, `refactor`, `test-writing` или
  `performance`.
- `skills.primary`: ровно один assignable id из taxonomy v3.
- `skills.secondary`: уникальные ids taxonomy; максимум три по schema v1.
- `prerequisites`: уже освоенные skills, необходимые до начала задачи.
- `new_concepts`: skills, явно вводимые этой встречей; максимум два в Phase 3A.
- `difficulty`: `introductory`, `basic`, `intermediate` или `advanced`.
- `estimated_time_minutes.min/max`: диапазон активной работы.
- `source`: registry id, repository, path, exact commit SHA и SPDX license.
- `transformation_type`: `original`, `adapted`, `translated-adaptation`,
  `combined` или `metadata-only`.
- `validation_commands`: детерминированные CI-команды legacy schema v1;
  `task-check` дополнительно использует фиксированную матрицу ADR 0004.
- `quality_status`: совместимое отображение lifecycle: `draft → draft`,
  `in-review → review`, `ready → ready`, `rejected → quarantined`, `retired →
  retired`. `validated` остаётся enum schema v1, но не lifecycle state.

Source id существует в registry; repository/license/commit совпадают с audited
entry. `reference-only` и `concepts-only` допускают только metadata-only
provenance и обязательный clean-room review. `original` использует внутреннюю
entry `go-practice-bank`; пока она не зарегистрирована с лицензией и canonical
commit, original-задача не может пройти promotion.

## Lifecycle и стабильность

Состояния: `draft`, `in-review`, `ready`, `rejected`, `retired`. Разрешённые
переходы и atomic promotion определены ADR 0004. Изменение schema v1, enum или
skill id требует отдельного ADR и миграционного плана.

Три ready packages, существовавшие до ADR 0004, являются legacy v1 и могут не
иметь новых authoring-файлов. В ФАЗЕ 3A они не мигрируются и не переписываются.
