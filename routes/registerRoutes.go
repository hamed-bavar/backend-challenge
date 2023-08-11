package routes

import (
	"challenge/controller"
	"challenge/repository"
	"challenge/service"
	"challenge/utils"

	"github.com/gorilla/mux"
)

func Register(router *mux.Router) {
	//create db instanc
	dbClient, _ := utils.GetDbClient()
	deviceRepo := repository.NewDeviceRepositoryDb(dbClient)
	deviceController := controller.DeviceController{Service: service.NewDeviceService(deviceRepo)}

	router.HandleFunc("/devices", deviceController.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", deviceController.GetDevice).Methods("GET")
}
