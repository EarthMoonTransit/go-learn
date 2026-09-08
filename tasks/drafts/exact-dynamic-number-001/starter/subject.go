package subject

type FancyNumber struct{ Text string }

func (n FancyNumber) Value() string { return n.Text }

type FancyNumberBox interface{ Value() string }

func Extract(box FancyNumberBox) int { return 0 }
