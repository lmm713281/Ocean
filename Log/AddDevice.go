package Log

import "github.com/SommerEngineering/Ocean/Log/Device"

/*
Registering the logging devices. Normally, it is not necessary to call this function. To enable or disable a logging device,
please use the configuration database instead. But if you create your own logging device, let say a e-mail logger, then you
are able to use this function to activate your own logging device. It is save to use this function at any time and it is
thread-save ;-)
*/
func AddLoggingDevice(device Device.Device) {

	newDevice := device
	go func() {
		mutexDevices.Lock()
		devices.PushBack(newDevice)
		mutexDevices.Unlock()
	}()
}
