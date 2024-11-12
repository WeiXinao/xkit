package stringx_test

import (
	"fmt"
	"github.com/WeiXinao/xkit/stringx"
)

func ExampleUnsafeToBytes() {
	str := "hello"
	val := stringx.UnsafeToBytes(str)
	fmt.Println(len(val))
	// Output:
	// 5
}

func ExampleUnsafeToString() {
	val := stringx.UnsafeToString([]byte("hello"))
	fmt.Println(val)
	// Output:
	// hello
}
