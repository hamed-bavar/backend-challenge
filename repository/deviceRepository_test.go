package repository

import (
	"challenge/domain"
	"challenge/lib/errors"
	"challenge/mocks"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-playground/assert/v2"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateDeviceOnSuccess(t *testing.T) {
	//Arrange
	mockDevice := getMockDevice()
	mockDB := mocks.DynamoDBAPI{}
	//we pass same device that we got from client to the client.so it does not matter what the inputs and outputs are
	mockDB.On("PutItemWithContext", mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)
	repo := NewDeviceRepositoryDb(dynamo.NewFromIface(&mockDB))

	//Assert
	createdDevice, _ := repo.Create(&mockDevice)

	//Act
	assert.Equal(t, *createdDevice, mockDevice)
	mockDB.AssertExpectations(t)
}
func TestCreateDeviceOnServerError(t *testing.T) {
	//Arrange
	mockDevice := getMockDevice()
	mockDB := mocks.DynamoDBAPI{}
	e := errorObject{}
	expectedError := errors.InternalServerError("Internal Server Error")
	//this error happens if we have a server error.so we expect the error to be returned.but we does not care what the error is
	mockDB.On("PutItemWithContext", mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{}, e)
	repo := NewDeviceRepositoryDb(dynamo.NewFromIface(&mockDB))

	//Act
	_, err := repo.Create(&mockDevice)

	//Assertion
	assert.Equal(t, expectedError, *err)
	mockDB.AssertExpectations(t)
}

func getMockDevice() domain.Device {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "nice car",
		Serial:      "8765432",
	}
	return mockDevice
}
func TestGetDeviceByIdOnSuccess(t *testing.T) {
	//Arrange
	dynamodbOutPut := getDynamodbOutput()
	mockDevice := getMockDevice()
	mockDB := mocks.DynamoDBAPI{}
	mockDB.On("GetItemWithContext", mock.Anything, &dynamodb.GetItemInput{
		TableName: aws.String("Device"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1234"),
			},
		},
	}).Return(dynamodbOutPut, nil)
	repo := NewDeviceRepositoryDb(dynamo.NewFromIface(&mockDB))

	//Act
	actualDevice, _ := repo.FindById("1234")
	fmt.Println(actualDevice)

	//Assert
	assert.Equal(t, mockDevice, *actualDevice)
}

func TestGetDeviceByIdOnNotFound(t *testing.T) {
	//Arrange
	mockDB := mocks.DynamoDBAPI{}
	expectedError := errors.NotFoundError("Device not found")
	e := errorObject{}
	mockDB.On("GetItemWithContext", mock.Anything, &dynamodb.GetItemInput{
		TableName: aws.String("Device"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("12"),
			},
		},
	}).Return(&dynamodb.GetItemOutput{}, e)
	repo := NewDeviceRepositoryDb(dynamo.NewFromIface(&mockDB))

	//Act
	_, err := repo.FindById("12")

	//Assert
	assert.Equal(t, expectedError.Message, err.Message)
}

type error interface {
	Error() string
}
type errorObject struct {
}

func (e errorObject) Error() string {
	return "This is an error"
}
func getDynamodbOutput() *dynamodb.GetItemOutput {
	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1234"),
			},
			"DeviceModel": {
				S: aws.String("mercedes"),
			},
			"Name": {
				S: aws.String("your car"),
			},
			"Note": {
				S: aws.String("nice car"),
			},
			"Serial": {
				S: aws.String("8765432"),
			},
		},
	}
}
