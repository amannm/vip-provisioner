package main

import (
	"fmt"
	"os"
)

func main() {
	exitCode := doMain()
	os.Exit(exitCode)
}

func doMain() int {
	fmt.Printf("Hello World\n")
	return 0
}
