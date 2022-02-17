package validate

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Get() *validator.Validate {
	return validate
}
