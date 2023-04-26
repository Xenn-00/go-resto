package controllers

import (
	"github.com/Xenn-00/go-resto/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var InvoiceCollection *mongo.Collection = database.CreateCollection(database.Client, "Invoices")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}