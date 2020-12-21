package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart represents cart as a common model.
type Cart struct {
	ID    string     `json:"_id"   bson:"_id,omitempty"`
	Items []CartItem `json:"items" bson:"items"`
}

// MongoCart represents cart as a mongo doc based on common model.
type MongoCart struct {
	ID    primitive.ObjectID `json:"_id"   bson:"_id,omitempty"`
	Items []CartItem         `json:"items" bson:"items"`
}

// MongoObject converts common Cart object to mongo version.
func (c Cart) MongoObject() (*MongoCart, error) {
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

// DefaultObject converts mongo Cart object to common version.
func (mc MongoCart) DefaultObject(c *Cart) *Cart {
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
