package grpc_server_test

import (
	"context"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpc_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024
var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	grpcPb.RegisterGRPCServiceServer(s, &grpc_server.GrpcServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetOrderDetailPass(t *testing.T) {
	var orderNumber uint32
	orderNumber = 20
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.OrderDetailsRequest{
		OrderId: orderNumber,
	}
	_, err = oc.GetOrderDetails(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetOrderDetailFail(t *testing.T) {
	var orderNumber uint32
	orderNumber = 1001
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.OrderDetailsRequest{
		OrderId: orderNumber,
	}
	_, err = oc.GetOrderDetails(context.Background(), req)
	if err == nil {
		t.Fatalf("Error was expected due to orderId out of bounds")
	}
}

func TestGetOrdersCountPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.OrdersCountRequest{}
	_, err = oc.GetOrdersCount(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

//func TestGetPopularDishPass(t *testing.T) {
//	cityName := "SanFrancisco"
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//	oc := grpcPb.NewGRPCServiceClient(conn)
//	req := &grpcPb.PopularDishRequest{
//		CityName: cityName,
//	}
//	_, err = oc.GetPopularDish(context.Background(), req)
//	if err != nil {
//		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
//	}
//}

//func TestGetPopularDishFail(t *testing.T) {
//	cityName := "UnknownCity"
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//	oc := grpcPb.NewGRPCServiceClient(conn)
//	req := &grpcPb.PopularDishRequest{
//		CityName: cityName,
//	}
//	_, err = oc.GetPopularDish(context.Background(), req)
//	if err == nil {
//		t.Fatalf("Error was expected as an unknow city is entered")
//	}
//}
