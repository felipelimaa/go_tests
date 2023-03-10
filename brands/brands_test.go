package brands_test

import (
	"testing"
	. "tests/brands"
)

type testScenario struct {
	carName       string
	expectedBrand string
}

func TestBrandName(t *testing.T) {
	scenarios := []testScenario{
		{"Chevrolet Tracker", "Chevrolet"},
		{"Fiat Toro", "Fiat"},
		{"Volkswagen Amarok", "Volkswagen"},
		{"Toyota Hilux", "Invalid Brand."},
	}

	for _, scenario := range scenarios {
		responseBrand := BrandName(scenario.carName)

		if responseBrand != scenario.expectedBrand {
			t.Errorf("Was expected %s but was returned %s.", scenario.expectedBrand, responseBrand)
		}
	}
}
