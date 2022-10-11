package mutex

import (
	"sync"
)

type RWMutex struct {
	sync.RWMutex
}

func (r *RWMutex) GetReader() {
}


func (r *RWMutex) GetWriter() {

}