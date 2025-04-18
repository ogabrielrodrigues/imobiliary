package test

import (
	"testing"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
)

func TestPlan(t *testing.T) {
	t.Run("should be able to create a free plan", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindFree, 0, 0, 0)

		if new_plan.Price != plan.PlanFreePrice {
			t.Errorf("expected price: %f, got %f", plan.PlanFreePrice, new_plan.Price)
		}
	})

	t.Run("should be able to create a free plan", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindPro, -1, 0, -1)

		if new_plan.Price != plan.PlanProPrice {
			t.Errorf("expected price: %f, got %f", plan.PlanProPrice, new_plan.Price)
		}
	})

	t.Run("should be able to update free plan to pro", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindFree, 0, 0, 0)

		err := new_plan.Upgrade()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should not be able to update free plan to pro if already pro", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindPro, -1, 0, -1)

		err := new_plan.Upgrade()
		if err == nil {
			t.Errorf("expected error: %s, got %s", plan.ERR_PLAN_ALREADY_UPGRADED, err.Error())
		}
	})

	t.Run("should be able to downgrade pro plan to free", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindPro, -1, 0, -1)

		err := new_plan.Downgrade()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should not be able to downgrade pro plan to free if already free", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindFree, 0, 0, 0)

		err := new_plan.Downgrade()
		if err == nil {
			t.Errorf("expected error: %s, got %s", plan.ERR_PLAN_ALREADY_DOWNGRADED, err.Error())
		}
	})

	t.Run("should not be able to downgrade pro plan to free if have more than 30 properties", func(t *testing.T) {
		new_plan := plan.New(plan.PlanKindPro, -1, 31, -1)

		err := new_plan.Downgrade()
		if err == nil {
			t.Errorf("expected error: %s, got %s", plan.ERR_PLAN_CANNOT_DOWNGRADE, err.Error())
		}
	})
}
