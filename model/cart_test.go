package model_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/roflanKisel/cart-go/model"
)

func TestCartMongoObject(t *testing.T) {
	c := model.Cart{
		ID:    id,
		Items: []model.CartItem{},
	}

	mc, err := c.MongoObject()
	if err != nil {
		t.Error(err)
		return
	}

	if mc.ID.Hex() != id || len(mc.Items) != len(c.Items) {
		t.Errorf("MongoObject(): expected %v, actual %v", c, mc)
	}
}

func TestCartDefaultObject(t *testing.T) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatal("DefaultObject(): Cannot convert ID to ObjectID")
		return
	}

	mc := model.MongoCart{
		ID:    objID,
		Items: []model.CartItem{},
	}

	c := mc.DefaultObject(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if c.ID != id || len(c.Items) != len(mc.Items) {
		t.Errorf("MongoObject(): expected %v, actual %v", c, mc)
		return
	}
}
