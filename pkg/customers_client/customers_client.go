package customers_client

import (
	"context"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/services/grpcPb"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/grpcUtil"
	"github.com/varungupte/BootCamp_Team3/pkg/auth"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
)

// AddOrderPaths adds GET and POST API paths for gin.
func AddCustomerAPIs(router *gin.Engine) {
	customers := router.Group("/customer")

	// gets the number of customers in the database
	customers.POST("/new", AddCustomer)

	// gets the number of customers in the database
	customers.GET("/count", GetCustomerCount)

	// gets the details of all the customers in the database
	//customers.GET("/all", GetAllCustomers)

	// gets the details of a particular customer
	customers.GET("/id/:customerId", GetCustomer)

	// deletes turn the activeStatus of a customer to false
	customers.DELETE("/id/:customerId", DeleteCustomer)
}


// GetCustomerCount is the handler for /customers/count API.
// It displays the number of customers in the database.
func GetCustomerCount(c *gin.Context) {
	err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	grpcClient := grpcPb.NewGRPCServiceClient(conn)

	req := &grpcPb.CustomersCountRequest{}
	res, err := grpcClient.GetCustomersCount(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetCustomersCount: %v", err)
	}

	c.JSON(200, gin.H{
		"Number of Customers" : res.Count,
	})
}

// AddCustomer is the handler for /customers/new API.
// It adds a new customer in the database and displays a success or failure message.
func AddCustomer (c *gin.Context) {
	err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	errorutil.CheckError(err, "Sorry No Content:")

	conn, err := grpc.Dial(grpcUtil.GRPC_target_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		c.JSON(http.StatusBadGateway, gin.H {
			"Error Message" : "Connection failed with gRPC server",
		})
		return
	}
	defer conn.Close()

	oc := grpcPb.NewGRPCServiceClient(conn)

	req := &grpcPb.AddCustomerRequest{
		NewCustomer: string(content),
	}
	resp, err := oc.AddCustomer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetOrderDetail : %v ", err)
		c.JSON (http.StatusForbidden, gin.H{
			"Customer Status": "Issue while updating....",
		})
		return
	}else {
		c.JSON(http.StatusOK, gin.H{
			"Customer Status" : resp.Status,
		})
	}
}

// GetCustomer is the handler for /customers/id/:customerId API.
// It displays the customer details of a particular customerId
func GetCustomer (c *gin.Context) {
	err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	customerId := c.Param("customerId")

	conn, err := grpc.Dial(grpcUtil.GRPC_target_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		return
	}
	defer conn.Close();

	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.CustomerRequest {
		CustomerId: customerId,
	}
	res, err := oc.GetCustomer(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Customer Data" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Customer Data" : res.CustomerData,
	})
}

// DeleteCustomer is the handler for /customers/id/:customerId API.
// It displays the customer details of a particular customerId
func DeleteCustomer (c *gin.Context) {
	err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	customerId := c.Param("customerId")

	conn, err := grpc.Dial(grpcUtil.GRPC_target_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
		return
	}
	defer conn.Close();

	oc := grpcPb.NewGRPCServiceClient(conn)
	req := &grpcPb.CustomerRequest {
		CustomerId: customerId,
	}
	res, err := oc.DeleteCustomer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling DeleteCustomer : %v ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Customer Data" : res.CustomerData,
	})
}