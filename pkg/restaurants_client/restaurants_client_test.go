package restaurants_client

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/varungupte/BootCamp_Team3/pkg/auth"
	"github.com/varungupte/BootCamp_Team3/pkg/services/grpcPb"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRestaurant(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET ("/restaurant/1", GetRestaurant)
	req, _ := http.NewRequest("GET", "/restaurant/1", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}

func TestDeleteRestaurant(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.DELETE ("/restaurant/1", DeleteRestaurant)
	req, _ := http.NewRequest("DELETE", "/restaurant/1", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
func TestGetRestaurantCount(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET("/count/restaurant", GetRestaurantCount)
	req, _ := http.NewRequest("GET", "/count/restaurant", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}

func TestDeleteRestaurantItem(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.DELETE ("/restaurant/1/Coffee", DeleteRestaurant)
	req, _ := http.NewRequest("DELETE", "/restaurant/1/Coffee", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
func TestGetRestaurantItemsInRange(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET("/restaurant/3/items/0/100", GetRestaurantItemsInRange)
	req, _ := http.NewRequest("GET", "/restaurant/3/items/0/100", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
var addrUtil = &grpcPb.Address{
	City:    "Test2",
	HouseNo: "Test3",
	Street:  "Test4",
	PIN:     "Test5",
}
var itemsUtil = []*grpcPb.Item{
	{
		Id:       1,
		Name:     "test6",
		Quantity: 1,
		Cost:     123,
		Cuisine:  "Indian",
	},
	{
		Id:       1,
		Name:     "test7",
		Quantity: 1,
		Cost:     123,
		Cuisine:  "Indian",
	},
}
var restaurantUtil = grpcPb.PostRestaurantRequest{
	Name:              "Test1",
	Id:                101,
	Status:            true,
	RestaurantAddress: addrUtil,
	Items:             itemsUtil,
}

func TestCreateRestaurant(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	str, _ := json.Marshal(restaurantUtil)
	token := getToken()
	r.POST("/restaurant/", GetRestaurantItemsInRange)
	req, _ := http.NewRequest("POST", "/restaurant/",bytes.NewBuffer(str) )

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
func TestGetRestaurantItems(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET("/restaurant/3/items", GetRestaurantItemsInRange)
	req, _ := http.NewRequest("GET", "/restaurant/3/items", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
func TestUpdateRestaurantItem(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.Default()
	item:=grpcPb.Item{
		Id:       1,
		Name:     "test6",
		Quantity: 1,
		Cost:     123,
		Cuisine:  "Indian",
	}
	str, _ := json.Marshal(item)
	token := getToken()
	r.PUT("/restaurant/1", UpdateRestaurantItem)
	req, _ := http.NewRequest("PUT", "/restaurant/1",bytes.NewBuffer(str) )

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
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