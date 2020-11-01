package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SarunasBucius/tables-management/config"
	"github.com/SarunasBucius/tables-management/tables"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri, err := config.GetDbURI("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	client := connectToMongo(uri)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Mount("/api/tables", tables.SetRoutes(client))

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
