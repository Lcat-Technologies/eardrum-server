package model


// returns the device's id
func (n NewDevice) GetDeviceID() string {
	return n.DeviceID
}

// returns the device's model
func (n NewDevice) GetDeviceModel() string {
	return n.Model
}