package provider

import (
	"path/filepath"
)

type InMemoryAvatarRepository struct {
	path             string
	default_filename string
}

func NewInMemoryAvatarRepository(path string) *InMemoryAvatarRepository {
	return &InMemoryAvatarRepository{
		path:             filepath.Join(path),
		default_filename: "avatar.png",
	}
}
