package ICCC

import "fmt"
import "reflect"
import "strconv"

func Data2Message(target interface{}, data map[string][]string) (channel, command string, obj interface{}) {
	if data == nil || len(data) == 0 {
		channel = ``
		command = ``
		obj = nil
		return
	}

	element := reflect.ValueOf(target)
	element = element.Elem()
	elementType := element.Type()

	channel = data[`channel`][0]
	command = data[`command`][0]
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

	obj = target
	return
}
