package tree

import (
	"errors"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/tree"
)

var (
	errRBTreeComparatorIsNull = errors.New("xkit: RBTree 的 Comparator 不能为 nil")
)

// RBTree 简单的封装一下红黑树
type RBTree[K any, V any] struct {
	rbTree *tree.RBTree[K, V] // 红黑树本体
}

func NewRBTree[K any, V any](compare xkit.Comparator[K]) (*RBTree[K, V], error) {
	if nil == compare {
		return nil, errRBTreeComparatorIsNull
	}

	return &RBTree[K, V]{
		rbTree: tree.NewRBTree[K, V](compare),
	}, nil
}

// Add 增加节点
func (rb *RBTree[K, V]) Add(key K, value V) error {
	return rb.rbTree.Add(key, value)
}

// Delete 删除节点
func (rb *RBTree[K, V]) Delete(key K) (V, bool) {
	return rb.rbTree.Delete(key)
}

// Set 修改节点
func (rb *RBTree[K, V]) Set(key K, value V) error {
	return rb.rbTree.Set(key, value)
}

// Find 查找节点
func (rb *RBTree[K, V]) Find(key K) (V, error) {
	return rb.rbTree.Find(key)
}

// Size 返回红黑树节点的个数
func (rb *RBTree[K, V]) Size() int {
	return rb.rbTree.Size()
}

// KeyValues 获取红黑树所有的节点 K, V
func (rb *RBTree[K, V]) KeyValues() ([]K, []V) {
	return rb.rbTree.KeyValues()
}
