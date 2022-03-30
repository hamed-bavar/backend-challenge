package repository

import (
	"challenge/domain"
	"challenge/lib/errors"
	"github.com/guregu/dynamo"
)

//go:generate mockgen -destination=../mocks/mockDeviceRepository/mockDeviceRepository.go -package=mockDeviceRepository . DeviceRepository
type DeviceRepository interface {
	Create(device *domain.Device) (*domain.Device, *errors.AppError)
	FindById(id string) (*domain.Device, *errors.AppError)
}

type DeviceRepositoryDb struct {
	db *dynamo.DB
}

func (d DeviceRepositoryDb) Create(device *domain.Device) (*domain.Device, *errors.AppError) {
	//marshaledDevice, _ := dynamodbattribute.MarshalMap(device)
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
	return &domain.Device{Id: "1", Name: "test"}, nil
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
	return &domain.Device{Id: "1", Name: "test"}, nil
}
func NewDeviceRepositoryDb(dbClient *dynamo.DB) DeviceRepositoryDb {
	return DeviceRepositoryDb{dbClient}
}
