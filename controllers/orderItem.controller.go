package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderItemCollection *mongo.Collection = database.CreateCollection(database.Client, "Order-Items")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
