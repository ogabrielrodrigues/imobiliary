package plan

import "github.com/google/uuid"

type DTO struct {
	ID                       uuid.UUID `json:"id"`
	Kind                     PlanKind  `json:"kind"`
	Price                    float64   `json:"price"`
	PropertiesTotalQuota     int       `json:"properties_total_quota"`
	PropertiesUsedQuota      int       `json:"properties_used_quota"`
	PropertiesRemainingQuota int       `json:"properties_remaining_quota"`
}
