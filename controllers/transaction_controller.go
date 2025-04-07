package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ledger-service/database"
	"ledger-service/models"
)

func CreateTransaction(c *gin.Context) {
	var txn models.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	txn.Timestamp = time.Now()
	txn.ID = primitive.NewObjectID() // Generate ObjectID before inserting

	client := database.DB
	sess, err := client.StartSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start DB session"})
		return
	}
	defer sess.EndSession(ctx)

	result, err := sess.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		customerCol := database.GetCollection("customers")
		txnCol := database.GetCollection("transactions")

		var customer models.Customer
		err := customerCol.FindOne(sessCtx, bson.M{"_id": txn.CustomerID}).Decode(&customer)
		if err != nil {
			return nil, err
		}

		if txn.Type == "debit" && customer.Balance < txn.Amount {
			return nil, mongo.CommandError{Message: "Insufficient balance"}
		}

		newBalance := customer.Balance
		if txn.Type == "credit" {
			newBalance += txn.Amount
		} else {
			newBalance -= txn.Amount
		}

		// Explicit insert with all fields including customer_id
		_, err = txnCol.InsertOne(sessCtx, bson.M{
			"_id":         txn.ID,
			"customer_id": txn.CustomerID,
			"amount":      txn.Amount,
			"type":        txn.Type,
			"timestamp":   txn.Timestamp,
		})
		if err != nil {
			return nil, err
		}

		_, err = customerCol.UpdateOne(sessCtx, bson.M{"_id": txn.CustomerID}, bson.M{"$set": bson.M{"balance": newBalance}})
		if err != nil {
			return nil, err
		}

		return newBalance, nil
	}, options.Transaction())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": txn.ID,
		"status":         "success",
		"balance":        result,
	})
}
func GetTransactionHistory(c *gin.Context) {
	customerID := c.Param("customer_id")
	objID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Parse pagination parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}}) // latest first

	cursor, err := database.GetCollection("transactions").Find(
		ctx,
		bson.M{"customer_id": objID},
		findOptions,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	defer cursor.Close(ctx)

	var transactions []models.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor parse failed"})
		return
	}

	// Struct without customer_id
	type PublicTransaction struct {
		ID        primitive.ObjectID `json:"_id"`
		Amount    float64            `json:"amount"`
		Type      string             `json:"type"`
		Timestamp time.Time          `json:"timestamp"`
	}

	var response []PublicTransaction
	for _, txn := range transactions {
		response = append(response, PublicTransaction{
			ID:        txn.ID,
			Amount:    txn.Amount,
			Type:      txn.Type,
			Timestamp: txn.Timestamp,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"page":         page,
		"limit":        limit,
		"transactions": response,
	})
}
