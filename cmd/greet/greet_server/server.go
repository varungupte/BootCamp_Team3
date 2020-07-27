package main

import (
	"context"
	"fmt"
	"github.com/elgs/gojq"
	"log"
	"net"
	"github.com/varungupte/BootCamp_Team3/cmd/greet/greetpb"
	"google.golang.org/grpc"
	"strconv"
)

type server struct {

}

func (*server) Greet (ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Function called... ")

	orderNumber:= req.GetGreeting().GetOrderNumber()
	totalOrder:= req.GetGreeting().GetTotalOrder()

	parser, _ := gojq.NewStringQuery(totalOrder)
	ord,_ := strconv.Atoi(orderNumber)
	ord=ord-1
	quer := "["+strconv.Itoa(ord)+"]"
	order_detail,_:=parser.Query(quer)

	result := fmt.Sprint(order_detail)

	res := &greetpb.GreetResponse {
		Result: result,
	}
	return res, nil

}



func main() {
	fmt.Println("Hello World From Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}
	s:= grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if s.Serve(lis); err != nil  {
		log.Fatalf("failed to Serve %v", err)
	}

}