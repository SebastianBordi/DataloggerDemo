package controller

import (
	"errors"

	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

type temperatureController struct {
	dataContext database.IContext
}

var temperatureControllerInstance *temperatureController

func GetTemperatureController() *temperatureController {
	if temperatureControllerInstance == nil {
		temperatureControllerInstance = &temperatureController{}
	}
	return temperatureControllerInstance
}

func (temperatureController) InitTemperatureController(context database.IContext) {
	controller := GetTemperatureController()
	controller.dataContext = context
}
func (controller *temperatureController) GetAll() (*[]model.Temperature, error) {
	context := controller.dataContext.GetContext()

	var result []model.Temperature
	err := context.Find(&result).Error
	return &result, err
}

func (controller *temperatureController) GetById(id int) (*model.Temperature, error) {
	context := controller.dataContext.GetContext()

	var result model.Temperature
	err := context.First(&result, id).Error
	return &result, err
}

func (controller *temperatureController) Update(entity *model.Temperature) (*model.Temperature, error) {

	return nil, errors.New("the temeperature registers couldn't be modified")
}

func (controller *temperatureController) Delete(id int) (*model.Temperature, error) {
	context := controller.dataContext.GetContext()

	entityBase, err := controller.GetById(id)
	if err != nil {
		return nil, err
	}

	err = context.Delete(entityBase, id).Error

	return entityBase, err
}
