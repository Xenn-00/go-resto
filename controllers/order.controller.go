package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderCollection *mongo.Collection = database.OrderData(database.Client, "Orders")
