package controller

import (
	"encoding/json"
	"fmt"
	"backend/initialize" 
	"backend/service"

	validator "gopkg.in/go-playground/validator.v9"
)
// ErrorResponse get error
func ErrorResponse(err error) service.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := initialize.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := initialize.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return service.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return service.ParamErr("JSON types don't match", err)
	}

	return service.ParamErr("param error", err)
}
