package subject

var shared int

func NewCounter(initial int) *int { shared = initial; return &shared }
