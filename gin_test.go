package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/testgin/routers"
)

var g *gin.Engine

func init() {
	g = routers.SetupRouter()
}

func TestURLEncode(t *testing.T) {
	req, err := http.NewRequest("POST", "/testgin/urlencode?id=1234&page=1",
		strings.NewReader("username=alice&password=1234"))
	if err != nil {
		t.Fatal(err)
	}
	rw := httptest.NewRecorder()
	g.ServeHTTP(rw, req)

	if rw.Code != 200 {
		t.Fatal(`should receive "200"`, rw.Code)
	}
}

func TestJson(t *testing.T) {
	req, err := http.NewRequest("POST", "/testgin/json",
		strings.NewReader(`{"id":"123","page":"1","username":"alice","password":1234}`))
	if err != nil {
		t.Fatal(err)
	}
	rw := httptest.NewRecorder()
	g.ServeHTTP(rw, req)

	if rw.Code != 200 {
		t.Fatal(`should receive "200"`, rw.Code)
	}
}
