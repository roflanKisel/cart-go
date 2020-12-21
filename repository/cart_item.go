package repository

import (
	"context"

	"github.com/roflanKisel/cart-go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CartItemKeeper is the interface that wraps basic CRUD operation over the CartItem.
type CartItemKeeper interface {
	Create(ctx context.Context, c *model.CartItem) (*model.CartItem, error)
	All(ctx context.Context) ([]*model.CartItem, error)
	ByID(ctx context.Context, id string) (*model.CartItem, error)
	UpdateByID(ctx context.Context, id string, c *model.CartItem) error
	DeleteByID(ctx context.Context, id string) error
	ByCartID(ctx context.Context, id string) ([]model.CartItem, error)
}

// NewMongoCartItemRepository returns mongo repository for CartItem which uses passed database.
func NewMongoCartItemRepository(db *mongo.Database) *MongoCartItemRepository {
	return &MongoCartItemRepository{collection: db.Collection("cart_items")}
}

// MongoCartItemRepository is the implementation of CartItemKeeper interface based on MondoDB.
type MongoCartItemRepository struct {
	collection *mongo.Collection
}

// Create inserts a CartItem object into database.
func (m MongoCartItemRepository) Create(ctx context.Context, c *model.CartItem) (*model.CartItem, error) {
	mc, err := c.MongoObject()
	if err != nil {
		return nil, err
	}

	res, err := m.collection.InsertOne(ctx, mc)
	if err != nil {
		return nil, err
	}

	c.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return c, nil
}

// All returns an array of CartItem objects from database.
func (m MongoCartItemRepository) All(ctx context.Context) ([]*model.CartItem, error) {
	var cartItems []*model.CartItem

	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem model.MongoCartItem
		if err = cur.Decode(&elem); err != nil {
			return nil, err
		}

		cartItems = append(cartItems, elem.DefaultObject(nil))
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

// ByID returns a CartItem object from database that matches passed id.
func (m MongoCartItemRepository) ByID(ctx context.Context, id string) (*model.CartItem, error) {
	var cartItem model.MongoCartItem

	idFilter, err := idFilter(id)
	if err != nil {
		return nil, err
	}

	err = m.collection.FindOne(ctx, idFilter).Decode(&cartItem)
	if err != nil {
		return nil, err
	}

	return cartItem.DefaultObject(nil), nil
}

// UpdateByID updates a CartItem object in database that matches passed id.
func (m MongoCartItemRepository) UpdateByID(ctx context.Context, id string, c *model.CartItem) error {
	idFilter, err := idFilter(id)
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

	_, err = m.collection.UpdateOne(ctx, idFilter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes CartItem object from database that matches passed id.
func (m MongoCartItemRepository) DeleteByID(ctx context.Context, id string) error {
	idFilter, err := idFilter(id)
	if err != nil {
		return err
	}

	_, err = m.collection.DeleteOne(ctx, idFilter)
	if err != nil {
		return err
	}

	return nil
}

// ByCartID returns an array of CartItem objects from database that match passed id and cart_id property.
func (m *MongoCartItemRepository) ByCartID(ctx context.Context, id string) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	cartObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	query := bson.M{"cart_id": cartObjectID}

	cur, err := m.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem model.MongoCartItem
		if err = cur.Decode(&elem); err != nil {
			return nil, err
		}

		cartItems = append(cartItems, *elem.DefaultObject(nil))
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}
