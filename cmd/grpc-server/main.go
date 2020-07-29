package main

import (
	"fmt"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_server"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}
	s := grpc.NewServer()
	orderspb.RegisterOrdersServiceServer(s, &orders_server.OrdersServer{})
	orders_server.GenerateOrdersJSON("Orders.csv")
	fmt.Println("Orders Server starting...")
	if s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}
}
