package grpc_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/varungupte/BootCamp_Team3/pkg/dynamoDB/types"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"github.com/varungupte/BootCamp_Team3/pkg/services/restaurentService"
	"log"
	"strconv"
)

func getDBInstance() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))
	return dynamodb.New(sess)
}

type GrpcServer struct {}

//func (*GrpcServer) GetPopularDish(ctx context.Context,req *grpcPb.PopularDishRequest) (*grpcPb.PopularDishResponse, error) {
//	//Using gojq library https://github.com/elgs/gojq#gojq
//	parser, _ := gojq.NewStringQuery(gJsonData)
//	cityName := req.CityName
//	//Popular Dish Areawise (In a particular User City, which is the dish maximum ordered)
//	var m = make(map[string]int)
//	for i := 0; i < 1000; i++ {
//		var f string
//		f = "[" + strconv.Itoa(i) + "].User.City"
//		q, _ := parser.Query(f)
//		if q == cityName {
//			var d string
//			d = "[" + strconv.Itoa(i) + "].DishName"
//			dishName, _ := parser.Query(d)
//			m[dishName.(string)] = m[dishName.(string)] + 1
//		}
//	}
//
//	// Iterating map
//	var name string
//	maxres := -1
//	for i, p := range m {
//		if p > maxres {
//			name = i
//			maxres = p
//		}
//	}
//	res := &grpcPb.PopularDishResponse{}
//	if maxres == -1 {
//		return res, errors.New("City doesn't exist in the database")
//	}
//	res.DishName = name
//	return res, nil
//}

func (*GrpcServer) CreateOrder (ctx context.Context, req *grpcPb.CreateOrderRequest) (*grpcPb.CreateOrderResponse, error) {
	var order types.Order
	var Items []types.Item
	var item types.Item

	order.ResId = req.GetResId()
	order.CustId = req.GetCustId()
	order.Discount = req.GetDiscount()
	order.Id = uuid.New().ID()

	order.DeliveryAddr = types.Address {
		HouseNo: req.GetAddress().GetHouseNo(),
		Street:  req.GetAddress().GetStreet(),
		City:    req.GetAddress().GetCity(),
		PIN:     req.GetAddress().GetPin(),
	}

	for _, v := range req.GetItems() {
		item = types.Item{
			Id:       v.GetId(),
			Name:     v.GetName(),
			Cuisine:  v.GetCuisine(),
			Cost:     v.GetCost(),
			Quantity: v.GetQuantity(),
		}
		Items = append(Items, item)
	}
	order.Items = Items

	log.Println("order create ", order)

	db := getDBInstance()

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

	return &grpcPb.CreateOrderResponse{
		Status:  true,
		Message: "fine",
		OrderId: order.Id,
	}, nil
}

func (*GrpcServer) UpdateOrderItem (ctx context.Context, req *grpcPb.UpdateOrderItemRequest) (*grpcPb.UpdateOrderItemResponse, error) {
	orderId := req.GetOrderId()
	itemId := req.GetItemId()
	quantity := req.GetQuantity()
	customerId := req.GetCustId()

	db := getDBInstance()

	params := &dynamodb.GetItemInput{
		TableName:aws.String("OrdersT3"),
		Key:map[string]*dynamodb.AttributeValue{
			"OrderId" :{
				S:aws.String(strconv.Itoa(int(orderId))),
			},
			"CustomerId" :{
				S:aws.String(strconv.Itoa(int(customerId))),
			},
		},
	}
	resp, err := db.GetItem(params)

	//if err != nil {
	//
	//}

	var order = types.Order{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &order)

	//if err != nil {
	//
	//}

	log.Println(order)
	var item types.Item
	j := 0

	for i, v := range order.Items {
		item = v
		if item.Id == itemId {
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

	return &grpcPb.UpdateOrderItemResponse{
		Status:true,
		Message:"finez",
	}, nil
}

func (*GrpcServer) GetOrdersCount(ctx context.Context, req *grpcPb.OrdersCountRequest) (*grpcPb.OrdersCountResponse, error)  {
	db := getDBInstance()

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("OrdersT3"),
	}
	resp, err := db.DescribeTable(params)
	if err != nil {
		return &grpcPb.OrdersCountResponse{
			Count: 0,
		}, nil
	}

	return &grpcPb.OrdersCountResponse{
		Count: int64(*resp.Table.ItemCount),
	}, nil
}

func (*GrpcServer) PostRestaurant(ctx context.Context, req *grpcPb.PostRestaurantRequest) (*grpcPb.GenericResponse, error) {
	//Id           int
	//Name         string
	//Items        []ItemEntity
	//Address      AddressEntity
	//ActiveStatus bool
	restaurant := restaurentService.RestaurantEntity{
		Id:    req.Id,
		Name:  req.Name,
		Items: getItemEntityFromItem(req.Items),
		Address: restaurentService.AddressEntity{
			HouseNo: req.GetRestaurantAddress().HouseNo,
			Street:  req.GetRestaurantAddress().Street,
			City:    req.GetRestaurantAddress().City,
			Pin:     req.GetRestaurantAddress().Pin,
		},
		ActiveStatus: req.GetStatus(),
	}
	res, err := restaurentService.SaveRestaurant(restaurant)
	fmt.Println("Successfully Inserted Restaurant", res)
	if err != nil {
		return &grpcPb.GenericResponse{
			Status:  "404",
			Message: "Unsuccessful",
		}, err
	}
	return &grpcPb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Inserted ", res.Name, " in Table"),
	}, nil
}

