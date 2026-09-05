# Progressive hints

1. A trimming cutset describes individual characters; a prefix describes one complete sequence at the beginning.
2. Check both the remaining string and the boolean. An empty prefix is a successful match even though nothing changes.
3. Compare the contracts of strings.TrimPrefix and strings.CutPrefix. A repeated leading prefix still requires only one removal.
