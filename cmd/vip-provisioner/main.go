package main

import (
	"fmt"
	"os"
	"vip-provisioner/internal"
)

func main() {
	exitCode := doMain()
	os.Exit(exitCode)
}

func doMain() int {
	err := internal.Provision()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
		return 1
	}
	return 0
}
