package errors

import "fmt"

type ArgumentError struct {
	Arg   interface{}
	Issue string
}

func (e *ArgumentError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}
