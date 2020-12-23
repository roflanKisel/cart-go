package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/roflanKisel/cart-go/config"
	"github.com/roflanKisel/cart-go/db"
	"github.com/roflanKisel/cart-go/repository"
	"github.com/roflanKisel/cart-go/router"
	"github.com/roflanKisel/cart-go/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := db.NewClient(ctx, config.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Error disconnecting database client: %s", err)
		}
	}()

	database, err := db.NewDB(ctx, client, config.DB())
	if err != nil {
		log.Fatal(err)
	}

	cartRepository := repository.NewMongoCartRepository(database)
	cartItemRepository := repository.NewMongoCartItemRepository(database)

	cartService := service.NewCartService(cartRepository, cartItemRepository)
	cartItemService := service.NewCartItemService(cartItemRepository)

	r := router.NewRouter(cartService, cartItemService)
	log.Fatal(http.ListenAndServe(config.AppPort(), r))
}
