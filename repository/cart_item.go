package repository

import (
	"context"

	"github.com/roflanKisel/cart-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CartItemRepository interface
type CartItemRepository interface {
	Create(c *model.CartItem) (*model.CartItem, error)
	FindAll() ([]*model.CartItem, error)
	FindByID(id string) (*model.CartItem, error)
	UpdateByID(id string, c *model.CartItem) error
	DeleteByID(id string) error
	FindByCartID(id string) ([]model.CartItem, error)
}

// MongoCartItemRepository provides methods for managing
// cart items using mongo database
type MongoCartItemRepository struct {
	Db *mongo.Database
}

const cartItemCollectionName = "cart_items"

// Create will insert cart item into database
func (m *MongoCartItemRepository) Create(c *model.CartItem) (*model.CartItem, error) {
	collection := m.Db.Collection(cartItemCollectionName)

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

// FindAll will return all the cart items from database
func (m *MongoCartItemRepository) FindAll() ([]*model.CartItem, error) {
	collection := m.Db.Collection(cartItemCollectionName)
	var cartItems []*model.CartItem

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem model.MongoCartItem
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, elem.GetDefaultObject(nil))
	}

	return cartItems, nil
}

// FindByID will return cart item from database that matches provided id
func (m *MongoCartItemRepository) FindByID(id string) (*model.CartItem, error) {
	collection := m.Db.Collection(cartItemCollectionName)
	var cartItem model.MongoCartItem

	idFilter, err := getIDFilter(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), idFilter).Decode(&cartItem)
	if err != nil {
		return nil, err
	}

	return cartItem.GetDefaultObject(nil), nil
}

// UpdateByID will update cart item in database that matches provided id
func (m *MongoCartItemRepository) UpdateByID(id string, c *model.CartItem) error {
	collection := m.Db.Collection(cartItemCollectionName)

	idFilter, err := getIDFilter(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"product":  c.Product,
			"quantity": c.Quantity,
			"cart_id":  c.CartID,
		},
	}

	_, err = collection.UpdateOne(context.TODO(), idFilter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID will delete cart item from database that matches provided id
func (m *MongoCartItemRepository) DeleteByID(id string) error {
	collection := m.Db.Collection(cartItemCollectionName)
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

// FindByCartID will return cart items matching given ID
func (m *MongoCartItemRepository) FindByCartID(id string) ([]model.CartItem, error) {
	collection := m.Db.Collection(cartItemCollectionName)
	var cartItems []model.CartItem

	cartObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	query := bson.M{"cart_id": cartObjectID}

	cur, err := collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem model.MongoCartItem
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, *elem.GetDefaultObject(nil))
	}

	return cartItems, nil
}
