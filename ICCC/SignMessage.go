package ICCC

import (
	"crypto/sha512"
	"fmt"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/url"
	"sort"
)

// Sign a message to secure it.
func signMessage(data map[string][]string) (result url.Values) {
	// Create the hash generator:
	hash := sha512.New()

	// To the start, we hash the password:
	fmt.Fprintf(hash, "%s", Tools.InternalCommPassword())

	// Because the order of a map is random, we have to sort
	// the keys beforehand. Next, we can use the ordered keys
	// to access the data:
	keys := []string{}

	// Get all keys:
	for key := range data {
		keys = append(keys, key)
	}

	// Sort the keys:
	sort.Strings(keys)

	// Now, loop over all the data:
	for _, key := range keys {
		// Get the value:
		value := data[key]

		// Hash each key and value:
		fmt.Fprintf(hash, "key=%s :: value=%s\n", key, value)
	}

	// Create the result:
	result = url.Values(data)

	// Append the sign:
	result.Add(`checksum`, fmt.Sprintf("%x", hash.Sum(nil)))
	return
}
