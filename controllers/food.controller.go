package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Xenn-00/go-resto/database"
	"github.com/Xenn-00/go-resto/helpers"
	"github.com/Xenn-00/go-resto/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FoodCollection *mongo.Collection = database.CreateCollection(database.Client, "Foods")
var Validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 5
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, _ = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{
			{
				Key:   "$match",
				Value: bson.D{{}},
			},
		}
		groupStage := bson.D{
			{
				Key: "$group",
				Value: bson.D{
					{
						Key: "_id",
						Value: bson.D{
							{
								Key:   "_id",
								Value: "null",
							},
						},
					},
					{
						Key: "total_count",
						Value: bson.D{{
							Key:   "$sum",
							Value: 1,
						}},
					},
					{
						Key: "data",
						Value: bson.D{{
							Key:   "$push",
							Value: "$$ROOT",
						}},
					},
				},
			},
		}
		projectStage := bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "total_count", Value: 1},
					{Key: "food_items", Value: bson.D{
						{
							Key:   "$slice",
							Value: []interface{}{"$data", startIndex, recordPerPage},
						},
					}},
				},
			},
		}
		result, err := FoodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		}, options.Aggregate())

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"Foods": allFoods[0],
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodId := c.Param("food_id")
		var food *models.Food

		err := FoodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "error occured while fetching the food",
			})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var food *models.Food
		var menu *models.Menu

		if err := c.ShouldBindJSON(&food); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := Validate.Struct(food)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = MenuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_Id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "menu was not found",
			})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Food_Id = food.ID.Hex()

		var price = helpers.ToFixed(*food.Price, 2)
		food.Price = &price

		result, err := FoodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to create a food item",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu *models.Menu
		var food *models.Food

		foodId := c.Param("food_id")

		if err := c.ShouldBindJSON(&food); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var updateObj primitive.D
		if food.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: food.Name})
		}
		if food.Price != nil {
			updateObj = append(updateObj, bson.E{Key: "price", Value: food.Price})
		}
		if food.Food_Image != nil {
			updateObj = append(updateObj, bson.E{Key: "food_image", Value: food.Food_Image})
		}
		if food.Menu_Id != nil {
			err := MenuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_Id}).Decode(&menu)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "menu not found",
					"error":   err.Error(),
				})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "menu_id", Value: food.Price})
		}
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: food.Updated_at})

		upsert := true
		filter := bson.M{"food_id": foodId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := FoodCollection.UpdateOne(ctx, filter, bson.D{
			{
				Key:   "$set",
				Value: updateObj,
			},
		}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update food item",
				"error":   err.Error(),
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})

	}
}
