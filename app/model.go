package app

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Exercise is struct for data
type Exercise struct {
	ID          interface{} `bson:"_id,omitempty"`
	Title       string      `bson:"title"`
	Description string      `bson:"description"`
	Testcase    string      `bson:"testcase"`
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

func getDBExer(id string) (*Exercise, error) {
	var result Exercise
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &result, nil
}

func updateDBExer(id string, exer *Exercise) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: exer.Title},
			{Key: "description", Value: exer.Description},
			{Key: "testcase", Value: exer.Testcase},
		}},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
