package queue

import "sync/atomic"

const capacity = 1024
const mask = capacity - 1

type RingQueue[T any] struct {
	buffer   [capacity]T
	sequence [capacity]uint64
	head     uint64
	_pad1    [7]uint64
	tail     uint64
	_pad2    [7]uint64
}

func NewRingQueue[T any]() *RingQueue[T] {
	q := &RingQueue[T]{}
	for i := range capacity {
		q.sequence[i] = uint64(i)
	}
	return q
}

func (q *RingQueue[T]) Enqueue(v T) bool {
	for {
		tail := atomic.LoadUint64(&q.tail)
		head := atomic.LoadUint64(&q.head)
		if tail-head >= capacity {
			return false
		}

		idx := tail & mask
		seq := atomic.LoadUint64(&q.sequence[idx])

		if seq == tail {
			if atomic.CompareAndSwapUint64(&q.tail, tail, tail+1) {
				q.buffer[idx] = v
				atomic.StoreUint64(&q.sequence[idx], tail+1)
				return true
			}
		}
	}
}

func (q *RingQueue[T]) Dequeue() (T, bool) {
	var zero T
	for {
		head := atomic.LoadUint64(&q.head)
		tail := atomic.LoadUint64(&q.tail)
		if head >= tail {
			return zero, false // empty
		}

		idx := head & mask
		seq := atomic.LoadUint64(&q.sequence[idx])

		if seq == head+1 {
			if atomic.CompareAndSwapUint64(&q.head, head, head+1) {
				val := q.buffer[idx]
				q.buffer[idx] = zero
				atomic.StoreUint64(&q.sequence[idx], head+capacity)
				return val, true
			}
		}
	}
}
