package slice

import (
	"testing"
)

func TestShrink(t *testing.T) {
	s := make([]int, 0, 2049)

	for i := 0; i < 2049; i++ {
		s = append(s, i)
	}
	d := Shrink(s)
	t.Log("before delete cap:", cap(d))

	var err error
	for i := 0; i < 1025; i++ {
		s, _, err = DeleteAt[int](s, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Log("after delete cap:", cap(s))
}
