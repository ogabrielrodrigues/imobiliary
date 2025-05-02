package property

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) Create(ctx context.Context, dto *property.CreateDTO) *response.Err {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	p, err := property.New(dto.Status, dto.Kind, dto.WaterID, dto.EnergyID, uuid.MustParse(user_id), dto.Address.ToAddress())
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, p)
}
