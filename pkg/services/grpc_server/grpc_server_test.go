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

var testItem = &grpcPb.Item{
	Id:21,
	Name:"cake",
	Cuisine:"ind",
	Cost:23.43,
	Quantity:2,
}

var testItems = []*grpcPb.Item{testItem}

var testAddress = &grpcPb.Address{
	HouseNo:"12",
	Street:"mg",
	City:"mumbai",
	PIN:"444002",
}

var testOrder = &grpcPb.CreateOrderRequest{
	ResId:1,
	CustId:2,
	Items:testItems,
	Discount:32.2,
	Address:testAddress,

}

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
	orderNumber = 1937215807
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
		//t.Fatalf("Error was expected due to orderId out of bounds")
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
	resp, err := oc.GetOrdersCount(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrdersCount : %v ", err)
	}

	if resp.Count == 0 {
		t.Fatalf("No orders in DB : %v ", err)
	}
}

func TestGrpcServer_CreateOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := testOrder
	resp, err := oc.CreateOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}

	if resp.Status != true{
		t.Fatalf("Failed to create order : %v ", err)
	}
}

func TestGrpcServer_UpdateOrderItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.UpdateOrderItemRequest{
		OrderId:1937215807,
		CustId:100,
		ItemId:100,
		Quantity:2,
	}
	resp, err := oc.UpdateOrderItem(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}

	if resp.Status != true{
		t.Fatalf("Failed to update order : %v ", err)
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
