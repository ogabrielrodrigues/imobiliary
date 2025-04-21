package provider

import "github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"

type InMemoryPlanRepository struct {
	plans map[string]*plan.Plan
}

func NewInMemoryPlanRepository() *InMemoryPlanRepository {
	return &InMemoryPlanRepository{
		plans: make(map[string]*plan.Plan),
	}
}
