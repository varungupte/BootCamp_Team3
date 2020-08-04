package customers_client

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

func TestGetCustomer(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET ("/customer/id/1", GetCustomer)
	req, _ := http.NewRequest("GET", "/customer/id/1", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}

func TestGetCustomerCount(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.GET ("/customer/count", GetCustomerCount)
	req, _ := http.NewRequest("GET", "/customer/count", nil)

	bearer := "Bearer "+token
	req.Header.Add("Authorization", bearer)

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("failed ")
		log.Println(w.Result())
	}
}
func TestDeleteCustomer(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	token := getToken()
	r.DELETE ("/customer/id/1", DeleteCustomer)
	req, _ := http.NewRequest("DELETE", "/customer/id/1", nil)

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