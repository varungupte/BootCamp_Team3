package main
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)
type Order struct {
	Id int
	Quantity int
	Amount int
	DishName string
	Emp Employee
	Restau Restaurant
	DeliveryTime string
}
type Employee struct {
	Id int
	Name string
	Street string
	City string
	Rating int
}
type Restaurant struct {
	Id int
	Name string
	Street string
	City string
	Rating int
}

func getOrderByResId(resId int, orders []Order) Order{
	var orderNil Order

	for _,order := range orders {
		if order.Restau.Id == resId {
			return order
		}
	}

	return orderNil
}

func getInfo() {
	jsonFilePath := "data.json"
	orderJsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("failed to open file")
	}
	defer orderJsonFile.Close()

	byteValue, _ := ioutil.ReadAll(orderJsonFile)

	var orderList []Order
	json.Unmarshal(byteValue, &orderList)

	resId := 18

	order := getOrderByResId(resId, orderList)
	jsonData, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Failed to convert to json")
	}
	fmt.Println(string(jsonData))
}

func main() {

	orderFile, err := os.Open("Order.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer orderFile.Close()

	userFile, err := os.Open("User.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer userFile.Close()

	restaurantFile, err := os.Open("Restaurant.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer restaurantFile.Close()

	reader := csv.NewReader(userFile)
	reader.FieldsPerRecord = -1

	userData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, each := range userData {
		emp.Id,_ = strconv.Atoi(each[0])
		emp.Name = each[1]
		emp.Street= each[2]
		emp.City= each[3]
		emp.Rating,_=strconv.Atoi(each[4])
		employees = append(employees, emp)
	}

	reader = csv.NewReader(restaurantFile)
	reader.FieldsPerRecord = -1

	restaurantData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res Restaurant
	var restaurants []Restaurant

	for _, each := range restaurantData {
		res.Id,_ = strconv.Atoi(each[0])
		res.Name = each[1]
		res.Street= each[2]
		res.City= each[3]
		res.Rating,_=strconv.Atoi(each[4])
		restaurants = append(restaurants, res)
	}

	reader = csv.NewReader(orderFile)
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
		ord.Emp = employees[userid-1]
		var resid int
		resid,_ = strconv.Atoi(each[5])
		ord.Restau = restaurants[resid-1]
		ord.DeliveryTime = each[6]
		orders= append(orders, ord)
	}


	// Convert to JSON
	jsonData, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	getInfo()

}
