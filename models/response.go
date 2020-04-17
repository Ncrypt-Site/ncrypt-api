package models

type Response struct {
	Code    int
	Message string
	Data    interface{}
	Error   error
	Meta    interface{}
}
