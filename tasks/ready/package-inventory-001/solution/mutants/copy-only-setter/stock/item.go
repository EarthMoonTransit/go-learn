package stock

type Item struct {
	Name         string
	Weight       int
	onHand, sold int
}

func New(name string, weight int) Item { return Item{Name: name, Weight: weight} }
func (i *Item) SetOnHand(n int)        { copyValue := *i; copyValue.onHand = n }
func (i *Item) SetSold(n int)          { i.sold = n }
func (i Item) OnHand() int             { return i.onHand }
func (i Item) Sold() int               { return i.sold }
