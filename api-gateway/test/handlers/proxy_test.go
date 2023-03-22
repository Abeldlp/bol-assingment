package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abeldlp/bol-assingment/api-gateway/handlers"
	"github.com/gin-gonic/gin"
)

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func (r *TestResponseRecorder) closeClient() {
	r.closeChannel <- true
}

func TestProxyToSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test" {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	rr := CreateTestResponseRecorder()
	_, r := gin.CreateTestContext(rr)
	defer rr.closeClient()

	r.GET("/v1/api/*proxyPath", handlers.ProxyTo(server.URL))

	request, _ := http.NewRequest(http.MethodGet, "/v1/api/test", http.NoBody)
	r.ServeHTTP(rr, request)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestProxyTo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	}))
	defer server.Close()

	rr := CreateTestResponseRecorder()
	_, r := gin.CreateTestContext(rr)
	defer rr.closeClient()

	r.GET("/v1/api/*proxyPath", handlers.ProxyTo(server.URL))

	request, _ := http.NewRequest(http.MethodGet, "/v1/api/non-existant", http.NoBody)
	r.ServeHTTP(rr, request)
	if rr.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestProxyToPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test" {
			if r.Method == http.MethodPost {
				w.WriteHeader(http.StatusCreated)
			}
		}
	}))
	defer server.Close()

	rr := CreateTestResponseRecorder()
	_, r := gin.CreateTestContext(rr)
	defer rr.closeClient()

	r.POST("/v1/api/*proxyPath", handlers.ProxyTo(server.URL))

	request, _ := http.NewRequest(http.MethodPost, "/v1/api/test", http.NoBody)
	r.ServeHTTP(rr, request)
	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}
