package service

import (
	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
)

// CartItemService handles CartItem operations
type CartItemService struct {
	R repository.CartItemRepository
}

// ErrNotMatchCartID will be fired when cart item ID does not match correct CartID
type ErrNotMatchCartID struct{}

func (e *ErrNotMatchCartID) Error() string {
	return "Item does not match Cart ID"
}

// CreateCartItem will create Cart
func (cis *CartItemService) CreateCartItem(cartID string, product string, quantity int) (*model.CartItem, error) {
	// TODO: Put validator here

	cartItem := &model.CartItem{
		CartID:   cartID,
		Product:  product,
		Quantity: quantity,
	}

	ci, err := cis.R.Create(cartItem)
	if err != nil {
		return nil, err
	}

	return ci, nil
}

// RemoveCartItem will remove cart with given ID
func (cis *CartItemService) RemoveCartItem(cartID string, id string) error {
	ci, err := cis.R.FindByID(id)
	if err != nil {
		return err
	}

	if ci.CartID != cartID {
		return &ErrNotMatchCartID{}
	}

	err = cis.R.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
