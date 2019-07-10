package app

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Exercise is struct for data
type Exercise struct {
	Title       string
	Description string
	Testcase    string
}

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	collection = client.Database("valexer").Collection("exercise")
}

func createDBExer(exer *Exercise) error {
	_, err := collection.InsertOne(context.TODO(), exer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getDBExers() ([]*Exercise, error) {
	var results []*Exercise
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	defer cur.Close(context.TODO())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem Exercise
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}
