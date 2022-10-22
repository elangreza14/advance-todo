package config

import (
	"errors"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
)

type (
	// IValidator is validator interface
	IValidator interface {
		apply(*Configuration) error
		Struct(interface{}) []error
	}

	iValidator struct{}
)

func newValidator() IValidator {
	return &iValidator{}
}

// WithValidator is validator
func WithValidator() Option {
	return newValidator()
}

func (iv *iValidator) apply(conf *Configuration) error {
	conf.Validator = iv
	return nil
}

func (iv *iValidator) Struct(data interface{}) []error {
	validate := validator.New()
	trans, err := iv.registerTranslation(validate)
	if err != nil {
		return []error{err}
	}

	errs := []error{}
	if err = validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, errors.New(err.Translate(*trans)))
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (iv *iValidator) registerTranslation(validate *validator.Validate) (*ut.Translator, error) {
	uni := ut.New(en.New(), en.New())
	trans, _ := uni.GetTranslator("en")

	if err := validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, iv.translateFunc("required")); err != nil {
		return nil, err
	}

	if err := validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must have a valid email!", true)
	}, iv.translateFunc("email")); err != nil {
		return nil, err
	}

	if err := validate.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} must greeter than {1}!", true)
	}, iv.translateFunc("gte")); err != nil {
		return nil, err
	}

	return &trans, nil
}

func (iv *iValidator) translateFunc(types string) validator.TranslationFunc {
	switch types {
	case "gte":
		return func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(types, fe.Field(), fe.Param())
			return t
		}
	default: // handle email and required
		return func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(types, fe.Field())
			return t
		}
	}
}
