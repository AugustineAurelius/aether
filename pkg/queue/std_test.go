package queue_test

import (
	"testing"

	"github.com/AugustineAurelius/aether/pkg/queue"
)

func TestStdQueueSafe(t *testing.T) {
	q := queue.NewStdQueueSafe[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %d", val)
	}
}

func BenchmarkS(b *testing.B) {
	q := queue.NewStdQueueSafe[int]()
	b.Run("Enqueue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Enqueue(1)

			}
		})

	})
	b.Run("Dequeue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Dequeue()
			}
		})
	})
}
