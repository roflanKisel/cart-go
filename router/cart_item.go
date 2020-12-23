package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roflanKisel/cart-go/model"
	"github.com/roflanKisel/cart-go/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewCartItemRouter returns router for CartItem which uses passed services.
func NewCartItemRouter(cSvc *service.CartService, ciSvc *service.CartItemService) *CartItemRouter {
	return &CartItemRouter{cSvc: cSvc, ciSvc: ciSvc}
}

// CartItemRouter provides possibility to register HTTP handlers for CartItem model.
type CartItemRouter struct {
	cSvc  *service.CartService
	ciSvc *service.CartItemService
}

// RegisterCartItemHandlers registers HTTP handlers for CartItem model.
func (cir CartItemRouter) RegisterCartItemHandlers(r *mux.Router) {
	r.HandleFunc("/carts/{cartID}/items", cir.createCartItem).Methods("POST")
	r.HandleFunc("/carts/{cartID}/items/{id}", cir.deleteCartItem).Methods("DELETE")
}

func (cir CartItemRouter) createCartItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	cartID, ok := vars["cartID"]
	if !ok {
		fmt.Println("id not found in URL")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cartItemBody model.CartItem
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cartItemBody); err != nil {
		_, err = w.Write([]byte(fmt.Sprintf("Invalid request payload: %s", err.Error())))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	_, err := cir.cSvc.CartByID(ctx, cartID)
	if err != nil {
		if err == mongo.ErrNilDocument {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = w.Write([]byte(fmt.Sprintf("An error ocurred when getting cart by ID: %s", err.Error())))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ci, err := cir.ciSvc.CreateCartItem(ctx, cartID, cartItemBody.Product, cartItemBody.Quantity)
	if err != nil {
		_, err = w.Write([]byte(fmt.Sprintf("An error ocurred when creating cart item: %s", err.Error())))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(ci)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cir CartItemRouter) deleteCartItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	cartID, ok := vars["cartID"]
	if !ok {
		fmt.Println("cartID not found in URL")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, ok := vars["id"]
	if !ok {
		fmt.Println("id not found in URL")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := cir.ciSvc.RemoveCartItem(ctx, cartID, id)
	if err != nil {
		if err == mongo.ErrNilDocument || err == service.ErrNotMatchCartID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = w.Write([]byte(fmt.Sprintf("An error ocurred when removing cart item: %s", err.Error())))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(struct{}{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
