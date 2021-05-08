package router

import (
	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/view"
)

func GetRouter(baseURL string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(baseURL+"/sensor", view.CreateSensor).Methods("POST")
	router.HandleFunc(baseURL+"/sensor", view.GetSensors).Methods("GET")
	router.HandleFunc(baseURL+"/sensor/{id}", view.GetSensorById).Methods("GET")
	router.HandleFunc(baseURL+"/sensor", view.UpdateSensor).Methods("PUT")
	router.HandleFunc(baseURL+"/sensor/{id}", view.DeleteSensor).Methods("DELETE")

	router.HandleFunc(baseURL+"/temperature", view.CreateTemperature).Methods("POST")
	router.HandleFunc(baseURL+"/temperature", view.GetTemperatures).Methods("GET")
	router.HandleFunc(baseURL+"/temperature/{id}", view.GetTemperatureById).Methods("GET")
	//router.HandleFunc(baseURL+"/temperature", view.UpdateTemperature).Methods("PUT")
	router.HandleFunc(baseURL+"/temperature/{id}", view.DeleteTemperature).Methods("DELETE")

	return router
}
