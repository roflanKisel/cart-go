package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CartItem model
type CartItem struct {
	ID       string `json:"_id"      bson:"_id,omitempty"`
	CartID   string `json:"cart_id"  bson:"cart_id,omitempty"`
	Product  string `json:"product"  bson:"product"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

// MongoCartItem is a mongo version of CartItem model
type MongoCartItem struct {
	ID       primitive.ObjectID `json:"_id"      bson:"_id,omitempty"`
	CartID   primitive.ObjectID `json:"cart_id"  bson:"cart_id,omitempty"`
	Product  string             `json:"product"  bson:"product"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

// GetMongoObject converts default CartItem to mongo version
func (ci *CartItem) GetMongoObject() (*MongoCartItem, error) {
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

// GetDefaultObject converts MongoCartItem to default CartItem struct
func (mci *MongoCartItem) GetDefaultObject(ci *CartItem) *CartItem {
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
