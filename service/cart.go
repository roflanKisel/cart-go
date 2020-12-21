package service

import (
	"context"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
)

// NewCartService returns a service for Cart manipulations over the passed CartKeeper.
func NewCartService(ck repository.CartKeeper, cik repository.CartItemKeeper) *CartService {
	return &CartService{ck: ck, cik: cik}
}

// CartService handles Cart operations over the CartKeeper.
type CartService struct {
	ck  repository.CartKeeper
	cik repository.CartItemKeeper
}

// CreateEmpty creates a Cart in CartKeeper without any items.
func (cs CartService) CreateEmpty(ctx context.Context) (*model.Cart, error) {
	c, err := cs.ck.Create(ctx, &model.Cart{Items: []model.CartItem{}})
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CartByID returns a Cart that matches passed id with corresponding cart items.
func (cs CartService) CartByID(ctx context.Context, id string) (*model.Cart, error) {
	c, err := cs.ck.ByID(ctx, id)
	if err != nil {
		return nil, err
	}

	items, err := cs.cik.ByCartID(ctx, c.ID)
	if err != nil {
		return nil, err
	}

	c.Items = items
	return c, nil
}
