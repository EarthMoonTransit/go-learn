# Task packages

Готовая задача хранится как изолированный пакет:

```text
tasks/ready/<task-id>/
├── go.mod
├── task.yaml
├── README.md
├── ATTRIBUTION.md
├── starter/
├── tests/
│   ├── public/
│   └── grader/
├── solution/
│   ├── reference/
│   └── incorrect/
└── provenance.yaml
```

`task.yaml` валидируется по `schemas/task.schema.json`. Каждый task package —
отдельный Go module: public tests запускаются против `starter`, а reference и
намеренно неверная реализация проверяют силу public/grader suites. Физические
границы артефактов показаны в `tasks/_template/`; его README-файлы не являются
тестами или решением.

Для решения откройте `README.md` выбранной задачи, меняйте только `starter/` и
запускайте из каталога задачи:

```bash
go test ./tests/public
```
