package provider

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
)

type InMemoryUserRepository struct {
	users []*user.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}
