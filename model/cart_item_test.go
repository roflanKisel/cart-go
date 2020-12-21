package model_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/roflanKisel/cart-go/model"
)

const (
	id     = "5fd78878266f92cf3670b9a5"
	cartID = "5fd78878266f92cf3670b9a5"
)

func TestCartItemMongoObject(t *testing.T) {
	ci := model.CartItem{
		ID:       id,
		CartID:   cartID,
		Product:  "Test",
		Quantity: 10,
	}

	mci, err := ci.MongoObject()
	if err != nil {
		t.Error(err)
		return
	}

	if mci.Product != ci.Product || mci.Quantity != ci.Quantity {
		t.Errorf("MongoObject(): expected %v, actual %v", ci, mci)
		return
	}
}

func TestCartItemDefaultObject(t *testing.T) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatal("DefaultObject(): Cannot convert ID to ObjectID")
		return
	}

	objCartID, err := primitive.ObjectIDFromHex(cartID)
	if err != nil {
		t.Fatal("DefaultObject(): Cannot convert CartID to ObjectID")
		return
	}

	mci := model.MongoCartItem{
		ID:       objID,
		CartID:   objCartID,
		Product:  "Test",
		Quantity: 10,
	}

	ci := mci.DefaultObject(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if ci.Product != mci.Product || ci.Quantity != mci.Quantity || ci.ID != id || ci.CartID != cartID {
		t.Errorf("MongoObject(): expected %v, actual %v", ci, mci)
		return
	}
}
