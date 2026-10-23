# Acknowledged snapshots

Implement Exchange(values []int, out chan<- *Packet, ack <-chan struct{}). Packet contains Value and Twice. Use one packet for the entire call. For each input, fill both fields, send its pointer, and wait for exactly one acknowledgement before modifying it again. Close out after the final acknowledgement, including empty input. The caller reads or copies the packet before acknowledging; it retains no right to read it afterward. Channels are non-nil and callers obey that protocol.

Write a regression test with distinct zero, negative and positive values. Read both fields before acknowledging, preserve each snapshot independently, and join the sender before completing the test. Collect mismatches without leaving the sender blocked. Explain the publication and acknowledgement happens-before edges using the actual source. Interpret the supplied unsafe native race reports and compare the same probes after repair. Separately show that a race-free implementation with incorrect Twice still fails value assertions. Scope every clean detector result to the paths actually exercised.

