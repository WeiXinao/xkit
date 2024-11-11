package list

import (
	"github.com/WeiXinao/xkit/internal/errs"
	"github.com/WeiXinao/xkit/internal/slice"
	"sync"
)

// CopyOnWriteArrayList 基于切片的简单封装 写时加锁，读不加锁，适用于读多写少的场景
type CopyOnWriteArrayList[T any] struct {
	vals  []T
	mutex *sync.Mutex
}

func NewCopyOnWriteArrayList[T any]() *CopyOnWriteArrayList[T] {
	m := &sync.Mutex{}
	return &CopyOnWriteArrayList[T]{
		vals:  make([]T, 0),
		mutex: m,
	}
}

// NewCopyOnWriteArrayListOf 直接使用ts，会执行复制
func NewCopyOnWriteArrayListOf[T any](ts []T) *CopyOnWriteArrayList[T] {
	items := make([]T, len(ts))
	copy(items, ts)
	m := &sync.Mutex{}
	return &CopyOnWriteArrayList[T]{
		vals:  items,
		mutex: m,
	}
}

func (c *CopyOnWriteArrayList[T]) Get(index int) (T, error) {
	var (
		t T
		e error
	)
	l := c.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return c.vals[index], e
}

// Append 往 CopyOnWriteArrayList 里追加数据
func (c *CopyOnWriteArrayList[T]) Append(ts ...T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n := len(c.vals)
	newItems := make([]T, n, n+len(ts))
	copy(newItems, c.vals)
	newItems = append(newItems, ts...)
	c.vals = newItems
	return nil
}

// Add 在 CopyOnWriteArrayList 下标为 index 的位置插入一个元素
// 当 index 等于 CopyOnWriteArrayList长度时，等同于 append
func (c *CopyOnWriteArrayList[T]) Add(index int, t T) error {
	var err error
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n := len(c.vals)
	newItems := make([]T, n, n+1)
	copy(newItems, c.vals)
	newItems, err = slice.Add(newItems, t, index)
	c.vals = newItems
	return err
}

// Set 设置在 CopyOnWriteArrayList 里 index 的位置的值为 t
func (c *CopyOnWriteArrayList[T]) Set(index int, t T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n := len(c.vals)
	if index >= n || index < 0 {
		return errs.NewErrIndexOutOfRange(n, index)
	}
	newItems := make([]T, n)
	copy(newItems, c.vals)
	newItems[index] = t
	c.vals = newItems
	return nil
}

// Delete 这里不涉及缩容，每次都是当前内容长度申请的数组容量
func (c *CopyOnWriteArrayList[T]) Delete(index int) (T, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var res T
	n := len(c.vals)
	if index >= n || index < 0 {
		return res, errs.NewErrIndexOutOfRange(n, index)
	}
	newItems := make([]T, n-1)
	item := 0
	for i, v := range c.vals {
		if i == index {
			res = v
			continue
		}
		newItems[item] = v
		item++
	}
	c.vals = newItems
	return res, nil
}

func (c *CopyOnWriteArrayList[T]) Len() int {
	return len(c.vals)
}

func (c *CopyOnWriteArrayList[T]) Cap() int {
	return cap(c.vals)
}

func (c *CopyOnWriteArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range c.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (c *CopyOnWriteArrayList[T]) AsSlice() []T {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	res := make([]T, len(c.vals))
	copy(res, c.vals)
	return res
}
