package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/owner"

	"github.com/google/uuid"
)

type FindAllOwner struct {
	repository owner.Repository
}

func NewFindAllOwner(repository owner.Repository) *FindAllOwner {
	return &FindAllOwner{repository}
}

func (fa *FindAllOwner) Execute(ctx context.Context, managerID uuid.UUID) ([]response.OwnerDTO, *httperr.HttpError) {
	owners, err := fa.repository.FindAll(ctx, managerID)
	if err != nil {
		return nil, err
	}

	var ownersDTO []response.OwnerDTO

	for _, owner := range owners {
		ownersDTO = append(ownersDTO, response.OwnerDTO{
			ID:            owner.ID.String(),
			ManagerID:     owner.ManagerID.String(),
			Fullname:      owner.Fullname,
			CPF:           owner.CPF.Value(),
			RG:            owner.RG.Value(),
			Email:         owner.Email.Value(),
			Phone:         owner.Phone.Value(),
			Occupation:    owner.Occupation,
			MaritalStatus: owner.MaritalStatus,
			Address: response.AddressDTO{
				FullAddress:  owner.Address.FullAddress,
				Street:       owner.Address.Street,
				Number:       owner.Address.Number,
				Complement:   owner.Address.Complement,
				Neighborhood: owner.Address.Neighborhood,
				City:         owner.Address.City,
				State:        owner.Address.State,
				ZipCode:      owner.Address.ZipCode,
			},
		})
	}

	return ownersDTO, nil
}
