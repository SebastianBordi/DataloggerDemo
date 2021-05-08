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

func CreateSensor(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetSensorController()
	var sensor model.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		log.Println(err)
		basicResponse(&w, 400, "error decoding the body")
		return
	}
	newEntity, err := controller.Create(&sensor)

	if err != nil {
		log.Println(err)
		basicResponse(&w, 500, "internal server error")
		return
	}
	json.NewEncoder(w).Encode(newEntity)
}

func GetSensors(w http.ResponseWriter, r *http.Request) {
	controller := controller.GetSensorController()
	sensors, err := controller.GetAll()

	if err != nil {
		if err.Error() == "record not found" {
			basicResponse(&w, 404, "No sensor were found")
		} else {
			log.Println(err)
			basicResponse(&w, 500, "Internal server error")
		}
		return
	}
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
	controller := controller.GetSensorController()
	var sensor model.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		log.Println(err)
		basicResponse(&w, 400, "error decoding the body")
		return
	}
	newEntity, err := controller.Update(&sensor)

	if err != nil {
		log.Println(err)
		basicResponse(&w, 500, "internal server error")
		return
	}
	json.NewEncoder(w).Encode(newEntity)
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
	json.NewEncoder(w).Encode(sensor)
}
