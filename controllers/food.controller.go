package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var FoodCollection *mongo.Collection = database.CreateCollection(database.Client, "Foods")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
