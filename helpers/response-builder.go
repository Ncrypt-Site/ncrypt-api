package helpers

import "ncrypt-api/models"

func BuildResponse(c int, m string, d interface{}, e error, meta interface{}) models.Response {
	return models.Response{
		Code:    c,
		Message: m,
		Data:    d,
		Error:   e,
		Meta:    meta,
	}
}
