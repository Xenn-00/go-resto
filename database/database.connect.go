package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetup() *mongo.Client {
	uri := fmt.Sprintf("mongodb+srv://Xenn:%s@cluster0.ycwzsug.mongodb.net/restaurant-management?retryWrites=true&w=majority", os.Getenv("MONGO_PASSWORD"))
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("failed to connect to database")
		return nil
	}
	fmt.Println("Successfully connect to database")
	return client
}

var Client *mongo.Client = DBSetup()

func FoodData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}

func InvoiceData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}

func MenuData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}

func TableData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}

func OrderData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}
func OrderItemData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant-management").Collection(collectionName)
	return collection
}