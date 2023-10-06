package enums

type Collection string
type Response struct {
	Status  int
	Message string
	Data    interface{}
	Error   error
}

type ValidationError struct {
	Field string
	Tag   string
}

// type ResponseBody struct {
// 	Status  int
// 	Message string
// 	Data    interface{}
// }

// type Response struct {
// 	Message string
// 	Data    interface{}
// }
