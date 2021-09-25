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
	err := printRequest()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
		return 1
	}
	return 0
}
