package service

import (
	"context"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
)

// NewCartItemService returns a service for CartItem manipulations over the passed CartItemKeeper.
func NewCartItemService(cik repository.CartItemKeeper) *CartItemService {
	return &CartItemService{cik: cik}
}

// CartItemService handles CartItem operations over the CartItemKeeper.
type CartItemService struct {
	cik repository.CartItemKeeper
}

// ErrNotMatchCartID fires when cart item ID does not match correct CartID.
type ErrNotMatchCartID struct{}

// Error returns ErrNotMatchCartID error description.
func (e ErrNotMatchCartID) Error() string {
	return "Item does not match Cart ID"
}

// CreateCartItem creates a CartItem in CartItemKeeper based on passed properties.
func (cis CartItemService) CreateCartItem(ctx context.Context, cartID string, product string, quantity int) (*model.CartItem, error) {
	// TODO: Put validator here

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
		return &ErrNotMatchCartID{}
	}

	err = cis.cik.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
