package constants

import enums "todolist-app/common/types"

const (
	// ENVIRONMENT VARIABLES
	APP_PORT = 3000
	// Db configurations
	DATABASE_NAME enums.Collection = "user-management"
	MONGO_DB_URL  enums.Collection = "mongodb://localhost:27017"

	// schema names
	USER_SCHEMA_NAME enums.Collection = "users"

	// custom messages
	USER_CREATED_SUCCESSFULLY string = "User created successfully."
	USER_FETCHED_SUCCESSFULLY string = "User Details Fetched successfully."
	USER_UPDATED_SUCCESSFULLY string = "User Updated successfully."
	USER_DELETED_SUCCESSFULLY string = "User Deleted successfully."
	INTERNAL_SERVER_ERROR     string = "Internal server error occurred while processing your request."
	USER_NOT_FOUND            string = "The user you are looking for does not exist."
	FORM_VALIDTION_ERROR      string = "Bad Request."
	// custom status codes
	SUCCESS      int = 200
	CREATED      int = 201
	NOT_FOUND    int = 404
	SERVER_ERROR int = 500
	BAD_REQUEST  int = 400
)
