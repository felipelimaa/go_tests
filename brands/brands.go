package brands

import "strings"

// BrandName receive a car name and return brand name
func BrandName(carName string) string {
	validBrands := []string{"Chevrolet", "Volkswagen", "Fiat"}

	brandCar := strings.Split(carName, " ")[0]

	hasValidBrand := false

	for _, brand := range validBrands {
		if strings.ToLower(brand) == strings.ToLower(brandCar) {
			hasValidBrand = true
		}
	}

	if hasValidBrand {
		return strings.Title(brandCar)
	}

	return "Invalid Brand."
}