func (*GrpcServer) DeleteItem(ctx context.Context, req *grpcPb.DeleteItemRequest) (*grpcPb.GenericResponse, error) {
	err := restaurentService.DeleteItemFromRestaurant(req.RestaurantId, req.ItemName)
	if err != nil {
		return nil, err
	}
	return &grpcPb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Deleted ", req.ItemName, " in Table"),
	}, nil
}

func (*GrpcServer) DeleteRestaurant(ctx context.Context, req *grpcPb.RestaurantRequest) (*grpcPb.GenericResponse, error) {
	err := restaurentService.DeleteRestaurant(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &grpcPb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Deleted ", req.RestaurantId, " from Table"),
	}, nil
}

func (*GrpcServer) UpdateItem(ctx context.Context, req *grpcPb.UpdateItemRequest) (*grpcPb.GenericResponse, error) {
	itemEntity := restaurentService.ItemEntity{
		Name:     req.ItemToBeUpdates.Name,
		Cuisine:  req.ItemToBeUpdates.Cuisine,
		Cost:     req.ItemToBeUpdates.Cost,
		Quantity: req.ItemToBeUpdates.Quantity,
	}
	err := restaurentService.UpdateItemInRestaurant(req.RestaurantId, itemEntity)
	if err != nil {
		return nil, err
	}
	return &grpcPb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Updated  ", req.RestaurantId, " from Table"),
	}, nil
}

func (*GrpcServer) GetCountOfRestaurant(ctx context.Context, req *grpcPb.OrdersCountRequest) (*grpcPb.OrdersCountResponse, error) {
	count, err := restaurentService.GetRestaurantCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("Count of Restaurant",*count)
	return &grpcPb.OrdersCountResponse{
		Count: *count,
	}, nil
}

func (*GrpcServer) GetRestaurant(ctx context.Context, req *grpcPb.RestaurantRequest) (*grpcPb.PostRestaurantRequest, error) {
	restaurant, err := restaurentService.GetRestaurant(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &grpcPb.PostRestaurantRequest{
		Name:   restaurant.Name,
		Status: restaurant.ActiveStatus,
		Id:     restaurant.Id,
		RestaurantAddress: &grpcPb.Address{
			Street:  restaurant.Address.Street,
			HouseNo: restaurant.Address.HouseNo,
			Pin:     restaurant.Address.Pin,
			City:    restaurant.Address.City,
		},
		Items: getItemFromItemEntity(restaurant.Items),
	}, nil
}

func getItemFromItemEntity(itemEntities []restaurentService.ItemEntity) []*grpcPb.Item {
	items := make([]*grpcPb.Item, 0, 5)
	for _, val := range itemEntities {
		temp := &grpcPb.Item{
			Name:     val.Name,
			Cuisine:  val.Cuisine,
			Cost:     val.Cost,
			Quantity: val.Quantity,
		}
		items = append(items, temp)
	}
	return items
}

func (*GrpcServer) GetItemsOfRestaurant(ctx context.Context, req *grpcPb.RestaurantRequest) (*grpcPb.ItemsListResponse, error) {
	items, err := restaurentService.GetRestaurantItems(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &grpcPb.ItemsListResponse{
		Items: getItemFromItemEntity(items),
	}, nil
}

func (*GrpcServer) GetItemsInRange(ctx context.Context, req *grpcPb.ItemsInRangeRequest) (*grpcPb.ItemsListResponse, error) {
	items, err := restaurentService.GetItemsBetweenRange(req.MinRange, req.MaxRange, req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &grpcPb.ItemsListResponse{
		Items: getItemFromItemEntity(items),
	}, nil
}

func getItemEntityFromItem(items []*grpcPb.Item) []restaurentService.ItemEntity {
	itemEntities := make([]restaurentService.ItemEntity, 0, 5)
	for _, val := range items {
		temp := restaurentService.ItemEntity{
			Name:     val.Name,
			Cuisine:  val.Cuisine,
			Cost:     val.Cost,
			Quantity: val.Quantity,
		}
		itemEntities = append(itemEntities, temp)
	}
	return itemEntities
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