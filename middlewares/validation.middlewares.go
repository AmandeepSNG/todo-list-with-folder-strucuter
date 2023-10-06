package validationMiddleware

import (
	"regexp"
	"strings"
	helperFunctions "todolist-app/common/helper"
	enums "todolist-app/common/types"

	"todolist-app/modals"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateUserValidations(c *fiber.Ctx) error {
	var requestBody *modals.User
	var errorArray []*enums.ValidationError

	if err := c.BodyParser(&requestBody); err != nil {
		if err.Error() == "Unprocessable Entity" {
			return helperFunctions.GenerateValidationErrorResponse(c, fiber.Map{
				"message": "All fields are required",
			})
		}
		return helperFunctions.GenerateValidationErrorResponse(c, err)
	}
	validate := validator.New()

	validationErrors := validate.Struct(requestBody)

	if validationErrors != nil {
		for _, err := range validationErrors.(validator.ValidationErrors) {
			var singleError enums.ValidationError
			singleError.Field = strings.Split(err.StructNamespace(), ".")[1]
			singleError.Tag = err.Tag()
			errorArray = append(errorArray, &singleError)
		}
	}

	if len(errorArray) != 0 {
		return helperFunctions.GenerateValidationErrorResponse(c, errorArray)
	}
	return c.Next()
}

func ValidateEmail(field validator.FieldLevel) bool {
	email := field.Field().String()
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func ValidateMobileNumber(field validator.FieldLevel) bool {
	mobileNumber := field.Field().String()
	regex := `^[6789]\d{9}$`
	// match, _ := regexp.MatchString(regex, mobileNumber)
	return regexp.MustCompile(regex).MatchString(mobileNumber)
}
