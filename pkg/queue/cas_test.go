package queue_test

import (
	"testing"

	"github.com/AugustineAurelius/aether/pkg/queue"
)

func TestStdCasQueue(t *testing.T) {
	q := queue.NewRingQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %d", val)
	}
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("expected 2, got %d", val)
	}
	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
	_, ok = q.Dequeue()
	if ok {
		t.Errorf("expected false, got true")
	}
}

func BenchmarkVyukovQueue(b *testing.B) {
	q := queue.NewRingQueue[int]()
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
