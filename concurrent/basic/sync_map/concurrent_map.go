package sync_map

import (
	"sync"
)

var SHARED_COUNT = 32

type ConcurrentMap []*ConcurrentMapShared

type ConcurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}



func NewConcurrentMap() ConcurrentMap {
	m := make(ConcurrentMap, SHARED_COUNT)
	for i := 0; i < SHARED_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

func (c ConcurrentMap) GetShared(key string) *ConcurrentMapShared {
	return c[uint(fnv32(key))%uint(SHARED_COUNT)]
}

func (c ConcurrentMap) Set(key string, val interface{}) {
	shared := c.GetShared(key)
	shared.Lock()
	shared.items[key] = val
	shared.Unlock()
}

func (c ConcurrentMap) Get(key string) (interface{}, bool) {
	shared := c.GetShared(key)
	shared.RLock()
	defer shared.RUnlock()
	val, ok := shared.items[key]
	return val, ok
}


// fnv32 hash函数
func fnv32(key string) uint32 {
	// 著名的fnv哈希函数，由 Glenn Fowler、Landon Curt Noll和 Kiem-Phong Vo 创建
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
	   hash *= prime32
	   hash ^= uint32(key[i])
	}
	return hash
 }