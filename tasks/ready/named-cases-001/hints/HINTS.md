# Progressive hints

1. t.Run creates the child test and supplies its *testing.T.
2. The ordinary loop waits for each non-parallel subtest.
3. Pass the callback to t.Run instead of calling it with the parent.
