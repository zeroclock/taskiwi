package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	AggregateCondition struct {
		Tags []string `json:"tags" validate:"required"`
		Start string `json:"start" validate:"required,is_date"`
		End string `json:"end" validate:"required,is_date"`
	}

	TaskByDateCondition struct {
		Date string `json:"date" validate:"required,is_date"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func DateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}
