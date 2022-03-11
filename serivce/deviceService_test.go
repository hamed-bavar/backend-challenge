package service

import (
	"challenge/domain"
	"challenge/lib/errors"
	"challenge/mocks/mockDeviceRepository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeviceService_CreateDevice(t *testing.T) {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "nice car",
		Serial:      "8765432",
	}
	tests := []struct {
		name   string
		err    *errors.AppError
		status int
		input  domain.Device
		output domain.Device
	}{
		{name: "validation error with status 400", input: domain.Device{
			Id:          "1234",
			DeviceModel: "mercedes",
			Name:        "your car",
		}, status: 400, err: errors.ValidationError("some fields are missing")},
		{name: "create device correctly", status: 201, input: mockDevice, output: mockDevice},
		{name: "server error while creating device", status: 500, input: mockDevice, err: errors.InternalServerError("Internal Server Error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			var service DefaultDeviceService

			if test.status != 400 {
				mockRepo := mockDeviceRepository.NewMockDeviceRepository(ctrl)
				mockRepo.EXPECT().Create(&test.input).Return(&test.output, test.err)
				service = NewDeviceService(mockRepo)
			} else {
				service = NewDeviceService(nil)
			}

			createdDevice, err := service.CreateDevice(&test.input)
			if err != nil {
				require.Equal(t, *err, *test.err)
			}
			if createdDevice != nil {
				require.Equal(t, *createdDevice, test.output)
			}
		})
	}
}
func TestDeviceService_GetDevice(t *testing.T) {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "nice car",
		Serial:      "8765432",
	}
	tests := []struct {
		name   string
		err    *errors.AppError
		status int
		input  string
		output domain.Device
	}{
		{name: "get device with status 200", input: "1234", status: 400, output: mockDevice},
		{name: "device not found with status 404", status: 404, input: "1234", err: errors.NotFoundError("Device not found")},
		{name: "server error while getting device", status: 500, input: "1234", err: errors.InternalServerError("Internal Server Error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mockDeviceRepository.NewMockDeviceRepository(ctrl)
			mockRepo.EXPECT().FindById(test.input).Return(&test.output, test.err)
			mockService := NewDeviceService(mockRepo)

			createdDevice, err := mockService.GetDevice(test.input)
			if err != nil {
				require.Equal(t, *err, *test.err)
			}
			if createdDevice != nil {
				require.Equal(t, *createdDevice, test.output)
			}
		})
	}
}
