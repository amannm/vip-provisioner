package internal

import (
	"log"
	"testing"
)

func TestRequestGeneration(t *testing.T) {
	requestXml, err := GenerateRequest()
	if err != nil {
		log.Panicf("error: %s", err)
	}
	log.Printf("%s\n", requestXml)
}
