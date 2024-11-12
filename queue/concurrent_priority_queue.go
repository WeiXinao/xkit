package queue

import (
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/queue"
	"sync"
)

type ConcurrentPriorityQueue[T any] struct {
	pq queue.PriorityQueue[T]
	m  sync.RWMutex
}

// NewConcurrentPriorityQueue 创建优先队列 capacity <= 0 时，为无界队列
func NewConcurrentPriorityQueue[T any](capacity int, compare xkit.Comparator[T]) *ConcurrentPriorityQueue[T] {
	return &ConcurrentPriorityQueue[T]{
		pq: *queue.NewPriorityQueue[T](capacity, compare),
	}
}

// Cap 无界队列返回 0，有界队列返回创建队列时设置的值
func (c *ConcurrentPriorityQueue[T]) Cap() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.pq.Cap()
}

func (c *ConcurrentPriorityQueue[T]) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.pq.Len()
}

func (c *ConcurrentPriorityQueue[T]) Peek() (T, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.pq.Peek()
}

func (c *ConcurrentPriorityQueue[T]) Enqueue(t T) error {
	c.m.Lock()
	defer c.m.Unlock()
	return c.pq.Enqueue(t)
}

func (c *ConcurrentPriorityQueue[T]) Dequeue() (T, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.pq.Dequeue()
}
