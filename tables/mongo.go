package tables

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getTableByID(c config, ID string) (table, error) {
	db := c.mongo.Database("content_management")
	coll := db.Collection("tables")

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		c.log.Println(err)
		return table{}, err
	}

	var t table
	err = coll.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&t)
	if err != nil {
		c.log.Println(err)
		return table{}, err
	}

	return t, nil
}
