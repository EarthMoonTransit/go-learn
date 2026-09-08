package subject

import "task.local/package-inventory-001/stock"

import (
	"fmt"
)

func Describe(item stock.Item) string {
	return fmt.Sprintf("%s:%d:%d:%d", item.Name, item.Weight, item.Sold(), item.OnHand())
}
