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

var addrUtil = &grpcPb.Address{
	City:    "Test2",
	HouseNo: "Test3",
	Street:  "Test4",
	PIN:     "Test5",
}
var itemsUtil = []*grpcPb.Item{
	{
		Id:       1,
		Name:     "test6",
		Quantity: 1,
		Cost:     123,
		Cuisine:  "Indian",
	},
	{
		Id:       1,
		Name:     "test7",
		Quantity: 1,
		Cost:     123,
		Cuisine:  "Indian",
	},
}
var restaurantUtil = grpcPb.PostRestaurantRequest{
	Name:              "Test1",
	Id:                101,
	Status:            true,
	RestaurantAddress: addrUtil,
	Items:             itemsUtil,
}

func TestPostRestaurantPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &restaurantUtil
	response, err := oc.PostRestaurant(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	if response.Status != "SuccessFul" {
		t.Error("Faled as didn't Got Successful Message")
	}
}

func TestDeleteItemPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.DeleteItemRequest{
		RestaurantId: 101,
		ItemName:     "test7",
	}
	response, err := oc.DeleteItem(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	if response.Status != "SuccessFul" {
		t.Error("Faled as didn't Got Successful Message")
	}
}
func TestUpdateItemPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.UpdateItemRequest{
		RestaurantId: 101,
		ItemToBeUpdated: &grpcPb.Item{
			Id:       1,
			Name:     "test6",
			Quantity: 1,
			Cost:     1233141,
			Cuisine:  "Indian",
		},
	}
	response, err := oc.UpdateItem(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	if response.Status != "SuccessFul" {
		t.Error("Faled as didn't Got Successful Message")
	}
}
func TestGetCountOfRestaurantPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.OrdersCountRequest{}
	_, err = oc.GetCountOfRestaurant(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}
func TestGetRestaurantPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.RestaurantRequest{
		RestaurantId: 101,
	}
	response, err := oc.GetRestaurant(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	if response.Name != "Test1" {
		t.Error("Faled as didn't Got Successful Message")
	}
}
func TestGetItemsOfRestaurantPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.RestaurantRequest{
		RestaurantId: 101,
	}
	_, err = oc.GetItemsOfRestaurant(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}
func TestGetItemsInRangePass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.ItemsInRangeRequest{
		RestaurantId: 101,
		MaxRange:     1111111110,
		MinRange:     0,
	}
	_, err = oc.GetItemsInRange(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}
func TestDeleteRestaurantPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.RestaurantRequest{
		RestaurantId: 101,
	}
	response, err := oc.DeleteRestaurant(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	if response.Status != "SuccessFul" {
		t.Error("Faled as didn't Got Successful Message")
	}
}

func TestGetCustomersCountPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.CustomersCountRequest{}
	_, err = oc.GetCustomersCount(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGetCustomersPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.CustomerRequest{
		CustomerId: "1",
	}
	_, err = oc.GetCustomer(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestGDeleteCustomerPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.CustomerRequest{
		CustomerId: "2",
	}
	_, err = oc.DeleteCustomer(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}

func TestAddCustomerPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.AddCustomerRequest{}
	_, err = oc.AddCustomer(context.Background(), req)
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
