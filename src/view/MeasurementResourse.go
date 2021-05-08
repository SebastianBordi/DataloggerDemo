package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/controller"
)

func CreateMeasurement(w http.ResponseWriter, r *http.Request) {

}

func GetMeasurements(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetMeasurementController()
	measurements, err := controller.GetAll()

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No measurement were found")
		} else {
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurements)
}

func GetMeasurementById(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetMeasurementController()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		basicResponse(&w, 400, fmt.Sprintf("Invalid id %s", params["id"]))
		return
	}

	measurement, err := controller.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No measurement were found")
		} else {
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}

// func UpdateMeasurement(w http.ResponseWriter, r *http.Request) {

// }

func DeleteMeasurement(w http.ResponseWriter, r *http.Request) {

}
