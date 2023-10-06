package helperFunctions

import (
	constants "todolist-app/common/shared-constants"
	enums "todolist-app/common/types"

	"github.com/gofiber/fiber/v2"
)

/*
Function to return response from every layer except controller
*/
func Response(data interface{}, statusCode int, message string, err error) enums.Response {
	var response enums.Response

	response.Data = data
	response.Status = statusCode
	response.Message = message
	response.Error = err

	return response
}

/*
Function to return the response as API response
*/
func GenerateResponse(fiberContext *fiber.Ctx, statusCode int, responseMessage string, responseData interface{}) error {
	return fiberContext.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": responseMessage,
		"data":    responseData,
	})
}

func GenerateValidationErrorResponse(fiberContext *fiber.Ctx, err interface{}) error {
	return fiberContext.Status(constants.BAD_REQUEST).JSON(fiber.Map{
		"status":  constants.BAD_REQUEST,
		"message": constants.FORM_VALIDTION_ERROR,
		"error":   err,
	})
}

func GetCurrentEnvironment(env string) string {
	switch env {
	case "staging":
		return "staging"
	case "prod":
		return "prod"
	default:
		return "dev"
	}
}
