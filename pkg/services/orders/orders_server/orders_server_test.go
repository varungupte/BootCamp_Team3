package orders_server_test

import (
	"context"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_server"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
)

const bufSize = 1024 * 1024
var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	orderspb.RegisterOrdersServiceServer(s, &orders_server.Orders_server{})
	orders_server.GenerateOrdersJSON(string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/Orders.csv")
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
	orderNumber := "20"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.OrderDetailRequest{
		OrderNumber: orderNumber,
	}
	_, err = oc.GetOrderDetail(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetOrderDetailFail(t *testing.T) {
	orderNumber := "1001"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.OrderDetailRequest{
		OrderNumber: orderNumber,
	}
	_, err = oc.GetOrderDetail(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetOrdersCountPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.OrdersCountRequest{}
	_, err = oc.GetOrdersCount(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetPopularDishPass(t *testing.T) {
	cityName := "SanFrancisco"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.PopularDishRequest{
		CityName: cityName,
	}
	_, err = oc.GetPopularDish(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetPopularDishFail(t *testing.T) {
	cityName := "UnknownCity"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.PopularDishRequest{
		CityName: cityName,
	}
	_, err = oc.GetPopularDish(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestUpdateDishPass(t *testing.T) {
	orderIdStr := "20"
	updatedDish := "Pizza"
	orderId, _ := strconv.Atoi(orderIdStr)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.UpdateDishRequest{
		OrderId: int64(orderId),
		UpdatedDish: updatedDish,
	}
	_, err = oc.UpdateDish(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling UpdateDish : %v ", err)
	}
}

func TestUpdateDishFail(t *testing.T) {
	orderIdStr := "1001"
	updatedDish := "Pizza"
	orderId, _ := strconv.Atoi(orderIdStr)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.UpdateDishRequest{
		OrderId: int64(orderId),
		UpdatedDish: updatedDish,
	}
	_, err = oc.UpdateDish(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling UpdateDish : %v ", err)
	}
}
