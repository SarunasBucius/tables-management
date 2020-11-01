package tables

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func getTables(c config) ([]table, error) {
	db := c.mongo.Database("content_management")
	coll := db.Collection("tables")

	res, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		c.log.Println(err)
		return nil, err
	}

	var tables []table
	err = res.All(context.Background(), &tables)
	if err != nil {
		log.Fatal(err)
	}

	return tables, nil
}
