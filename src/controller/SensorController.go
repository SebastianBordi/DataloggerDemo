package controller

import (
	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

type sensorController struct {
	dataContext database.IContext
}

var sensorControllerInstance *sensorController

func GetSensorController() *sensorController {
	if sensorControllerInstance == nil {
		sensorControllerInstance = &sensorController{
			dataContext: context,
		}
	}
	return sensorControllerInstance
}

func InitSensorController(context database.IContext) {
	controller := GetSensorController()
	controller.dataContext = context
}

func (controller *sensorController) GetAll() (*[]model.Sensor, error) {
	context := controller.dataContext.GetContext()

	var result []model.Sensor
	err := context.Find(&result).Error
	return &result, err
}

func (controller *sensorController) GetById(id int) (*model.Sensor, error) {
	context := controller.dataContext.GetContext()

	var result model.Sensor
	err := context.First(&result, id).Error
	return &result, err
}

func (controller *sensorController) Update(entity *model.Sensor) (*model.Sensor, error) {
	context := controller.dataContext.GetContext()

	entityBase, err := controller.GetById(entity.IDSensor)
	if err != nil {
		return nil, err
	}
	entityBase.Mac = entity.Mac
	entityBase.Password = entity.Password
	entityBase.Description = entity.Description

	err = context.Save(entityBase).Error
	return entityBase, err
}

func (controller *sensorController) Delete(id int) (*model.Sensor, error) {
	context := controller.dataContext.GetContext()

	entityBase, err := controller.GetById(id)
	if err != nil {
		return nil, err
	}

	err = context.Delete(entityBase, id).Error

	return entityBase, err
}
