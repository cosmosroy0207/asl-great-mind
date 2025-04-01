package main

import (
	"asl-great-mind/auth/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// config := configs.Config{
	// 	DBUser:     "root",
	// 	DBPassword: "asl1234",
	// 	DBHost:     "mariadb",
	// 	DBPort:     "93306",
	// 	DBName:     "greatmind01",
	// 	JWTSecret:  "your-secret-key",
	// }

	// Temporarily skip database connection
	// Uncomment these lines when the database is set up
	/*
	   db.ConnectDB(config)
	   db.DB.AutoMigrate(&models.User{})
	*/

	r := gin.Default()
	routes.SetupAuthRoutes(r)
	r.Run(":8081")
}
