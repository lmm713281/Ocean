package NumGen

import (
	"time"
)

// Function to generate a unique ID.
func GetUniqueID() (id int64) {

	// Exklusive access:
	genLock.Lock()
	defer genLock.Unlock()
	incremented := false

Retry:

	// Read the current time
	t1 := time.Now().UTC()
	t1Milliseconds := int(float64(t1.Nanosecond()) / 1000000.0)

	// Calculate the difference to the last generated number:
	diff := t1.Sub(genCurrentTime)

	// Case 1: Huge difference?
	if diff.Seconds() > 1.0 {
		genCurrentTime = t1
		genCurrentMillisecond = t1Milliseconds
		genCurrentMillisecondCounter = 0
		id = generateUniqueID(genCurrentMillisecondCounter)
		return
	}

	/*
		The fist case is necessary, because the same count of milliseconds
		can occur every second, etc. ;-) Therefore, a check only by the
		milliseconds is not sufficient.
	*/

	// Case 2: Small difference?
	if t1Milliseconds != genCurrentMillisecond {
		genCurrentTime = t1
		genCurrentMillisecond = t1Milliseconds
		genCurrentMillisecondCounter = 0
		id = generateUniqueID(genCurrentMillisecondCounter)
		return
	}

	//
	// Case 3: Another number must be generated at the same millisecond!
	//

	if incremented == false {
		genCurrentMillisecondCounter++
		incremented = true
	}

	// Case 3.1: More than 1000 numbers are generated at one millisecond?
	if genCurrentMillisecondCounter > 999 {
		// This is not possible with the current algorithm!
		// Therefore, we force this and all other request to wait some time:
		time.Sleep(1 * time.Millisecond)

		// Try it again:
		goto Retry
		// Case 3.2: Less than 1000 numbers are generated at one millisecond?
	} else {
		// This case is fine:
		id = generateUniqueID(genCurrentMillisecondCounter)
		return
	}
}
