package controller

import (
	"github.com/sebastianbordi/DataloggerDemo/model"
)

type SensorController struct {
}

func (SensorController) GetAll() (*[]model.Sensor, error) {

}

func (SensorController) GetById(id int) (*model.Sensor, error) {

}

func (SensorController) Update(*model.Sensor) (*model.Sensor, error) {

}

func (SensorController) Delete(id int) (*model.Sensor, error) {

}
