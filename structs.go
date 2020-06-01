package main

import "github.com/go-playground/validator/v10"

//Validator custom validator for echo
type Validator struct {
	Validator *validator.Validate
}
