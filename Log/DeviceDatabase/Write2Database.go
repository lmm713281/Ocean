package DeviceDatabase

import (
	"fmt"
)

// Function to write a logging event to the database.
func write2Database(entry LogDBEntry) {
	// Try to write the event to the database:
	if err := logDBCollection.Insert(entry); err != nil {
		// Case: Error!
		// Cannot log here to prevent endless loop (consumer is also producer)
		fmt.Printf("Was not able to write a logging event to the database: '%s'. The log entry was: '%s'.\n", err.Error(), entry.Format())
	}
}
