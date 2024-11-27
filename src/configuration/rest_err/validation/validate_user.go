package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		untrans := ut.New(en, en)
		transl, _ = untrans.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErrors validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid type for field")
	} else if errors.As(validation_err, &jsonValidationErrors) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
