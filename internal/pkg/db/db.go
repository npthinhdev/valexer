package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// TypeMongoDB is type of mongodb
	TypeMongoDB = "mongodb"
	// TypeMySQL is type of mysql
	TypeMySQL = "mysql"
)

type (
	// Connections all supported types of database connections
	Connections struct {
		Type    string
		MongoDB *mongo.Client
	}
)

// Close close all underlying connections
func (c *Connections) Close() error {
	switch c.Type {
	case TypeMongoDB:
		if c.MongoDB != nil {
			err := c.MongoDB.Disconnect(context.TODO())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
