package repository

import (
	"challenge/domain"
	"challenge/lib/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DeviceRepositoryDb struct {
	db dynamodbiface.DynamoDBAPI
}

func (d DeviceRepositoryDb) Create(device *domain.Device) (*domain.Device, *errors.AppError) {
	marshaledDevice, _ := dynamodbattribute.MarshalMap(device)
	dynamoDbItem := &dynamodb.PutItemInput{
		Item:      marshaledDevice,
		TableName: aws.String("Device"),
	}
	_, err := d.db.PutItem(dynamoDbItem)
	if err != nil {
		return nil, errors.InternalServerError("Internal Server Error")
	}
	return device, nil
}
func (d DeviceRepositoryDb) FindById(id string) (*domain.Device, *errors.AppError) {
	dynamoDbItem := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("Device"),
	}
	result, err := d.db.GetItem(dynamoDbItem)
	if err != nil {
		return nil, errors.InternalServerError("Internal Server Error")
	}
	if result.Item == nil {
		return nil, errors.NotFoundError("Device not found")
	}
	device := &domain.Device{}
	err = dynamodbattribute.UnmarshalMap(result.Item, device)
	if err != nil {
		return nil, errors.InternalServerError("Internal Server Error")
	}
	return device, nil
}
func NewDeviceRepositoryDb(dbClient dynamodbiface.DynamoDBAPI) DeviceRepositoryDb {
	return DeviceRepositoryDb{dbClient}
}
