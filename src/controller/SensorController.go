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
<<<<<<< HEAD
		sensorControllerInstance = &sensorController{}
=======
		sensorControllerInstance = &sensorController{
			dataContext: context,
		}
>>>>>>> 4b4af171392f56cdea87281e99db27059bc23977
	}
	return sensorControllerInstance
}

<<<<<<< HEAD
func (sensorController) InitSensorController(context database.IContext) {
=======
func InitSensorController(context database.IContext) {
>>>>>>> 4b4af171392f56cdea87281e99db27059bc23977
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
