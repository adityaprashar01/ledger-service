package tests

import (
	"context"
	"testing"
	"time"

	"ledger-service/database"
	"ledger-service/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	database.DB = client
}

func TestCreateTransactionLogic(t *testing.T) {
	customer := models.Customer{
		ID:      primitive.NewObjectID(),
		Name:    "Test User",
		Balance: 5000,
	}
	_, err := database.GetCollection("customers").InsertOne(context.TODO(), customer)
	assert.Nil(t, err)

	txn := models.Transaction{
		ID:         primitive.NewObjectID(),
		CustomerID: customer.ID,
		Type:       "credit",
		Amount:     1000,
		Timestamp:  time.Now(),
	}

	_, err = database.GetCollection("transactions").InsertOne(context.TODO(), txn)
	assert.Nil(t, err)

	_, err = database.GetCollection("customers").UpdateOne(context.TODO(),
		bson.M{"_id": txn.CustomerID},
		bson.M{"$inc": bson.M{"balance": txn.Amount}})
	assert.Nil(t, err)

	var updatedCustomer models.Customer
	err = database.GetCollection("customers").FindOne(context.TODO(),
		bson.M{"_id": customer.ID}).Decode(&updatedCustomer)
	assert.Nil(t, err)

	assert.Equal(t, float64(6000), updatedCustomer.Balance)
}
