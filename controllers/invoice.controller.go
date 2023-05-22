package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Xenn-00/go-resto/database"
	"github.com/Xenn-00/go-resto/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var InvoiceCollection *mongo.Collection = database.CreateCollection(database.Client, "Invoices")

type InvoiceFormat struct {
	Invoice_id       string
	Order_id         string
	Table_number     interface{}
	Order_details    interface{}
	Payment_method   string
	Payment_status   string
	Payment_due      interface{}
	Payment_due_date time.Time
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := InvoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"Invoices": allInvoices,
		})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoice_id := c.Param("invoice_id")

		var invoice *models.Invoice
		err := InvoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoice_id}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var invoiceFormat InvoiceFormat
		allOrderItems, _ := ItemsByOrder(invoice.Order_Id)

		invoiceFormat.Order_id = invoice.Order_Id
		invoiceFormat.Payment_due_date = invoice.Payment_due_date
		invoiceFormat.Payment_method = *invoice.Payment_method
		invoiceFormat.Invoice_id = invoice.Invoice_Id
		invoiceFormat.Payment_status = *invoice.Payment_status

		invoiceFormat.Payment_due = allOrderItems[0]["payment_due"]
		invoiceFormat.Table_number = allOrderItems[0]["table_number"]
		invoiceFormat.Order_details = allOrderItems[0]["order_details"]

		c.JSON(http.StatusOK, gin.H{
			"Invoice": invoiceFormat,
		})
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice *models.Invoice
		var order *models.Order

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := OrderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_Id}).Decode(&order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Cannot find the order id",
			})
			return
		}
		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_Id = invoice.ID.Hex()

		err = Validate.Struct(invoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		result, err := InvoiceCollection.InsertOne(ctx, invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Cannot insert invoice item ",
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice *models.Invoice
		var updateObj primitive.D

		invoice_id := c.Param("invoice_id")
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		filter := bson.M{"invoice_id": invoice_id}

		if invoice.Payment_status != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_status", Value: invoice.Payment_status})
		}

		if invoice.Payment_method != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_method", Value: invoice.Payment_method})
		}

		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: invoice.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := InvoiceCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Cannot update invoice item",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{
			"message": "update invoice successful",
			"result":  result,
		})
	}
}
