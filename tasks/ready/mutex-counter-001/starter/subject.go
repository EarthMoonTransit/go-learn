package subject

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter    { return &Counter{} }
func (c *Counter) Inc()       {}
func (c *Counter) Value() int { return c.value }
