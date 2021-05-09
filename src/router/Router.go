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

	router.HandleFunc(baseURL+"/measurement", view.CreateMeasurement).Methods("POST")
	router.HandleFunc(baseURL+"/measurement", view.GetMeasurements).Methods("GET")
	router.HandleFunc(baseURL+"/measurement/{id}", view.GetMeasurementById).Methods("GET")
	//router.HandleFunc(baseURL+"/temperature", view.UpdateMeasurement).Methods("PUT")
	router.HandleFunc(baseURL+"/measurement/{id}", view.DeleteMeasurement).Methods("DELETE")

	return router
}
