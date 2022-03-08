package serivce

import (
	"test/domain"
	"test/lib/errors"
)

type DeviceService interface {
	CreateDevice(domain.Device) (domain.Device, *errors.AppError)
}
type DefaultDeviceService struct {
	repo domain.DeviceRepository
}

func (s DefaultDeviceService) CreateDevice(device domain.Device) (*domain.Device, *errors.AppError) {
	d, err := s.repo.Create(device)
	if err != nil {
		return nil, err
	}
	return d, nil
}
func NewDeviceService(repo domain.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{repo}
}
