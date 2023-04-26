package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderCollection *mongo.Collection = database.CreateCollection(database.Client, "Orders")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}