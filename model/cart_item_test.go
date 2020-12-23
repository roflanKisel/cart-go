package model_test

import (
	"testing"

	"github.com/roflanKisel/cart-go/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	id = "5fd78878266f92cf3670b9a5"
)

func TestCartItemMongoObject(t *testing.T) {
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Cannot convert id to ObjectID")
	}

	cartItems := []struct {
		name     string
		ci       model.CartItem
		expected model.MongoCartItem
	}{
		{"With correct values", model.CartItem{
			ID:       id,
			CartID:   id,
			Product:  "Test",
			Quantity: 10,
		}, model.MongoCartItem{
			ID:       mID,
			CartID:   mID,
			Product:  "Test",
			Quantity: 10,
		}},
		{"With only IDs", model.CartItem{
			ID:     id,
			CartID: id,
		}, model.MongoCartItem{
			ID:       mID,
			CartID:   mID,
			Product:  "",
			Quantity: 0,
		}},
		{"Without IDs", model.CartItem{}, model.MongoCartItem{}},
	}

	for _, ci := range cartItems {
		t.Run(ci.name, func(t *testing.T) {
			mci, err := ci.ci.MongoObject()
			if err != nil {
				if ci.ci.ID == ci.ci.CartID && ci.ci.ID == "" {
					return
				}

				t.Fatalf("Error during convertion: %s", err)
			}

			assert.Equal(t, ci.expected, *mci, "should be equal")
		})
	}
}

func TestCartItemDefaultObject(t *testing.T) {
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Cannot convert id to ObjectID")
	}

	mCartItems := []struct {
		name     string
		m        model.MongoCartItem
		expected model.CartItem
	}{
		{"With correct values", model.MongoCartItem{
			ID:       mID,
			CartID:   mID,
			Product:  "Test",
			Quantity: 10,
		}, model.CartItem{
			ID:       id,
			CartID:   id,
			Product:  "Test",
			Quantity: 10,
		}},
		{"With only IDs", model.MongoCartItem{
			ID:     mID,
			CartID: mID,
		}, model.CartItem{
			ID:       id,
			CartID:   id,
			Product:  "",
			Quantity: 0,
		}},
		{"Without IDs", model.MongoCartItem{}, model.CartItem{
			ID:     "000000000000000000000000", // primitive.ObjectID.Hex() result when ObjectID has default value
			CartID: "000000000000000000000000",
		}},
	}

	for _, mci := range mCartItems {
		t.Run(mci.name, func(t *testing.T) {
			ci := mci.m.DefaultObject(nil)
			assert.Equal(t, mci.expected, *ci, "should be equal")
		})
	}
}
