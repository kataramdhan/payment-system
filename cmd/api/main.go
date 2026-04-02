package main

import (
	"payment-system/internal/handler"
	"payment-system/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.NewPostgresDB()
	defer db.Close()

	transactionRepo := repository.NewTransactionRepository(db)
	queue := repository.NewRedisClient()
	defer queue.Close()
	transactionHandler := handler.NewTransactionHandler(transactionRepo, queue)
	userRepo := repository.NewUserRepository(db)
	authHandler := handler.NewAuthHandler(userRepo)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	r.GET("/health", handler.HealthCheck)
	//r.POST("/transactions", transactionHandler.Create)

	auth := r.Group("/")
	auth.Use(handler.AuthMiddleware())

	auth.POST("/transactions", transactionHandler.Create)
	auth.GET("/transactions", transactionHandler.List)

	//auth
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}
