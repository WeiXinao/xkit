package list_test

import (
	"fmt"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/list"
)

func ExampleNewSkipList() {
	l := list.NewSkipList[int](xkit.ComparatorRealNumber[int])
	l.Insert(123)
	val, _ := l.Get(0)
	fmt.Println(val)
	// Output:
	// 123
}
