package queue

// Queue 普通队列
// 参考 BlockingQueue 阻塞队列
// 一个队列是否遵循 FIFO 取决于具体实现
type Queue[T any] interface {
	// Enqueue 将元素放入队列，如果此队列已经满了，那么返回错误
	Enqueue(t T) error
	// Dequeue 从队首获得一个元素
	// 如果一个队列里面没有元素，那么返回错误
	Dequeue() (T, error)
}
