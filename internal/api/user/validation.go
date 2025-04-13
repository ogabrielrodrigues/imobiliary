package user

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	ERR_FULLNAME_EMPTY    string = "user fullname is empty"
	ERR_FULLNAME_INVALID  string = "user fullname must be between 10 and 100 characters"
	ERR_EMAIL_EMPTY       string = "user email is empty"
	ERR_EMAIL_INVALID     string = "user email is invalid format"
	ERR_CRECIID_EMPTY     string = "user creci_id is empty"
	ERR_CRECIID_INVALID   string = "user creci_id is invalid format"
	ERR_CELLPHONE_EMPTY   string = "user cellphone is empty"
	ERR_CELLPHONE_INVALID string = "user cellphone is invalid format"
	ERR_PASSWORD_EMPTY    string = "user password is empty"
	ERR_PASSWORD_INVALID  string = "user password is invalid must be at least 8 characters"
)

func (u *User) validate() error {
	if u.Fullname == "" {
		return errors.New(ERR_FULLNAME_EMPTY)
	}

	if len(u.Fullname) < 10 || len(u.Fullname) > 100 {
		return errors.New(ERR_FULLNAME_INVALID)
	}

	if u.Email == "" {
		return errors.New(ERR_EMAIL_EMPTY)
	}

	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, u.Email); !match {
		return errors.New(ERR_EMAIL_INVALID)
	}

	if u.CreciID == "" {
		return errors.New(ERR_CRECIID_EMPTY)
	}

	if len(u.CreciID) != 7 {
		return errors.New(ERR_CRECIID_INVALID)
	}

	if u.Cellphone == "" {
		return errors.New(ERR_CELLPHONE_EMPTY)
	}

	if match, _ := regexp.MatchString(`^\(\d{2}\)\s\d{4,5}-\d{4}$`, u.Cellphone); !match {
		fmt.Println(u.Cellphone)
		fmt.Println(match)
		return errors.New(ERR_CELLPHONE_INVALID)
	}

	if u.password == "" {
		return errors.New(ERR_PASSWORD_EMPTY)
	}

	if len(u.password) < 8 {
		return errors.New(ERR_PASSWORD_INVALID)
	}

	return nil
}
