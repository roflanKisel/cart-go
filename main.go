package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/roflanKisel/cart-go/repository"
	"github.com/roflanKisel/cart-go/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/roflanKisel/cart-go/router"
)

const (
	connectionString = "mongodb://localhost:27017"
	databaseName     = "cart_go"
	port             = ":3000"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	defer client.Disconnect(context.TODO())
	defer fmt.Println("Disconnected")

	if err != nil {
		log.Fatal("Error connecting database", err)
	}

	database := client.Database(databaseName)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	cartRepository := &repository.MongoCartRepository{Db: database}
	cartItemRepository := &repository.MongoCartItemRepository{Db: database}

	cartService := &service.CartService{R: cartRepository, Cir: cartItemRepository}
	cartItemService := &service.CartItemService{R: cartItemRepository}

	cr := &router.CartRouter{Cs: cartService}
	cir := &router.CartItemRouter{Cis: cartItemService, Cs: cartService}

	cr.RegisterCartHandlers(r)
	cir.RegisterCartItemHandlers(r)

	http.ListenAndServe(port, r)
}
