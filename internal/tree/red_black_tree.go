package tree

import (
	"errors"
	"github.com/WeiXinao/xkit"
)

type color bool

const (
	Red   color = false
	Black color = true
)

var (
	ErrRBTreeSameRBNode = errors.New("xkit: RBTree 不能添加重复节点Key")
	ErrRBTreeNotRBNode  = errors.New("xkit: RBTree 不存在节点key")
)

type rbNode[K any, V any] struct {
	color               color
	key                 K
	value               V
	left, right, parent *rbNode[K, V]
}

func (node *rbNode[K, V]) setNode(v V) {
	if node == nil {
		return
	}
	node.value = v
}

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare xkit.Comparator[K]
	size    int
}

func newRBNode[K any, V any](key K, value V) *rbNode[K, V] {
	return &rbNode[K, V]{
		color:  Red,
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func NewRBTree[K any, V any](compare xkit.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		compare: compare,
		root:    nil,
	}
}

func (rb *RBTree[K, V]) Size() int {
	if rb == nil {
		return 0
	}
	return rb.size
}

// Add 添加节点
func (rb *RBTree[K, V]) Add(key K, value V) error {
	return rb.addNode(newRBNode(key, value))
}

// Delete 删除节点
func (rb *RBTree[K, V]) Delete(key K) (V, bool) {
	if node := rb.findNode(key); node != nil {
		value := node.value
		rb.deleteNode(node)
		return value, true
	}
	var v V
	return v, false
}

// Find 查找节点
func (rb *RBTree[K, V]) Find(key K) (V, error) {
	var v V
	if node := rb.findNode(key); node != nil {
		return node.value, nil
	}
	return v, ErrRBTreeNotRBNode
}

func (rb *RBTree[K, V]) Set(key K, value V) error {
	if node := rb.findNode(key); node != nil {
		node.setNode(value)
		return nil
	}
	return ErrRBTreeNotRBNode
}

// KeyValues 获取红黑树所有节点 K, V
func (rb *RBTree[K, V]) KeyValues() ([]K, []V) {
	keys := make([]K, 0, rb.size)
	values := make([]V, 0, rb.size)
	if rb.root == nil {
		return keys, values
	}
	rb.inOrderTraversal(func(node *rbNode[K, V]) {
		keys = append(keys, node.key)
		values = append(values, node.value)
	})
	return keys, values
}

// inOrderTraversal 中序遍历
func (rb *RBTree[K, V]) inOrderTraversal(visit func(node *rbNode[K, V])) {
	stack := make([]*rbNode[K, V], 0, rb.size)
	curr := rb.root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visit(curr)
		curr = curr.right
	}
}

