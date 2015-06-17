package main

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/System"
)

/*
This is the entry point of Ocean in case of using it as e.g. messaging broker
or logging service, etc. This function does not matter if Ocean is used as
framework.
*/
func main() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Ocean is starting.`)
	System.InitHandlers()
	System.StartAndBlockForever()
}
