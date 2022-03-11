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

func Test_Should_Return_Create_Device(t *testing.T) {
	dynamoDbItem, mockDevice := getMockDevice()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().PutItem(dynamoDbItem).Return(nil, nil)
	repo := NewDeviceRepositoryDb(mockDB)
	createdDevice, _ := repo.Create(&mockDevice)
	assert.Equal(t, mockDevice, *createdDevice)
}
func Test_Should_Return_Server_Error_While_Creating_Device(t *testing.T) {
	dynamoDbItem, mockDevice := getMockDevice()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().PutItem(dynamoDbItem).Return(nil, &MyError{})
	repo := NewDeviceRepositoryDb(mockDB)
	_, repoError := repo.Create(&mockDevice)

	assert.Equal(t, "Internal Server Error", (*repoError).Message)
}
func Test_Should_Return_Device_With_Specific_Id(t *testing.T) {
	_, expectedDevice := getMockDevice()
	dynamoDbItem := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1234"),
			},
		},
		TableName: aws.String("Device"),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().GetItem(dynamoDbItem).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1234"),
			},
			"deviceModel": {
				S: aws.String("mercedes"),
			},
			"name": {
				S: aws.String("your car"),
			},
			"note": {
				S: aws.String("nice car"),
			},
			"serial": {
				S: aws.String("8765432"),
			},
		},
	}, nil)
	repo := NewDeviceRepositoryDb(mockDB)
	actualDevice, _ := repo.FindById("1234")
	assert.Equal(t, expectedDevice, *actualDevice)
}
func Test_Should_Return_Server_Error_While_Getting_Device(t *testing.T) {
	dynamoDbItem := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1234"),
			},
		},
		TableName: aws.String("Device"),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDynamoDBAPI.NewMockDynamoDBAPI(ctrl)
	mockDB.EXPECT().GetItem(dynamoDbItem).Return(nil, &MyError{})
	repo := NewDeviceRepositoryDb(mockDB)
	_, err := repo.FindById("1234")
	assert.Equal(t, "Internal Server Error", (*err).Message)
}

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}
func getMockDevice() (*dynamodb.PutItemInput, domain.Device) {
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
	return dynamoDbItem, mockDevice
}
