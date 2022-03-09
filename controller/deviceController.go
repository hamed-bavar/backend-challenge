package controller

import (
	"challenge/domain"
	service "challenge/serivce"
	"challenge/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type DeviceController struct {
	Service service.DeviceService
}

func (dc *DeviceController) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device domain.Device
	//validate request
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "invalid fields")
		return
	}
	validate := validator.New()
	validationError := validate.Struct(device)
	if validationError != nil {
		utils.WriteResponse(w, http.StatusBadRequest, validationError.Error())
		return
	}
	response, appError := dc.Service.CreateDevice(&device)
	if appError != nil {
		utils.WriteResponse(w, appError.Code, appError.Message)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, response)
}
