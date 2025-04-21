package plan

import "github.com/google/uuid"

type DTO struct {
	ID                       uuid.UUID `json:"id"`
	Kind                     PlanKind  `json:"kind"`
	Price                    float32   `json:"price"`
	PropertiesTotalQuota     int       `json:"properties_total_quota"`
	PropertiesUsedQuota      int       `json:"properties_used_quota"`
	PropertiesRemainingQuota int       `json:"properties_remaining_quota"`
}

func (p *DTO) ToPlan() *Plan {
	return &Plan{
		ID:                       p.ID,
		Kind:                     p.Kind,
		Price:                    p.Price,
		PropertiesTotalQuota:     p.PropertiesTotalQuota,
		PropertiesUsedQuota:      p.PropertiesUsedQuota,
		PropertiesRemainingQuota: p.PropertiesRemainingQuota,
	}
}
