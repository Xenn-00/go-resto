package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var TableCollection *mongo.Collection = database.CreateCollection(database.Client, "Tables")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
