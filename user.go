package todo

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id        string    `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UpdateUserInput struct {
	Name     *string `json:"name"`
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required"`
}

func (u *User) ValidateUserInput() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 25)),
		validation.Field(&u.Username, validation.Required, validation.Length(2, 25),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]*$")).Error(
				"only alphanumeric characters and underscore allowed")),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 100),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]*$")).Error(
				"only alphanumeric characters and underscore allowed")),
	)
}

func (u *UpdateUserInput) ValidateUserUpdate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 25)),
		validation.Field(&u.Username, validation.Required, validation.Length(2, 25),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]*$")).Error(
				"only alphanumeric characters and underscore allowed")),
	)
}
