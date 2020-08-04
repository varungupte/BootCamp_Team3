package orders_client

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Count struct {
	TotalOrders int `json:"total_orders"`
}

type UpdateOrder struct {
	OrderId uint32
	CustId uint32
	ItemId uint32
	Quantity uint32
}

type UpdateOrderResp struct {
	Status bool
	Message string
}

func TestGetOrderDetails(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	r.GET("/order_details/125183710", GetOrderDetails)

	req, _ := http.NewRequest("GET", "/order_details/125183710", nil)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}

func TestUpdateOrderItem(t *testing.T) {
	var order = UpdateOrder {
		OrderId:1,
		CustId:2,
		ItemId:3,
		Quantity:4,
	}

	str, _ := json.Marshal(order)

	w := httptest.NewRecorder()
	r := gin.Default()

	r.POST("/updateOrderItem", UpdateOrderItem)

	req, _ := http.NewRequest("POST", "/updateOrderItem", bytes.NewBuffer(str))

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}

	var resp UpdateOrderResp

	p, err := ioutil.ReadAll(w.Body)
	if err != nil {

	}
	json.Unmarshal(p, &resp)
	if resp.Status == false {
		t.Fatalf("failed to update")
	}
}

func TestOrderCount(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	r.GET("/count", OrderCount)
	req, _ := http.NewRequest("GET", "/count", nil)
	r.ServeHTTP(w, req)

	p, err := ioutil.ReadAll(w.Body)
	if err != nil {

	}

	var c Count
	json.Unmarshal(p, &c)

	total_order_in_db := 11

	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Body)
	}
	if c.TotalOrders != total_order_in_db {
		t.Fatalf("failed no of orders do not match")
	}
}