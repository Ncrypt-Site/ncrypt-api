package models

//Response data model struct
type Response struct {
	Code    int
	Message string
	Data    interface{}
	Error   []string
	Meta    interface{}
}
