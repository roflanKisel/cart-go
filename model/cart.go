package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cart model
type Cart struct {
	ID    string     `json:"_id"   bson:"_id,omitempty"`
	Items []CartItem `json:"items" bson:"items"`
}

// MongoCart is a mongo version of Cart model
type MongoCart struct {
	ID    primitive.ObjectID `json:"_id"   bson:"_id,omitempty"`
	Items []CartItem         `json:"items" bson:"items"`
}

// GetMongoObject converts default Cart to MongoCart
func (c *Cart) GetMongoObject() (*MongoCart, error) {
	var id primitive.ObjectID

	if c.ID != "" {
		mid, err := primitive.ObjectIDFromHex(c.ID)
		if err != nil {
			return nil, err
		}

		id = mid
	}

	return &MongoCart{
		ID:    id,
		Items: c.Items,
	}, nil
}

// GetDefaultObject converts MongoCartItem to Cart
func (mc *MongoCart) GetDefaultObject(c *Cart) *Cart {
	if c == nil {
		return &Cart{
			ID:    mc.ID.Hex(),
			Items: mc.Items,
		}
	}

	c.ID = mc.ID.Hex()
	c.Items = mc.Items

	return c
}
