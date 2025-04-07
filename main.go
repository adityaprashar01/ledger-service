// @title Ledger Service API
// @version 1.0
// @description Backend service to manage customer ledger with transaction support.
// @host ledger-service.onrender.com
// @BasePath /
// @schemes https

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ledger-service/database"
	_ "ledger-service/docs" // Swagger docs
	"ledger-service/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing with system environment variables")
	}

	database.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
