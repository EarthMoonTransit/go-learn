# Progressive hints

1. The supplied gatedWork helper lets work start and remain blocked. Use synctest.Wait before observing a phase, and defer release even when an assertion fails.
2. Distinguish a completion observed while work is held from one observed after release. A closed channel reports ok=false; a sent value reports ok=true.
3. Hold two callbacks independently. Compare their completion handles and release both gates on every exit; completing one call must not complete the other.
