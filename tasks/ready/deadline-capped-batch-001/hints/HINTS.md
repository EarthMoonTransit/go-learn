# Hints
- Compare the three possible upper bounds before choosing which context to pass onward.
- A later callback inherits time already spent; consider where the relative budget begins.
- Account for every cancellation function you create, including an intermediate helper invisible to callbacks.
- Callback results and the context's status have separate contracts.
