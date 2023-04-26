package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Invoice_Id       string             `bson:"invoice_id" json:"invoice_id"`
	Order_Id         string             `bson:"order_id" json:"order_id"`
	Payment_method   *string            `bson:"payment_method" json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `bson:"payment_status" json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `bson:"payment_due_date" json:"payment_due_date"`
	Created_at       time.Time          `bson:"created_at" json:"created_at"`
	Updated_at       time.Time          `bson:"updated_at" json:"updated_at"`
}
