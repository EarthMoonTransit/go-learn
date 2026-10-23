# Hints

- Identify which context owns the next operation. A previous operation's released scope cannot be its parent.
- Place the cleanup boundary after one synchronous callback, before moving to another.
- Give your private helper an explicit context argument first. Keep a returned value even when its paired error is non-nil.
