package list

import (
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/list"
)

type SkipList[T any] struct {
	skipList *list.SkipList[T]
}

func NewSkipList[T any](compare xkit.Comparator[T]) *SkipList[T] {
	pq := &SkipList[T]{}
	pq.skipList = list.NewSkipList[T](compare)
	return pq
}

func (sl *SkipList[T]) Search(target T) bool {
	return sl.skipList.Search(target)
}

func (sl *SkipList[T]) AsSlice() []T {
	return sl.skipList.AsSlice()
}

func (sl *SkipList[T]) Len() int {
	return sl.skipList.Len()
}

func (sl *SkipList[T]) Cap() int {
	return sl.Len()
}

func (sl *SkipList[T]) Insert(val T) {
	sl.skipList.Insert(val)
}

func (sl *SkipList[T]) DeleteElement(target T) bool {
	return sl.skipList.DeleteElement(target)
}
