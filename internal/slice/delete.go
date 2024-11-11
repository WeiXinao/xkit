package slice

import (
	"github.com/WeiXinao/xkit/internal/errs"
)

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	var delElem T
	if index < 0 || index >= length {
		return nil, delElem, errs.NewErrIndexOutOfRange(length, index)
	}
	delElem = src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	return src[:length-1], delElem, nil
}
