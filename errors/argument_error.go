package errs

import "fmt"

type ArgumentError struct {
	Arg   interface{}
	Issue string
}

func (e *ArgumentError) Error() string {
	return fmt.Sprintf("%s - %s", e.Arg, e.Issue)
}
