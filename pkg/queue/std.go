package queue

import (
	"container/list"
	"sync"
)

type StdQueueSafe[T any] struct {
	list *list.List
	mu   sync.Mutex
}

func NewStdQueueSafe[T any]() *StdQueueSafe[T] {
	return &StdQueueSafe[T]{
		list: list.New(),
	}
}

func (q *StdQueueSafe[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushBack(item)
}

func (q *StdQueueSafe[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if item := q.list.Front(); item != nil {
		q.list.Remove(item)
		return item.Value.(T), true
	}
	var zero T
	return zero, false
}

func (q *StdQueueSafe[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.list.Len()
}
