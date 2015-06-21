package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"reflect"
	"strconv"
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
		field := element.Field(i)
		switch field.Kind().String() {

		case `int64`:
			mapName := fmt.Sprintf(`int:%s`, elementType.Field(i).Name)
			mapValue := data[mapName][0]
			v, _ := strconv.ParseInt(mapValue, 10, 64)
			field.SetInt(v)

		case `string`:
			mapName := fmt.Sprintf(`str:%s`, elementType.Field(i).Name)
			mapValue := data[mapName][0]
			field.SetString(mapValue)

		case `float64`:
			mapName := fmt.Sprintf(`f64:%s`, elementType.Field(i).Name)
			mapValue := data[mapName][0]
			v, _ := strconv.ParseFloat(mapValue, 64)
			field.SetFloat(v)

		case `bool`:
			mapName := fmt.Sprintf(`bool:%s`, elementType.Field(i).Name)
			mapValue := data[mapName][0]
			v, _ := strconv.ParseBool(mapValue)
			field.SetBool(v)

		case `uint8`:
			mapName := fmt.Sprintf(`ui8:%s`, elementType.Field(i).Name)
			mapValue := data[mapName][0]
			v, _ := strconv.ParseUint(mapValue, 16, 8)
			field.SetUint(v)

		// Case: Arrays...
		case `slice`:
			sliceInterface := field.Interface()
			sliceKind := reflect.ValueOf(sliceInterface).Type().String()

			switch sliceKind {
			case `[]uint8`: // bytes
				mapName := fmt.Sprintf(`ui8[]:%s`, elementType.Field(i).Name)
				mapValues := data[mapName]
				fieldLen := len(mapValues)
				fieldData := make([]uint8, fieldLen, fieldLen)
				for n, mapValue := range mapValues {
					v, _ := strconv.ParseUint(mapValue, 16, 8)
					fieldData[n] = byte(v)
				}

				fieldDataValue := reflect.ValueOf(fieldData)
				field.Set(fieldDataValue)

			case `[]int64`:
				mapName := fmt.Sprintf(`int[]:%s`, elementType.Field(i).Name)
				mapValues := data[mapName]
				fieldLen := len(mapValues)
				fieldData := make([]int64, fieldLen, fieldLen)
				for n, mapValue := range mapValues {
					v, _ := strconv.ParseInt(mapValue, 10, 64)
					fieldData[n] = v
				}

				fieldDataValue := reflect.ValueOf(fieldData)
				field.Set(fieldDataValue)

			case `[]bool`:
				mapName := fmt.Sprintf(`bool[]:%s`, elementType.Field(i).Name)
				mapValues := data[mapName]
				fieldLen := len(mapValues)
				fieldData := make([]bool, fieldLen, fieldLen)
				for n, mapValue := range mapValues {
					v, _ := strconv.ParseBool(mapValue)
					fieldData[n] = v
				}

				fieldDataValue := reflect.ValueOf(fieldData)
				field.Set(fieldDataValue)

			case `[]string`:
				mapName := fmt.Sprintf(`str[]:%s`, elementType.Field(i).Name)
				mapValues := data[mapName]
				fieldDataValue := reflect.ValueOf(mapValues)
				field.Set(fieldDataValue)

			case `[]float64`:
				mapName := fmt.Sprintf(`f64[]:%s`, elementType.Field(i).Name)
				mapValues := data[mapName]
				fieldLen := len(mapValues)
				fieldData := make([]float64, fieldLen, fieldLen)
				for n, mapValue := range mapValues {
					v, _ := strconv.ParseFloat(mapValue, 64)
					fieldData[n] = v
				}

				fieldDataValue := reflect.ValueOf(fieldData)
				field.Set(fieldDataValue)
			}
		}
	}

	obj = element.Interface()
	return
}
