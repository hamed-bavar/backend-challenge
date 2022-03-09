package controller

import (
	"challenge/domain"
	service "challenge/serivce"
	"challenge/utils"
	"encoding/json"
	"net/http"
)

type DeviceController struct {
	Service service.DeviceService
}

func (dc *DeviceController) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device domain.Device
	//validate request
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response, appError := dc.Service.CreateDevice(&device)
	if appError != nil {
		utils.WriteResponse(w, appError.Code, appError.Message)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, response)
}
