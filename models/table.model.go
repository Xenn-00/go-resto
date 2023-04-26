package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Number_of_guests *int               `bson:"number_of_guest" json:"number_of_guest" validate:"required"`
	Table_number     *int               `bson:"table_number" json:"table_number" validate:"required"`
	Created_at       time.Time          `bson:"created_at" json:"created_at"`
	Updated_at       time.Time          `bson:"updated_at" json:"updated_at"`
	Table_Id         string             `bson:"table_id" json:"table_id"`
}
