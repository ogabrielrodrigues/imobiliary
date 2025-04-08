package entity

import (
	"errors"
	"regexp"
)

const (
	ERR_FULLNAME_EMPTY   string = "user fullname is empty"
	ERR_FULLNAME_INVALID string = "user fullname is invalid"
	ERR_EMAIL_EMPTY      string = "user email is empty"
	ERR_EMAIL_INVALID    string = "user email is invalid"
)

func (u *user) validate() error {
	if u.fullname == "" {
		return errors.New(ERR_FULLNAME_EMPTY)
	}

	if len(u.fullname) < 10 || len(u.fullname) > 100 {
		return errors.New(ERR_FULLNAME_INVALID)
	}

	if u.email == "" {
		return errors.New(ERR_EMAIL_EMPTY)
	}

	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, u.email); !match {
		return errors.New(ERR_EMAIL_INVALID)
	}

	return nil
}
