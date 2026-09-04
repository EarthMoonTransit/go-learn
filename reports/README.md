# Generated audit reports

Файлы в этом каталоге генерируются детерминированно командой
`go run ./cmd/catalog-report -root . -source <source-id>`. Они содержат только
метаданные, hashes и curator findings; upstream-условия, код и решения не
копируются.

Для Exercism отчёты также используют versioned concept map, reviewed skill
evidence и зарегистрированные exact revisions `problem-specifications`. Их
повторная генерация с теми же входами должна быть byte-identical.
