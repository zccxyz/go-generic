package errs

import (
	"errors"
	"fmt"
)

var ErrLength = errors.New("长度为 0")

func NewErrIndexOutOfRange(length, index int) error {
	return fmt.Errorf("下标超出范围，长度：%d，下标：%d", length, index)
}
