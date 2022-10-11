package main

import (
	"context"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

func main() {
	// taskSchedul()
	TestWaterFactory()
}

func taskSchedul() {
	const chanNum = 4
	chs := make([]chan struct{}, chanNum)
	for i := 0; i < chanNum; i++ {
		chs[i] = make(chan struct{}, 1)
	}

	for i := 0; i < chanNum; i++ {
		go task(chs[i], chs[(i+1)%chanNum], i)
	}
	chs[0] <- struct{}{}
	select {}
}

func task(recv, send chan struct{}, msg int) {
	for {
		<-recv
		println(msg+1)
		time.Sleep(time.Second)
		send <- struct{}{}
	}
}

// 水分子制造工厂 利用信号量和循环栅栏 

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b     cyclicbarrier.CyclicBarrier
}

func NewH2O() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)
	releaseHydrogen()
	h2o.b.Await(context.Background())
	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	releaseOxygen()
	h2o.b.Await(context.Background())
	h2o.semaO.Release(1)
}

func TestWaterFactory() {
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}

	releaseOxygen := func() {
		ch <- "O"
	}

	var N = 100
	ch = make(chan string, N*3)
	h2o := NewH2O()

	var wg sync.WaitGroup
	wg.Add(N * 3)

	for i := 0; i < N*2; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}

	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	wg.Wait()
	if len(ch) != N*3 {
		panic("11")
	}
	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <- ch
		s[1] = <- ch
		s[2] = <- ch
		sort.Strings(s)
		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			panic("22222"+water)
		}
	}
}
