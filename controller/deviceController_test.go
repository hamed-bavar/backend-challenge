package controller

import (
	"bytes"
	"challenge/domain"
	"challenge/lib/errors"
	service "challenge/mocks/mockDeviceService"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeviceController_CreateDevice(t *testing.T) {
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
		{name: "invalid data with status 400", input: domain.Device{Id: "1"}, status: http.StatusBadRequest, err: errors.ValidationError("some fields are missing"), output: nil},
		{name: "server error with status 500", input: mockDevice, status: http.StatusInternalServerError, err: errors.InternalServerError("internal server error"), output: nil},
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
			router.HandleFunc("/devices", dc.CreateDevice).Methods("POST")
			request, _ := http.NewRequest(http.MethodPost, "/devices", bytes.NewReader(marshaledDevice))
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

func TestDeviceController_GetDevice(t *testing.T) {
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
		input  string
		status int
		output interface{}
	}{
		{name: "get device with id 123", input: "123", status: http.StatusOK, output: mockDevice},
		{name: "invalid data with status 400", input: "123", status: http.StatusNotFound, err: errors.NotFoundError("Device not found"), output: nil},
		{name: "server error with status 500", input: "123", status: http.StatusInternalServerError, err: errors.InternalServerError("internal server error"), output: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := service.NewMockDeviceService(ctrl)
			dc := DeviceController{
				Service: mockService,
			}
			dummyDevice := mockDevice
			//marshaledDevice, _ := json.Marshal(dummyDevice)
			mockService.EXPECT().GetDevice(test.input).Return(&dummyDevice, test.err)
			router := mux.NewRouter()
			router.HandleFunc("/devices/{id}", dc.GetDevice).Methods("GET")
			request, _ := http.NewRequest(http.MethodGet, "/devices/"+test.input, nil)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			require.Equal(t, test.status, recorder.Code)

			if recorder.Code == http.StatusOK {
				marshaledOutput, _ := json.Marshal(mockDevice)
				require.JSONEq(t, string(marshaledOutput), recorder.Body.String())
			}
		})
	}
}
