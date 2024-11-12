package slice

// FilterMap 执行过滤并转化
// 如果第二个返回值是 false, 那么我们会忽略第一个返回值
// 即使第二个返回值是 false, 后续的元素依旧会被遍历
func FilterMap[Src any, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := m(i, s)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}

func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}

// ToMap
// 将 []Ele 映射到 map[Key]Ele
// 从 Ele 中提取 Key 的函数 fn 由使用者提供
//
// 注意：
// 如果出现 i < j
// 设：
//
// key_i := fn(element[i])
// key_j := fn(element[j])
//
// 满足 key_i = key_j 的情况，则在返回结果的 resultMap 中
// resultMap[key_i] = val_j
//
// 即使传入的字符串为 nil， 也要保证返回的 map 是一个空 map 而不是 nil
func ToMap[Ele any, Key comparable](
	elements []Ele,
	fn func(element Ele) Key,
) map[Key]Ele {
	return ToMapV(elements, func(element Ele) (Key, Ele) {
		return fn(element), element
	})
}

// ToMapV
// 将 []Ele 映射到 map[Key]Val
// 从 Ele 中提取 Key 和 Val 的函数 fn 由使用者提供
//
// 注意：
// 如果出现 i < j
// 设：
//
// key_i, val_i := fn(element[i])
// key_j, val_j := fn(element[j])
//
// 满足 key_i = key_j 的情况，则在返回结果的 resultMap 中
// resultMap[key_i] = val_j
//
// 即使传入的字符串为 nil， 也要保证返回的 map 是一个空 map 而不是 nil
func ToMapV[Ele any, Key comparable, Val any](
	elements []Ele,
	fn func(element Ele) (Key, Val),
) (resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k, v := fn(element)
		resultMap[k] = v
	}
	return
}

// 构造 map
func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		dataMap[v] = struct{}{}
	}
	return dataMap
}

func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make([]T, 0, len(data))
	for k, v := range data {
		if !ContainsFunc[T](data[k+1:], func(src T) bool {
			return equal(src, v)
		}) {
			newData = append(newData, v)
		}
	}
	return newData
}

func deduplicate[T comparable](data []T) []T {
	dataMap := toMap[T](data)
	var newData = make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
