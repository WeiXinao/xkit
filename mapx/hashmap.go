package mapx

import (
	"github.com/WeiXinao/xkit/syncx"
)

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
	nodePool *syncx.Pool[*node[T, ValType]]
}

func (m *HashMap[T, ValType]) newNode(key T, val ValType) *node[T, ValType] {
	newNode := m.nodePool.Get()
	newNode.value = val
	newNode.key = key
	return newNode
}

func NewHashMap[T Hashable, ValType any](size int) *HashMap[T, ValType] {
	return &HashMap[T, ValType]{
		nodePool: syncx.NewPool[*node[T, ValType]](func() *node[T, ValType] {
			return &node[T, ValType]{}
		}),
		hashmap: make(map[uint64]*node[T, ValType], size),
	}
}

func (m *HashMap[T, ValType]) Put(key T, val ValType) error {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	if !ok {
		hash = key.Code()
		newNode := m.newNode(key, val)
		m.hashmap[hash] = newNode
		return nil
	}
	pre := root
	for root != nil {
		if root.key.Equals(key) {
			root.value = val
			return nil
		}
		pre = root
		root = root.next
	}
	newNode := m.newNode(key, val)
	pre.next = newNode
	return nil
}

func (m *HashMap[T, ValType]) Get(key T) (ValType, bool) {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	var val ValType
	if !ok {
		return val, false
	}
	for root != nil {
		if root.key.Equals(key) {
			return root.value, true
		}
		root = root.next
	}
	return val, false
}

// Delete 第一个返回值为删除 key 的值，第二个是 hashmap 是否真的有这个 key
func (m *HashMap[T, ValType]) Delete(key T) (ValType, bool) {
	root, ok := m.hashmap[key.Code()]
	if !ok {
		var t ValType
		return t, false
	}
	pre := root
	num := 0
	for root != nil {
		if root.key.Equals(key) {
			if num == 0 && root.next == nil {
				delete(m.hashmap, key.Code())
			} else if num == 0 && root.next != nil {
				m.hashmap[key.Code()] = root.next
			} else {
				pre.next = root.next
			}
			val := root.value
			root.formatting()
			m.nodePool.Put(root)
			return val, true
		}
		num++
		pre = root
		root = root.next
	}
	var t ValType
	return t, false
}

func (n *node[T, ValType]) formatting() {
	var val ValType
	var t T
	n.key = t
	n.value = val
	n.next = nil
}

// Keys 返回 Hashmap 里面所有的 key。
// 注意：key 的顺序是随机的。
func (m *HashMap[T, ValType]) Keys() []T {
	res := make([]T, 0)
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.key)
			curNode = curNode.next
		}
	}
	return res
}

// Values 返回 HashMap 里面所有的 value。
// 注意：value 的顺序是随机的。
func (m *HashMap[T, ValType]) Values() []ValType {
	res := make([]ValType, 0)
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.value)
			curNode = curNode.next
		}
	}
	return res
}

// TODO: 有问题
func (m *HashMap[T, ValType]) Len() int64 {
	return int64(len(m.hashmap))
}
