package domain

import "test/lib/errs"

type Device struct {
	Id          string `json:"id" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}
type DeviceRepository interface {
	findById(string) (*Device, errs.AppError)
	create(Device) (Device, errs.AppError)
}