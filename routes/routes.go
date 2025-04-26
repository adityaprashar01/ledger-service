package routes

import (
	"ledger-service/controllers"
	"ledger-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
	}

	customers := r.Group("/customers")
	customers.Use(middleware.AuthMiddleware()) // protect customers routes
	{
		customers.POST("", controllers.CreateCustomer)
		customers.GET("/:customer_id/balance", controllers.GetCustomerBalance)
		customers.GET("/:customer_id/transactions", controllers.GetTransactionHistory)
	}

	r.POST("/transactions", middleware.AuthMiddleware(), controllers.CreateTransaction)
}
