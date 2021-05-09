package view

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/controller"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

func CreateMeasurement(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetMeasurementController()
	var measurementDto model.MeasurementPostDto
	err := json.NewDecoder(r.Body).Decode(&measurementDto)
	if err != nil {
		log.Println(err)
		basicResponse(&w, 400, "can't decode body request")
		return
	}
	measurement, err := controller.CreateFromPostDTO(&measurementDto)

	if err != nil {
		if err.Error() == "bad password" || err.Error() == "mac not found" {
			basicResponse(&w, 401, "error identifying sensor")
			return
		}
		log.Println(err)
		basicResponse(&w, 400, "can't decode body request")
		return
	}
	json.NewEncoder(w).Encode(measurement)
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
	controller := controller.GetMeasurementController()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		basicResponse(&w, 400, fmt.Sprintf("Invalid id %s", params["id"]))
		return
	}
	measurement, err := controller.Delete(id)

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
