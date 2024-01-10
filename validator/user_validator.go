package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required."),
			is.Email.Error("Email is not valid."),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required."),
			validation.RuneLength(6, 30).Error("Limited min 6 characters and max 30 characters."),
		),
	)
}
