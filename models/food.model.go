package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name       *string            `bson:"name" json:"name" validate:"required,min=3.max=100"`
	Price      *float64           `bson:"price" json:"price" validate:"required"`
	Food_Image *string            `bson:"food_image" json:"food_image" validate:"required"`
	Created_at time.Time          `bson:"created_at" json:"created_at"`
	Updated_at time.Time          `bson:"updated_at" json:"updated_at"`
	Food_Id    string             `bson:"food_id" json:"food_id"`
	Menu_Id    *string            `bson:"menu_id" json:"menu_id" validate:"required"`
}
