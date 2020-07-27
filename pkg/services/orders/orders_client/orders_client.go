package orders_client

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func AddOrderPaths(router *gin.Engine) {
	order:= router.Group("/order",gin.BasicAuth(gin.Accounts{
		"user1": "gupte",//username:password
		"user2": "gupte",
		"user3": "gupte",
	}))
	order.GET("/", HomePage)
	order.GET("/count", OrderCount)
	order.GET("/populardish/city/:city", AnalyticsPopularDIsh)
	//order.GET("/order_details/order_id/:ordernumber", OrderDetail)
	//order.GET("/order_details/tillorder/:tillorder", OrderDetailAll)
	//order.POST("/add_order", PostOrder)
	//order.POST("/updateOrderDish", UpdateOrderDish)
}


//func PostOrder (c *gin.Context) {
//	body := c.Request.Body
//
//	content, err := ioutil.ReadAll(body)
//	errorutil.CheckError(err, "Sorry No Content:")
//
//	fmt.Println(string(content))
//
//	//unmarshalling orders
//	var orders []Order
//	err = json.Unmarshal([]byte(gJsonData), &orders)
//	errorutil.CheckError(err, "unmarshalling orders")
//
//	//unmarshalling content
//	var orderData2 Order
//	err = json.Unmarshal([]byte(content), &orderData2)
//	errorutil.CheckError(err, "")
//
//	//appending new order
//	orders = append(orders, orderData2)
//
//	// Convert to JSON
//	updatedData, err4 := json.Marshal(orders)
//	errorutil.CheckError(err4, "")
//
//	gJsonData = string(updatedData)
//	c.JSON(http.StatusCreated, gin.H {
//		"message" :string(updatedData),
//	})
//}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi there !... This is analytics tool to find popular dish based on various parameters.",
	})
}

func OrderCount(c *gin.Context) {
	//unmarshalling orders
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrdersServiceClient(conn)

	req := &orderspb.OrdersCountRequest{}
	res, err := oc.GetOrdersCount(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrdersCount: %v", err)
	}

	c.JSON(200, gin.H{
		"Number of orders": res.Count,
	})
}

//func OrderDetailAll(c *gin.Context) {
//	//c.JSON(200, gin.H{
//	//	"order_details": gJsonData,
//	//})
//	tillorder := c.Param("tillorder")
//	//unmarshalling orders
//	var orders []Order
//	err := json.Unmarshal([]byte(gJsonData), &orders)
//	errorutil.CheckError(err, "unmarshalling orders")
//
//	var neworders []Order
//	i,_:= strconv.Atoi(tillorder)
//	for _,v := range orders{
//		if i==0{
//			break
//		}
//		neworders=append(neworders,v)
//		i=i-1
//	}
//	// Convert to JSON
//	updatedData, err4 := json.Marshal(neworders)
//	errorutil.CheckError(err4, "")
//
//	c.JSON(http.StatusOK, gin.H {
//		"message" :string(updatedData),
//	})
//
//
//}
//
//func OrderDetail (c *gin.Context) {
//	ordernumber := c.Param("ordernumber")
//	//Using gojq library https://github.com/elgs/gojq#gojq
//	parser, _ := gojq.NewStringQuery(gJsonData)
//	ord,_ := strconv.Atoi(ordernumber)
//	ord=ord-1
//	quer := "["+strconv.Itoa(ord)+"]"
//	order_detail,_:=parser.Query(quer)
//	fmt.Println(order_detail)
//	c.JSON(http.StatusOK, gin.H{
//		"Order Details":order_detail,
//	})
//}
//
//
func AnalyticsPopularDIsh (c *gin.Context) {
	cityName := c.Param("city")
	conn,err:=grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrdersServiceClient(conn)

	req := &orderspb.City{
		CityName: cityName,
	}
	fmt.Println("Calling grpc Server")
	res, err := oc.GetPopularDish(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetPopularDish: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Dish Name":res.DishName,
		"Most Popular dish in :-" :cityName,
	})
}
//
//func UpdateOrderDish (c *gin.Context) {
//	order_id_str :=  c.DefaultQuery("order_id", "0")
//	updated_dish := c.Query("dish")
//	order_id, _ := strconv.Atoi(order_id_str)
//
//	jsonFilePath := "../../pkg/orders/orders.json"
//	orderList, err := parseJsonFile(jsonFilePath)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"message": "Failed to open file",
//		})
//	}
//
//	for _,order := range orderList {
//		if order.Id == order_id {
//			prev_dish := order.DishName
//			order.DishName = updated_dish
//			c.JSON(200, gin.H{
//				"message": "Successfully updated",
//				"previous": prev_dish,
//				"updated_dish": updated_dish,
//			})
//			return
//		}
//	}
//
//	c.JSON(200, gin.H{
//		"message": "No order found with this order_id",
//	})
//}
//
//func parseJsonFile(jsonFilePath string) ([]Order, error){
//	orderJsonFile, err := os.Open(jsonFilePath)
//	var orderList []Order
//
//	if err != nil {
//		return orderList, err
//	}
//	defer orderJsonFile.Close()
//
//	byteValue, _ := ioutil.ReadAll(orderJsonFile)
//	json.Unmarshal(byteValue, &orderList)
//
//	return orderList, nil
//}
