package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/repository"
	"github.com/roflanKisel/cart-go/service"

	"github.com/stretchr/testify/assert"
)

var (
	mockCart = model.Cart{}
)

type FakeCartKeeper struct {
	WithCreateError bool
	WithByIDError   bool
}

func (f FakeCartKeeper) Create(ctx context.Context, c *model.Cart) (*model.Cart, error) {
	if f.WithCreateError {
		return nil, fmt.Errorf("Mock Create error")
	}

	return &mockCart, nil
}

func (f FakeCartKeeper) All(ctx context.Context) ([]*model.Cart, error) {
	return nil, nil
}

func (f FakeCartKeeper) ByID(ctx context.Context, id string) (*model.Cart, error) {
	if f.WithByIDError {
		return nil, fmt.Errorf("Mock ByID error")
	}

	return &mockCart, nil
}

func (f FakeCartKeeper) UpdateByID(ctx context.Context, id string, c *model.Cart) error {
	return nil
}

func (f FakeCartKeeper) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func TestCreateEmptyCart(t *testing.T) {
	ctx := context.TODO()

	services := []struct {
		name     string
		ck       repository.CartKeeper
		cik      repository.CartItemKeeper
		expected *model.Cart
	}{
		{"Without errors", FakeCartKeeper{}, FakeCartItemKeeper{}, &mockCart},
		{"With Create error", FakeCartKeeper{WithCreateError: true}, FakeCartItemKeeper{}, nil},
	}

	for _, svc := range services {
		t.Run(svc.name, func(t *testing.T) {
			s := service.NewCartService(svc.ck, svc.cik)

			c, err := s.CreateEmpty(ctx)
			if err != nil {
				assert.Nil(t, svc.expected)
				return
			}

			assert.Equal(t, *svc.expected, *c, "should be equal")
		})
	}
}

func TestCartByID(t *testing.T) {
	id := "**MockObjectID**"
	ctx := context.TODO()

	services := []struct {
		name     string
		ck       repository.CartKeeper
		cik      repository.CartItemKeeper
		expected *model.Cart
	}{
		{"Without errors", FakeCartKeeper{}, FakeCartItemKeeper{}, &mockCart},
		{"With ByID error", FakeCartKeeper{WithByIDError: true}, FakeCartItemKeeper{}, nil},
		{"With CartItem error", FakeCartKeeper{}, FakeCartItemKeeper{WithError: true}, nil},
	}

	for _, svc := range services {
		t.Run(svc.name, func(t *testing.T) {
			s := service.NewCartService(svc.ck, svc.cik)

			c, err := s.CartByID(ctx, id)
			if err != nil {
				assert.Nil(t, svc.expected, "should return an error")
				return
			}

			assert.Equal(t, *svc.expected, *c, "should be equal")
		})
	}
}
