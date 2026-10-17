# Hints

1. Relative child budgets and an earlier parent deadline must coexist.
2. The owner of the derived context also owns its cancellation function, even when the operation succeeds quickly.
3. The operation owns its result and cancellation response. Context expiration is not permission to leave its cleanup running after the wrapper returns.
