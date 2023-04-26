package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var MenuCollection *mongo.Collection = database.CreateCollection(database.Client, "Menus")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
