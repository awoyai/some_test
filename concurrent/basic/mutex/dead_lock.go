package mutex

import (
	"fmt"
	"sync"
)

func deadLock() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello")
}