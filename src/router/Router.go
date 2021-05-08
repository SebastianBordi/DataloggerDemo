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

	router.HandleFunc(baseURL+"/measure", view.CreateMeasurement).Methods("POST")
	router.HandleFunc(baseURL+"/measure", view.GetMeasurements).Methods("GET")
	router.HandleFunc(baseURL+"/measure/{id}", view.GetMeasurementById).Methods("GET")
	//router.HandleFunc(baseURL+"/temperature", view.UpdateMeasurement).Methods("PUT")
	router.HandleFunc(baseURL+"/measure/{id}", view.DeleteMeasurement).Methods("DELETE")

	return router
}
