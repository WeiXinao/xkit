package list

import (
	"errors"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/errs"
	"math/rand"
)

// 跳表 skip list

const (
	FactorP  = float32(0.25) // level i 上的结点 有 FactorP 的比例出现在 level i + 1上
	MaxLevel = 32
)

type skipListNode[T any] struct {
	Val     T
	Forward []*skipListNode[T] // 跳表结点在每一层的后继结点
}

type SkipList[T any] struct {
	header  *skipListNode[T]
	level   int // SkipList 为空时，level 为 1
	compare xkit.Comparator[T]
	size    int
}

// 实例化一个跳表节点
func newSkipListNode[T any](Val T, level int) *skipListNode[T] {
	return &skipListNode[T]{
		Val, make([]*skipListNode[T], level),
	}
}

func (sl *SkipList[T]) AsSlice() []T {
	curr := sl.header
	slice := make([]T, 0, sl.size)
	for curr.Forward[0] != nil {
		slice = append(slice, curr.Forward[0].Val)
		curr = curr.Forward[0]
	}
	return slice
}

func NewSkipListFromSlice[T any](slice []T, compare xkit.Comparator[T]) *SkipList[T] {
	sl := NewSkipList[T](compare)
	for _, n := range slice {
		sl.Insert(n)
	}
	return sl
}

func NewSkipList[T any](compare xkit.Comparator[T]) *SkipList[T] {
	return &SkipList[T]{
		header: &skipListNode[T]{
			Forward: make([]*skipListNode[T], MaxLevel),
		},
		level:   1,
		compare: compare,
	}
}

// levels 的生成和跳表中元素个数无关
func (sl *SkipList[T]) randomLevel() int {
	level := 1
	p := FactorP
	for (rand.Int31() & 0xFFFF) < int32(p*0xFFFF) {
		level++
	}
	if level < MaxLevel {
		return level
	}
	return MaxLevel
}

func (sl *SkipList[T]) Search(target T) bool {
	curr, _ := sl.traverse(target, sl.level)
	curr = curr.Forward[0] // 第 1 层 包含所有元素
	return curr != nil && sl.compare(curr.Val, target) == 0
}

func (sl *SkipList[T]) traverse(Val T, level int) (*skipListNode[T], []*skipListNode[T]) {
	update := make([]*skipListNode[T], MaxLevel) // update[i] 包含位于 level i 的插入/删除位置左侧的指针
	curr := sl.header
	for i := level - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && sl.compare(curr.Forward[i].Val, Val) < 0 {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}
	return curr, update
}

func (sl *SkipList[T]) Insert(Val T) {
	_, update := sl.traverse(Val, sl.level)
	level := sl.randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}

	//	插入新节点
	newNode := newSkipListNode[T](Val, level)
	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	sl.size += 1
}

func (sl *SkipList[T]) Len() int {
	return sl.size
}

func (sl *SkipList[T]) DeleteElement(target T) bool {
	curr, update := sl.traverse(target, sl.level)
	node := curr.Forward[0]
	if node == nil || sl.compare(node.Val, target) != 0 {
		return true
	}
	for i := 0; i < sl.level && update[i].Forward[i] == node; i++ {
		update[i].Forward[i] = node.Forward[i]
	}

	// 更新层级
	for sl.level > 1 && sl.header.Forward[sl.level-1] == nil {
		sl.level--
	}
	sl.size -= 1
	return true
}

func (sl *SkipList[T]) Peek() (T, error) {
	curr := sl.header
	curr = curr.Forward[0]
	var zero T
	if curr == nil {
		return zero, errors.New("跳表为空")
	}
	return curr.Val, nil
}

func (sl *SkipList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= sl.size {
		return zero, errs.NewErrIndexOutOfRange(sl.size, index)
	}
	curr := sl.header
	for i := 0; i <= index; i++ {
		curr = curr.Forward[0]
	}
	return curr.Val, nil
}
