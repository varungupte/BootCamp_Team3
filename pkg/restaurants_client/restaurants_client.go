package restaurants_client

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

// AddRestaurantAPIs adds GET and POST APIs for gin.
func AddRestaurantAPIs(router *gin.Engine) {
	restaurant := router.Group("/restaurant")
	itemRouter := router.Group("/item")

	restaurant.GET ("/:id", GetRestaurant)
	restaurant.POST("/", CreateRestaurant)
	restaurant.DELETE("/:id", DeleteRestaurant)

	itemRouter.DELETE("/:restId/:itemName", DeleteRestaurantItem)
	itemRouter.PUT("/:id", UpdateRestaurantItem)

	router.GET("/count/restaurant", GetRestaurantCount)
	restaurant.GET("/:id/item", GetRestaurantItems)

	//https://example.org/?a=1&a=2&b=&=3&&&&
	restaurant.GET("/:id/items/:priceMin/:priceMax", GetRestaurantItemsInRange)
}

func GetRestaurant(c *gin.Context) {
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
}

func CreateRestaurant(c *gin.Context) {
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

}

func DeleteRestaurant (c *gin.Context) {
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
}

func DeleteRestaurantItem (c *gin.Context) {
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
}

func UpdateRestaurantItem(c *gin.Context) {
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
		ItemToBeUpdated: &item,
	}
	res, err := oc.UpdateItem(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
}

func GetRestaurantCount(c *gin.Context) {
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
}

func GetRestaurantItems(c *gin.Context) {
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
		"status":  "Successful",
		"message": res.Items,
	})
}

func GetRestaurantItemsInRange(c *gin.Context) {
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
		"status":  "Successful",
		"message": res.Items,
	})
}