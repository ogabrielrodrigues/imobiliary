package plan

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type PlanKind string

const (
	PlanKindFree PlanKind = "free"
	PlanKindPro  PlanKind = "pro"

	PlanFreePrice float32 = 0.00
	PlanProPrice  float32 = 15.99

	ERR_PLAN_ALREADY_UPGRADED   = "plan already upgraded"
	ERR_PLAN_ALREADY_DOWNGRADED = "plan already downgraded"
	ERR_PLAN_CANNOT_DOWNGRADE   = "cannot downgrade your plan if you have more than 30 properties"
)

type Plan struct {
	ID                       uuid.UUID
	Kind                     PlanKind
	Price                    float32
	PropertiesTotalQuota     int
	PropertiesUsedQuota      int
	PropertiesRemainingQuota int
}

func (p *Plan) Upgrade() *response.Err {
	if p.Kind == PlanKindPro {
		return response.NewErr(http.StatusBadRequest, ERR_PLAN_ALREADY_UPGRADED)
	}

	p.Kind = PlanKindPro
	p.Price = PlanProPrice
	p.PropertiesTotalQuota = -1     // unlimited
	p.PropertiesRemainingQuota = -1 // unlimited

	return nil
}

func (p *Plan) Downgrade() *response.Err {
	if p.Kind == PlanKindFree {
		return response.NewErr(http.StatusBadRequest, ERR_PLAN_ALREADY_DOWNGRADED)
	}

	if p.PropertiesUsedQuota > 30 {
		return response.NewErr(http.StatusBadRequest, ERR_PLAN_CANNOT_DOWNGRADE)
	}

	p.Kind = PlanKindFree
	p.Price = PlanFreePrice
	p.PropertiesTotalQuota = 30
	p.PropertiesRemainingQuota = (p.PropertiesTotalQuota - p.PropertiesUsedQuota)

	return nil
}

func New(kind PlanKind, propertiesTotalQuota, propertiesUsedQuota, propertiesRemainingQuota int) *Plan {
	var price float32

	switch kind {
	case PlanKindFree:
		price = PlanFreePrice
	case PlanKindPro:
		price = PlanProPrice
	default:
		price = PlanFreePrice
	}

	return &Plan{
		ID:                       uuid.New(),
		Kind:                     kind,
		Price:                    price,
		PropertiesTotalQuota:     propertiesTotalQuota,
		PropertiesUsedQuota:      propertiesUsedQuota,
		PropertiesRemainingQuota: propertiesRemainingQuota,
	}
}

func (p *Plan) AddProperty() {
	if p.PropertiesUsedQuota == p.PropertiesTotalQuota {
		return
	}

	p.PropertiesUsedQuota++
	p.PropertiesRemainingQuota--
}

func (p *Plan) ToDTO() *DTO {
	return &DTO{
		ID:                       p.ID,
		Kind:                     p.Kind,
		Price:                    p.Price,
		PropertiesTotalQuota:     p.PropertiesTotalQuota,
		PropertiesUsedQuota:      p.PropertiesUsedQuota,
		PropertiesRemainingQuota: p.PropertiesRemainingQuota,
	}
}
