package subject_test

import (
	subject "task.local/transfer-shared-values-001/subject"
	"testing"
)

func TestTransferSmallStateSpace(t *testing.T) {
	for d := 0; d <= 8; d++ {
		for s := 0; s <= 8; s++ {
			for limit := -1; limit <= 10; limit++ {
				dst, src := d, s
				amount := limit
				if amount < 0 {
					amount = 0
				}
				if amount > s {
					amount = s
				}
				got := subject.Transfer(&dst, &src, limit)
				if got != amount || dst != d+amount || src != s-amount {
					t.Fatalf("state=%d,%d limit=%d: got %d,%d,%d; want %d,%d,%d", d, s, limit, got, dst, src, amount, d+amount, s-amount)
				}
			}
		}
	}
}
