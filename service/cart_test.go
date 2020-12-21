package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/service"
)

type FakeCartKeeper struct {
	WithError bool
}

func (f FakeCartKeeper) Create(ctx context.Context, c *model.Cart) (*model.Cart, error) {
	if f.WithError {
		return nil, fmt.Errorf("Fake Cart Create Error")
	}

	return &model.Cart{}, nil
}

func (f FakeCartKeeper) All(ctx context.Context) ([]*model.Cart, error) {
	return nil, nil
}

func (f FakeCartKeeper) ByID(ctx context.Context, id string) (*model.Cart, error) {
	if f.WithError {
		return nil, fmt.Errorf("Fake Cart ByID Error")
	}

	return &model.Cart{}, nil
}

func (f FakeCartKeeper) UpdateByID(ctx context.Context, id string, c *model.Cart) error {
	return nil
}

func (f FakeCartKeeper) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func TestCreateEmptyCart(t *testing.T) {
	ctx := context.TODO()

	ck := &FakeCartKeeper{WithError: false}
	cik := &FakeCartItemKeeper{WithError: false}

	svc := service.NewCartService(ck, cik)

	_, err := svc.CreateEmpty(ctx)
	if err != nil {
		t.Errorf("CreateEmpty(): %v", err)
		return
	}

	ck.WithError = true
	_, err = svc.CreateEmpty(ctx)
	if err == nil {
		t.Error("CreateEmpty(): Should return an error")
	}
}

func TestCartByID(t *testing.T) {
	id := "**MockObjectID**"
	ctx := context.TODO()

	ck := &FakeCartKeeper{WithError: false}
	cik := &FakeCartItemKeeper{WithError: false}

	svc := service.NewCartService(ck, cik)

	_, err := svc.CartByID(ctx, id)
	if err != nil {
		t.Errorf("CartByID(): %v", err)
		return
	}

	ck.WithError = true

	_, err = svc.CartByID(ctx, id)
	if err == nil {
		t.Error("CartByID(): Should return an error")
		return
	}

	ck.WithError = false
	cik.WithError = true

	_, err = svc.CartByID(ctx, id)
	if err == nil {
		t.Error("CartByID(): Should return an error")
	}
}
