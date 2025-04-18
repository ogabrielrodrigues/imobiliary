package test

import (
	"path/filepath"
	"testing"

	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
)

func TestUser(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	t.Run("should be able to create a user", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err != nil {
			t.Errorf("expected an nil, got error %v", err)
		}
	})

	t.Run("should not be able to create a user with empty fullname", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_FULLNAME_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_EMPTY, err)
		}
	})

	t.Run("should not be able to create a user if fullname less than 10 characters", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_INVALID, err)
		}
	})

	t.Run("should not be able to create a user if fullname more than 100 characters", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe ofSilva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_FULLNAME_INVALID, err)
		}
	})

	t.Run("should not be able to create a user with empty email", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"",
			"password",
		)

		if err.Message != user.ERR_EMAIL_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_EMAIL_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_EMAIL_EMPTY, err)
		}
	})

	t.Run("should not be able to create a user with invalid email", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe.com",
			"password",
		)

		if err.Message != user.ERR_EMAIL_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_EMAIL_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_EMAIL_INVALID, err)
		}
	})

	t.Run("should not be able to create a user with empty creci_id", func(t *testing.T) {
		_, err := user.New(
			"",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_CRECIID_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CRECIID_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CRECIID_EMPTY, err)
		}
	})

	t.Run("should not be able to create a user with invalid creci_id", func(t *testing.T) {
		_, err := user.New(
			"1234F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_CRECIID_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CRECIID_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CRECIID_INVALID, err)
		}
	})

	t.Run("should not be able to create a user with empty cellphone", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_CELLPHONE_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CELLPHONE_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CELLPHONE_EMPTY, err)
		}
	})

	t.Run("should not be able to create a user with invalid cellphone", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"01 1236789",
			"johndoe@example.com",
			"password",
		)

		if err.Message != user.ERR_CELLPHONE_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CELLPHONE_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_CELLPHONE_INVALID, err)
		}
	})

	t.Run("should not be able to create a user with empty password", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"",
		)

		if err.Message != user.ERR_PASSWORD_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_PASSWORD_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_PASSWORD_EMPTY, err)
		}
	})

	t.Run("should not be able to create a user with invalid password", func(t *testing.T) {
		_, err := user.New(
			"12345-F",
			"John Doe of Silva",
			"(01) 12345-6789",
			"johndoe@example.com",
			"pwd",
		)

		if err.Message != user.ERR_PASSWORD_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_PASSWORD_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", user.ERR_PASSWORD_INVALID, err)
		}
	})
}
