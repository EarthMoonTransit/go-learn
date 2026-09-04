# Phase 3A.1 adversarial test report

Date: 2026-09-04

## Воспроизведённые pre-hardening outcomes

Synthetic reproducer запускает старый authored-binding pattern с заведомо
неверной implementation. Пустой reference test завершается успешно, не вызывая
grader. Mutant test с unconditional `t.Fatal` завершается ошибкой независимо от
grader. Именно эти process outcomes старый checker принимал соответственно за
reference pass и mutant detection, поэтому обе проверки можно было подделать.

Новый checker отклоняет оба файла уже package contract gate и никогда не
исполняет solution-authored bindings.

## Adversarial fixtures

| Fixture / attack | Expected hardened result |
| --- | --- |
| no-op reference binding | package contract fail |
| unconditional-fail mutant binding | package contract fail |
| no-op public test | starter-public fail: starter unexpectedly passes |
| no-op grader test | mutant-coverage fail: realistic mutant survives |
| correct reference | shared public+grader pass twice |
| realistic off-by-one mutant | detected by canonical grader; survives no-op grader |
| `curl` / `env` | legacy parser rejection |
| command substitution | legacy parser rejection |
| output redirection / pipeline | legacy parser rejection |
| traversal / outside marker write | preflight rejection; marker absent |

## Additional regression coverage

- minimal environment does not contain a synthetic secret and roots
  `HOME`, `TMPDIR`, `GOCACHE`, `GOMODCACHE` inside temp;
- complete legacy command list is parsed before any command is run;
- source checkout remains unchanged after direct Go execution in a copy;
- broken draft and broken in-review remain visible to `CheckAll` and fail;
- malformed task directory and cross-state duplicate fail discovery;
- aggregate output is deterministic and includes technical/review/starter/
  reference/mutant/race evidence;
- stale report comparison fails without rewriting the other artifact;
- promotion tests cover backdated dates, future lifecycle mismatch, collision,
  concurrent lock, successful `draft → in-review → ready` and exact rollback;
- toolchain tests cover policy loading, legacy exception, forbidden dependency
  directives and forbidden `toolchain` directive.

All fixtures are synthetic and live under `internal/testdata`; no real task was
created or promoted.
