package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/service"
)

type FakeCartItemKeeper struct {
	WithError       bool
	WithDeleteError bool
}

func (f FakeCartItemKeeper) Create(ctx context.Context, c *model.CartItem) (*model.CartItem, error) {
	if f.WithError {
		return nil, fmt.Errorf("Fake CartItem Create Error")
	}

	return c, nil
}

func (f FakeCartItemKeeper) All(ctx context.Context) ([]*model.CartItem, error) {
	return nil, nil
}

func (f FakeCartItemKeeper) ByID(ctx context.Context, id string) (*model.CartItem, error) {
	if f.WithError {
		return nil, fmt.Errorf("Fake CartItem ByID Error")
	}

	return &model.CartItem{ID: "**mongoID**", CartID: "**mongoCartID**"}, nil
}

func (f FakeCartItemKeeper) UpdateByID(ctx context.Context, id string, c *model.CartItem) error {
	return nil
}

func (f FakeCartItemKeeper) DeleteByID(ctx context.Context, id string) error {
	if f.WithDeleteError {
		return fmt.Errorf("Fake CartItem DeleteByID Error")
	}

	return nil
}

func (f FakeCartItemKeeper) ByCartID(ctx context.Context, id string) ([]model.CartItem, error) {
	if f.WithError {
		return nil, fmt.Errorf("Fake CartItem ByCartID Error")
	}

	return []model.CartItem{}, nil
}

func TestCreateCartItem(t *testing.T) {
	c := model.CartItem{
		CartID:   "**mongoID**",
		Product:  "Product",
		Quantity: 1,
	}

	ctx := context.TODO()
	cik := &FakeCartItemKeeper{WithError: false}

	svc := service.NewCartItemService(cik)

	_, err := svc.CreateCartItem(ctx, c.CartID, c.Product, c.Quantity)
	if err != nil {
		t.Errorf("CreateCartItem(): %v", err)
		return
	}

	cik.WithError = true

	_, err = svc.CreateCartItem(ctx, c.CartID, c.Product, c.Quantity)
	if err == nil {
		t.Error("CreateCartItem(): Should return an error")
	}
}

func TestRemoveCartItem(t *testing.T) {
	c := model.CartItem{
		ID:       "**mongoID**",
		CartID:   "**mongoCartID**",
		Product:  "Product",
		Quantity: 1,
	}

	ctx := context.TODO()
	cik := &FakeCartItemKeeper{WithError: false, WithDeleteError: false}

	svc := service.NewCartItemService(cik)

	err := svc.RemoveCartItem(ctx, c.CartID, c.ID)
	if err != nil {
		t.Errorf("RemoveCartItem(): %v", err)
		return
	}

	cik.WithError = true

	err = svc.RemoveCartItem(ctx, c.CartID, c.ID)
	if err == nil {
		t.Error("CreateCartItem(): Should return an error")
		return
	}

	cik.WithError = false
	cik.WithDeleteError = true

	err = svc.RemoveCartItem(ctx, c.CartID, c.ID)
	if err == nil {
		t.Error("CreateCartItem(): Should return an error")
	}
}
