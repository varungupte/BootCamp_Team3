package orders_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// AddOrderAPIs adds GET and POST API paths for gin.
func AddOrderAPIs(router *gin.Engine) {
	//BasicAuth module for authorization while hitting the api
	order := router.Group("/order", gin.BasicAuth(gin.Accounts{
		"user1": "gupte", //username:password
		"user2": "gupte",
		"user3": "gupte",
	}))
	order.GET("/", HomePage)

	//will return the total orders so far
	order.GET("/count", OrderCount)

	//Popular Dish Areawise(In a particular User City, which is the dish maximum ordered)
	//order.GET("/populardish/city/:city", PopularDish)

	// will return json object containing info about the order with orderid "ordernumber"
	//order.GET("/order_details/order_id/:ordernumber", OrderDetail)

	// will return slice for orders till orderid "tillorder"
	// order.GET("/order_details/tillorder/:tillorder", OrderDetailAll)

	// POST request to add orders
    //order.POST("/add_order", PostOrder)

	order.POST("/updateOrderItem", UpdateOrderItem)

	order.POST("/createOrder", CreateOrder)

	restaurant := router.Group("/restaurant")
	itemRouter:=router.Group("/item")
	restaurant.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		n, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.RestaurantRequest{
			RestaurantId: n,
		}
		res, err := oc.GetRestaurant(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"Restaurant Details": res,
		})
	})
	restaurant.POST("/", func(c *gin.Context) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()
		body := c.Request.Body
		content, err := ioutil.ReadAll(body)
		var restaurant grpcPb.PostRestaurantRequest
		err = json.Unmarshal(content, &restaurant)
		errorutil.CheckError(err, "unmarshalling orders")
		oc := grpcPb.NewGRPCServiceClient(conn)
		res, err := oc.PostRestaurant(context.Background(), &restaurant)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})

	})
	restaurant.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		n, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.RestaurantRequest{
			RestaurantId: n,
		}
		res, err := oc.DeleteRestaurant(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
	})
	itemRouter.DELETE("/:restId/:itemName", func(c *gin.Context) {
		Id1 := c.Param("restId")
		itemName := c.Param("itemName")
		resId, err := strconv.ParseInt(Id1, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.DeleteItemRequest{
			RestaurantId: resId,
			ItemName:     itemName,
		}
		res, err := oc.DeleteItem(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
	})
	itemRouter.PUT("/:id", func(c *gin.Context) {
		body := c.Request.Body
		content, err := ioutil.ReadAll(body)
		id := c.Param("id")
		n, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		var item grpcPb.Item
		err = json.Unmarshal(content, &item)
		errorutil.CheckError(err, "unmarshalling orders")
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.UpdateItemRequest{
			RestaurantId:    n,
			ItemToBeUpdates: &item,
		}
		res, err := oc.UpdateItem(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
	})
	router.GET("/count/restaurant", func(c *gin.Context) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.OrdersCountRequest{}
		res, err := oc.GetCountOfRestaurant(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"count": res.Count,
		})
	})
	restaurant.GET("/:id/item", func(c *gin.Context) {
		id := c.Param("id")
		n, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		req := &grpcPb.RestaurantRequest{
			RestaurantId: n,
		}
		res, err := oc.GetItemsOfRestaurant(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Successfull",
			"message": res.Items,
		})
	})
	//https://example.org/?a=1&a=2&b=&=3&&&&
	restaurant.GET("/:id/items/:priceMin/:priceMax", func(c *gin.Context) {
		id := c.Param("id")
		n, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			fmt.Println("Unable to convert string to int 64")
		}
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		}
		defer conn.Close()

		oc := grpcPb.NewGRPCServiceClient(conn)
		value, err := strconv.ParseFloat(c.Param("priceMin"), 32)
		if err != nil {
			// do something sensible
		}
		min := float32(value)
		value, err = strconv.ParseFloat(c.Param("priceMax"), 32)
		if err != nil {
			// do something sensible
		}
		max := float32(value)
		req := &grpcPb.ItemsInRangeRequest{
			RestaurantId: n,
			MinRange:     min,
			MaxRange:     max,
		}
		res, err := oc.GetItemsInRange(context.Background(), req)
		if err != nil {
			log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Successfull",
			"message": res.Items,
		})
	})
}

// OrderDetail is the handler for /order_details/order_id/:ordernumber API.
// It displays the order detail by a particular orderId
/*
func OrderDetail(c *gin.Context) {
	ordernumber := c.Param("ordernumber")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.OrderDetailRequest {
		OrderNumber: ordernumber,
	}
	res, err := oc.GetOrderDetail(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Order Details": res.OrderDetail,
	})

}


 */

// PostOrder adds a new order to the database.
// It displays a success or failure message.
/*
func PostOrder(c *gin.Context) {
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	errorutil.CheckError(err, "Sorry No Content:")

	fmt.Println(string(content))

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"Error Message": "Connection failed with gRPC server",
		})
		return
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)

	req := &grpcPb.PostOrderRequest{
		NewOrder: string(content),
	}
	res, err := oc.PostOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		c.JSON(http.StatusForbidden, gin.H{
			"Order Status": "Issue while updating....",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Order Status" : res.Status,
		})
	}
	fmt.Println(res.Status)
}

 */

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
	log.Println("ghhh")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)

	req := &grpcPb.OrdersCountRequest{}
	res, err := oc.GetOrdersCount(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrdersCount: %v", err)
	}

	c.JSON(200, gin.H{
		"total_orders": res.Count,
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
//this function is deleted so commenting
/*
func PopularDish(c *gin.Context) {
	cityName := c.Param("city")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.PopularDishRequest{
		CityName: cityName,
	}
	res, err := oc.GetPopularDish(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetPopularDish: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Dish Name":               res.DishName,
		"Most Popular dish in :-": cityName,
	})
}


 */
func UpdateOrderItem (c *gin.Context) {
	var req grpcPb.UpdateOrderItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(req)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)

	res, err := oc.UpdateOrderItem(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(200, gin.H{
			"Status": res.Status,
		})
		return
	}

	c.JSON(200, gin.H{
		"Status":res.Status,
		"Message":res.Message,
	})
}

func CreateOrder (c *gin.Context) {
	var req grpcPb.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := grpcPb.NewGRPCServiceClient(conn)

	res, err := oc.CreateOrder(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling CreateOrder : %v ", err)
		c.JSON(200, gin.H{
			"Status": res.Status,
		})
		return
	}

	c.JSON(200, gin.H{
		"Status":res.Status,
		"Message":res.Message,
		"OrderId":res.OrderId,
	})
}
