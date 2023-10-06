package userControllers

import (
	helperFunctions "todolist-app/common/helper"
	"todolist-app/modals"
	userService "todolist-app/services/user-services"

	"github.com/gofiber/fiber/v2"
)

func GetUserList(c *fiber.Ctx) error {
	listOfUsers := []modals.User{}
	response := userService.GetUserList()
	if response.Error != nil {
		return helperFunctions.GenerateResponse(c, response.Status, response.Message, nil)
	}
	if response.Data != nil {
		return helperFunctions.GenerateResponse(c, response.Status, response.Message, response.Data)
	}
	return helperFunctions.GenerateResponse(c, response.Status, response.Message, listOfUsers)
}

func GetUserById(c *fiber.Ctx) error {
	response := userService.GetUserById(c.Params("userId"))
	if response.Error != nil {
		return helperFunctions.GenerateResponse(c, response.Status, response.Message, response.Data)
	}
	return helperFunctions.GenerateResponse(c, response.Status, response.Message, response.Data)
}

func CreateUser(c *fiber.Ctx) error {

	response := userService.CreateUser(c)
	if response.Error != nil {
		return helperFunctions.GenerateResponse(c, response.Status, response.Message, nil)
	}
	return helperFunctions.GenerateResponse(c, response.Status, response.Message, response.Data)
	// return nil
}

func UpdateUser(c *fiber.Ctx) error {
	updatedUserResponse := userService.UpdateUser(c, c.Params("userId"))

	if updatedUserResponse.Error != nil {
		return helperFunctions.GenerateResponse(c, updatedUserResponse.Status, updatedUserResponse.Message, nil)
	}
	return helperFunctions.GenerateResponse(c, updatedUserResponse.Status, updatedUserResponse.Message, updatedUserResponse.Data)
}

func DeleteUser(c *fiber.Ctx) error {
	response := userService.DeleteUser(c.Params("userId"))
	return helperFunctions.GenerateResponse(c, response.Status, response.Message, nil)
}
