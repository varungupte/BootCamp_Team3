package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"github.com/varungupte/BootCamp_Team3/pkg/users"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

type Order struct {
	Id int
	Quantity int
	Amount float64
	DishName string
	User users.User
	Restau restaurants.Restaurant
	DeliveryTime string
}

var gJsonData string

func convertToJSON(orders []Order)  {
	jsonData, err := json.Marshal(orders)
	errorutil.CheckError(err, "")

	jsonFile, err := os.Create("./orders.json")
	errorutil.CheckError(err, "")

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	gJsonData = string(jsonData)
}

func GenerateOrdersJSON(filename string) {
	usrs := users.GetUsers("User.csv")
	rests := restaurants.GetRestaurants("Restaurant.csv")

	orderFile, err := os.Open(filename)
	errorutil.CheckError(err, "")

	defer orderFile.Close()

	reader := csv.NewReader(orderFile)
	reader.FieldsPerRecord = -1

	orderData, err := reader.ReadAll()
	errorutil.CheckError(err, "")

	var ord Order
	var orders []Order

	for _, each := range orderData {
		ord.Id,_ = strconv.Atoi(each[0])
		ord.Amount,_ = strconv.ParseFloat((each[1]),32)
		ord.Quantity,_ = strconv.Atoi(each[2])
		ord.DishName= each[3]
		var userid int
		userid,_= strconv.Atoi(each[4])
		ord.User = usrs[userid-1]
		var resid int
		resid,_ = strconv.Atoi(each[5])
		ord.Restau = rests[resid-1]
		ord.DeliveryTime = each[6]
		orders= append(orders, ord)
	}
	convertToJSON(orders)
}

type orders_server struct {

}

func main()  {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}

	s := grpc.NewServer()

	orderspb.RegisterOrdersServiceServer(s, &orders_server{})
	GenerateOrdersJSON("Orders.csv")

	fmt.Println("Orders Server starting...")
	if s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}
}

func (*orders_server) GetOrdersCount(ctx context.Context, req *orderspb.OrdersCountRequest) (*orderspb.OrdersCountResponse, error)  {
	var orders []Order
	err := json.Unmarshal([]byte(gJsonData), &orders)
	if err != nil {
		return nil, err
	}
	res := &orderspb.OrdersCountResponse{
		Count: strconv.Itoa(len(orders)),
	}
	return res, nil
}

func (*orders_server) GetOrderDetail (ctx context.Context, req *orderspb.OrderDetailRequest) (*orderspb.OrderDetailResponse, error) {
	orderNumber:= req.OrderNumber

	parser, _ := gojq.NewStringQuery(gJsonData)
	ord,_ := strconv.Atoi(orderNumber)
	ord = ord-1
	quer := "["+strconv.Itoa(ord)+"]"
	orderDetail, _ := parser.Query(quer)

	result := fmt.Sprint(orderDetail)
	res := &orderspb.OrderDetailResponse{
		OrderDetail: result,
	}
	return res, nil
}