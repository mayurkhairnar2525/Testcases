package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
		request, err := http.NewRequest("GET", "/", nil)
		if err!= nil {
			t.Fatalf("couldn't create request %v",err)
	}
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
		assert.Equal(t,"Hello world",response.Body.String(),"Incorrect Body found ")
	}
