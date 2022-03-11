package controller

import (
	"bytes"
	"challenge/domain"
	"challenge/lib/errors"
	service "challenge/mocks"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDeviceController(t *testing.T) {
	mockDevice := domain.Device{
		Id:          "1234",
		DeviceModel: "mercedes",
		Name:        "your car",
		Note:        "this is your car",
		Serial:      "123456789",
	}
	tests := []struct {
		name   string
		err    *errors.AppError
		input  domain.Device
		status int
		output interface{}
	}{
		{name: "create device with status 201", input: mockDevice, status: http.StatusCreated, output: mockDevice},
		{name: "invalid data with status 400", input: domain.Device{Id: "1"}, status: 400, err: errors.ValidationError("some fields are missing"), output: nil},
		{name: "server error with status 500", input: mockDevice, status: 500, err: errors.InternalServerError("internal server error"), output: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := service.NewMockDeviceService(ctrl)
			dc := DeviceController{
				Service: mockService,
			}
			dummyDevice := test.input
			marshaledDevice, _ := json.Marshal(dummyDevice)
			mockService.EXPECT().CreateDevice(&dummyDevice).Return(&dummyDevice, test.err)
			router := mux.NewRouter()
			router.HandleFunc("/devices", dc.CreateDevice)
			request, _ := http.NewRequest(http.MethodGet, "/devices", bytes.NewReader(marshaledDevice))
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			assert.Equal(t, test.status, recorder.Code)
			if test.status == http.StatusCreated {
				marshaledOutput, _ := json.Marshal(test.output)
				require.JSONEq(t, string(marshaledOutput), recorder.Body.String())
			}
		})
	}
}

//func Test_should_return_error_when_device_not_found(t *testing.T) {
//	// Given
//	device := &models.Device{
//		ID: "1",
//	}
//	deviceRepository := &mocks.DeviceRepository{}
//	deviceRepository.On("FindByID", device.ID).Return(nil, errors.New("error"))
//	deviceController := NewDeviceController(deviceRepository)
//
//	// When
//	err := deviceController.Delete(device)
//
//	// Then
//	assert.NotNil(t, err)
//	assert.Equal(t, "error", err.Error())
//}
