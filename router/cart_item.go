package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roflanKisel/cart-go/model"

	"github.com/gorilla/mux"
	"github.com/roflanKisel/cart-go/service"
)

// CartItemRouter provides CartItem handlers
type CartItemRouter struct {
	Cis *service.CartItemService
	Cs  *service.CartService
}

// RegisterCartItemHandlers registers HTTP handlers for CartItem model
func (cir *CartItemRouter) RegisterCartItemHandlers(r *mux.Router) {
	r.HandleFunc("/carts/{cartID}/items", cir.createCartItem).Methods("POST")
	r.HandleFunc("/carts/{cartID}/items/{id}", cir.deleteCartItem).Methods("DELETE")
}

func (cir *CartItemRouter) createCartItem(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln("Invalid request payload", err)))
		return
	}

	defer r.Body.Close()

	cart, err := cir.Cs.GetCartByID(cartID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if cart == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There is no cart with given ID"))
		return
	}

	ci, err := cir.Cis.CreateCartItem(cartID, cartItemBody.Product, cartItemBody.Quantity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ci)
}

func (cir *CartItemRouter) deleteCartItem(w http.ResponseWriter, r *http.Request) {
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

	err := cir.Cis.RemoveCartItem(cartID, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{}{})
}
