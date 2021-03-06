package Log

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// Converts a list with logging events to an array.
func logEntryListToArray(data *list.List) (result []Meta.Entry) {
	count := data.Len()
	result = make([]Meta.Entry, count, count)
	position := 0
	for entry := data.Front(); entry != nil; entry = entry.Next() {
		result[position] = entry.Value.(Meta.Entry)
		position++
	}

	return
}
