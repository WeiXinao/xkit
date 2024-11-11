package mapx_test

import (
	"fmt"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/mapx"
)

func ExampleNewTreeMap() {
	m, _ := mapx.NewTreeMap[int, int](xkit.ComparatorRealNumber[int])
	_ = m.Put(1, 11)
	val, _ := m.Get(1)
	fmt.Println(val)
	// Output:
	// 11
}
