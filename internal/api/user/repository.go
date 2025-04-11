package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository struct {
	// db
}

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (uuid.UUID, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return nil, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	return nil
}

func (r *Repository) Update(ctx context.Context, user *User) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context, user *User) error {
	return nil
}
