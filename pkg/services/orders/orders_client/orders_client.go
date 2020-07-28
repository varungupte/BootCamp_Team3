package orders_client

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orderspb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// AddOrderPaths adds GET and POST API paths for gin.
func AddOrderPaths(router *gin.Engine) {
	
	//BasicAuth module for authorization while hitting the api
	order:= router.Group("/order",gin.BasicAuth(gin.Accounts{
		"user1": "gupte",//username:password
		"user2": "gupte",
		"user3": "gupte",
	}))
	order.GET("/", HomePage)
	
	//will return the total orders so far
	order.GET("/count", OrderCount)
	
	//Popular Dish Areawise(In a particular User City, which is the dish maximum ordered)
	order.GET("/populardish/city/:city", PopularDish)
	
	// will return json object containing info about the order with orderid "ordernumber"
	order.GET("/order_details/order_id/:ordernumber", OrderDetail)
	
	// will return slice for orders till orderid "tillorder"
    // order.GET("/order_details/tillorder/:tillorder", OrderDetailAll)
	
	// POST request to add orders
    order.POST("/add_order", PostOrder)
	
	order.POST("/updateOrderDish", UpdateOrderDish)
}

// OrderDetail is the handler for /order_details/order_id/:ordernumber API.
// It displays the order detail by a particular orderId
func OrderDetail (c *gin.Context) {
	ordernumber := c.Param("ordernumber")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.OrderDetailRequest{
		OrderNumber: ordernumber,
	}
	res, err := oc.GetOrderDetail(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Order Details" : res.OrderDetail,
	})
}

// PostOrder adds a new order to the database.
// It displays a success or failure message.
func PostOrder (c *gin.Context) {
	body := c.Request.Body

 	content, err := ioutil.ReadAll(body)
 	errorutil.CheckError(err, "Sorry No Content:")

 	fmt.Println(string(content))

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := orderspb.NewOrdersServiceClient(conn)

	req := &orderspb.PostOrderRequest{
		Neworder: string(content),
	}
	res, err := oc.PostOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		c.JSON(http.StatusOK, gin.H{
			"Order Status": "Issue while updating....",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"Order Status" : "Post call successfully executed.",
		})
	}
	fmt.Println(res.Updatedorders)
}

// HomePage is the handler for /order API.
// It displays an introductory message.
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi there !... This is analytics tool to find popular dish based on various parameters.",
	})
}

// OrderCount is the handler for /order/count API.
// It displays the total number of orders in the database.
func OrderCount(c *gin.Context) {
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

// func OrderDetailAll(c *gin.Context) {
// 	//c.JSON(200, gin.H{
// 	//	"order_details": gJsonData,
// 	//})
// 	tillorder := c.Param("tillorder")
// 	//unmarshalling orders
// 	var orders []Order
// 	err := json.Unmarshal([]byte(gJsonData), &orders)
// 	errorutil.CheckError(err, "unmarshalling orders")

// 	var neworders []Order
// 	i,_:= strconv.Atoi(tillorder)
// 	for _,v := range orders{
// 		if i==0{
// 			break
// 		}
// 		neworders=append(neworders,v)
// 		i=i-1
// 	}
// 	// Convert to JSON
// 	updatedData, err4 := json.Marshal(neworders)
// 	errorutil.CheckError(err4, "")

// 	c.JSON(http.StatusOK, gin.H {
// 		"message" :string(updatedData),
// 	})
// }

// PopularDish is the handler for /populardish/city/:city API.
// It displays the most popular dish of a particular city.
func PopularDish (c *gin.Context) {
	cityName := c.Param("city")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.PopularDishRequest{
		CityName: cityName,
	}
	res, err := oc.GetPopularDish(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetPopularDish: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Dish Name": res.DishName,
		"Most Popular dish in :-" : cityName,
	})
}

// UpdateOrderDish is the handler for /updateOrderDish API.
// It updates the dish name of a particular order.
func UpdateOrderDish (c *gin.Context) {
	//orderIdStr :=  c.DefaultQuery("order_id", "0")
	orderIdStr := c.DefaultPostForm("order_id", "0")
	//updatedDish := c.Query("dish")
	updatedDish := c.PostForm("dish")
	orderId, _ := strconv.Atoi(orderIdStr)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrdersServiceClient(conn)
	req := &orderspb.UpdateDishRequest{
		OrderId: int64(orderId),
		UpdatedDish: updatedDish,
	}
	res, err := oc.UpdateDish(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(200, gin.H{
			"Status": res.Status,
		})
		return
	}
	c.JSON(200, gin.H{
		"Status":      res.Status,
		"UpdatedDish": updatedDish,
	})
}

