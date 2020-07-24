package orders

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/restaurants"
	"github.com/varungupte/BootCamp_Team3/pkg/users"
	"io/ioutil"
	"net/http"
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

	jsonFile, err := os.Create("../../pkg/orders/orders.json")
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

func AddOrderPaths(router *gin.Engine) {
	order:= router.Group("/order")
	order.GET("/", HomePage)
	order.GET("/populardish/city/:city", AnalyticsPopularDIsh)
	order.GET("/order_details/:ordernumber",  OrderDetail)
	order.GET("/order_details/",  OrderDetailAll)
	order.POST("/add_order", PostOrder)
	order.POST("/updateOrderDish", UpdateOrderDish)
}

func PostOrder (c *gin.Context) {
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	errorutil.CheckError(err, "Sorry No Content:")

	fmt.Println(string(content))

	//unmarshalling orders
	var orders []Order
	err = json.Unmarshal([]byte(gJsonData), &orders)
	errorutil.CheckError(err, "unmarshalling orders")

	//unmarshalling content
	var orderData2 Order
	err = json.Unmarshal([]byte(content), &orderData2)
	errorutil.CheckError(err, "")

	//appending new order
	orders = append(orders, orderData2)

	// Convert to JSON
	updatedData, err4 := json.Marshal(orders)
	errorutil.CheckError(err4, "")

	gJsonData = string(updatedData)
	c.JSON(http.StatusCreated, gin.H {
		"message" :string(updatedData),
	})
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi there !... This is analytics tool to find popular dish based on various parameters.",
	})
}

func OrderDetailAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"order_details": gJsonData,
	})
}

func OrderDetail (c *gin.Context) {
	ordernumber := c.Param("ordernumber")
	//Using gojq library https://github.com/elgs/gojq#gojq
	parser, _ := gojq.NewStringQuery(gJsonData)
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
	parser, _ := gojq.NewStringQuery(gJsonData)

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

func UpdateOrderDish (c *gin.Context) {
	order_id_str :=  c.DefaultQuery("order_id", "0")
	updated_dish := c.Query("dish")
	order_id, _ := strconv.Atoi(order_id_str)

	jsonFilePath := "../../pkg/orders/orders.json"
	orderList, err := parseJsonFile(jsonFilePath)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Failed to open file",
		})
	}

	for _,order := range orderList {
		if order.Id == order_id {
			prev_dish := order.DishName
			order.DishName = updated_dish
			c.JSON(200, gin.H{
				"message": "Successfully updated",
				"previous": prev_dish,
				"updated_dish": updated_dish,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "No order found with this order_id",
	})
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
