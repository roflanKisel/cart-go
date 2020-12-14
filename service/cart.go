package service

import (
	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
)

// CartService handles Cart operations
type CartService struct {
	Cir repository.CartItemRepository
	R   repository.CartRepository
}

// CreateEmpty will create a cart without items
func (cs *CartService) CreateEmpty() (*model.Cart, error) {
	c, err := cs.R.Create(&model.Cart{Items: []model.CartItem{}})
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetCartByID will return a cart that matches given id
func (cs *CartService) GetCartByID(id string) (*model.Cart, error) {
	c, err := cs.R.FindByID(id)
	if err != nil {
		return nil, err
	}

	items, err := cs.Cir.FindByCartID(c.ID)
	if err != nil {
		return nil, err
	}

	c.Items = items
	return c, nil
}
