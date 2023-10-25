package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	// deprecated sin Go 1.20
	//rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
