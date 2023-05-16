package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Xenn-00/go-resto/database"
	"github.com/Xenn-00/go-resto/helpers"
	"github.com/Xenn-00/go-resto/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MenuCollection *mongo.Collection = database.CreateCollection(database.Client, "Menus")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := MenuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		var Menus []bson.M
		if err = result.All(ctx, &Menus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menuId := c.Param("menu_id")
		var menu *models.Menu

		err := MenuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "error occured while fetching the menu",
			})
			return
		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu *models.Menu

		if err := c.ShouldBindJSON(&menu); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := Validate.Struct(menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_Id = menu.ID.Hex()

		result, err := MenuCollection.InsertOne(ctx, menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to create menu item",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu *models.Menu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D
		if menu.Start_date != nil && menu.End_date != nil {
			if !helpers.InTimeSpan(*menu.Start_date, *menu.End_date, time.Now()) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "kindly retype the time",
				})
				defer cancel()
				return
			}
			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_date})

			if menu.Name != nil {
				updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
			}

			if menu.Category != nil {
				updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := MenuCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": "failed to update menu",
				})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)
		}
	}
}
