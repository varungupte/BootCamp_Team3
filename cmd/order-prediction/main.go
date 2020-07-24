package main

import (
	"encoding/json"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

var jsonData2 string

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

	jsonData2 = string(jsonData)

	//Popular restaurants(based on number of orders)
	analyticsPopularFood(string(jsonData))

	//Popular Dish Areawise(In a particular User City, which is the dish maximum ordered)
	analyticsPopularDishCitywise(string(jsonData),"SanFrancisco")


	//gin stuff (Popular Dish Areawise)
	router := gin.Default()
	api:= router.Group("/order")
	//api.GET("/",  HomePage)
	api.GET("/populardish/city/:city",AnalyticsPopularDIsh)
	api.GET("/order_details/:ordernumber",  OrderDetail)
	router.Run("localhost:5656")

}



func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi there !... This is analytics tool to find popular dish based on various parameters.",
	})
}


func OrderDetail (c *gin.Context) {
	ordernumber := c.Param("ordernumber")
	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(jsonData2)
	ord,_ := strconv.Atoi(ordernumber)
	ord=ord-1
	quer := "["+strconv.Itoa(ord)+"]"
	order_detail,_:=parser.Query(quer)
	fmt.Println(order_detail)
	c.JSON(http.StatusOK, gin.H{
		"Order Details":order_detail,
	})
}


func AnalyticsPopularDIsh (c *gin.Context) {
	cityName := c.Param("city")
	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(jsonData2)


	//Popular Dish Areawise(In a particular User City, which is the dish maximum ordered)
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

	c.JSON(http.StatusOK, gin.H{
		"Dish Name":res,
		"Most Popular dish in :-" :cityName,
	})
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


	//Popular Dish Areawise(In a particular User City, which is the dish maximum ordered)
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
