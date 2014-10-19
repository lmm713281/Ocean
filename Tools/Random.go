package Tools

import (
	"math/rand"
)

func RandomInteger(max int) (rnd int) {
	rnd = rand.Intn(max)
	return
}
