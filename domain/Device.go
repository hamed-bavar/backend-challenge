package domain

type Device struct {
	Id          string `json:"id" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}

type DeviceRepository interface {
	findById(string) ([]Device, error)
	create(Device) (Device, error)
}
