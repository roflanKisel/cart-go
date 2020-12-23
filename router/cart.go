package router

import (
	"encoding/json"
	"net/http"

	"github.com/roflanKisel/cart-go/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewCartRouter returns router for Cart which uses passed cart service.
func NewCartRouter(cSvc *service.CartService) *CartRouter {
	return &CartRouter{cSvc: cSvc}
}

// CartRouter provides possibility to register HTTP handlers for Cart model.
type CartRouter struct {
	cSvc *service.CartService
}

// RegisterCartHandlers registers HTTP handlers for Cart model.
func (cr CartRouter) RegisterCartHandlers(r *mux.Router) {
	r.HandleFunc("/carts/{id}", cr.cart).Methods("GET")
	r.HandleFunc("/carts", cr.createCart).Methods("POST")
}

func (cr CartRouter) cart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		cart, err := cr.cSvc.CartByID(ctx, id)
		if err != nil {
			if err == mongo.ErrNilDocument {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		err = json.NewEncoder(w).Encode(cart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (cr CartRouter) createCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cart, err := cr.cSvc.CreateEmpty(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
