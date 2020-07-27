package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
	"google.golang.org/grpc"
	"log"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_client"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();
	orders.AddOrderPaths(router,conn)
	orders_client.AddOrderPaths(router)
	router.Run("localhost:5656")
}

