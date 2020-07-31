package main

import (
	"fmt"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpc_server"
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

	grpcPb.RegisterGRPCServiceServer(s, &grpc_server.GrpcServer{})

	fmt.Println("Orders Server starting...")
	if s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}
}

