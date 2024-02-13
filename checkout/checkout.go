package checkout

import (
	"checkout-system/pricing"
	"fmt"
)

// ICheckout is the interface for the supermarket checkout system
type ICheckout interface {
	Scan(item Product)
	GetTotalPrice() int
}

type Product string

// Checkout represents the supermarket checkout system
type Checkout struct {
	pricingRules map[Product]pricing.PricingRule
	items        map[Product]int
}

// NewCheckout initializes a new checkout instance with pricing rules
func NewCheckout(pricingRules map[Product]pricing.PricingRule) *Checkout {
	return &Checkout{
		pricingRules: pricingRules,
		items:        make(map[Product]int),
	}
}

// Scan adds an item to the checkout
func (c *Checkout) Scan(item Product) {
	c.items[item]++
}

// GetTotalPrice calculates the total price of all scanned items
func (c *Checkout) GetTotalPrice() int {
	totalPrice := 0

	for item, count := range c.items {
		if rule, exists := c.pricingRules[item]; exists {
			totalPrice += c.calculateItemPrice(rule, count)
		} else {
			fmt.Printf("Warning: No pricing rule found for item %s\n", item)
		}
	}

	return totalPrice
}

// calculateItemPrice calculates the total price for a specific item based on the pricing rule
func (c *Checkout) calculateItemPrice(rule pricing.PricingRule, count int) int {
	totalPrice := 0

	// Apply special price if applicable
	if rule.SpecialAmount > 0 && count >= rule.SpecialAmount {
		totalPrice += (count / rule.SpecialAmount) * rule.SpecialPrice
		count %= rule.SpecialAmount
	}

	// Calculate remaining price for individual items
	totalPrice += count * rule.UnitPrice

	return totalPrice
}
