package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/auth"
	"github.com/varungupte/BootCamp_Team3/pkg/customers_client"
	"github.com/varungupte/BootCamp_Team3/pkg/orders_client"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants_client"
)

// Main launches the gin server at http://localhost:5657
func main() {
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	orders_client.AddOrderAPIs(router)
	customers_client.AddCustomerAPIs(router)
	restaurants_client.AddRestaurantAPIs(router)
	auth.AddAuthApis(router)

	// listen and serve on localhost:5657
	router.Run("localhost:5657")
}

