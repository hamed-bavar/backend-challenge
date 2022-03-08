package service

import (
	"challenge/domain"
	"challenge/lib/errors"
	"challenge/repository"
)

type DeviceService interface {
	CreateDevice(*domain.Device) (*domain.Device, *errors.AppError)
}
type DefaultDeviceService struct {
	repo repository.DeviceRepository
}

func (s DefaultDeviceService) CreateDevice(device *domain.Device) (*domain.Device, *errors.AppError) {
	d, err := s.repo.Create(device)
	if err != nil {
		return nil, err
	}
	return d, nil
}
func NewDeviceService(repo repository.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{repo}
}
