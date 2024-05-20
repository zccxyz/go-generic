package slice

import (
	"github.com/zccxyz/go-generic/internal/errs"
)

func Add[T any](src []T, value T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	var v T
	src = append(src, v)
	for i := length; i > index; i-- {
		src[i] = src[i-1]
	}
	src[index] = value

	return src, nil
}
