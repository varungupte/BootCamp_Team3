package orders_server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"github.com/varungupte/BootCamp_Team3/pkg/services/restaurentService"
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

func convertToJSON(orders []Order) {
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
		ord.Id, _ = strconv.Atoi(each[0])
		ord.Amount, _ = strconv.ParseFloat((each[1]), 32)
		ord.Quantity, _ = strconv.Atoi(each[2])
		ord.DishName = each[3]
		var userid int
		userid, _ = strconv.Atoi(each[4])
		ord.User = usrs[userid-1]
		var resid int
		resid, _ = strconv.Atoi(each[5])
		ord.Restau = rests[resid-1]
		ord.DeliveryTime = each[6]
		orders = append(orders, ord)
	}
	orderscurrent = orders
	convertToJSON(orders)
}

type Orders_server struct{}

// GetOrdersCount find the total number of orders in the database.
// It returns the OrderCountResponse and any error encountered.
func (*Orders_server) GetOrdersCount(ctx context.Context, req *orderspb.OrdersCountRequest) (*orderspb.OrdersCountResponse, error) {
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

func (*Orders_server) PostOrder(ctx context.Context, req *orderspb.PostOrderRequest) (*orderspb.PostOrderResponse, error) {
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

func (*Orders_server) GetPopularDish(ctx context.Context, req *orderspb.PopularDishRequest) (*orderspb.PopularDishResponse, error) {
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

func (*Orders_server) GetOrderDetail(ctx context.Context, req *orderspb.OrderDetailRequest) (*orderspb.OrderDetailResponse, error) {
	orderNumber := req.OrderNumber

	parser, _ := gojq.NewStringQuery(gJsonData)
	ord, _ := strconv.Atoi(orderNumber)
	ord = ord - 1
	res := &orderspb.OrderDetailResponse{}
	if ord >= 1000 {
		return res, errors.New("OrderId out of bounds")
	}
	quer := "[" + strconv.Itoa(ord) + "]"
	orderDetail, _ := parser.Query(quer)
	result := fmt.Sprint(orderDetail)
	res.OrderDetail = result
	return res, nil
}

func parseJsonFile(jsonFilePath string) ([]Order, error) {
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

func writeJsonFile(jsonFilePath string, ordersList []Order) error {
	jsonData, err := json.Marshal(ordersList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
func (*Orders_server) PostRestaurant(ctx context.Context, req *orderspb.PostRestaurantRequest) (*orderspb.GenericResponse, error) {
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
			HouseNo: req.GetRestaurantAddress().HounseNo,
			Street:  req.GetRestaurantAddress().Street,
			City:    req.GetRestaurantAddress().City,
			Pin:     req.GetRestaurantAddress().Pin,
		},
		ActiveStatus: req.GetStatus(),
	}
	res, err := restaurentService.SaveRestaurant(restaurant)
	fmt.Println("Successfully Inserted Restaurant", res)
	if err != nil {
		return &orderspb.GenericResponse{
			Status:  "404",
			Message: "Unsuccessful",
		}, err
	}
	return &orderspb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Inserted ", res.Name, " in Table"),
	}, nil
}
func (*Orders_server) DeleteItem(ctx context.Context, req *orderspb.DeleteItemRequest) (*orderspb.GenericResponse, error) {
	err := restaurentService.DeleteItemFromRestaurant(req.RestaurantId, req.ItemName)
	if err != nil {
		return nil, err
	}
	return &orderspb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Deleted ", req.ItemName, " in Table"),
	}, nil
}
func (*Orders_server) DeleteRestaurant(ctx context.Context, req *orderspb.RestaurantRequest) (*orderspb.GenericResponse, error) {
	err := restaurentService.DeleteRestaurant(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &orderspb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Deleted ", req.RestaurantId, " from Table"),
	}, nil
}
func (*Orders_server) UpdateItem(ctx context.Context, req *orderspb.UpdateItemRequest) (*orderspb.GenericResponse, error) {
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
	return &orderspb.GenericResponse{
		Status:  "SuccessFul",
		Message: fmt.Sprint("Updated  ", req.RestaurantId, " from Table"),
	}, nil
}
func (*Orders_server) GetCountOfRestaurant(ctx context.Context, req *orderspb.OrdersCountRequest) (*orderspb.OrdersCountResponse, error) {
	count, err := restaurentService.GetRestaurantCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("Count of Restaurant",*count)
	return &orderspb.OrdersCountResponse{
		Count: fmt.Sprint(*count),
	}, nil
}
func (*Orders_server) GetRestaurant(ctx context.Context, req *orderspb.RestaurantRequest) (*orderspb.PostRestaurantRequest, error) {
	restaurant, err := restaurentService.GetRestaurant(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &orderspb.PostRestaurantRequest{
		Name:   restaurant.Name,
		Status: restaurant.ActiveStatus,
		Id:     restaurant.Id,
		RestaurantAddress: &orderspb.Address{
			Street:   restaurant.Address.Street,
			HounseNo: restaurant.Address.HouseNo,
			Pin:      restaurant.Address.Pin,
			City:     restaurant.Address.City,
		},
		Items: getItemFromItemEntity(restaurant.Items),
	}, nil
}

func getItemFromItemEntity(itemEntities []restaurentService.ItemEntity) []*orderspb.Item {
	items := make([]*orderspb.Item, 0, 5)
	for _, val := range itemEntities {
		temp := &orderspb.Item{
			Name:     val.Name,
			Cuisine:  val.Cuisine,
			Cost:     val.Cost,
			Quantity: val.Quantity,
		}
		items = append(items, temp)
	}
	return items
}
func (*Orders_server) GetItemsOfRestaurant(ctx context.Context, req *orderspb.RestaurantRequest) (*orderspb.ItemsListResponse, error) {
	items, err := restaurentService.GetRestaurantItems(req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &orderspb.ItemsListResponse{
		Items: getItemFromItemEntity(items),
	}, nil
}
func (*Orders_server) GetItemsInRange(ctx context.Context, req *orderspb.ItemsInRangeRequest) (*orderspb.ItemsListResponse, error) {
	items, err := restaurentService.GetItemsBetweenRange(req.MinRange, req.MaxRange, req.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &orderspb.ItemsListResponse{
		Items: getItemFromItemEntity(items),
	}, nil
}
func getItemEntityFromItem(items []*orderspb.Item) []restaurentService.ItemEntity {
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

//This function updates the dish of order given order_id it opens file and write to it after update
func (*Orders_server) UpdateDish(ctx context.Context, req *orderspb.UpdateDishRequest) (*orderspb.UpdateDishResponse, error) {
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
	res = &orderspb.UpdateDishResponse{
		Status: "FAILURE: No order found with this orderId",
	}
	return res, errors.New("No order found with this orderId")
}
