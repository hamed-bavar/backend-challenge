package repository

import (
	"challenge/domain"
	"challenge/lib/errors"
	"challenge/lib/logger"
	"github.com/guregu/dynamo"
)

//go:generate mockgen -destination=../mocks/mockDeviceRepository/mockDeviceRepository.go -package=mockDeviceRepository . DeviceRepository
type DeviceRepository interface {
	Create(device *domain.Device) (*domain.Device, *errors.AppError)
	FindById(id string) (*domain.Device, *errors.AppError)
}

type DeviceRepositoryDb struct {
	table dynamo.Table
}

func (d DeviceRepositoryDb) Create(device *domain.Device) (*domain.Device, *errors.AppError) {
	err := d.table.Put(device).Run()
	if err != nil {
		logger.Error("Error while putting item in dynamoDb")
		return nil, errors.InternalServerError("Internal Server Error")
	}
	return device, nil
}

func (d DeviceRepositoryDb) FindById(id string) (*domain.Device, *errors.AppError) {
	var result domain.Device
	err := d.table.Get("id", id).
		One(&result)
	if err != nil {
		return nil, errors.NotFoundError("Device not found")
	}
	return &result, nil
}
func NewDeviceRepositoryDb(dbClient *dynamo.DB) DeviceRepositoryDb {
	var table = dbClient.Table("Device")
	return DeviceRepositoryDb{table}
}
