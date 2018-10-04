package errors

import (
	"fmt"
)

type StoreError struct {
	Problem string
	Keys  []string
	Err     error
}

func (e *StoreError) Error() string {
	return fmt.Sprintf("Problem: %s, Keys: %v, error: %s", e.Problem, e.Keys, e.Err.Error())
}
