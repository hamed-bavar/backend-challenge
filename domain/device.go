package domain

import "challenge/lib/errors"

type Device struct {
	Id          string `json:"id" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}

//go:generate mockgen -destination=../mocks/mockDeviceRepository.go -package=domain . DeviceRepository
type DeviceRepository interface {
	Create(device *Device) (*Device, *errors.AppError)
	FindById(id string) (*Device, *errors.AppError)
}
