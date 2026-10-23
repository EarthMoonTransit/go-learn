# Reference predictions made before commands

Producer parameter is send-only: send and close are permitted, receive is rejected.
Receiver parameter is receive-only: receive is permitted, send and close are rejected.
The owner retains the bidirectional channel and narrows it at both helper calls.
These directions restrict operations, not buffering or exclusive sender ownership.

| Input | Predicted readings (Value, OK, Ready) | Send panic | Second-close panic |
| --- | --- | --- | --- |
| nil/empty | (0,false,true), (0,false,true) | true | true |
| 0 | (0,true,true), (0,false,true), (0,false,true) | true | true |
| 3,0,-2 | (3,true,true), (0,true,true), (-2,true,true), (0,false,true), (0,false,true) | true | true |
| -4,-4,0,8 | (-4,true,true), (-4,true,true), (0,true,true), (8,true,true), (0,false,true), (0,false,true) | true | true |

The native producer receive probe should be rejected as a receive from a send-only
channel. Receiver send and receiver close should each be rejected against a
receive-only parameter. Removing only the inserted operation restores the correct
helper contract. Widening both parameters makes those operations compile, but
violates the stated API permissions even when ordinary behavior still passes.

Missing-close keeps queued values but gives Ready=false after exhaustion. The
nonblocking send attempt does not panic; the owner's subsequent first close does
not panic. Discarding the buffer gives closed-empty readings too soon. Treating
zero as closure makes the queued-zero OK flag false. Omitting owner attempts leaves
false flags. Hardcoding true flags passes ordinary observations but does not
execute either required post-close operation, so source/method review rejects it.

Each real post-close operation panics inside its own synchronous recovery call.
No goroutine, timer, unbounded receive or test timeout is needed for these predictions.
