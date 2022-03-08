package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"test/domain"
	"test/lib/errors"
)

type DeviceRepositoryDb struct {
	db dynamodbiface.DynamoDBAPI
}

func (d *DeviceRepositoryDb) create(device *domain.Device) (*domain.Device, *errors.AppError) {
	marshaledDevice, _ := dynamodbattribute.MarshalMap(device)
	dynamoDbItem := &dynamodb.PutItemInput{
		Item:      marshaledDevice,
		TableName: aws.String("Device"),
	}
	_, err := d.db.PutItem(dynamoDbItem)
	if err != nil {
		return nil, errors.InternalServerError("Unexpected database error")
	}
	return device, nil
}
func NewDeviceRepositoryDb(dbClient dynamodbiface.DynamoDBAPI) DeviceRepositoryDb {
	return DeviceRepositoryDb{dbClient}
}