func (rb *RBTree[K, V]) findNode(key K) *rbNode[K, V] {
	node := rb.root
	for node != nil {
		cmp := rb.compare(key, node.key)
		if cmp < 0 {
			node = node.left
		} else if cmp > 0 {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

// addNode 插入新节点
func (rb *RBTree[K, V]) addNode(node *rbNode[K, V]) error {
	var fixNode *rbNode[K, V]
	if rb.root == nil {
		rb.root = newRBNode[K, V](node.key, node.value)
		fixNode = rb.root
	} else {
		t := rb.root
		cmp := 0
		parent := &rbNode[K, V]{}
		for t != nil {
			parent = t
			cmp = rb.compare(node.key, t.key)
			if cmp < 0 {
				t = t.left
			} else if cmp > 0 {
				t = t.right
			} else if cmp == 0 {
				return ErrRBTreeSameRBNode
			}
		}
		fixNode = &rbNode[K, V]{
			color:  Red,
			key:    node.key,
			value:  node.value,
			parent: parent,
		}
		if cmp < 0 {
			parent.left = fixNode
		} else {
			parent.right = fixNode
		}
	}
	rb.size++
	rb.fixAfterAdd(fixNode)
	return nil
}

// deleteNode 红黑树删除方法
// 删除分两步，第一步取出后继节点，第二步着色旋转
// 取后继节点
//
//	case1: node 左右非空节点，通过 getSuccessor 获取后继节点
//	case2: node 左右只有一个非空子节点
//	case3: node 左右均为空节点
//
// 着色旋转
//
//	case1: 当删除结点非空且为黑色时，会违反红黑树任何路径黑节点个数相同的约束，所以需要重新平衡
//	case2：当删除红色节点时，不会破坏任何约束，所以不需要平衡
func (rb *RBTree[K, V]) deleteNode(tgt *rbNode[K, V]) {
	node := tgt
	if node.left != nil && node.right != nil {
		s := rb.findSuccessor(node)
		node.key = s.key
		node.value = s.value
		node = s
	}
	var replacement *rbNode[K, V]
	//	node 节点只有一个非空子节点
	if node.left != nil {
		replacement = node.left
	} else {
		replacement = node.right
	}
	if replacement != nil {
		replacement.parent = node.parent
		if node.parent == nil {
			rb.root = replacement
		} else if node == node.parent.left {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}
		node.left = nil
		node.right = nil
		node.parent = nil
		if node.getColor() {
			rb.fixAfterDelete(replacement)
		}
	} else if node.parent == nil {
		//	如果 node 节点无父节点，说明 node 节点为 root 节点
		rb.root = nil
	} else {
		//	node 的子节点均为空
		if node.getColor() {
			rb.fixAfterDelete(node)
		}
		if node.parent != nil {
			if node == node.parent.left {
				node.parent.left = nil
			} else if node == node.parent.right {
				node.parent.right = nil
			}
			node.parent = nil
		}
	}
	rb.size--
}

// fixAfterDelete 删除时着色旋转
// 根据 x 节点位置分为 fixAfterDeleteLeft, fixAfterDeleteRight 两种情况
func (rb *RBTree[K, V]) fixAfterDelete(x *rbNode[K, V]) {
	for x != rb.root && x.getColor() == Black {
		if x == x.parent.getLeft() {
			x = rb.fixAfterDeleteLeft(x)
		} else {
			x = rb.fixAfterDeleteRight(x)
		}
	}
	x.setColor(Black)
}

// fixAfterDeleteLeft 处理 x 为左子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteLeft(x *rbNode[K, V]) *rbNode[K, V] {
	sib := x.getParent().getRight()
	// 如果 x 的兄弟节点是红色（此时，x 的父母，sib 的孩子，均为黑色），则转化为 x 的兄弟节点是黑色的情况
	if sib.getColor() == Red {
		// 交换 parent 和 sib 的颜色
		// sib 置 黑
		sib.setColor(Black)
		// parent 置 红
		sib.getParent().setColor(Red)

		// 左旋 parent
		rb.rotateLeft(x.getParent())
		// 获取 x 的新 sib
		sib = x.getParent().getRight()
	}

	// x 的兄弟为黑的情况
	// 1. sib 的左右孩子均为黑的情况
	if sib.getLeft().getColor() == Black && sib.getRight().getColor() == Black {
		// 将 sib 节点置 红
		sib.setColor(Red)
		// TODO: 将 parent 置 黑
		x = x.getParent()
	} else {
		//	sib 的左孩子是红色的 或 sib 的右孩子是红色的
		//	2. sib 的左孩子是红色的情况, 转为 sib 的右孩子是红色的情况
		if sib.getRight().getColor() == Black {
			sib.getLeft().setColor(Black)
			sib.setColor(Red)
			rb.rotateRight(sib)
			sib = x.getParent().getRight()
		}
		// 3. sib 的右孩子是红色的情况
		// 交换 sib 和 parent 的颜色
		sib.setColor(x.getParent().getColor())
		x.getParent().setColor(Black)
		// sib 的 右黑子 置黑
		sib.getRight().setColor(Black)
		// 左旋 parent
		rb.rotateLeft(x.getParent())
		// 这一行是为了让修复过程结束
		x = rb.root
	}
	return x
}

// fixAfterDeleteRight 处理 x 为右子节点时的平衡处理
func (rb *RBTree[K, V]) fixAfterDeleteRight(x *rbNode[K, V]) *rbNode[K, V] {
	sib := x.getParent().getLeft()
	if sib.getColor() == Red {
		sib.setColor(Black)
		x.getParent().setColor(Red)
		rb.rotateRight(x.getParent())
		sib = x.getBrother()
	}
	if sib.getRight().getColor() == Black && sib.getLeft().getColor() == Black {
		sib.setColor(Red)
		x = x.getParent()
	} else {
		if sib.getLeft().getColor() == Black {
			sib.getRight().setColor(Black)
			sib.setColor(Red)
			rb.rotateLeft(sib)
			sib = x.getParent().getLeft()
		}
		sib.setColor(x.getParent().getColor())
		x.getParent().setColor(Black)
		sib.getLeft().setColor(Black)
		rb.rotateRight(x.getParent())
		x = rb.root
	}
	return x
}

// findSuccessor 寻找后继节点
// case1：node 节点存在右子节点，则右子树的最小节点是 node 的后继节点
// case2：node 节点不存在右子节点时，则其第一个为左节点祖先的父节点为 node 的后继节点 // TODO 存疑
func (rb *RBTree[K, V]) findSuccessor(node *rbNode[K, V]) *rbNode[K, V] {
	if node == nil {
		return nil
	} else if node.right != nil {
		p := node.right
		for p.left != nil {
			p = p.left
		}
		return p
	} else {
		p := node.parent
		ch := node
		for p != nil && ch == p.right {
			ch = p
			p = p.parent
		}
		return p
	}
}

func (rb *RBTree[K, V]) fixAfterAdd(x *rbNode[K, V]) {
	x.color = Red
	for x != nil && x != rb.root && x.getParent().getColor() == Red {
		uncle := x.getUncle()
		if uncle.getColor() == Red {
			x = rb.fixUncleRed(x, uncle)
			continue
		}
		if x.getParent() == x.getGrandParent().getLeft() {
			x = rb.fixAddLeftBlack(x)
			continue
		}
		x = rb.fixAddRightBlack(x)
	}
	rb.root.setColor(Black)
}

func (rb *RBTree[K, V]) fixAddRightBlack(x *rbNode[K, V]) *rbNode[K, V] {
	if x == x.getParent().getLeft() {
		x = x.getParent()
		rb.rotateRight(x)
	}
	x.getParent().setColor(Black)
	x.getGrandParent().setColor(Red)
	rb.rotateLeft(x.getGrandParent())
	return x
}

// fixAddLeftBlack 叔叔节点上黑色右节点，父节点是祖父节点是左节点
// 如果 x 节点是父节点的右节点，执行左旋，如果 x 为左节点则跳过左旋操作
// 此时 x 节点变为原 x 节点的父节点 a, 也就是左子节点。
// 对爷爷接点进行右旋，接着将 x 节点和爷节点染色（红变黑，黑变红），此时红黑树完成
func (rb *RBTree[K, V]) fixAddLeftBlack(x *rbNode[K, V]) *rbNode[K, V] {
	if x == x.getParent().getRight() {
		x = x.getParent()
		rb.rotateLeft(x)
	}
	x.getParent().setColor(Black)
	x.getGrandParent().setColor(Red)
	rb.rotateRight(x.getGrandParent())
	return x
}

func (rb *RBTree[K, V]) fixUncleRed(x *rbNode[K, V], y *rbNode[K, V]) *rbNode[K, V] {
	x.getParent().setColor(Black)
	y.setColor(Black)
	x.getGrandParent().setColor(Red)
	x = x.getGrandParent()
	return x
}

// rotateLeft 左旋
func (rb *RBTree[K, V]) rotateLeft(node *rbNode[K, V]) {
	if node == nil || node.getRight() == nil {
		return
	}
	r := node.right
	node.right = r.left
	if r.left != nil {
		r.left.parent = node
	}
	r.parent = node.parent
	if node.parent == nil {
		rb.root = r
	} else if node.parent.left == node {
		node.parent.left = r
	} else {
		node.parent.right = r
	}
	r.left = node
	node.parent = r
}

func (rb *RBTree[K, V]) rotateRight(node *rbNode[K, V]) {
	if node == nil || node.getLeft() == nil {
		return
	}
	l := node.left
	node.left = l.right
	if l.right != nil {
		l.right.parent = node
	}
	l.parent = node.parent
	if node.parent == nil {
		rb.root = l
	} else if node.parent.right == node {
		node.parent.right = l
	} else {
		node.parent.left = l
	}
	l.right = node
	node.parent = l
}

func (node *rbNode[K, V]) getColor() color {
	if node == nil {
		return Black
	}
	return node.color
}

func (node *rbNode[K, V]) setColor(color color) {
	if node == nil {
		return
	}
	node.color = color
}

func (node *rbNode[K, V]) getParent() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.parent
}

func (node *rbNode[K, V]) getLeft() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.left
}

func (node *rbNode[K, V]) getRight() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.right
}

func (node *rbNode[K, V]) getGrandParent() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.getParent().getParent()
}

func (node *rbNode[K, V]) getBrother() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	if node == node.parent.getLeft() {
		return node.getParent().getRight()
	}
	return node.getParent().getLeft()
}

func (node *rbNode[K, V]) getUncle() *rbNode[K, V] {
	if node == nil {
		return nil
	}
	return node.getParent().getBrother()
}
