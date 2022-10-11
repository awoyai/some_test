package cond

import (
	"log"
	"math/rand"
	"sync"
	"time"
)


// 需要在wait之前加锁

func CondTest() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func (i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()

			ready++
			c.L.Unlock()

			log.Println("111")
			c.Broadcast()
		}(i)
	}

	c.L.Lock()

	for ready != 10 {
		c.Wait()
		log.Println("唤醒一次")
	}

	c.L.Unlock()

	log.Println("all read")
}
