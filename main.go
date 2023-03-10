package main

import (
	"fmt"
	"tests/addresses"
	"tests/brands"
)

func main() {
	addressType := addresses.AddressType("Rua Paracati")
	fmt.Println(addressType)

	brandName := brands.BrandName("Chevrolet Tracker")
	fmt.Println(brandName)
}
