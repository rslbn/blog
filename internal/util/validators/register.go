package validators

import (
	"fmt"
	"log"
	"regexp"
	"unicode"

	customError "github.com/rslbn/blog/internal/errors"
	"github.com/rslbn/blog/internal/model"
)

// TODO: remove hardcoded message

type registerErrors customError.FieldErrors

const (
	// minimal pattern: [username]@[domain].com
	EmailPattern    = `^([a-zA-Z0-9_.]+@[a-z]+\.[com]+)$`
	PasswordPattern = `^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*()_+{}\[\]\\\|;:"',.<>/?])(?=\S+$).{8,70}$`
)

func ValidateRegisterRequest(r *model.RegisterRequest) error {
	errs := make(registerErrors)
	errs.validateNotBlank(r)
	errs.validatePattern("email", r.Email, EmailPattern)
	errs.validatePasswordPattern(r.Password)
	errs.validateSize(8, 70, "password", r.Password)
	if len(errs) == 0 {
		return nil
	}
	return (customError.FieldErrors)(errs)
}

func (errs registerErrors) validateNotBlank(r *model.RegisterRequest) {
	if r.Username == "" {
		errs["username"] = "username is required!"
	}
	if r.Email == "" {
		errs["email"] = "email is required!"
	}
	if r.Password == "" {
		errs["password"] = "password is required!"
	}
	if r.ConfirmationPassword == "" {
		errs["confirmation_password"] = "confirmation password is required"
	}
}

func (errs registerErrors) validatePattern(fieldname, value, regex string) {
	matched, err := regexp.MatchString(regex, value)
	if err != nil {
		log.Printf("Error matching regex for %s: %v", fieldname, err)
		return
	}
	if !matched {
		pattern := ""
		switch fieldname {
		case "password":
			pattern = `password must at least consist of one uppercase letter, one lowercase letter, one digit, and symbols (!@#$%^&*()[]\{}|;':"<>?,./), with min length of 8 chars`
		case "email":
			pattern = `invalid format :[username]@[domain].com`
		}
		errs.put(fieldname, pattern)
	}
}

func (errs registerErrors) validateSize(min, max int, field, value string) {
	if len(value) <= min || len(value) >= max {
		message := fmt.Sprintf("%v must be at least between %v and %v chars long", field, min, max)
		errs.put(field, message)
	}
}

func (errs registerErrors) put(key, value string) {
	v := errs[key]
	if v != "" {
		errs[key] = fmt.Sprintf("%v | %v", v, value)
	} else {
		errs[key] = value
	}
}

func (errs registerErrors) validatePasswordPattern(value string) {
	var (
		hasUpper     bool
		hasLower     bool
		hasNumber    bool
		hasSpecial   bool
		noWhitespace = true
		sizeValid    = len(value) >= 8
	)

ValidatePass:
	for _, c := range value {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			hasSpecial = true
		case c == ' ' || c == '\n' || c == '\t':
			noWhitespace = false
			break ValidatePass
		default:
			hasNumber, hasSpecial, hasUpper, hasLower, noWhitespace = false, false, false, false, false
		}
	}
	valid := hasNumber && hasSpecial && hasLower && hasUpper && noWhitespace && sizeValid
	if !valid {
		errs.put("password", `password must at least consist of one uppercase letter, one lowercase letter, one digit, and symbols (!@#$%^&*()[]\{}|;':"<>?,./), with min length of 8 chars`)
	}
}
