package repository

import (
	"challenge/domain"
	"challenge/mocks/mockDynamoDBAPI"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"testing"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}
func Test_Should_Return_Created_Device(t *testing.T) {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "nice car",
		Serial:      "8765432",
	}
	marshaledDevice, _ := dynamodbattribute.MarshalMap(mockDevice)
	dynamoDbItem := &dynamodb.PutItemInput{
		Item:      marshaledDevice,
		TableName: aws.String("Device"),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().PutItem(dynamoDbItem).Return(nil, nil)
	repo := NewDeviceRepositoryDb(mockDB)
	createdDevice, _ := repo.Create(&mockDevice)
	assert.Equal(t, mockDevice, *createdDevice)
}
func Test_Should_Return_Server_Error(t *testing.T) {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "nice car",
		Serial:      "8765432",
	}
	marshaledDevice, _ := dynamodbattribute.MarshalMap(mockDevice)
	dynamoDbItem := &dynamodb.PutItemInput{
		Item:      marshaledDevice,
		TableName: aws.String("Device"),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().PutItem(dynamoDbItem).Return(nil, &MyError{})
	repo := NewDeviceRepositoryDb(mockDB)
	_, repoError := repo.Create(&mockDevice)

	assert.Equal(t, (*repoError).Message, "Internal Server Error")
}
