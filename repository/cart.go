package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/roflanKisel/cart-go/model"
)

// CartRepository interface
type CartRepository interface {
	Create(c *model.Cart) (*model.Cart, error)
	FindAll() ([]*model.Cart, error)
	FindByID(id string) (*model.Cart, error)
	UpdateByID(id string, c *model.Cart) error
	DeleteByID(id string) error
}

// MongoCartRepository provides methods for managing
// carts using mongo database
type MongoCartRepository struct {
	Db *mongo.Database
}

const cartCollectionName = "carts"

// Create will insert cart into database
func (m *MongoCartRepository) Create(c *model.Cart) (*model.Cart, error) {
	collection := m.Db.Collection(cartCollectionName)

	mc, err := c.GetMongoObject()
	if err != nil {
		return nil, err
	}

	res, err := collection.InsertOne(context.TODO(), mc)
	if err != nil {
		return nil, err
	}

	c.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return c, nil
}

// FindAll will return all the carts from database
func (m *MongoCartRepository) FindAll() ([]*model.Cart, error) {
	collection := m.Db.Collection(cartCollectionName)
	var carts []*model.Cart

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem model.MongoCart
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		carts = append(carts, elem.GetDefaultObject(nil))
	}

	return carts, nil
}

// FindByID will return cart from database that matches provided id
func (m *MongoCartRepository) FindByID(id string) (*model.Cart, error) {
	collection := m.Db.Collection(cartCollectionName)
	var cart model.MongoCart

	idFilter, err := getIDFilter(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), idFilter).Decode(&cart)
	if err != nil {
		return nil, err
	}

	return cart.GetDefaultObject(nil), nil
}

// UpdateByID will update cart in database that matches provided id
func (m *MongoCartRepository) UpdateByID(id string, c *model.Cart) error {
	collection := m.Db.Collection(cartCollectionName)

	idFilter, err := getIDFilter(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"items": c.Items,
		},
	}

	_, err = collection.UpdateOne(context.TODO(), idFilter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID will delete cart from database that matches provided id
func (m *MongoCartRepository) DeleteByID(id string) error {
	collection := m.Db.Collection(cartCollectionName)
	idFilter, err := getIDFilter(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), idFilter)
	if err != nil {
		return err
	}

	return nil
}
