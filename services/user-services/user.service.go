package userService

import (
	helperFunctions "todolist-app/common/helper"
	constants "todolist-app/common/shared-constants"
	enums "todolist-app/common/types"
	"todolist-app/modals"
	userRepository "todolist-app/repository/user-repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserList() enums.Response {
	response := userRepository.GetUserList()

	if response.Error != nil {
		return helperFunctions.Response(nil, response.Status, response.Message, response.Error)
	}
	return response
}

func GetUserById(userId string) enums.Response {
	response := userRepository.GetUserById(userId)

	if response.Error != nil {
		if response.Status == constants.NOT_FOUND {
			response.Error = nil
			return response
		}
		return response
	}

	return response
}

func CreateUser(requestBody *fiber.Ctx) enums.Response {
	user := modals.User{
		Id:     primitive.NewObjectID(),
		UserId: (uuid.New()).String(),
	}
	if err := requestBody.BodyParser(&user); err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}

	response := userRepository.CreateUser(user)
	if response.Error != nil {
		return response
	}
	return helperFunctions.Response(&user, response.Status, response.Message, nil)
}

func UpdateUser(fiberContext *fiber.Ctx, userId string) enums.Response {
	user := modals.User{}
	// Need to implement uniqueness for email
	checkUser := userRepository.GetUserById(userId)
	if checkUser.Error != nil {
		if checkUser.Status == constants.NOT_FOUND {
			checkUser.Error = nil
			return checkUser
		}
		return checkUser
	}
	// Parse the request body into the user struct
	if err := fiberContext.BodyParser(&user); err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}

	// Perform the update operation by calling your userRepository.UpdateUser function
	updatedUserResponse := userRepository.UpdateUser(userId, user)
	if updatedUserResponse.Error != nil {
		return helperFunctions.Response(nil, updatedUserResponse.Status, updatedUserResponse.Message, updatedUserResponse.Error)
	}

	return updatedUserResponse
}

func DeleteUser(userId string) enums.Response {
	response := userRepository.DeleteUser(userId)
	return response
}
