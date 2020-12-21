package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CartItem represents item in cart as a common model.
type CartItem struct {
	ID       string `json:"_id"      bson:"_id,omitempty"`
	CartID   string `json:"cart_id"  bson:"cart_id,omitempty"`
	Product  string `json:"product"  bson:"product"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

// MongoCartItem represents item in cart as a mongo doc based on common model.
type MongoCartItem struct {
	ID       primitive.ObjectID `json:"_id"      bson:"_id,omitempty"`
	CartID   primitive.ObjectID `json:"cart_id"  bson:"cart_id,omitempty"`
	Product  string             `json:"product"  bson:"product"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

// MongoObject converts common CartItem object to mongo version.
func (ci CartItem) MongoObject() (*MongoCartItem, error) {
	var ciID primitive.ObjectID
	if ci.ID != "" {
		id, err := primitive.ObjectIDFromHex(ci.ID)
		if err != nil {
			return nil, err
		}

		ciID = id
	}

	var ciCartID primitive.ObjectID
	if ci.CartID != "" {
		cartID, err := primitive.ObjectIDFromHex(ci.CartID)
		if err != nil {
			return nil, err
		}

		ciCartID = cartID
	}

	return &MongoCartItem{
		ID:       ciID,
		CartID:   ciCartID,
		Product:  ci.Product,
		Quantity: ci.Quantity,
	}, nil
}

// DefaultObject converts mongo CartItem object to common version.
func (mci MongoCartItem) DefaultObject(ci *CartItem) *CartItem {
	if ci == nil {
		return &CartItem{
			ID:       mci.ID.Hex(),
			CartID:   mci.CartID.Hex(),
			Product:  mci.Product,
			Quantity: mci.Quantity,
		}
	}

	ci.ID = mci.ID.Hex()
	ci.CartID = mci.CartID.Hex()
	ci.Product = mci.Product
	ci.Quantity = mci.Quantity

	return ci
}
