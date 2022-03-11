package controller

import (
	"challenge/domain"
	"challenge/lib/logger"
	service "challenge/serivce"
	"challenge/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type DeviceController struct {
	Service service.DeviceService
}

func (dc *DeviceController) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device domain.Device
	//validate request
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		logger.Error("Error while decoding posted data" + err.Error())
		utils.WriteResponse(w, http.StatusBadRequest, "invalid fields")
		return
	}
	response, appError := dc.Service.CreateDevice(&device)
	if appError != nil {
		utils.WriteResponse(w, appError.Code, appError.Message)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, response)
}
func (dc *DeviceController) GetDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	response, appError := dc.Service.GetDevice(id)
	if appError != nil {
		utils.WriteResponse(w, appError.Code, appError.Message)
		return
	}
	utils.WriteResponse(w, http.StatusOK, response)
}
