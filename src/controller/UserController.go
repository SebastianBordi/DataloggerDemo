package controller

import (
	"github.com/sebastianbordi/DataloggerDemo/database"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

type userController struct {
	dataContext database.IContext
}

var userControllerInstance *userController

func GetUserController() *userController {
	if userControllerInstance == nil {
		userControllerInstance = &userController{}
	}
	return userControllerInstance
}

func (userController) InitUserController(context database.IContext) {
	controller := GetUserController()
	controller.dataContext = context
}

func (controller *userController) GetUserByName(name string) (*model.User, error) {
	var user model.User
	dbContext := controller.dataContext.GetContext()
	err := dbContext.Where("user = ?", name).First(&user).Error
	return &user, err
}
