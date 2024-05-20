package slice

import (
	"github.com/zccxyz/go-generic"
	"github.com/zccxyz/go-generic/internal/errs"
)

func Max[T go_generic.RealNumber](slice []T) (T, error) {
	if len(slice) == 0 {
		return 0, errs.ErrLength
	}

	maxVal := slice[0]
	for i := 1; i < len(slice); i++ {
		if maxVal < slice[i] {
			maxVal = slice[i]
		}
	}

	return maxVal, nil
}

func Min[T go_generic.RealNumber](slice []T) (T, error) {
	if len(slice) == 0 {
		return 0, errs.ErrLength
	}

	minVal := slice[0]
	for i := 1; i < len(slice); i++ {
		if minVal > slice[i] {
			minVal = slice[i]
		}
	}

	return minVal, nil
}

func Sum[T go_generic.Number](slice []T) T {
	var sumVal T
	for _, v := range slice {
		sumVal += v
	}

	return sumVal
}
