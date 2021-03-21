package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
	_ "testing"
	"encoding/json"
	"bytes"
)



func TestPlantGetRequest(t *testing.T) {
	fmt.Print("MY TEST FILE++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	request, _ := http.NewRequest("GET", "/plant", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestPlantCreateRequest(t *testing.T) {
    plant := &Plant{
        Ident: "112",
        Name: "Plant Test",
		Status: "available",
		Price: "4000",
    }
    jsonPlant, _ := json.Marshal(plant)
    request, _ := http.NewRequest("POST", "/plant", bytes.NewBuffer(jsonPlant))
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}


func TestPriceRequest(t *testing.T){
	inputData := &Order{
		Ident:"111",
		Start:"20",
		End:"22",
	}

	jsonData, _ := json.Marshal(inputData)
    request, _ := http.NewRequest("POST", "/price", bytes.NewBuffer(jsonData))
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestStatusRequest(t *testing.T){
	inputData := &Order{
		Ident:"111",
		Start:"20",
		End:"22",
	}

	jsonData, _ := json.Marshal(inputData)
    request, _ := http.NewRequest("POST", "/status", bytes.NewBuffer(jsonData))
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestRequestsUrl(t *testing.T){
	request, _ := http.NewRequest("GET", "/requests", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}