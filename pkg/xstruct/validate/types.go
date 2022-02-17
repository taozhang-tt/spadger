package validate

import "github.com/go-playground/validator/v10"

type (
	CustomTypeFunc               = validator.CustomTypeFunc
	FieldError                   = validator.FieldError
	FieldLevel                   = validator.FieldLevel
	FilterFunc                   = validator.FilterFunc
	Func                         = validator.Func
	FuncCtx                      = validator.FuncCtx
	InvalidValidationError       = validator.InvalidValidationError
	RegisterTranslationsFunc     = validator.RegisterTranslationsFunc
	StructLevel                  = validator.StructLevel
	StructLevelFunc              = validator.StructLevelFunc
	TagNameFunc                  = validator.TagNameFunc
	TranslationFunc              = validator.TranslationFunc
	Validate                     = validator.Validate
	ValidationErrors             = validator.ValidationErrors
	ValidationErrorsTranslations = validator.ValidationErrorsTranslations
)
