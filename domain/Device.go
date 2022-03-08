package domain

type Device struct {
	id          string `validate:"required"`
	deviceModel string `validate:"required"`
	name        string `validate:"required"`
	note        string `validate:"required"`
	serial      string `validate:"required"`
}

type DeviceRepository interface {
	findById(string) ([]Device, error)
	create(Device) (Device, error)
}
