package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"transaction_id"`
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	Type       string             `bson:"type" json:"type"`
	Amount     float64            `bson:"amount" json:"amount"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
}
