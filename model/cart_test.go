package model_test

import (
	"testing"

	"github.com/roflanKisel/cart-go/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCartMongoObject(t *testing.T) {
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Cannot convert id to ObjectID")
	}

	carts := []struct {
		name     string
		c        model.Cart
		expected model.MongoCart
	}{
		{"With correct ID", model.Cart{
			ID: id,
		}, model.MongoCart{
			ID: mID,
		}},
		{"Without ID", model.Cart{}, model.MongoCart{}},
	}

	for _, cart := range carts {
		t.Run(cart.name, func(t *testing.T) {
			mc, err := cart.c.MongoObject()
			if err != nil {
				if cart.c.ID == "" {
					return
				}

				t.Fatalf("Error during convertion: %s", err)
			}

			assert.Equal(t, cart.expected, *mc, "should be equal")
		})
	}
}

func TestCartDefaultObject(t *testing.T) {
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Cannot convert id to ObjectID")
	}

	mCarts := []struct {
		name     string
		mc       model.MongoCart
		expected model.Cart
	}{
		{"With correct ID", model.MongoCart{
			ID: mID,
		}, model.Cart{
			ID: id,
		}},
		{"Without ID", model.MongoCart{}, model.Cart{
			ID: "000000000000000000000000",
		}},
	}

	for _, mCart := range mCarts {
		t.Run(mCart.name, func(t *testing.T) {
			c := mCart.mc.DefaultObject(nil)
			assert.Equal(t, mCart.expected, *c, "should be equal")
		})
	}
}
