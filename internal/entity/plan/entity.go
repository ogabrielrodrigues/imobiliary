package plan

import "github.com/google/uuid"

type PlanKind string

const (
	PlanKindFree PlanKind = "free"
	PlanKindPro  PlanKind = "pro"
)

type Plan struct {
	ID                       uuid.UUID
	Kind                     PlanKind
	Price                    float64
	PropertiesTotalQuota     int
	PropertiesUsedQuota      int
	PropertiesRemainingQuota int
}

func New(kind PlanKind, price float64, propertiesTotalQuota, propertiesUsedQuota, propertiesRemainingQuota int) *Plan {
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
