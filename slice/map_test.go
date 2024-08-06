package slice

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4}
	res := Map[int, string](a, func(idx int, src int) string {
		return "str:" + strconv.Itoa(src)
	})
	t.Log(res)
}
