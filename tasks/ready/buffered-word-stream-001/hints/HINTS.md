# Hints

- Scanner can return a final word before exposing its reader error.
- A destination that already buffers data still belongs to the caller.
- Resolve write and flush errors before returning a scanning error.
