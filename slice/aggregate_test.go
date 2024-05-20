package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/zccxyz/go-generic/internal/errs"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		wantValue int
		wantErr   error
	}{
		{
			name:      "零个数据",
			slice:     []int{},
			wantValue: 0,
			wantErr:   errs.ErrLength,
		},
		{
			name:      "一个数据",
			slice:     []int{1},
			wantValue: 1,
		},
		{
			name:      "多个数据",
			slice:     []int{1, 2, 3, 9, 5, 45, 65, 0},
			wantValue: 65,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rs, err := Max(tc.slice)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantValue, rs)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		wantValue int
		wantErr   error
	}{
		{
			name:      "零个数据",
			slice:     []int{},
			wantValue: 0,
			wantErr:   errs.ErrLength,
		},
		{
			name:      "一个数据",
			slice:     []int{1},
			wantValue: 1,
		},
		{
			name:      "多个数据",
			slice:     []int{1, 2, 3, 9, 5, 45, 65, 0},
			wantValue: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rs, err := Min(tc.slice)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantValue, rs)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		wantValue int
	}{
		{
			name:      "零个数据",
			slice:     []int{},
			wantValue: 0,
		},
		{
			name:      "一个数据",
			slice:     []int{1},
			wantValue: 1,
		},
		{
			name:      "多个数据",
			slice:     []int{1, 2, 3, 9, 0},
			wantValue: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rs := Sum(tc.slice)
			assert.Equal(t, tc.wantValue, rs)
		})
	}
}
