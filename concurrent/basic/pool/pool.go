package pool

import (
	"bytes"
	"sync"
)

// 这段代码简单阐述了如何使用 Pool，但是不可以用在实际产品中
// 这段代码会有内存泄漏的问题

var buffers = sync.Pool {
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(bf *bytes.Buffer) {
	bf.Reset()
	buffers.Put(bf)
}