package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"customer_id"`
	Name    string             `json:"name"`
	Balance float64            `json:"initial_balance"`
}
