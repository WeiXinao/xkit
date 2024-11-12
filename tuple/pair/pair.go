package pair

import (
	"fmt"
)

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func (pair *Pair[K, V]) String() string {
	return fmt.Sprintf("<%#v, %#v>", pair.Key, pair.Value)
}

// Split 方法将 Key, Value 作为返回参数传出。
func (pair *Pair[K, V]) Split() (K, V) {
	return pair.Key, pair.Value
}

func NewPair[K any, V any](
	key K,
	value V,
) Pair[K, V] {
	return Pair[K, V]{
		Key:   key,
		Value: value,
	}
}

// NewPairs 需要传入两个长度相同并且均不为 nil 的数组 keys 和 values，
// 假设 keys 长度为 nil, 返回一个长度为 n 的 pair 数组。
// 保证：
//
// 返回的 pair 数组满足条件（设pair数组为p）：
//
//	对于所有的 0 <= i < n
//	p[i].key == keys[i], 并且 p[i].Value == values[i]
//
// 如果传入的 keys 或者 values 为 nil，会返回 error
//
// 如果传入的 keys 长度与 values 长度不同，会返回 error
func NewPairs[K any, V any](
	keys []K,
	values []V,
) ([]Pair[K, V], error) {
	if keys == nil || values == nil {
		return nil, fmt.Errorf("keys与values均不可为nil")
	}
	n := len(keys)
	if n != len(values) {
		return nil, fmt.Errorf("keys与values的长度不同, len(keys)=%d, len(values)=%d", n, len(values))
	}
	pairs := make([]Pair[K, V], n)
	for i := 0; i < n; i++ {
		pairs[i] = NewPair(keys[i], values[i])
	}
	return pairs, nil
}

// SplitPairs 需要传入一个[]Pair[K, V]，数组可以为 nil。
// 设 pairs 数组的长度为 n， 返回两个长度均为 n 的数组 keys，values。
// 如果 pairs 数组是 nil，那么返回的 keys 与 values 也均为 nil。
func SplitPairs[K any, V any](pairs []Pair[K, V]) (keys []K, values []V) {
	if pairs == nil {
		return nil, nil
	}
	n := len(pairs)
	keys = make([]K, n)
	values = make([]V, n)
	for i, pair := range pairs {
		keys[i], values[i] = pair.Split()
	}
	return
}

//	FlattenPairs 需要传入一个 []Pairs[K, V]，数字可以为 nil
//	如果 pairs 数组为 nil，返回的 flatPairs 数组也为 nil
//
// 设 pairs 数组长度为 n，保证返回的 flatPairs 数组长度为 2 * n 且满足：
//
//	 对于所有的 0 <= i < n
//		flatPairs[i * 2] == pairs[i].Key
//		flatPairs[i * 2 + 1] == pairs[i].Value
func FlattenPairs[K any, V any](pairs []Pair[K, V]) (flatPairs []any) {
	if pairs == nil {
		return nil
	}
	n := len(pairs)
	flatPairs = make([]any, 0, n*2)
	for _, pair := range pairs {
		flatPairs = append(flatPairs, pair.Key, pair.Value)
	}
	return
}

// PackPairs 需要传入一个长度为 2 * n 的数组 flatPairs，数组可以为 nil。
//
// 函数会返回一个长度为 n 的 pairs 数组，pairs 满足
//
//	对于所有的 0 <= i < n
//	pairs[i].Key == flatPairs[i * 2]
//	pairs[i].Value == flatPairs[i * 2 + 1]
//	如果 flatPairs 长度为 nil，则返回 pairs 也为 nil
//
// 入参 flatPairs 需要满足以下条件：
//
//	对于所有 0 <= i < n
//	flatPairs[i * 2] 的类型为 K
//	flatPairs[i * 2 + 1] 的类型为 V
func PackPairs[K any, V any](flatPairs []any) (pairs []Pair[K, V]) {
	if flatPairs == nil {
		return nil
	}
	n := len(flatPairs) / 2
	pairs = make([]Pair[K, V], n)
	for i := 0; i < n; i++ {
		pairs[i] = NewPair(flatPairs[i*2].(K), flatPairs[i*2+1].(V))
	}
	return
}
