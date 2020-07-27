package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"../orderproto"
	"google.golang.org/grpc"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
)

type server struct {

}


func (*server) UpdateDish (ctx context.Context, req *orderproto.UpdateDishRequest) (*orderproto.UpdateDishResponse, error) {

	fmt.Println("Update dish called")

	order_id := int(req.GetOrderId())
	updated_dish := req.GetUpdatedDish()

	message := orders.UpdateOrderDish(order_id, updated_dish)

	res := &orderproto.UpdateDishResponse {
		Message:message,
	}
	return res, nil

}

func main() {
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}

	s:= grpc.NewServer()

	orderproto.RegisterUpdateOrderServiceServer(s, &server{})


	if s.Serve(lis); err != nil  {
		log.Fatalf("failed to Serve %v", err)
	}

}