package mutex

import (
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}