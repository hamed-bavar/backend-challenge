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
	//dynamoDbItem := &dynamodb.PutItemInput{
	//	Item:      marshaledDevice,
	//	TableName: aws.String("Device"),
	//}
	//_, err := d.db.PutItem(dynamoDbItem)
	//if err != nil {
	//	logger.Error("Error while putting item in dynamoDb" + err.Error())
	//	return nil, errors.InternalServerError("Internal Server Error")
	//}
	//return device, nil
	err := d.table.Put(device).Run()
	if err != nil {
		logger.Error("Error while putting item in dynamoDb" + err.Error())
		return nil, errors.InternalServerError("Internal Server Error" + err.Error())
	}
	return device, nil
}

func (d DeviceRepositoryDb) FindById(id string) (*domain.Device, *errors.AppError) {
	//dynamoDbItem := &dynamodb.GetItemInput{
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"id": {
	//			S: aws.String(id),
	//		},
	//	},
	//	TableName: aws.String("Device"),
	//}
	var result domain.Device
	err := d.table.Get("id", id).
		One(&result)
	if err != nil {
		return nil, errors.NotFoundError("Device not found" + err.Error())
	}

	//result, err := d.db.GetItem(dynamoDbItem)
	//if err != nil {
	//	logger.Error("Error while getting item from dynamoDb" + err.Error())
	//	return nil, errors.InternalServerError("Internal Server Error")
	//}
	//if result.Item == nil {
	//	return nil, errors.NotFoundError("Device not found")
	//}
	//device := &domain.Device{}
	//err = dynamodbattribute.UnmarshalMap(result.Item, device)
	//if err != nil {
	//	logger.Error("Error while unmarshalling item from dynamoDb" + err.Error())
	//	return nil, errors.InternalServerError("Internal Server Error")
	//}
	return &result, nil
}
func NewDeviceRepositoryDb(dbClient *dynamo.DB) DeviceRepositoryDb {
	var table = dbClient.Table("Device")
	return DeviceRepositoryDb{table}
}
