package slice

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

func DeleteAt[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	var delElem T
	if index < 0 || index >= length {
		return nil, delElem, fmt.Errorf("ekit: %w, 下标超出范围，长度 %d, 下标 %d",
			ErrIndexOutOfRange, length, index)
	}
	delElem = src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	return src[:length-1], delElem, nil
}
