package orders_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/elgs/gojq"
	"github.com/google/uuid"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)
/*
type Order struct {
	Id int
	Quantity int
	Amount float64
	DishName string
	User users.User
	Restau restaurants.Restaurant
	DeliveryTime string
}
*/

type Item struct {
	Id int32
	Name string
	Cuisine string
	Cost float32
	Quantity int32
}

type Address struct {
	HouseNo int32
	Street string
	City string
	Pin int32
}

type Order struct {
	OrderId string
	ResId string
	CustomerId string
	Items []Item
	Discount float32
	Address Address
}

var gJsonData string
var orderscurrent []Order

func convertToJSON(orders []Order)  {
	jsonData, err := json.Marshal(orders)
	errorutil.CheckError(err, "")

	jsonFile, err := os.Create(string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/orders.json")
	errorutil.CheckError(err, "")

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	gJsonData = string(jsonData)
}

func GenerateOrdersJSON(filename string) {
	/*
	usrs := users.GetUsers(string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/User.csv")
	rests := restaurants.GetRestaurants(string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/Restaurant.csv")

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
	orderscurrent = orders
	convertToJSON(orders)
*/
}


type Orders_server struct {}

func (*Orders_server) PostOrder(ctx context.Context, req *orderspb.PostOrderRequest) (*orderspb.PostOrderResponse, error)  {
	////unmarshalling orders
	var orders []Order
	err := json.Unmarshal([]byte(gJsonData), &orders)
	errorutil.CheckError(err, "unmarshalling orders")

	////unmarshalling content
	var orderData2 Order
	err = json.Unmarshal([]byte(req.Neworder), &orderData2)
	errorutil.CheckError(err, "")

	////appending new order
	orders = append(orders, orderData2)

	convertToJSON(orders)
	res := &orderspb.PostOrderResponse{
		Updatedorders: "SUCCESS: Order updated",
	}
	return res, nil
}

func (*Orders_server) GetPopularDish(ctx context.Context,req *orderspb.PopularDishRequest) (*orderspb.PopularDishResponse, error) {
	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(gJsonData)
	cityName := req.CityName
	//Popular Dish Areawise (In a particular User City, which is the dish maximum ordered)
	var m = make(map[string]int)
	for i := 0; i < 1000; i++ {
		var f string
		f = "[" + strconv.Itoa(i) + "].User.City"
		q, _ := parser.Query(f)
		if q == cityName {
			var d string
			d = "[" + strconv.Itoa(i) + "].DishName"
			dishName, _ := parser.Query(d)
			m[dishName.(string)] = m[dishName.(string)] + 1
		}
	}

	// Iterating map
	var name string
	maxres := -1
	for i, p := range m {
		if p > maxres {
			name = i
			maxres = p
		}
	}
	res := &orderspb.PopularDishResponse{}
	if maxres == -1 {
		return res, errors.New("City doesn't exist in the database")
	}
	res.DishName = name
	return res, nil
}

func (*Orders_server) GetOrderDetail (ctx context.Context, req *orderspb.OrderDetailRequest) (*orderspb.OrderDetailResponse, error) {
	orderNumber:= req.OrderNumber

	parser, _ := gojq.NewStringQuery(gJsonData)
	ord,_ := strconv.Atoi(orderNumber)
	ord = ord-1
	res := &orderspb.OrderDetailResponse{}
	if ord >= 1000 {
		return res, errors.New("OrderId out of bounds")
	}
	quer := "["+strconv.Itoa(ord)+"]"
	orderDetail, _ := parser.Query(quer)
	result := fmt.Sprint(orderDetail)
	res.OrderDetail = result
	return res, nil
}

func parseJsonFile(jsonFilePath string) ([]Order, error){
	orderJsonFile, err := os.Open(jsonFilePath)
	var orderList []Order

	if err != nil {
		return orderList, err
	}
	defer orderJsonFile.Close()

	byteValue, _ := ioutil.ReadAll(orderJsonFile)
	json.Unmarshal(byteValue, &orderList)

	return orderList, nil
}

func writeJsonFile(jsonFilePath string, ordersList []Order) error{
	jsonData, err := json.Marshal(ordersList)
	if err!= nil {
		return err
	}

	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err!= nil {
		return err
	}

	return nil
}

//This function updates the dish of order given order_id it opens file and write to it after update

func (*Orders_server) UpdateDish (ctx context.Context, req *orderspb.UpdateDishRequest) (*orderspb.UpdateDishResponse, error) {
	/*
	order_id := int(req.GetOrderId())
	updated_dish := req.GetUpdatedDish()
	res := &orderspb.UpdateDishResponse{
		Status: "SUCCESS: Order updated",
	}

	jsonFilePath := string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/orders.json"
	orderList, err := parseJsonFile(jsonFilePath)
	if err != nil {
		res = &orderspb.UpdateDishResponse{
			Status: "Failed to write in file",
		}
		return res, err
	}

	for i, order := range orderList {
		if order.Id == order_id {
			log.Println(orderList[i].DishName)
			orderList[i].DishName = updated_dish
			err = writeJsonFile(jsonFilePath, orderList)
			if err != nil {
				res = &orderspb.UpdateDishResponse{
					Status: "Failed to write in file",
				}
			}
			return res, err
		}
	}

	 */
	res := &orderspb.UpdateDishResponse{
		Status: "FAILURE: No order found with this orderId",
	}
	return res, errors.New("No order found with this orderId")
}

// Below are new methods for order apis

func (*Orders_server) CreateOrder (ctx context.Context, req *orderspb.CreateOrderRequest) (*orderspb.CreateOrderResponse, error) {
	var order Order
	var Items []Item
	var item Item

	order.ResId = req.GetResId()
	order.CustomerId = req.GetCustomerId()
	order.Discount = req.GetDiscount()
	order.OrderId = uuid.New().String()

	order.Address = Address {
		HouseNo:req.GetAddress().GetHouseNo(),
		Street:req.GetAddress().GetStreet(),
		City:req.GetAddress().GetCity(),
		Pin:req.GetAddress().GetPin(),
	}

	for _, v := range req.GetItems() {
		item = Item{
			Id:v.GetId(),
			Name:v.GetName(),
			Cuisine:v.GetCuisine(),
			Cost:v.GetCost(),
			Quantity:v.GetQuantity(),
		}
		Items = append(Items, item)
	}
	order.Items = Items

	log.Println("order create ", order)

	db := DB()

	orderMap, err := dynamodbattribute.MarshalMap(order)
	log.Println("mapppp", orderMap)
	if err != nil {
		panic("Cannot map the values given in order struct...")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("OrdersT3"),
		Item: orderMap,
	}
	log.Println("pa", params)

	resp, err := db.PutItem(params)

	if err != nil {
		log.Fatalf("Some problem while inserting : %v", err)
	}

	log.Println(resp)

	return &orderspb.CreateOrderResponse{
		Status:true,
		Message:"fine",
		OrderId:order.OrderId,
	}, nil
}

func (*Orders_server) UpdateOrderItem (ctx context.Context, req *orderspb.UpdateOrderItemRequest) (*orderspb.UpdateOrderItemResponse, error) {
	order_id := req.GetOrderId()
	item_id := req.GetItemId()
	quantity := req.GetQuantity()
	customer_id := req.GetCustomerId()

	db := DB()

	params := &dynamodb.GetItemInput{
		TableName:aws.String("OrdersT3"),
		Key:map[string]*dynamodb.AttributeValue{
			"OrderId" :{
				S:aws.String(order_id),
			},
			"CustomerId" :{
				S:aws.String(customer_id),
			},
		},
	}

	resp, err := db.GetItem(params)

	if err!= nil {

	}

	var order = Order{}
	erre := dynamodbattribute.UnmarshalMap(resp.Item, &order)

	if erre!= nil {

	}

	log.Println(order)
	var item Item
	j := 0

	for i, v := range order.Items {
		item = v
		if item.Id == item_id {
			item.Quantity = quantity
			order.Items[i] = item
			break
		}
	}

	for _, v := range order.Items {
		item = v
		if item.Quantity > 0 {
			order.Items[j] = v
			j++
		}
	}

	order.Items = order.Items[:j]

	orderMap, err := dynamodbattribute.MarshalMap(order)
	param := &dynamodb.PutItemInput{
		TableName: aws.String("OrdersT3"),
		Item: orderMap,
	}

	_, err = db.PutItem(param)
	errorutil.CheckError(err, "Error while updating")

	return &orderspb.UpdateOrderItemResponse{
		Status:true,
		Message:"finez",
	}, nil
}

func (*Orders_server) GetOrdersCount(ctx context.Context, req *orderspb.GetOrdersCountRequest) (*orderspb.GetOrdersCountResponse, error)  {
	db := DB()

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("OrdersT3"),
	}
	resp, err := db.DescribeTable(params)
	if err != nil {
		return &orderspb.GetOrdersCountResponse{
			Count:0,
		}, nil
	}

	return &orderspb.GetOrdersCountResponse{
		Count:int64(*resp.Table.ItemCount),
	}, nil
}

func DB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))

	return dynamodb.New(sess)
}