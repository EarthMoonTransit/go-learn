# Task package template

Этот каталог задаёт шаблоны и физические границы Phase 3A package, но сам задачей
не является. Не копируйте его прямо в `ready`: `task-init` создаёт из `*.tmpl`
новый `draft`, присваивает ID и подставляет только provenance metadata.

Template намеренно не содержит учебного API, starter code или tests. README в
пустых каталогах описывают контракт и не копируются в создаваемый package.

`task-init` по умолчанию создаёт metadata/provenance v2. Конкретная физическая
раскладка зависит от выбранного verification profile и описана в
`docs/TASK_CONTRACT_V2.md`; этот каталог остаётся content-free общей основой.
