package ICCC

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"io"
	"net/url"
	"reflect"
	"strings"
)

// Function to convert an ICCC message to HTTP data.
func Message2Data(channel, command string, message interface{}) (data map[string][]string) {
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, fmt.Sprintf("Was not able to convert the message to HTTP values. %s", err))
			data = make(map[string][]string, 0)
			return
		}
	}()

	// Create the map:
	data = make(map[string][]string)

	// Add the meta information:
	data[`command`] = []string{command}
	data[`channel`] = []string{channel}

	if message == nil {
		return
	}

	// Use reflection to determine the types:
	element := reflect.ValueOf(message)
	elementType := element.Type()

	// Iterate over all fields of the data type.
	// Transform the data regarding the type.
	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		keyName := elementType.Field(i).Name

		// A buffer for the binary representation:
		buffer := new(bytes.Buffer)

		// For possible errors:
		var errConverter error = nil

		// The key for this element:
		key := ``

		// Look for the right data type:
		switch field.Kind().String() {

		case `int64`:
			key = fmt.Sprintf(`int:%s`, keyName)

			// Converts the value in a byte array:
			errConverter = binary.Write(buffer, binary.LittleEndian, field.Int())

		case `string`:
			key = fmt.Sprintf(`str:%s`, keyName)

			// URL encode the string and copy its bytes to the buffer:
			io.Copy(buffer, strings.NewReader(url.QueryEscape(field.String())))

		case `float64`:
			key = fmt.Sprintf(`f64:%s`, keyName)

			// Converts the value in a byte array:
			errConverter = binary.Write(buffer, binary.LittleEndian, field.Float())

		case `bool`:
			key = fmt.Sprintf(`bool:%s`, keyName)

			// Directly convert the bool in a byte:
			if field.Bool() {
				// Case: True
				buffer.WriteByte(0x1) // Write 1
			} else {
				// Case: False
				buffer.WriteByte(0x0) // Write 0
			}

		case `uint8`: // a byte
			key = fmt.Sprintf(`ui8:%s`, keyName)

			// uint8 is a byte, thus, write it directly in the buffer:
			buffer.WriteByte(byte(field.Uint()))

		// Case: Arrays...
		case `slice`:
			sliceLen := field.Len()
			if sliceLen > 0 {

				// Which kind of data is this?
				sliceKind := field.Index(0).Kind()

				// Select the right data type:
				switch sliceKind.String() {
				case `uint8`: // a byte array
					key = fmt.Sprintf(`ui8[]:%s`, keyName)
					values := field.Interface().([]uint8)

					// Directly write the bytes in the buffer:
					for _, val := range values {
						buffer.WriteByte(byte(val))
					}

				case `int64`:
					key = fmt.Sprintf(`int[]:%s`, keyName)
					values := field.Interface().([]int64)

					// Converts the array in a byte array:
					errConverter = binary.Write(buffer, binary.LittleEndian, values)

				case `bool`:
					key = fmt.Sprintf(`bool[]:%s`, keyName)
					values := field.Interface().([]bool)

					// Cannot convert bool to bytes by using binary.Write(). Thus,
					// convert it by ower own:

					// Loop over all values:
					for _, val := range values {
						if val {
							// If the value is true:
							buffer.WriteByte(0x1) // Write 1
						} else {
							// If the value is false:
							buffer.WriteByte(0x0) // Write 0
						}
					}

				case `string`:
					key = fmt.Sprintf(`str[]:%s`, keyName)
					values := field.Interface().([]string)

					// Mask every string by using a URL encoding.
					// This masks e.g. every new-line i.e \n, etc.
					// This allows us to combine later the strings by
					// using \n:
					masked := make([]string, len(values))

					// Loop over each string and convert it:
					for n, val := range values {
						masked[n] = url.QueryEscape(val)
					}

					// Join all masked strings by using \n and copy the byte array
					// representation in the buffer:
					io.Copy(buffer, strings.NewReader(strings.Join(masked, "\n")))

				case `float64`:
					key = fmt.Sprintf(`f64[]:%s`, keyName)
					values := field.Interface().([]float64)

					// Converts the array in a byte array:
					errConverter = binary.Write(buffer, binary.LittleEndian, values)
				}
			}
		}

		if errConverter != nil {
			// An error occurs:
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `It was not possible to convert an array for an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errConverter.Error())
		} else {
			// Convert the byte array to a base64 string for the transportation on wire:
			data[key] = []string{base64.StdEncoding.EncodeToString(buffer.Bytes())}
		}
	}

	return
}
