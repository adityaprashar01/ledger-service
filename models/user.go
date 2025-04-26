package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"` // store hashed password
	CreatedAt time.Time          `json:"created_at"`
}
