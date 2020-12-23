package service

import (
	"context"
	"errors"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
	"github.com/roflanKisel/cart-go/validator"
)

// NewCartItemService returns a service for CartItem manipulations over the passed CartItemKeeper.
func NewCartItemService(cik repository.CartItemKeeper) *CartItemService {
	return &CartItemService{cik: cik, v: *validator.NewCartItemValidator()}
}

// CartItemService handles CartItem operations over the CartItemKeeper.
type CartItemService struct {
	cik repository.CartItemKeeper
	v   validator.CartItemValidator
}

var (
	// ErrNotMatchCartID fires when cart item ID does not match correct CartID.
	ErrNotMatchCartID = errors.New("Item does not match Cart ID")
	// ErrInvalidQuantity fires when cart item quantity is invalid.
	ErrInvalidQuantity = errors.New("Quantity must be positive")
	// ErrInvalidProduct fires when cart item product name is invalid.
	ErrInvalidProduct = errors.New("Product should not be an empty string")
)

// CreateCartItem creates a CartItem in CartItemKeeper based on passed properties.
func (cis CartItemService) CreateCartItem(ctx context.Context, cartID string, product string, quantity int) (*model.CartItem, error) {
	if !cis.v.ValidateQuantity(quantity) {
		return nil, ErrInvalidQuantity
	}

	if !cis.v.ValidateProduct(product) {
		return nil, ErrInvalidProduct
	}

	cartItem := &model.CartItem{
		CartID:   cartID,
		Product:  product,
		Quantity: quantity,
	}

	ci, err := cis.cik.Create(ctx, cartItem)
	if err != nil {
		return nil, err
	}

	return ci, nil
}

// RemoveCartItem removes a CartItem from CartItemKeeper that matches passed properties.
func (cis CartItemService) RemoveCartItem(ctx context.Context, cartID string, id string) error {
	ci, err := cis.cik.ByID(ctx, id)
	if err != nil {
		return err
	}

	if ci.CartID != cartID {
		return ErrNotMatchCartID
	}

	err = cis.cik.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
