package router

import (
	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/view"
)

func GetRouter(baseURL string) *mux.Router {
	router := mux.NewRouter()

	//TODO: Create static file server for de web page

	//TODO: Create endpoint for web sockets

	router.HandleFunc(baseURL+"/login", view.Login).Methods("POST")
	router.HandleFunc(baseURL+"/sensor", view.JWTAuth(view.CreateSensor)).Methods("POST")
	router.HandleFunc(baseURL+"/sensor", view.GetSensors).Methods("GET")
	router.HandleFunc(baseURL+"/sensor/{id}", view.GetSensorById).Methods("GET")
	router.HandleFunc(baseURL+"/sensor", view.JWTAuth(view.UpdateSensor)).Methods("PUT")
	router.HandleFunc(baseURL+"/sensor/{id}", view.JWTAuth(view.DeleteSensor)).Methods("DELETE")

	router.HandleFunc(baseURL+"/measurement", view.CreateMeasurement).Methods("POST")
	router.HandleFunc(baseURL+"/measurement", view.GetMeasurements).Methods("GET")
	router.HandleFunc(baseURL+"/measurement/{id}", view.GetMeasurementById).Methods("GET")
	//router.HandleFunc(baseURL+"/temperature", view.UpdateMeasurement).Methods("PUT")
	router.HandleFunc(baseURL+"/measurement/{id}", view.JWTAuth(view.DeleteMeasurement)).Methods("DELETE")

	return router
}
