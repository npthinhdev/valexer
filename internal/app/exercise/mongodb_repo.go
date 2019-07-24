package exercise

import (
	"context"

	"github.com/npthinhdev/valexer/internal/app/types"
	"github.com/npthinhdev/valexer/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository is MongoDB implementation of repository
type MongoRepository struct {
	client *mongo.Client
}

// NewMongoRepository return new MongoDB repository
func NewMongoRepository(c *mongo.Client) *MongoRepository {
	return &MongoRepository{c}
}

func (r *MongoRepository) collection(s *mongo.Client) *mongo.Collection {
	return s.Database("valexer").Collection("exercise")
}

// FindAll return all exercise
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Exercise, error) {
	var exercise []types.Exercise
	cl := r.client
	collection := r.collection(cl)
	cur, err := collection.Find(ctx, bson.D{{}})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem types.Exercise
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		exercise = append(exercise, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return exercise, nil
}

// FindByID return exercise base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Exercise, error) {
	var exercise *types.Exercise
	cl := r.client
	collection := r.collection(cl)
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(&exercise)
	if err != nil {
		return nil, err
	}
	return exercise, nil
}

// Create a exercise
func (r *MongoRepository) Create(ctx context.Context, exercise types.Exercise) (string, error) {
	cl := r.client
	collection := r.collection(cl)
	exercise.ID = uuid.New()
	_, err := collection.InsertOne(ctx, exercise)
	return exercise.ID, err
}

// Update a exercise
func (r *MongoRepository) Update(ctx context.Context, exercise types.Exercise) error {
	cl := r.client
	collection := r.collection(cl)
	filter := bson.D{{Key: "_id", Value: exercise.ID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: exercise.Title},
			{Key: "description", Value: exercise.Description},
			{Key: "testcase", Value: exercise.Testcase},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// Delete a exercise
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	cl := r.client
	collection := r.collection(cl)
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := collection.DeleteOne(ctx, filter, nil)
	return err
}
