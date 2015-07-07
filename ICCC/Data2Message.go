package ICCC

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/url"
	"reflect"
	"strings"
)

// Function to convert the HTTP data back to a message.
func Data2Message(target interface{}, data map[string][]string) (channel, command string, obj interface{}) {
	defer func() {
		if err := recover(); err != nil {
			channel = ``
			command = ``
			obj = nil
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, fmt.Sprintf("Was not able to convert the HTTP values to a message. %s", err))
			return
		}
	}()

	if data == nil || len(data) == 0 {
		channel = ``
		command = ``
		obj = nil
		return
	}

	// By using reflection, determine the type:
	elementType := reflect.TypeOf(target)

	// Is it a pointer?
	if elementType.Kind() == reflect.Ptr {
		// Get the value behind the pointer:
		elementType = elementType.Elem()
	}

	// ICCC works with structs! If this is not a struct, its an error.
	if elementType.Kind() != reflect.Struct {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactMiddle, LM.MessageNamePARSE, `Was not able to transform HTTP data to a message, because the given data was not a struct.`)
		return
	}

	// By using reflection, create a new instance:
	element := reflect.New(elementType).Elem()
	channel = data[`channel`][0]
	command = data[`command`][0]

	// Use the order of the destination type's fields:
	for i := 0; i < element.NumField(); i++ {

		// Read the current field:
		field := element.Field(i)

		// Choose the right type for this field:
		switch field.Kind().String() {
		case `int64`:
			// The name of the field:
			mapName := fmt.Sprintf(`int:%s`, elementType.Field(i).Name)

			// The value of the field as string:
			mapValue := data[mapName][0]

			// The value of the field as bytes:
			bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)

			if errBase64 != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
			} else {
				// The destination:
				var destination int64

				// A reader for the bytes:
				buffer := bytes.NewReader(bytesArray)

				// Try to decode the bytes to an instance of the type:
				errBinary := binary.Read(buffer, binary.LittleEndian, &destination)
				if errBinary != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the binary data to the type of the ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBinary.Error())
				} else {
					// Finnaly, store the value in the message:
					field.SetInt(destination)
				}
			}

		case `string`:
			// The name of the field:
			mapName := fmt.Sprintf(`str:%s`, elementType.Field(i).Name)

			// The value of the field as string:
			mapValue := data[mapName][0]

			// The value of the field as bytes:
			bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)

			if errBase64 != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
			} else {
				// Decode the bytes as string:
				text := string(bytesArray)

				// Decode the URL encoded string:
				textFinal, errURL := url.QueryUnescape(text)
				if errURL != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode a URL encoded string.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errURL.Error())
				} else {
					field.SetString(textFinal)
				}
			}

		case `float64`:
			// The name of the field:
			mapName := fmt.Sprintf(`f64:%s`, elementType.Field(i).Name)

			// The value of the field as string:
			mapValue := data[mapName][0]

			// The value of the field as bytes:
			bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)

			if errBase64 != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
			} else {
				// The destination:
				var destination float64

				// A reader for the bytes:
				buffer := bytes.NewReader(bytesArray)

				// Try to decode the bytes to an instance of the type:
				errBinary := binary.Read(buffer, binary.LittleEndian, &destination)
				if errBinary != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the binary data to the type of the ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBinary.Error())
				} else {
					// Finnaly, store the value in the message:
					field.SetFloat(destination)
				}
			}

		case `bool`:
			// The name of the field:
			mapName := fmt.Sprintf(`bool:%s`, elementType.Field(i).Name)

			// The value of the field as string:
			mapValue := data[mapName][0]

			// The value of the field as bytes:
			bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)
			if errBase64 != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
			} else {
				// Store the value:
				if bytesArray[0] == 0x1 {
					field.SetBool(true)
				} else {
					field.SetBool(false)
				}
			}

		case `uint8`:
			// The name of the field:
			mapName := fmt.Sprintf(`ui8:%s`, elementType.Field(i).Name)

			// The value of the field as string:
			mapValue := data[mapName][0]

			// The value of the field as bytes:
			bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)

			if errBase64 != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
			} else {
				// Store the value:
				field.SetUint(uint64(bytesArray[0]))
			}

		// Case: Arrays...
		case `slice`:
			sliceInterface := field.Interface()
			sliceKind := reflect.ValueOf(sliceInterface).Type().String()

			switch sliceKind {
			case `[]uint8`: // a byte array
				// The name of the field:
				mapName := fmt.Sprintf(`ui8[]:%s`, elementType.Field(i).Name)

				// The value of the field as string:
				mapValue := data[mapName][0]

				// The value of the field as bytes:
				bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)

				if errBase64 != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
				} else {
					// Store the values in the message:
					fieldDataValue := reflect.ValueOf(bytesArray)
					field.Set(fieldDataValue)
				}

			case `[]int64`:
				// The name of the field:
				mapName := fmt.Sprintf(`int[]:%s`, elementType.Field(i).Name)

				// The value of the field as string:
				mapValue := data[mapName][0]

				// The value of the field as bytes:
				bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)
				if errBase64 != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
				} else {
					// The destination:
					var destination []int64

					// A reader for the bytes:
					buffer := bytes.NewReader(bytesArray)

					// Try to decode the bytes to an instance of the type:
					errBinary := binary.Read(buffer, binary.LittleEndian, &destination)
					if errBinary != nil {
						Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the binary data to the type of the ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBinary.Error())
					} else {
						// Finnaly, store the value in the message:
						fieldDataValue := reflect.ValueOf(destination)
						field.Set(fieldDataValue)
					}
				}

			case `[]bool`:
				// The name of the field:
				mapName := fmt.Sprintf(`bool[]:%s`, elementType.Field(i).Name)

				// The value of the field as string:
				mapValue := data[mapName][0]

				// The value of the field as bytes:
				bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)
				if errBase64 != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
				} else {
					fieldLen := len(bytesArray)
					fieldData := make([]bool, fieldLen, fieldLen)

					// Convert each byte in a bool:
					for n, value := range bytesArray {
						if value == 0x1 {
							fieldData[n] = true
						} else {
							fieldData[n] = false
						}
					}

					// Store the values in the message:
					fieldDataValue := reflect.ValueOf(fieldData)
					field.Set(fieldDataValue)
				}

			case `[]string`:
				// The name of the field:
				mapName := fmt.Sprintf(`str[]:%s`, elementType.Field(i).Name)

				// The value of the field as string:
				mapValue := data[mapName][0]

				// The value of the field as bytes:
				bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)
				if errBase64 != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
				} else {
					// Get the URL encoded string of all values:
					allStringsRAW := string(bytesArray)

					// Split now the different strings:
					allStrings := strings.Split(allStringsRAW, "\n")

					// A place where we store the final strings:
					data := make([]string, len(allStrings), len(allStrings))

					// Loop over all URL encoded strings and decode it:
					for n, element := range allStrings {
						elementFinal, errURL := url.QueryUnescape(element)
						if errURL != nil {
							Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode a base64 string for a string array.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errURL.Error())
						} else {
							data[n] = elementFinal
						}
					}

					// Store the values in the message:
					fieldDataValue := reflect.ValueOf(data)
					field.Set(fieldDataValue)
				}

			case `[]float64`:
				// The name of the field:
				mapName := fmt.Sprintf(`f64[]:%s`, elementType.Field(i).Name)

				// The value of the field as string:
				mapValue := data[mapName][0]

				// The value of the field as bytes:
				bytesArray, errBase64 := base64.StdEncoding.DecodeString(mapValue)
				if errBase64 != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the base64 data to an ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBase64.Error())
				} else {
					// The destination:
					var destination []float64

					// A reader for the bytes:
					buffer := bytes.NewReader(bytesArray)

					// Try to decode the bytes to an instance of the type:
					errBinary := binary.Read(buffer, binary.LittleEndian, &destination)
					if errBinary != nil {
						Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to decode the binary data to the type of the ICCC message.`, fmt.Sprintf("channel='%s'", channel), fmt.Sprintf("command='%s'", command), errBinary.Error())
					} else {
						// Finnaly, store the value in the message:
						fieldDataValue := reflect.ValueOf(destination)
						field.Set(fieldDataValue)
					}
				}
			}
		}
	}

	obj = element.Interface()
	return
}
