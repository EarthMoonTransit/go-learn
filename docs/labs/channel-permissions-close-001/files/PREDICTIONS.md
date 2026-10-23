# Predictions made before commands

Record the chosen helper parameter types and permitted/forbidden operations.

| Input | Predicted readings (Value, OK, Ready) | Send panic | Second-close panic |
| --- | --- | --- | --- |
| nil/empty | TODO | TODO | TODO |
| 0 | TODO | TODO | TODO |
| 3,0,-2 | TODO | TODO | TODO |
| -4,-4,0,8 | TODO | TODO | TODO |

For each compiler probe, predict the failing operation and explain why removing
that operation repairs it while widening the helper permissions violates TASK.
Predict missing-close, discarded-buffer and zero-as-closure observations.
Explain why unrestricted helpers and hardcoded panic flags can pass behavior yet
fail the declared method. Compare these predictions with the retained actual results.
