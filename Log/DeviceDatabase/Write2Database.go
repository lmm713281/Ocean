package DeviceDatabase

func write2Database(entry LogDBEntry) {
	if err := logDBCollection.Insert(entry); err != nil {
		// Can not log here to prevent endless loop (consumer is also producer)
	}
}
