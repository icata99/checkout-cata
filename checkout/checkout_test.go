package checkout

import (
	"checkout-system/pricing"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckout_GetTotalPrice(t *testing.T) {
	// Define pricing rules
	pricingRules := map[Product]pricing.PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}

	type Test struct {
		name     string
		scanned  map[Product]int
		expected int
	}

	tests := []Test{
		{
			name: "3A",
			scanned: map[Product]int{
				"A": 3,
			},
			expected: 130,
		},
		{
			name: "7A",
			scanned: map[Product]int{
				"A": 7,
			},
			expected: 310,
		},
		{
			name: "2B",
			scanned: map[Product]int{
				"B": 2,
			},
			expected: 45,
		},
		{
			name: "4B",
			scanned: map[Product]int{
				"B": 4,
			},
			expected: 90,
		},
		{
			name: "2A 2B",
			scanned: map[Product]int{
				"A": 2,
				"B": 2,
			},
			expected: 145,
		},
		{
			name: "2A 2B 2C 2D",
			scanned: map[Product]int{
				"A": 2,
				"C": 2,
				"B": 2,
				"D": 2,
			},
			expected: 215,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkout := NewCheckout(pricingRules)

			for item, counts := range test.scanned {
				for i := 0; i < counts; i++ {
					checkout.Scan(item)
				}
			}

			require.Equal(t, test.expected, checkout.GetTotalPrice())
		})
	}
}
