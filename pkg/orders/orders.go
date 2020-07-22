package orders

import (
	"encoding/csv"
	"fmt"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants"
	"github.com/varungupte/BootCamp_Team3/pkg/users"
	"os"
	"strconv"
)

type Order struct {
	Id int
	Quantity int
	Amount int
	DishName string
	User users.User
	Restau restaurants.Restaurant
	DeliveryTime string
}

func GetOrders(filename string) []Order {
	usrs := users.GetUsers("User.csv")
	rests := restaurants.GetRestaurants("Restaurant.csv")

	orderFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer orderFile.Close()

	reader := csv.NewReader(orderFile)
	reader.FieldsPerRecord = -1

	orderData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ord Order
	var orders []Order

	for _, each := range orderData {
		ord.Id,_ = strconv.Atoi(each[0])
		ord.Amount,_ = strconv.Atoi(each[1])
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
	return orders
}