package queue

import (
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/queue"
)

type PriorityQueue[T any] struct {
	priorityQueue *queue.PriorityQueue[T]
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	return pq.priorityQueue.Dequeue()
}

func (pq *PriorityQueue[T]) Enqueue(t T) error {
	return pq.priorityQueue.Enqueue(t)
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	return pq.priorityQueue.Peek()
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.priorityQueue.Len()
}

func NewPriorityQueue[T any](capacity int, compare xkit.Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	pq.priorityQueue = queue.NewPriorityQueue[T](capacity, compare)
	return pq
}
