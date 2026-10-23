# Hints

- A slow callback does not pause elapsed time or restart the total budget.
- Distinguish the time carried by a queued tick from the time a callback starts.
- Keep cleanup at the scope that created both clock objects.
