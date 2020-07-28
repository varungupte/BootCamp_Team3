package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_client"
)

// main launches the gin server at http://localhost:5657
func main() {
	gin.ForceConsoleColor()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	orders_client.AddOrderPaths(router)
	// listen and serve on localhost:5657
	router.Run("localhost:5657")
}

