package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Text       string             `bson:"text" json:"text"`
	Title      string             `bson:"title" json:"title"`
	Created_at time.Time          `bson:"created_at" json:"created_at"`
	Updated_at time.Time          `bson:"updated_at" json:"updated_at"`
	Note_Id    string             `bson:"note_id"`
}
