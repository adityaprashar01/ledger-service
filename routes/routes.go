package routes

import (
	"ledger-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/customers", controllers.CreateCustomer)
	router.GET("/customers/:customer_id/balance", controllers.GetCustomerBalance)

	router.POST("/transactions", controllers.CreateTransaction)
	router.GET("/customers/:customer_id/transactions", controllers.GetTransactionHistory)
}
