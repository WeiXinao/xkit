package list

import (
	"github.com/WeiXinao/xkit/internal/errs"
	"github.com/WeiXinao/xkit/internal/slice"
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	vals []T
}

// NewArrayList 初始化一个 len 为 0，cap 为 cap 的 ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

// NewArrayListOf 直接使用 ts，而不会执行复制
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	var t T
	l := a.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.vals[index], nil
}

// Append 往ArrayList 里追加数据
func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

// Add 在 ArrayList 下标为 index 的位置插入一个元素
// 当 index 等于 ArrayList 长度等同于 append
func (a *ArrayList[T]) Add(index int, t T) error {
	var err error
	a.vals, err = slice.Add(a.vals, t, index)
	return err
}

// Set 设置 ArrayList 里 index 位置的值为 t
func (a *ArrayList[T]) Set(index int, t T) error {
	length := len(a.vals)
	if index >= length || index < 0 {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = t
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量的一半，那么就会缩容为原来的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原来的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很小，所以没必要消耗 CPU 去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(a.vals, index)
	if err != nil {
		return t, err
	}
	a.vals = res
	a.shrink()
	return t, nil
}

// shrink 数组缩容
func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink(a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.vals))
	copy(res, a.vals)
	return res
}
