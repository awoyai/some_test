package mutex

import (
	"sync"
	"testing"
)

func TestDeadLock(t *testing.T) {
	deadLock()
}

func TestWaitGroup(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(&c, &wg)
	}
	wg.Wait()
	println(c.Count())
}
