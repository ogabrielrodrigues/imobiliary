package test

import (
	"testing"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types"
)

func TestOwner(t *testing.T) {
	t.Run("should be able to create a owner", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err != nil {
			t.Errorf("expected an nil, got error %v", err)
		}
	})

	t.Run("should not be able to create a owner with empty fullname", func(t *testing.T) {
		_, err := owner.New(
			"",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_FULLNAME_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_EMPTY, err)
		}
	})

	t.Run("should not be able to create a owner if fullname less than 10 characters", func(t *testing.T) {
		_, err := owner.New(
			"John Doe",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_FULLNAME_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner if fullname more than 100 characters", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of Silva",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_FULLNAME_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_FULLNAME_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner with empty cpf", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_CPF_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CPF_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CPF_EMPTY, err)
		}
	})

	t.Run("should not be able to create a owner if invalid cpf", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"12345676845",
			"123456X",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_CPF_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CPF_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CPF_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner with empty rg", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_RG_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_EMPTY, err)
		}
	})

	t.Run("should not be able to create a owner if rg less than 5 characters", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.768-45",
			"1234",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_RG_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner if rg more than 15 characters", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.768-45",
			"1234323232323232",
			"johndoe@example.com",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_RG_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_RG_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner with empty email", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"123456X",
			"",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_EMAIL_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_EMAIL_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_EMAIL_EMPTY, err)
		}
	})

	t.Run("should not be able to create a owner with invalid email", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"123456X",
			"johndoeexample",
			"(01) 12345-6789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_EMAIL_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_EMAIL_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_EMAIL_INVALID, err)
		}
	})

	t.Run("should not be able to create a owner with empty cellphone", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)
		if err.Message != owner.ERR_CELLPHONE_EMPTY {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CELLPHONE_EMPTY, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CELLPHONE_EMPTY, err)
		}
	})

	t.Run("should not be able to create a owner with invalid cellphone", func(t *testing.T) {
		_, err := owner.New(
			"John Doe of Silva",
			"123.456.789-10",
			"123456X",
			"johndoe@example.com",
			"01123456789",
			"Comerciante",
			types.MaritalStatusSolteiro,
			types.NewAddress("Rua das Flores", "123", "", "Centro", "São Paulo", "SP", "12345678"),
		)

		if err.Message != owner.ERR_CELLPHONE_INVALID {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CELLPHONE_INVALID, err)
		}

		if err == nil {
			t.Errorf("expected an err: %s\ngot error %s", owner.ERR_CELLPHONE_INVALID, err)
		}
	})
}
