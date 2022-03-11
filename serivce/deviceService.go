package service

import (
	"challenge/domain"
	"challenge/lib/errors"
	"challenge/lib/logger"
	"challenge/repository"
	"github.com/go-playground/validator/v10"
)

//go:generate mockgen -destination=../mocks/mockDeviceService/mockDeviceService.go -package=mockDeviceService . DeviceService
type DeviceService interface {
	CreateDevice(*domain.Device) (*domain.Device, *errors.AppError)
	GetDevice(string) (*domain.Device, *errors.AppError)
}
type DefaultDeviceService struct {
	repo repository.DeviceRepository
}

func (s DefaultDeviceService) CreateDevice(device *domain.Device) (*domain.Device, *errors.AppError) {
	validate := validator.New()
	validationError := validate.Struct(device)
	if validationError != nil {
		logger.Error("Error while validating data" + validationError.Error())
		return nil, errors.ValidationError("some fields are missing")
	}
	d, err := s.repo.Create(device)
	if err != nil {
		return nil, err
	}
	return d, nil
}
func (s DefaultDeviceService) GetDevice(id string) (*domain.Device, *errors.AppError) {
	d, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return d, nil
}
func NewDeviceService(repo repository.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{repo}
}
