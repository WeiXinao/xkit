package mapx

import "sync"

type Hashable interface {
	//	Code 返回该元素的哈希值
	//	注意：哈希值应该尽可能的均匀避免冲突
	Code() uint64
	// Equals 比较两个元素是否相等。如果返回 true, 那么我们会认为两个键是一样的。
	Equals(key any) bool
}

type node[T Hashable, ValType any] struct {
	key   T
	value ValType
	next  *node[T, ValType]
}

type HashMap[T Hashable, ValType any] struct {
	hashmap  map[uint64]*node[T, ValType]
	nodePool *sync.Pool[*node[T, ValType]]
}

func (t *HashMap[T, ValType]) newNode() {

}
