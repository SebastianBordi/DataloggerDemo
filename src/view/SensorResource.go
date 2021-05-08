package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/controller"
)

func CreateSensor(w http.ResponseWriter, r *http.Request) {

}

func GetSensors(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetSensorController()
	sensors, err := controller.GetAll()

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No sensor were found")
		} else {
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensors)
}

func GetSensorById(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetSensorController()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		basicResponse(&w, 400, fmt.Sprintf("Invalid id %s", params["id"]))
		return
	}

	sensor, err := controller.GetById(id)

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No sensor were found")
		} else {
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}

func UpdateSensor(w http.ResponseWriter, r *http.Request) {

}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetSensorController()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		basicResponse(&w, 400, fmt.Sprintf("Invalid id %s", params["id"]))
		return
	}
	sensor, err := controller.Delete(id)

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No measurement were found")
		} else {
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}
