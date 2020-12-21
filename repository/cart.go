package repository

import (
	"context"

	"github.com/roflanKisel/cart-go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CartKeeper is the interface that wraps basic CRUD operation over the Cart.
type CartKeeper interface {
	Create(ctx context.Context, c *model.Cart) (*model.Cart, error)
	All(ctx context.Context) ([]*model.Cart, error)
	ByID(ctx context.Context, id string) (*model.Cart, error)
	UpdateByID(ctx context.Context, id string, c *model.Cart) error
	DeleteByID(ctx context.Context, id string) error
}

// NewMongoCartRepository returns mongo repository for Cart which uses passed database.
func NewMongoCartRepository(db *mongo.Database) *MongoCartRepository {
	return &MongoCartRepository{collection: db.Collection("carts")}
}

// MongoCartRepository is the implementation of CartKeeper interface based on MondoDB.
type MongoCartRepository struct {
	collection *mongo.Collection
}

// Create inserts a Cart object into database.
func (m MongoCartRepository) Create(ctx context.Context, c *model.Cart) (*model.Cart, error) {
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

// All returns an array of Cart objects from database.
func (m MongoCartRepository) All(ctx context.Context) ([]*model.Cart, error) {
	var carts []*model.Cart

	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem model.MongoCart
		if err = cur.Decode(&elem); err != nil {
			return nil, err
		}

		carts = append(carts, elem.DefaultObject(nil))
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

// ByID returns a Cart object from database that matches passed id.
func (m MongoCartRepository) ByID(ctx context.Context, id string) (*model.Cart, error) {
	var cart model.MongoCart

	idFilter, err := idFilter(id)
	if err != nil {
		return nil, err
	}

	err = m.collection.FindOne(ctx, idFilter).Decode(&cart)
	if err != nil {
		return nil, err
	}

	return cart.DefaultObject(nil), nil
}

// UpdateByID updates a Cart object in database that matches passed id.
func (m MongoCartRepository) UpdateByID(ctx context.Context, id string, c *model.Cart) error {
	idFilter, err := idFilter(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"items": c.Items,
		},
	}

	_, err = m.collection.UpdateOne(ctx, idFilter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes Cart object from database that matches passed id.
func (m MongoCartRepository) DeleteByID(ctx context.Context, id string) error {
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
