package NumGen

import (
	"math/big"
)

// Internal function to generate a unique ID with the given sequenceNumber.
func generateUniqueID(sequenceNumber int) (id int64) {

	// Ensure, that the sequenceNumber is not smaller than 0:
	if sequenceNumber < 0 {
		sequenceNumber = sequenceNumber * -1
	}

	// Ensure, that the sequenceNumber is not bigger than 999:
	if sequenceNumber > 999 {
		sequenceNumber = 999
	}

	// Convert the sequenceNumber to a 64 bit integer:
	seq := int64(sequenceNumber)

	// Generate the first and second part of the ID:
	timeID := generateTimeID()       // First part
	machineID := generateMachineID() // Second part

	// Convert all numbers to big integers:
	bigSEQ := big.NewInt(seq)
	bigMachineID := big.NewInt(machineID)
	bigTimeID := big.NewInt(timeID)

	// Add the parts to get the result:
	bigResult := bigTimeID.Add(bigTimeID, bigMachineID)
	bigResult = bigResult.Add(bigResult, bigSEQ)

	// The result as 64 bit integer:
	id = bigResult.Int64()
	return
}
