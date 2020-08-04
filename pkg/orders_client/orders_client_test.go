package orders_client

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/auth"
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
	token := getToken()

	r.GET("/order_details/125183710", GetOrderDetails)

	req, _ := http.NewRequest("GET", "/order_details/125183710", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

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
	token := getToken()

	r.POST("/updateOrderItem", UpdateOrderItem)

	req, _ := http.NewRequest("POST", "/updateOrderItem", bytes.NewBuffer(str))
	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

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
	token := getToken()

	w1 := httptest.NewRecorder()
	r1 := gin.Default()

	r1.GET("/count", OrderCount)
	req1, _ := http.NewRequest("GET", "/count", nil)
	bearer := "Bearer "+token
	req1.Header.Add("Authorization", bearer)
	r1.ServeHTTP(w1, req1)

	resp, err := ioutil.ReadAll(w1.Body)
	if err != nil {

	}

	var c Count
	json.Unmarshal(resp, &c)

	total_order_in_db := 10

	if w1.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w1.Body)
	}
	if c.TotalOrders != total_order_in_db {
		t.Fatalf("failed no of orders do not match")
	}
}

func getToken() string{
	w := httptest.NewRecorder()
	r := gin.Default()

	reqBody, err := json.Marshal(map[string]string{
		"username":"username",
		"password":"password",
	})

	r.POST("/auth/login", auth.Login)
	req, _ := http.NewRequest("POST", "/auth/login",bytes.NewBuffer(reqBody))
	r.ServeHTTP(w, req)

	p, err := ioutil.ReadAll(w.Body)
	p = p[1 : len(p)-1]
	if err != nil {

	}

	return string(p)
}