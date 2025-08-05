package helpers

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	if v == nil || v.validator == nil {
		return nil
	}

	return v.validator.Struct(i)
}
