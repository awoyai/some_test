package mutex

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int64
}

func (r *RecursiveMutex) Lock() {
	goId := GoId()
	if atomic.LoadInt64(&r.owner) == goId {
		r.recursion++
		return
	}
	r.Mutex.Lock()

	atomic.StoreInt64(&r.owner, goId)
	r.recursion = 1
}

func (r *RecursiveMutex) Unlock() {
	goId := GoId()
	if atomic.LoadInt64(&r.owner) != goId {
		panic(fmt.Sprintf("wrong owner %v", r.owner))
	}
	r.recursion--
	if r.recursion != 0 {
		return
	}
	atomic.StoreInt64(&r.owner, -1)
	r.Mutex.Unlock()
}

func GoId() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idFields := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idFields)
	if err != nil {
		panic(fmt.Sprintf("can't Get GoId, err:%v", err))
	}
	return int64(id)
}
