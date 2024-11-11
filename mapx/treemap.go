package mapx

import (
	"errors"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/tree"
)

var errTreeMapComparatorIsNull = errors.New("xkit: Comparator不能为nil")

// TreeMap 是基于红黑树实现的 Map
type TreeMap[K any, V any] struct {
	tree *tree.RBTree[K, V]
}

// NewTreeMapWithMap TreeMap 构造方法
// 支持通过传入的 map 构造生成 TreeMap
func NewTreeMapWithMap[K comparable, V any](compare xkit.Comparator[K], m map[K]V) (*TreeMap[K, V], error) {
	treeMap, err := NewTreeMap[K, V](compare)
	if err != nil {
		return treeMap, err
	}
	putAll(treeMap, m)
	return treeMap, nil
}

// NewTreeMap TreeMap 构造方法，创建一个 TreeMap
// 需要注意比较器 compare 不能为 nil
func NewTreeMap[K any, V any](compare xkit.Comparator[K]) (*TreeMap[K, V], error) {
	if compare == nil {
		return nil, errTreeMapComparatorIsNull
	}
	return &TreeMap[K, V]{
		tree: tree.NewRBTree[K, V](compare),
	}, nil
}

// putAll 将 map 传入 TreeMap
// 需注意如果 map 中的 key 已存在，value 将被替换
func putAll[K comparable, V any](treeMap *TreeMap[K, V], m map[K]V) {
	for k, v := range m {
		_ = treeMap.Put(k, v)
	}
}

// Put 在 TreeMap 插入指定值
// 需注意如果 TreeMap 已存在该key， 那么原值会被替换
func (t *TreeMap[K, V]) Put(key K, val V) error {
	err := t.tree.Add(key, val)
	if err == tree.ErrRBTreeSameRBNode {
		return t.tree.Set(key, val)
	}
	return nil
}

// Get 在 TreeMap 找到指定 key 的节点，返回 Val
// TreeMap 未找到指定节点会返回 false
func (t *TreeMap[K, V]) Get(key K) (V, bool) {
	v, err := t.tree.Find(key)
	return v, err == nil
}

// Delete TreeMap 中删除指定 key 的节点
func (t *TreeMap[K, V]) Delete(k K) (V, bool) {
	return t.tree.Delete(k)
}

// Keys 返回了全部的键
// 目前我们是按照中序遍历来返回的数据，但是你不能依赖这个特性
func (t *TreeMap[K, V]) Keys() []K {
	keys, _ := t.tree.KeyValues()
	return keys
}

// Values 返回了全部的值
// 目前我们是按照中序遍历来返回的数据，但是你不能依赖这个特性
func (t *TreeMap[K, V]) Values() []V {
	_, vals := t.tree.KeyValues()
	return vals
}

// Len 返回了键值对的数量
func (t *TreeMap[K, V]) Len() int64 {
	return int64(t.tree.Size())
}
