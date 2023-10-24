package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserRequest(validation_err error) *exception.Exception {

	var jsonException *json.UnmarshalTypeError
	var jsonExceptionError validator.ValidationErrors

	if errors.As(validation_err, &jsonException) {
		return exception.BadRequestException("Invalid request type")
	}

	if errors.As(validation_err, &jsonExceptionError) {
		errorsCauses := []exception.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := exception.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return exception.BadRequestValidationException("Some fields are invalid", errorsCauses)
	}

	return exception.BadRequestException("Error trying to convert fields")
}
