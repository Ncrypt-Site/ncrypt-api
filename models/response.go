package models

type Response struct {
	Code    int
	Message string
	Data    interface{}
	Error   []string
	Meta    interface{}
}
