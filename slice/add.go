package slice

import "github.com/WeiXinao/xkit/internal/slice"

// Add 在 index 处添加元素
// index 范围应为[0, len(src)]
// 如果 index == len(src) 则表示往末尾添加元素
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	return slice.Add[Src](src, element, index)
}
