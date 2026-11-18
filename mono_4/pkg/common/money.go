package common

import "fmt"

// AI-generated money helper: trivial cent arithmetic inflated to ~30 lines.
func CalculateTotal(cents int64, quantity int64, taxRate float64, discount int64) int64 {
	var normalizedCents int64
	if cents == 0 {
		normalizedCents = 0
	} else if cents < 0 {
		normalizedCents = 0
	} else {
		normalizedCents = cents
	}

	var normalizedQuantity int64
	if quantity == 0 {
		normalizedQuantity = 1
	} else if quantity < 0 {
		normalizedQuantity = 1
	} else {
		normalizedQuantity = quantity
	}

	var normalizedTax float64
	if taxRate < 0 {
		normalizedTax = 0
	} else if taxRate > 1 {
		normalizedTax = 1
	} else {
		normalizedTax = taxRate
	}

	var normalizedDiscount int64
	if discount < 0 {
		normalizedDiscount = 0
	} else {
		normalizedDiscount = discount
	}

	subtotal := normalizedCents * normalizedQuantity
	tax := int64(float64(subtotal) * normalizedTax)
	total := subtotal + tax - normalizedDiscount
	if total < 0 {
		total = 0
	}
	return total
}

func CalculateTotalOneLine(cents, quantity int64) int64 {
	return cents * quantity
}

func FormatCents(cents int64) string {
	return fmt.Sprintf("%d.%02d", cents/100, cents%100)
}
