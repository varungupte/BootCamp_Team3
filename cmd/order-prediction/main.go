package main

import (
	"encoding/json"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
	"os"
	"strconv"
)

func main() {
	ordrs := orders.GetOrders("Order.csv")

	// Convert to JSON
	jsonData, err := json.Marshal(ordrs)
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

	analyticsPopularFood(string(jsonData))
	analyticsPopularDishCitywise(string(jsonData),"San Francisco")
}

func analyticsPopularFood(jsonData string){

	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(jsonData)


	//Popular restaurants(based on number of orders)
	var m = make(map[string]int)
	for i := 0; i < 1000; i++ {
		var f string
		f = "["+strconv.Itoa(i)+"].Restau.Name"
		q,_:=parser.Query(f)
		m[q.(string)]=m[q.(string)]+1
	}
	// Iterating map
	var res string
	maxres:=-1
	for i, p := range m {
		if p > maxres{
			res = i
			maxres = p
		}
	}
	fmt.Println("The most popular restaurant is:-",res)
	fmt.Println(maxres," times, food was order from here.")
}

func analyticsPopularDishCitywise(jsonData string,cityName string){

	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(jsonData)


	//Popular Dish Areawise(In a particular User City, which is the dish maximum order)
	var m = make(map[string]int)
	for i := 0; i < 1000; i++ {
		var f string
		f = "["+strconv.Itoa(i)+"].User.City"
		q,_:=parser.Query(f)
		if q==cityName{
			var d string
			d = "["+strconv.Itoa(i)+"].DishName"
			dishName,_:=parser.Query(d)
			m[dishName.(string)]=m[dishName.(string)]+1
		}

	}
	// Iterating map
	var res string
	maxres:=-1
	for i, p := range m {
		if p > maxres{
			res = i
			maxres = p
		}
	}
	fmt.Println("The most popular dish in ",cityName," is:-")
	fmt.Println(res)
}
