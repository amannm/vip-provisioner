package main

import "testing"

func TestBasicMain(t *testing.T) {
	exitCode := doMain()
	if exitCode != 0 {
		t.Fail()
	}
}
