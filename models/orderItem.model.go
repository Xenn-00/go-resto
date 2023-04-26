package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Quantity      *string            `bson:"quantity" json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price    *float64           `bson:"unit_price" json:"unit_price" validate:"required"`
	Created_at    time.Time          `bson:"created_at" json:"created_at"`
	Updated_at    time.Time          `bson:"updated_at" json:"updated_at"`
	Food_Id       *string            `bson:"food_id" json:"food_id" validate:"required"`
	Order_Item_Id string             `bson:"order_item_id" json:"order_item_id"`
	Order_Id      string             `bson:"order_id" json:"order_id" validate:"required"`
}
