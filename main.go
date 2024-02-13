package main

import (
	"checkout-system/checkout"
	"checkout-system/pricing"
	"fmt"
)

func main() {
	// Define pricing rules
	pricingRules := map[checkout.Product]pricing.PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}

	// Create a new checkout instance
	var checkoutInterface checkout.ICheckout = checkout.NewCheckout(pricingRules)

	// Scan items
	checkoutInterface.Scan("A")
	checkoutInterface.Scan("B")
	checkoutInterface.Scan("C")
	checkoutInterface.Scan("B")

	// Get the total price
	totalPrice := checkoutInterface.GetTotalPrice()

	// Print the total price
	fmt.Printf("Total Price: %d\n", totalPrice)
}
