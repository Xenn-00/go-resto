package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var MenuCollection *mongo.Collection = database.MenuData(database.Client, "Menus")
