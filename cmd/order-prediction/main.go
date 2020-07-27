package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_client"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	orders_client.AddOrderPaths(router)
	router.Run("localhost:5656")
}

