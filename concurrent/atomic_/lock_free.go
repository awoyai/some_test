package main

import (
	"sync/atomic"
	"unsafe"
)

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type Node struct {
	value any
	next  unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	n := unsafe.Pointer(&Node{})
	return &LockFreeQueue{head: n, tail: n}
}

func (q *LockFreeQueue) EnQueue(val any) {
	n := &Node{value: val}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) { // 尾还是尾
			if next == nil { // 还没有新数据入队
				if cas(&tail.next, next, n) { // 增加到队尾
					cas(&q.tail, tail, n) // 入队成功，移动尾巴指针
					return
				}
			} else { // 已有新数据加到队列后面， 需要移动尾指针
				cas(&q.tail, tail, next)
			}
		}
	}
}

func (q *LockFreeQueue) DeQueue() any {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&tail.next)
		if head == load(&q.head) {
			if head == tail {
				if next == nil {
					return nil
				}
				cas(&q.tail, tail, next)
			} else {
				v := next.value
				if cas(&q.head, head, next) {
					return v
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) *Node {
	return (*Node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *Node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
