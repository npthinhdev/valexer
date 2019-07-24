package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// Config hold MongoDB configuration information
	Config struct {
		Addrs    string        `env:"MONGODB_ADDRS" default:"127.0.0.1:27017"`
		Database string        `env:"MONGGO_DATABASE" default:"valexer"`
		Username string        `env:"MONGODB_USERNAME"`
		Password string        `env:"MONGODB_PASSWORD"`
		Timeout  time.Duration `env:"MONGODB_TIMEOUT" default:"10s"`
	}
)

// Dial dial to target server with Monotonic mode
func Dial(conf *Config) (*mongo.Client, error) {
	conf.Addrs = "127.0.0.1:27017"
	conf.Database = "valexer"

	log.Printf("dialing to target MongoDB at: %v, database: %v", conf.Addrs, conf.Database)
	mongoURL := fmt.Sprintf("mongodb://%s/%s", conf.Addrs, conf.Database)
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("successfully dialing to MongoDB at", conf.Addrs)
	return client, nil
}
