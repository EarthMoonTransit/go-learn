# Progressive hints

1. Decode once for the object and once more to verify EOF.
2. A pointer destination distinguishes null from an empty object.
3. Discard partial decoded state on every rejection.
