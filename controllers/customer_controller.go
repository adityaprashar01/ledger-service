package controllers

import (
	"context"
	"net/http"
	"time"

	"ledger-service/database"
	"ledger-service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCustomer(c *gin.Context) {
	var input struct {
		Name           string  `json:"name"`
		InitialBalance float64 `json:"initial_balance"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:    input.Name,
		Balance: input.InitialBalance,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := database.GetCollection("customers").InsertOne(ctx, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	customer.ID = res.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusOK, customer)
}

func GetCustomerBalance(c *gin.Context) {
	customerID := c.Param("customer_id")
	objID, _ := primitive.ObjectIDFromHex(customerID)

	var customer models.Customer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := database.GetCollection("customers").FindOne(ctx, bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customer_id": customer.ID,
		"balance":     customer.Balance,
	})
}
