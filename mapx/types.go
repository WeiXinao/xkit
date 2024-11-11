package mapx

type mapi[K any, V any] interface {
	Put(key K, val V) error

	Get(key K) (V, bool)

	// Delete 删除
	// 第一个返回值是被删除的 key 对应的值
	// 第二个返回值代表是否真的删除了
	Delete(k K) (V, bool)

	// Keys 返回所有的键

}
