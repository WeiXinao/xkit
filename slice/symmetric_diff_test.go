package slice

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymmetricDiffSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "no inter",
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "part inter",
			src:  []int{1, 2, 3},
			dst:  []int{3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "src contain dst",
			src:  []int{1, 2, 3},
			dst:  []int{2, 3},
			want: []int{1},
		},
		{
			name: "dst contain src",
			src:  []int{4},
			dst:  []int{4, 5, 6},
			want: []int{5, 6},
		},
		{
			name: "equal",
			src:  []int{1, 2, 3},
			dst:  []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "dst empty",
			src:  []int{1, 2, 3},
			dst:  []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "src empty",
			src:  []int{},
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "all empty",
			src:  []int{},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "dst nil",
			src:  []int{4, 5, 6},
			dst:  nil,
			want: []int{4, 5, 6},
		},
		{
			name: "both nil",
			src:  nil,
			dst:  nil,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SymmetricDiffSet[int](tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "no inter",
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "part inter",
			src:  []int{1, 2, 3},
			dst:  []int{3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "src contain dst",
			src:  []int{1, 2, 3},
			dst:  []int{2, 3},
			want: []int{1},
		},
		{
			name: "dst contain src",
			src:  []int{4},
			dst:  []int{4, 5, 6},
			want: []int{5, 6},
		},
		{
			name: "equal",
			src:  []int{1, 2, 3},
			dst:  []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "dst empty",
			src:  []int{1, 2, 3},
			dst:  []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "src empty",
			src:  []int{},
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "all empty",
			src:  []int{},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "dst nil",
			src:  []int{4, 5, 6},
			dst:  nil,
			want: []int{4, 5, 6},
		},
		{
			name: "both nil",
			src:  nil,
			dst:  nil,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SymmetricDiffSetFunc[int](tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func ExampleSymmetricDiffSet() {
	res := SymmetricDiffSet[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}

func ExampleSymmetricDiffSetFunc() {
	res := SymmetricDiffSetFunc[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}
