package routes

import (
	"challenge/controller"
	"challenge/repository"
	"challenge/serivce"
	"challenge/utils"
	"github.com/gorilla/mux"
)

func Register(router *mux.Router) {
	//create db instance
	dbClient, _ := utils.GetDbClient()
	//create device repo
	deviceRepo := repository.NewDeviceRepositoryDb(dbClient)
	deviceController := controller.DeviceController{Service: service.NewDeviceService(deviceRepo)}
	router.HandleFunc("/device", deviceController.CreateDevice).Methods("POST")
}
