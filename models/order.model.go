package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Order_date time.Time          `bson:"order_date" json:"order_date"`
	Created_at time.Time          `bson:"created_date" json:"created_date"`
	Updated_at time.Time          `bson:"updated_at" json:"updated_at"`
	Order_Id   string             `bson:"order_id" json:"order_id"`
	Table_Id   *string            `bson:"table_id" json:"table_id"`
}
