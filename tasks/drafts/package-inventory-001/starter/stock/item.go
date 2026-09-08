package stock

type Item struct {
	Name         string
	Weight       int
	onHand, sold int
}

func New(name string, weight int) Item { return Item{} }
func (i *Item) SetOnHand(n int)        {}
func (i *Item) SetSold(n int)          {}
func (i Item) OnHand() int             { return 0 }
func (i Item) Sold() int               { return 0 }
