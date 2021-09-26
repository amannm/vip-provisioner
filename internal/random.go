package internal

import (
	"crypto/rand"
	"math/big"
)

func randomString(choiceString string, outputLength int) string {
	numChoices := len(choiceString)
	output := make([]byte, outputLength)
	for i := 0; i < outputLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(numChoices)))
		if err != nil {
			panic(err)
		}
		output[i] = choiceString[randomIndex.Int64()]
	}
	return string(output)
}

func randomInt(min int, max int) int {
	offset, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}
	return min + int(offset.Int64())
}
