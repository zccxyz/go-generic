package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/zccxyz/go-generic/internal/errs"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addValue  int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 2, 3},
			addValue:  10,
			index:     0,
			wantSlice: []int{10, 1, 2, 3},
		},
		{
			name:      "index 中间",
			slice:     []int{1, 2, 3, 4, 5},
			addValue:  10,
			index:     2,
			wantSlice: []int{1, 2, 10, 3, 4, 5},
		},
		{
			name:      "index 等于长度",
			slice:     []int{1, 2, 3, 4, 5},
			addValue:  10,
			index:     5,
			wantSlice: []int{1, 2, 3, 4, 5, 10},
		},
		{
			name:     "index 超出长度",
			slice:    []int{1, 2, 3, 4, 5},
			addValue: 10,
			index:    10,
			wantErr:  errs.NewErrIndexOutOfRange(5, 10),
		},
		{
			name:     "index 小于 0",
			slice:    []int{1, 2, 3, 4, 5},
			addValue: 10,
			index:    -1,
			wantErr:  errs.NewErrIndexOutOfRange(5, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice, err := Add(tc.slice, tc.addValue, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, slice)
		})
	}
}
