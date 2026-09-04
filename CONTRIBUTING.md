# Разработка через pull request

Server-side branch protection для приватного репозитория недоступна на текущем
GitHub plan. Поэтому до смены visibility или тарифа действует обязательный
локальный и процессный gate.

## Один раз для каждого checkout

```bash
git config --local core.hooksPath .githooks
```

Tracked `pre-push` hook после этого запрещает обновлять remote `master` из
локального Git. Это защита от случайной ошибки, а не security boundary: GitHub
UI, другой checkout и `--no-verify` технически могут её обойти.

## Workflow batch

```bash
git switch master
git pull --ff-only
git switch -c codex/<topic>

# внести изменение и выполнить acceptance из AGENTS.md

git push -u origin codex/<topic>
gh pr create --base master --head codex/<topic>
gh pr checks --watch
```

PR сливается только после успешного check `validate`. Красный или pending check
запрещает merge по правилам проекта, даже если GitHub UI технически оставляет
кнопку доступной. После merge push-event повторно проверяет новый `master`.

Workflow запускается для PR в `master` и для post-merge push в `master`.
Обычный push feature branch без открытого PR не создаёт дублирующий CI run.
