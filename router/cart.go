package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roflanKisel/cart-go/service"

	"github.com/gorilla/mux"
)

// CartRouter provides cart handlers
type CartRouter struct {
	Cs *service.CartService
}

// RegisterCartHandlers registers HTTP handlers for Cart model
func (cr *CartRouter) RegisterCartHandlers(r *mux.Router) {
	r.HandleFunc("/carts/{id}", cr.getCart).Methods("GET")
	r.HandleFunc("/carts", cr.createCart).Methods("POST")
}

func (cr *CartRouter) getCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		cart, err := cr.Cs.GetCartByID(id)
		if err != nil {
			fmt.Println("Cannot get cart with id", id, ": ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cart)
		return
	}

	fmt.Println("id not found in URL")
	w.WriteHeader(http.StatusBadRequest)
}

func (cr *CartRouter) createCart(w http.ResponseWriter, r *http.Request) {
	cart, err := cr.Cs.CreateEmpty()
	if err != nil {
		fmt.Println("Cannot create empty cart", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}
