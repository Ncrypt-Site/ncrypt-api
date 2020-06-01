package helpers

import "ncrypt-api/models"

//BuildResponse is a helper to build response in a defined format
func BuildResponse(c int, m string, d interface{}, e []string, meta interface{}) models.Response {
	return models.Response{
		Code:    c,
		Message: m,
		Data:    d,
		Error:   e,
		Meta:    meta,
	}
}
