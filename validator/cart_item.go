package validator

// CartItemValidator is a validator for CartItem model.
type CartItemValidator struct{}

// NewCartItemValidator returns a CartItemValidator instance.
func NewCartItemValidator() *CartItemValidator {
	return &CartItemValidator{}
}

// ValidateQuantity validates quantity of CartItem
func (v CartItemValidator) ValidateQuantity(q int) bool {
	return q > 0
}

// ValidateProduct validates product name of cart item
func (v CartItemValidator) ValidateProduct(s string) bool {
	return len(s) > 0
}
