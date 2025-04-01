package main

import (
	"asl-great-mind/core/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// config := configs.Config{
	//     DBUser:     "root",
	//     DBPassword: "asl1234",
	//     DBHost:     "mariadb",
	//     DBPort:     "3306",
	//     DBName:     "greatmind01",
	// }

	// Temporarily skip database connection
	// Uncomment this line when the database is set up
	/*
	   db.ConnectDB(config)
	*/

	r := gin.Default()
	routes.SetupCoreRoutes(r)
	r.Run(":8080")
}
