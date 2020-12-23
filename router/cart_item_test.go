package router_test

import (
	"context"
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

type FakeCartItemKeeper struct {
	ByIDNotFoundError bool
	DocID             string
	CartDocID         string
}

func (f FakeCartItemKeeper) Create(ctx context.Context, c *model.CartItem) (*model.CartItem, error) {
	return &model.CartItem{}, nil
}

func (f FakeCartItemKeeper) All(ctx context.Context) ([]*model.CartItem, error) {
	return []*model.CartItem{}, nil
}

func (f FakeCartItemKeeper) ByID(ctx context.Context, id string) (*model.CartItem, error) {
	if f.ByIDNotFoundError {
		return nil, mongo.ErrNoDocuments
	}

	return &model.CartItem{ID: f.DocID, CartID: f.CartDocID}, nil
}

func (f FakeCartItemKeeper) UpdateByID(ctx context.Context, id string, c *model.CartItem) error {
	return nil
}

func (f FakeCartItemKeeper) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func (f FakeCartItemKeeper) ByCartID(ctx context.Context, id string) ([]model.CartItem, error) {
	return []model.CartItem{}, nil
}

func TestCartItemCreate(t *testing.T) {
	id := "5fd78878266f92cf3670b9a5"

	routers := []struct {
		name         string
		body         model.CartItem
		cSvc         *service.CartService
		ciSvc        *service.CartItemService
		expectedCode int
	}{
		{
			"Success",
			model.CartItem{},
			service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{}),
			service.NewCartItemService(FakeCartItemKeeper{}),
			http.StatusOK,
		},
		{
			"Cart Not Found",
			model.CartItem{},
			service.NewCartService(FakeCartKeeper{NotFoundError: true}, FakeCartItemKeeper{}),
			service.NewCartItemService(FakeCartItemKeeper{}),
			http.StatusNotFound,
		},
	}

	for _, r := range routers {
		t.Run(r.name, func(t *testing.T) {
			mr := mux.NewRouter()
			cr := router.NewCartItemRouter(r.cSvc, r.ciSvc)
			cr.RegisterCartItemHandlers(mr)

			server := httptest.NewServer(mr)
			defer server.Close()

			e := httpexpect.New(t, server.URL)

			e.POST("/carts/" + id + "/items").WithJSON(r.body).Expect().Status(r.expectedCode)
		})
	}
}

func TestCartItemDelete(t *testing.T) {
	id := "5fd78878266f92cf3670b9a5"
	cartID := "5fd78878266f92cf3670b9a6"

	routers := []struct {
		name         string
		cSvc         *service.CartService
		ciSvc        *service.CartItemService
		expectedCode int
	}{
		{
			"Success",
			service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{DocID: id, CartDocID: cartID}),
			service.NewCartItemService(FakeCartItemKeeper{DocID: id, CartDocID: cartID}),
			http.StatusOK,
		},
		{
			"Cart Item Not Found",
			service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{ByIDNotFoundError: true}),
			service.NewCartItemService(FakeCartItemKeeper{ByIDNotFoundError: true}),
			http.StatusNotFound,
		},
		{
			"Cart Item does not match Cart ID",
			service.NewCartService(FakeCartKeeper{}, FakeCartItemKeeper{DocID: id, CartDocID: "not_correct"}),
			service.NewCartItemService(FakeCartItemKeeper{DocID: id, CartDocID: "not_correct"}),
			http.StatusNotFound,
		},
	}

	for _, r := range routers {
		t.Run(r.name, func(t *testing.T) {
			mr := mux.NewRouter()
			cr := router.NewCartItemRouter(r.cSvc, r.ciSvc)
			cr.RegisterCartItemHandlers(mr)

			server := httptest.NewServer(mr)
			defer server.Close()

			e := httpexpect.New(t, server.URL)

			e.DELETE("/carts/" + cartID + "/items/" + id).Expect().Status(r.expectedCode)
		})
	}
}
