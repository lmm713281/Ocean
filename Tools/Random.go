package Tools

import (
	"github.com/twinj/uuid"
	"math/rand"
)

func RandomInteger(max int) (rnd int) {
	rnd = rand.Intn(max)
	return
}

func RandomGUID() (guidString string) {
	guidString = uuid.NewV4().String()
	guidString = guidString[1 : len(guidString)-1]
	return
}
