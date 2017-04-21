package main

import (
	"fmt"
	"github.com/deneshshan/electronic_life/display"
	"os"
)

func main() {

	handle_error(display.Start())

	os.Exit(0)
}

func handle_error(result display.RenderingResult) {
	if result.Success != true {
		fmt.Fprintf(os.Stderr, "{}: %s\n", result.Reason, result.Err)
		os.Exit(result.ReturnValue)
	}
}
