package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.CreateCollection(database.Client, "Users")
