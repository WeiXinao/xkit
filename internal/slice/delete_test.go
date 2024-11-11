package slice

import "testing"

func TestDeleteAt(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s2, e, err := Delete[int](s, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("new slice: %v, deleted element: %v", s2, e)
}
