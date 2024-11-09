package list

import (
	"errors"
	"github.com/WeiXinao/xkit"
	"github.com/WeiXinao/xkit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	testCases := []struct {
		name       string
		compare    xkit.Comparator[int]
		level      int
		wantHeader *skipListNode[int]
		wantLevel  int
		wantSlice  []int
		wantErr    error
		wantSize   int
	}{
		{
			name:       "new skip list",
			compare:    xkit.ComparatorRealNumber[int],
			wantLevel:  1,
			wantHeader: newSkipListNode[int](0, MaxLevel),
			wantSlice:  []int{},
			wantSize:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList(tc.compare)
			assert.Equal(t, tc.wantLevel, sl.level)
			assert.Equal(t, tc.wantHeader, sl.header)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
			assert.Equal(t, tc.wantSize, sl.size)
		})
	}
}

func TestNewSkipListFromSlice(t *testing.T) {
	testCases := []struct {
		name    string
		compare xkit.Comparator[int]
		slice   []int

		wantSlice []int
		wantErr   error
		wantSize  int
	}{
		{
			name:    "new skip list",
			compare: xkit.ComparatorRealNumber[int],
			slice:   []int{1, 2, 3},

			wantSlice: []int{1, 2, 3},
			wantSize:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipListFromSlice[int](tc.slice, tc.compare)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
			assert.Equal(t, tc.wantSize, sl.size)
		})
	}
}

func TestSkipList_DeleteElement(t *testing.T) {
	testCases := []struct {
		name      string
		skipList  *SkipList[int]
		value     int
		wantSlice []int
		wantSize  int
		wantRes   bool
	}{
		{
			name:      "delete 2 from [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{3},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [1]",
			skipList:  NewSkipListFromSlice[int]([]int{1}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{},
			wantSize:  0,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [2]",
			skipList:  NewSkipListFromSlice[int]([]int{2}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{2},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "delete 3 from [1, 2, 3, 4, 5, 6, 7]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7}, xkit.ComparatorRealNumber[int]),
			value:     3,
			wantSlice: []int{1, 2, 4, 5, 6, 7},
			wantSize:  6,
			wantRes:   true,
		},
		{
			name:      "delete 8 from [1, 2, 3, 4, 5, 6, 7]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7}, xkit.ComparatorRealNumber[int]),
			value:     8,
			wantSlice: []int{1, 2, 3, 4, 5, 6, 7},
			wantSize:  7,
			wantRes:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.skipList.DeleteElement(tc.value)
			assert.Equal(t, tc.wantSize, tc.skipList.size)
			assert.Equal(t, tc.wantSlice, tc.skipList.AsSlice())
			assert.Equal(t, tc.wantRes, ok)
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name      string
		skipList  *SkipList[int]
		value     int
		wantSlice []int
		wantSize  int
	}{
		{
			name:      "insert 2 into [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 3},
			wantSize:  3,
		},
		{
			name:      "insert 1 into []",
			skipList:  NewSkipListFromSlice[int]([]int{}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1},
			wantSize:  1,
		},
		{
			name:      "insert 2 into [1, 2, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 2, 3},
			wantSize:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.skipList.Insert(tc.value)
			assert.Equal(t, tc.wantSize, tc.skipList.size)
			assert.Equal(t, tc.wantSlice, tc.skipList.AsSlice())
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	testCases := []struct {
		name      string
		skipList  *SkipList[int]
		value     int
		wantSlice []int
		wantSize  int
		wantRes   bool
	}{
		{
			name:      "search 2 from [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   false,
		},
		{
			name:      "search 1 from [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   true,
		},
		{
			name:      "search 1 from [1]",
			skipList:  NewSkipListFromSlice[int]([]int{1}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "search 1 from [2]",
			skipList:  NewSkipListFromSlice[int]([]int{2}, xkit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{2},
			wantSize:  1,
			wantRes:   false,
		},
		{
			name:      "search 3 from [1, 2, 3, 4, 5, 6]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6}, xkit.ComparatorRealNumber[int]),
			value:     3,
			wantSlice: []int{1, 2, 3, 4, 5, 6},
			wantSize:  6,
			wantRes:   true,
		},
		{
			name:      "search 8 from [1, 2, 3, 4, 5, 6]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6}, xkit.ComparatorRealNumber[int]),
			value:     8,
			wantSlice: []int{1, 2, 3, 4, 5, 6},
			wantSize:  6,
			wantRes:   false,
		},
		{
			name:      "search 2 from [1, 2, 2, 3, 3, 4, 5, 6]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 2, 3, 3, 4, 5, 6}, xkit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 2, 3, 3, 4, 5, 6},
			wantSize:  8,
			wantRes:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.skipList.Search(tc.value)
			assert.Equal(t, tc.wantSize, tc.skipList.size)
			assert.Equal(t, tc.wantSlice, tc.skipList.AsSlice())
			assert.Equal(t, tc.wantRes, ok)
		})
	}
}

func TestSkipList_randomLevel(t *testing.T) {
	sl := NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int])
	t.Log(sl.randomLevel())
}

func TestSkipList_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		skipList  *SkipList[int]
		wantVal   int
		wantError error
	}{
		{
			name:      "peek [1, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 3}, xkit.ComparatorRealNumber[int]),
			wantVal:   1,
			wantError: nil,
		},
		{
			name:      "peek []",
			skipList:  NewSkipListFromSlice[int]([]int{}, xkit.ComparatorRealNumber[int]),
			wantVal:   0,
			wantError: errors.New("跳表为空"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.skipList.Peek()
			assert.Equal(t, tc.wantError, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestSkipList_Get(t *testing.T) {
	testCases := []struct {
		name     string
		skipList *SkipList[int]
		index    int
		wantVal  int
		wantErr  error
	}{
		{
			name:     "get index -1 [1, 2, 3]",
			skipList: NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			index:    -1,
			wantVal:  0,
			wantErr:  errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:     "get index 3 [1, 2, 3]",
			skipList: NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			index:    3,
			wantVal:  0,
			wantErr:  errs.NewErrIndexOutOfRange(3, 3),
		},
		{
			name:     "get index 0 [1, 2, 3]",
			skipList: NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			index:    0,
			wantVal:  1,
			wantErr:  nil,
		},
		{
			name:     "get index 1 [1, 2, 3]",
			skipList: NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			index:    1,
			wantVal:  2,
			wantErr:  nil,
		},
		{
			name:     "get index 2 [1, 2, 3]",
			skipList: NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			index:    2,
			wantVal:  3,
			wantErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.skipList.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestSkipList_AsSlice(t *testing.T) {
	testCases := []struct {
		name      string
		skipList  *SkipList[int]
		wantSlice []int
	}{
		{
			name:      "[1, 2, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "[3, 2, 1]",
			skipList:  NewSkipListFromSlice[int]([]int{3, 2, 1}, xkit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "[1, 2, 3]",
			skipList:  NewSkipListFromSlice[int]([]int{1, 2, 3}, xkit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantSlice, tc.skipList.AsSlice())
		})
	}
}
