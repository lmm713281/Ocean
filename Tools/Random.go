package Tools

import (
	"github.com/twinj/uuid"
	"math/rand"
)

// Gets a random integer between 0 and max-1.
func RandomInteger(max int) (rnd int) {
	rnd = rand.Intn(max)
	return
}

// Gets a random 64 bit float between 0.0 and 1.0 (but without 1.0).
func RandomFloat64() (rnd float64) {
	rnd = rand.Float64()
	return
}

// Gets a random UUID (v4).
func RandomGUID() (guidString string) {
	guidString = uuid.NewV4().String()
	return
}
