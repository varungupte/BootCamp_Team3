package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
)

func main() {
	orders.GenerateOrdersJSON("Order.csv")

	router := gin.Default()
	orders.AddOrderPaths(router)
	router.Run("localhost:5656")
}

