package tables

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type config struct {
	mongo *mongo.Client
	log   *log.Logger
}

type table struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Columns []column           `bson:"columns"`
}

type column struct {
	ID         primitive.ObjectID `bson:"_id"`
	Index      int                `bson:"index"`
	Name       string             `bson:"name"`
	IsHidden   bool               `bson:"isHidden"`
	InputType  string             `bson:"inputType"`
	IsReadOnly bool               `bson:"isReadOnly"`
}
