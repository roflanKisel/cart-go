package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/roflanKisel/cart-go/repository"
	"github.com/roflanKisel/cart-go/router"
	"github.com/roflanKisel/cart-go/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	if err != nil {
		log.Fatal("Error connecting database", err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal("Error disconnecting database client")
		}
	}()

	database := client.Database(databaseName)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Connection to database failed: %s", err.Error()))
	}

	r := mux.NewRouter()

	cartRepository := repository.NewMongoCartRepository(database)
	cartItemRepository := repository.NewMongoCartItemRepository(database)

	cartService := service.NewCartService(cartRepository, cartItemRepository)
	cartItemService := service.NewCartItemService(cartItemRepository)

	cartRouter := router.NewCartRouter(cartService)
	cartItemRouter := router.NewCartItemRouter(cartService, cartItemService)

	cartRouter.RegisterCartHandlers(r)
	cartItemRouter.RegisterCartItemHandlers(r)

	log.Fatal(http.ListenAndServe(port, r))
}
