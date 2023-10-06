package main

import (
	"log"
	"os"

	// "strconv"
	// constants "todolist-app/common/shared-constants"
	helperFunctions "todolist-app/common/helper"
	userControllers "todolist-app/controllers/user-controllers"
	validationMiddleware "todolist-app/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	// Load environment variables from a .env file (default method)
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }
	env := os.Getenv("ENV")

	// In case we have multiple files
	currentEnvironment := helperFunctions.GetCurrentEnvironment(env)
	envFile := currentEnvironment + ".env"
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
	// test Route
	app.Get("/test", test)
	// validate := validator.New()
	// validate.RegisterValidation("email", validationMiddleware.ValidateEmail)
	// validate.RegisterValidation("mobileNumber", validationMiddleware.ValidateMobileNumber)
	// API to fetch list of users
	app.Get("/users", userControllers.GetUserList)
	app.Get("/users/:userId", userControllers.GetUserById)
	app.Post("/users", validationMiddleware.CreateUserValidations, userControllers.CreateUser)
	app.Patch("/users/:userId", userControllers.UpdateUser)
	app.Delete("/users/:userId", userControllers.DeleteUser)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}

func test(c *fiber.Ctx) error {
	return c.SendString("Heelo from Test API.")
}
