package controller

import (
	"errors"

	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

type measurementController struct {
	dataContext database.IContext
}

var measurementControllerInstance *measurementController

func GetMeasurementController() *measurementController {
	if measurementControllerInstance == nil {
		measurementControllerInstance = &measurementController{}
	}
	return measurementControllerInstance
}

func (measurementController) InitMeasurementController(context database.IContext) {
	controller := GetMeasurementController()
	controller.dataContext = context
}

func (controller *measurementController) Create(entity *model.Measurement) (*model.Measurement, error) {
	context := controller.dataContext.GetContext()

	err := context.Create(&entity).Error
	return entity, err
}
func (controller *measurementController) GetAll() (*[]model.Measurement, error) {
	context := controller.dataContext.GetContext()

	var result []model.Measurement
	err := context.Joins("sensors").Find(&result).Error
	return &result, err
}
func (controller *measurementController) GetById(id int) (*model.Measurement, error) {
	context := controller.dataContext.GetContext()

	var result model.Measurement
	err := context.Joins("sensors").First(&result, id).Error
	return &result, err
}
func (controller *measurementController) Update(entity *model.Measurement) (*model.Measurement, error) {

	return nil, errors.New("the temeperature registers couldn't be modified")
}
func (controller *measurementController) Delete(id int) (*model.Measurement, error) {
	context := controller.dataContext.GetContext()

	entityBase, err := controller.GetById(id)
	if err != nil {
		return nil, err
	}

	err = context.Delete(entityBase, id).Error

	return entityBase, err
}
