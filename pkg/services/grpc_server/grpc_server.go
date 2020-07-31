package grpc_server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/elgs/gojq"
	"github.com/varungupte/BootCamp_Team3/pkg/dynamoDB/types"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"github.com/varungupte/BootCamp_Team3/pkg/users"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Order struct {
	Id           int
	Quantity     int
	Amount       float64
	DishName     string
	User         users.User
	Restau       restaurants.Restaurant
	DeliveryTime string
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
}


func getDBInstance() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))
	return dynamodb.New(sess)
}

type GrpcServer struct {}

// GetOrdersCount find the total number of orders in the database.
// It returns the OrderCountResponse and any error encountered.
func (*GrpcServer) GetOrdersCount(ctx context.Context, req *grpcPb.OrdersCountRequest) (*grpcPb.OrdersCountResponse, error)  {
	var orders []Order
	err := json.Unmarshal([]byte(gJsonData), &orders)
	if err != nil {
		return nil, err
	}
	res := &grpcPb.OrdersCountResponse{
		Count: strconv.Itoa(len(orders)),
	}
	return res, nil
}

func (*GrpcServer) PostOrder(ctx context.Context, req *grpcPb.PostOrderRequest) (*grpcPb.PostOrderResponse, error)  {
	// unmarshalling orders
	var orders []Order
	err := json.Unmarshal([]byte(gJsonData), &orders)
	errorutil.CheckError(err, "unmarshalling orders")

	// unmarshalling content
	var orderData2 Order
	err = json.Unmarshal([]byte(req.NewOrder), &orderData2)
	errorutil.CheckError(err, "")

	// appending new order
	orders = append(orders, orderData2)

	convertToJSON(orders)
	res := &grpcPb.PostOrderResponse{
		Status: "SUCCESS: Order updated",
	}
	return res, nil
}

func (*GrpcServer) GetPopularDish(ctx context.Context,req *grpcPb.PopularDishRequest) (*grpcPb.PopularDishResponse, error) {
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
	res := &grpcPb.PopularDishResponse{}
	if maxres == -1 {
		return res, errors.New("City doesn't exist in the database")
	}
	res.DishName = name
	return res, nil
}

func (*GrpcServer) GetOrderDetail (ctx context.Context, req *grpcPb.OrderDetailRequest) (*grpcPb.OrderDetailResponse, error) {
	orderNumber:= req.OrderNumber

	parser, _ := gojq.NewStringQuery(gJsonData)
	ord,_ := strconv.Atoi(orderNumber)
	ord = ord-1
	res := &grpcPb.OrderDetailResponse{}
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
func (*GrpcServer) UpdateDish (ctx context.Context, req *grpcPb.UpdateDishRequest) (*grpcPb.UpdateDishResponse, error) {
	order_id := int(req.GetOrderId())
	updated_dish := req.GetUpdatedDish()
	res := &grpcPb.UpdateDishResponse{
		Status: "SUCCESS: Order updated",
	}

	jsonFilePath := string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/orders.json"
	orderList, err := parseJsonFile(jsonFilePath)
	if err != nil {
		res = &grpcPb.UpdateDishResponse{
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
				res = &grpcPb.UpdateDishResponse{
					Status: "Failed to write in file",
				}
			}
			return res, err
		}
	}
	res = &grpcPb.UpdateDishResponse{
		Status: "FAILURE: No order found with this orderId",
	}
	return res, errors.New("No order found with this orderId")
}

func (*GrpcServer) GetCustomersCount (ctx context.Context, req *grpcPb.CustomersCountRequest) (*grpcPb.CustomersCountResponse, error) {
	db := getDBInstance()
	// create the api params
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("T3_Customer"),
	}
	// get the table description
	resp, err := db.DescribeTable(params)
	if err != nil {
			  return nil, err
			  }
	countResp := &grpcPb.CustomersCountResponse{
		Count: aws.Int64Value(resp.Table.ItemCount),
	}
	return countResp, nil
}

func (*GrpcServer) AddCustomer(ctx context.Context, req *grpcPb.AddCustomerRequest) (*grpcPb.StatusResponse, error)  {
	// unmarshalling content
	var customerData types.Customer
	err := json.Unmarshal([]byte(req.NewCustomer), &customerData)
	errorutil.CheckError(err, "")

	db := getDBInstance()

	customerMap, err := dynamodbattribute.MarshalMap(customerData)
	if err != nil {
		panic("Cannot map the values given in customer struct...")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("T3_Customer"),
		Item: customerMap,
	}

	_, err = db.PutItem(params)

	if err != nil {
		log.Fatalf("Some problem while inserting : %v", err)
	}

	res := &grpcPb.StatusResponse{
		Status: "SUCCESS: New customer added",
	}
	return res, nil
}

func (*GrpcServer) GetCustomer (ctx context.Context, req *grpcPb.CustomerRequest) (*grpcPb.CustomerResponse, error) {
	customerId := req.CustomerId
	resp := &grpcPb.CustomerResponse{}

	db := getDBInstance()

	params := &dynamodb.GetItemInput{
		TableName: aws.String("T3_Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(customerId),
			},
		},
	}

	itemOutput, err := db.GetItem(params)
	if err != nil {
		return resp, err
	}
	if (len(itemOutput.Item) == 0) {
		resp.CustomerData = ""
		return resp, errors.New("FAILURE: Customer not found")
	}
	resp.CustomerData = fmt.Sprintf("%s", itemOutput.Item)
	return resp, nil
}

func (*GrpcServer) DeleteCustomer (ctx context.Context, req *grpcPb.CustomerRequest) (*grpcPb.CustomerResponse, error) {
	customerId := req.CustomerId
	resp := &grpcPb.CustomerResponse{}

	db := getDBInstance()

	// update active status to false
	activeStatus := false

	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("T3_Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(customerId),
			},
		},
		UpdateExpression: aws.String("set ActiveStatus=:as"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue {
			":as": {BOOL: aws.Bool(activeStatus)},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}
	// update the item
	itemOutput, err := db.UpdateItem(params)
	if err != nil {
		return resp, err
	}
	resp.CustomerData = fmt.Sprintf("%s", itemOutput.Attributes)
	return resp, nil
}