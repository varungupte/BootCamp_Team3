package restaurants_client

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