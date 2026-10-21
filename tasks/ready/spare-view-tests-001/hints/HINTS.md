# Hints

A view can allow some appends to reuse storage before forcing a later append to detach. Keep the original backing slice long enough to observe both boundaries.
