package internal

import (
	"log"
	"testing"
)

func TestProvision(t *testing.T) {
	err := Provision()
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
