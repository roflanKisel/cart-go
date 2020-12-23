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
	mockCartItem = model.CartItem{
		ID:       "**mongoID**",
		CartID:   "**mongoCartID**",
		Product:  "Product",
		Quantity: 1,
	}
	errCartItemCreate     = fmt.Errorf("Fake CartItem Create Error")
	errCartItemDeleteByID = fmt.Errorf("Fake CartItem DeleteByID Error")
)

type FakeCartItemKeeper struct {
	WithError       bool
	WithDeleteError bool
}

func (f FakeCartItemKeeper) Create(ctx context.Context, c *model.CartItem) (*model.CartItem, error) {
	if f.WithError {
		return nil, errCartItemCreate
	}

	return &mockCartItem, nil
}

func (f FakeCartItemKeeper) All(ctx context.Context) ([]*model.CartItem, error) {
	return nil, nil
}

func (f FakeCartItemKeeper) ByID(ctx context.Context, id string) (*model.CartItem, error) {
	if f.WithError {
		return nil, &service.ErrNotMatchCartID{}
	}

	return &model.CartItem{ID: "**mongoID**", CartID: "**mongoCartID**"}, nil
}

func (f FakeCartItemKeeper) UpdateByID(ctx context.Context, id string, c *model.CartItem) error {
	return nil
}

func (f FakeCartItemKeeper) DeleteByID(ctx context.Context, id string) error {
	if f.WithDeleteError {
		return errCartItemDeleteByID
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
	ctx := context.TODO()

	services := []struct {
		name     string
		cik      repository.CartItemKeeper
		ci       model.CartItem
		expected *model.CartItem
	}{
		{"Without errors", FakeCartItemKeeper{}, mockCartItem, &mockCartItem},
		{"With error", FakeCartItemKeeper{WithError: true}, mockCartItem, nil},
	}

	for _, svc := range services {
		t.Run(svc.name, func(t *testing.T) {
			s := service.NewCartItemService(svc.cik)

			ci, err := s.CreateCartItem(ctx, svc.ci.CartID, svc.ci.Product, svc.ci.Quantity)
			if err != nil {
				assert.Nil(t, svc.expected)
				return
			}

			assert.Equal(t, *svc.expected, *ci, "should be equal")
		})
	}
}

func TestRemoveCartItem(t *testing.T) {
	ctx := context.TODO()

	services := []struct {
		name     string
		cik      repository.CartItemKeeper
		expected error
	}{
		{"Without errors", FakeCartItemKeeper{}, nil},
		{"CartID not valid", FakeCartItemKeeper{WithError: true}, &service.ErrNotMatchCartID{}},
		{"With Delete error", FakeCartItemKeeper{WithDeleteError: true}, errCartItemDeleteByID},
	}

	for _, svc := range services {
		t.Run(svc.name, func(t *testing.T) {
			s := service.NewCartItemService(svc.cik)

			err := s.RemoveCartItem(ctx, mockCartItem.CartID, mockCartItem.ID)
			if svc.expected != nil {
				assert.EqualError(t, err, svc.expected.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}
