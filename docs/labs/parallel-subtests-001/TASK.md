# Observe parallel subtest boundaries

Prerequisites: testing.subtests and goroutines.launch-lifecycle. First encounter: testing.parallel.

Edit only parallel_test.go; fixture helpers are supplied.

- Author TestParallelTable with cases alpha→11, beta→-2 and gamma→0 using t.Run and t.Parallel.
- Give each case its own mutable fixture. Call checkFixture with that case's key and value; never hold a fixture lock across t.Parallel.
- Use parent-local atomic markers to assert that resumed child bodies see the parent-body end marker. Record each completed child once.
- After registering children, assert zero completed children in the parent body. Set the end marker immediately before returning. Parent Cleanup must observe three completed children.
- Log parent-body end, each resumed child and parent Cleanup. Predict the partial order, then compare native PAUSE/CONT and actual markers without prescribing sibling order.
- Run the finite reference and wrong specimens in PROCEDURE.md; explain early parent reads and shared-fixture failures.

Submit your test, outputs and explanation. No sleeps, parent-body waiting, sibling barrier or timing grade.
