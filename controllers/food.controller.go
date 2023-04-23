package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var FoodCollection *mongo.Collection = database.FoodData(database.Client, "Foods")