package service

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ValidationService struct {
	validate *validator.Validate
}

func NewValidationService() (ValidationService, error) {
	validate := validator.New()
	err := validate.RegisterValidation("before_or_today", func(fl validator.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		return !date.IsZero() && !date.After(time.Now())
	})
	if err != nil {
		return ValidationService{}, err
	}
	return ValidationService{
		validate: validate,
	}, nil
}

func (s *ValidationService) Validate(data interface{}) error {
	return s.validate.Struct(data)
}
