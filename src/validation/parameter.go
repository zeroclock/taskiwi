package validation

import (
	"github.com/go-playground/validator/v10"
)

type (
	TagsToSearch struct {
		Tags []string `json:"tags" validate:"required"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
