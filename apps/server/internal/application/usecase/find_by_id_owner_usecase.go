package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/owner"

	"github.com/google/uuid"
)

type FindByIDOwner struct {
	repository owner.Repository
}

func NewFindByIDOwner(repository owner.Repository) *FindByIDOwner {
	return &FindByIDOwner{repository}
}

func (fo *FindByIDOwner) Execute(ctx context.Context, ownerID, managerID uuid.UUID) (*response.OwnerDTO, *httperr.HttpError) {
	owner, err := fo.repository.FindByID(ctx, ownerID, managerID)
	if err != nil {
		return nil, err
	}

	ownerDTO := response.OwnerDTO{
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
	}

	return &ownerDTO, nil
}
