package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderItemCollection *mongo.Collection = database.OrderItemData(database.Client, "Order-Items")
