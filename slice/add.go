package slice

import "github.com/zccxyz/go-generic/internal/slice"

func Add[T any](src []T, value T, index int) ([]T, error) {
	return slice.Add(src, value, index)
}
