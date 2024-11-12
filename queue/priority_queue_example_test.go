package queue_test

import (
	"fmt"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/queue"
)

func ExampleNewPriorityQueue() {
	// 容量，并且队列里面放的是 int
	pq := queue.NewPriorityQueue(10, xkit.ComparatorRealNumber[int])
	_ = pq.Enqueue(10)
	_ = pq.Enqueue(9)
	val, _ := pq.Dequeue()
	fmt.Println(val)
	// Output:
	// 9
}
