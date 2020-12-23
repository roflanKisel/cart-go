package router_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/router"
	"github.com/roflanKisel/cart-go/service"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type FakeCartKeeper struct {
	NotFoundError bool
	CreateError   bool
}

func (f FakeCartKeeper) Create(ctx context.Context, c *model.Cart) (*model.Cart, error) {
	if f.CreateError {
		return nil, fmt.Errorf("Fake Cart create error")
	}

	return &model.Cart{}, nil
}

func (f FakeCartKeeper) All(ctx context.Context) ([]*model.Cart, error) {
	return nil, nil
}

func (f FakeCartKeeper) ByID(ctx context.Context, id string) (*model.Cart, error) {
	if f.NotFoundError {
		return nil, mongo.ErrNoDocuments
	}

	return &model.Cart{}, nil
}

func (f FakeCartKeeper) UpdateByID(ctx context.Context, id string, c *model.Cart) error {
	return nil
}

func (f FakeCartKeeper) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func TestCart(t *testing.T) {
	id := "5fd78878266f92cf3670b9a5"

	routers := []struct {
		name         string
		svc          *service.CartService
		expectedCode int
	}{
		{"Default result", service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{}), http.StatusOK},
		{"Cannot find Cart By ID", service.NewCartService(FakeCartKeeper{NotFoundError: true}, FakeCartItemKeeper{}), http.StatusNotFound},
	}

	for _, r := range routers {
		t.Run(r.name, func(t *testing.T) {
			mr := mux.NewRouter()
			cr := router.NewCartRouter(r.svc)
			cr.RegisterCartHandlers(mr)

			server := httptest.NewServer(mr)
			defer server.Close()

			e := httpexpect.New(t, server.URL)

			e.GET("/carts/" + id).Expect().Status(r.expectedCode)
		})
	}
}

func TestCartCreate(t *testing.T) {
	routers := []struct {
		name         string
		svc          *service.CartService
		expectedCode int
	}{
		{"Empty cart", service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{}), http.StatusOK},
		{"Error creating cart", service.NewCartService(FakeCartKeeper{CreateError: true}, FakeCartItemKeeper{}), http.StatusInternalServerError},
	}

	for _, r := range routers {
		t.Run(r.name, func(t *testing.T) {
			mr := mux.NewRouter()
			cr := router.NewCartRouter(r.svc)
			cr.RegisterCartHandlers(mr)

			server := httptest.NewServer(mr)
			defer server.Close()

			e := httpexpect.New(t, server.URL)

			e.POST("/carts").Expect().Status(r.expectedCode)
		})
	}
}
