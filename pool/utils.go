package pool

import "fmt"

func NewError(msg string) error {
	return fmt.Errorf("%s", msg)
}
