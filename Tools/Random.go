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

// Gets a random UUID (v4).
func RandomGUID() (guidString string) {
	guidString = uuid.NewV4().String()
	guidString = guidString[1 : len(guidString)-1]
	return
}
