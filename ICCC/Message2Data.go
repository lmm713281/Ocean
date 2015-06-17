package ICCC

import (
	"fmt"
	"reflect"
	"strconv"
)

// Function to convert an ICCC message to HTTP data.
func message2Data(channel, command string, message interface{}) (data map[string][]string) {

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

		switch field.Kind().String() {

		case `int64`:
			key := fmt.Sprintf(`int:%s`, keyName)
			data[key] = []string{strconv.FormatInt(field.Int(), 10)}

		case `string`:
			key := fmt.Sprintf(`str:%s`, keyName)
			data[key] = []string{field.String()}

		case `float64`:
			key := fmt.Sprintf(`f64:%s`, keyName)
			data[key] = []string{strconv.FormatFloat(field.Float(), 'f', 9, 64)}

		case `bool`:
			key := fmt.Sprintf(`bool:%s`, keyName)
			data[key] = []string{strconv.FormatBool(field.Bool())}

		case `uint8`: // byte
			key := fmt.Sprintf(`ui8:%s`, keyName)
			data[key] = []string{strconv.FormatUint(field.Uint(), 16)}

		// Case: Arrays...
		case `slice`:
			sliceLen := field.Len()
			if sliceLen > 0 {
				sliceKind := field.Index(0).Kind()
				key := ``
				dataValues := make([]string, sliceLen, sliceLen)
				switch sliceKind.String() {
				case `uint8`: // bytes
					key = fmt.Sprintf(`ui8[]:%s`, keyName)
					values := field.Interface().([]uint8)
					for index, value := range values {
						dataValues[index] = strconv.FormatUint(uint64(value), 16)
					}

				case `int64`:
					key = fmt.Sprintf(`int[]:%s`, keyName)
					values := field.Interface().([]int64)
					for index, value := range values {
						dataValues[index] = strconv.FormatInt(value, 10)
					}

				case `bool`:
					key = fmt.Sprintf(`bool[]:%s`, keyName)
					values := field.Interface().([]bool)
					for index, value := range values {
						dataValues[index] = strconv.FormatBool(value)
					}

				case `string`:
					key = fmt.Sprintf(`str[]:%s`, keyName)
					values := field.Interface().([]string)
					for index, value := range values {
						dataValues[index] = value
					}

				case `float64`:
					key = fmt.Sprintf(`f64[]:%s`, keyName)
					values := field.Interface().([]float64)
					for index, value := range values {
						dataValues[index] = strconv.FormatFloat(value, 'f', 9, 64)
					}
				}

				data[key] = dataValues
			}
		}
	}

	return
}
