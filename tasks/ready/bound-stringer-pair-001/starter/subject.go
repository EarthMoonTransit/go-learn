package subject

import "fmt"

type Label struct{ Text string }
type Counter struct{ Value int }

func (l Label) String() string                                        { return "" }
func (c *Counter) String() string                                     { return "" }
func Bind(label Label, counter *Counter) (fmt.Stringer, fmt.Stringer) { return nil, nil }
