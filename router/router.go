package router

import (
	"github.com/roflanKisel/cart-go/service"

	"github.com/gorilla/mux"
)

// NewRouter returns a router based on passed services.
func NewRouter(cartService *service.CartService, cartItemService *service.CartItemService) *mux.Router {
	r := mux.NewRouter()

	cartRouter := NewCartRouter(cartService)
	cartItemRouter := NewCartItemRouter(cartService, cartItemService)

	cartRouter.RegisterCartHandlers(r)
	cartItemRouter.RegisterCartItemHandlers(r)

	return r
}
