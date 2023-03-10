package addresses_test

import (
	"testing"
	. "tests/addresses"
)

type testScenario struct {
	address      string
	expectedType string
}

func TestAddressType(t *testing.T) {
	scenarios := []testScenario{
		{"Rua Paracati", "Rua"},
		{"Avenida Mira Mangue", "Avenida"},
		{"Estrada da Sé", "Estrada"},
		{"Rodovia BR-101", "Rodovia"},
		{"Praça das Flores", "Tipo inválido"},
		{"RUA AGRESTINA", "Rua"},
		{"AVENIDA DOS XAVANTES", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, scenario := range scenarios {
		responseType := AddressType(scenario.address)

		if responseType != scenario.expectedType {
			t.Errorf("The response Type is different from expected Type. Was expected %s but was returned %s.",
				scenario.expectedType,
				responseType,
			)
		}
	}
}
