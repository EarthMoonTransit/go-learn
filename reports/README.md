# Generated audit reports

Файлы в этом каталоге генерируются детерминированно командой
`go run ./cmd/catalog-report -root . -source <source-id>`. Они содержат только
метаданные, hashes и curator findings; upstream-условия, код и решения не
копируются.

Для Exercism отчёты также используют versioned concept map, reviewed skill
evidence и зарегистрированные exact revisions `problem-specifications`. Их
повторная генерация с теми же входами должна быть byte-identical.

Pedagogical triage reports генерируются отдельно командой
`go run ./cmd/candidate-report -root .`. Их нормативные входы —
`catalog/reviews/exercism-go.yaml`, `catalog/clusters/exercism-go.yaml` и
`catalog/seed-plans/exercism-foundation-pilot.yaml`; отчёты не являются
источником решений и не превращают seeds в готовые задачи.
