package orders_client

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

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