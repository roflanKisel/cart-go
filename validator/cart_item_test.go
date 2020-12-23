package validator_test

import (
	"testing"

	"github.com/roflanKisel/cart-go/validator"

	"github.com/stretchr/testify/assert"
)

func TestValidateQuantity(t *testing.T) {
	v := validator.NewCartItemValidator()

	quantities := []struct {
		name     string
		val      int
		expected bool
	}{
		{"Negative quantity", -1, false},
		{"Zero quantity", 0, false},
		{"Positive quantity", 1, true},
	}

	for _, q := range quantities {
		t.Run(q.name, func(t *testing.T) {
			assert.Equal(t, q.expected, v.ValidateQuantity(q.val))
		})
	}
}

func TestValidateProduct(t *testing.T) {
	v := validator.NewCartItemValidator()

	products := []struct {
		name     string
		val      string
		expected bool
	}{
		{"Empty product name", "", false},
		{"Non-empty product name", "Test", true},
	}

	for _, p := range products {
		t.Run(p.name, func(t *testing.T) {
			assert.Equal(t, p.expected, v.ValidateProduct(p.val))
		})
	}
}
