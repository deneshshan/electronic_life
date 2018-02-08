package main

import (
	"fmt"
	"github.com/deneshshan/electronic_life/display"
	"os"
)

func main() {

	display.NewDisplay()

	fmt.Println("Bye!!")
	os.Exit(0)
}
