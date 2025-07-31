package validators

import (
	customError "github.com/rslbn/blog/internal/errors"
	"github.com/rslbn/blog/internal/model"
)

func ValidateLoginRequest(r *model.LoginRequest) error {
	errs := make(customError.FieldErrors)
	if r.Username == "" {
		errs["username"] = "username is required!"
	}
	if r.Password == "" {
		errs["password"] = "password is required!"
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}
