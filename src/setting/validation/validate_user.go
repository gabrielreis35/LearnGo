package validation

import (
	"LearnGo/src/setting/rest_err"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type", nil)
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestError("Invalid field value", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields", nil)
	}
}
