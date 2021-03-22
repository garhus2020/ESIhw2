package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	_ "net/http"
	"testing"
	_ "testing"
	"encoding/json"
	"bytes"
	"github.com/garhus2020/ESIhw2/plant/pkg/domain"
)



func TestPlantGetRequest(t *testing.T) {
	fmt.Print("/plant URL testing")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/plant", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")
    response, err := client.Do(request)
	if err != nil {
        fmt.Println("Request error: ", err.Error())
    }
	// response := httptest.NewRecorder()
	// Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.StatusCode, "OK response is expected")
	defer response.Body.Close()
    
}

func TestPlantCreateRequest(t *testing.T) {
    plant := &domain.Plant{
        Ident: "112",
        Name: "Plant Test",
		Status: "available",
		Price: "4000",
    }
    jsonPlant, _ := json.Marshal(plant)
	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/plant", bytes.NewBuffer(jsonPlant))
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")
	response, err := client.Do(request)
    // response := httptest.NewRecorder()
    // Router().ServeHTTP(response, request)
	if err != nil {
        fmt.Println("Request error: ", err.Error())
    }
    assert.Equal(t, 201, response.StatusCode, "OK response is expected")
	defer response.Body.Close()
}


// func TestPriceRequest(t *testing.T){
// 	inputData := &domain.Order{
// 		Ident:"111",
// 		Start:"20",
// 		End:"22",
// 	}

// 	jsonData, _ := json.Marshal(inputData)
// 	client := &http.Client{}
// 	request, _ := http.NewRequest("POST", "http://localhost:8080/price", bytes.NewBuffer(jsonData))
// 	// request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")
// 	response, err := client.Do(request)
// 	if err != nil {
//         fmt.Println("Request error: ", err.Error())
//     }
//     assert.Equal(t, 201, response.StatusCode, "OK response is expected")
// 	defer response.Body.Close()
// }

// func TestStatusRequest(t *testing.T){
// 	inputData := &domain.Order{
// 		Ident:"111",
// 		Start:"20",
// 		End:"22",
// 	}

// 	jsonData, _ := json.Marshal(inputData)
// 	client := &http.Client{}
// 	request, err := http.NewRequest("POST", "http://localhost:8080/status", bytes.NewBuffer(jsonData))
// 	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")
// 	response, err := client.Do(request)
// 	if err != nil {
//         fmt.Println("Request error: ", err.Error())
//     }
//     // response := httptest.NewRecorder()
//     // Router().ServeHTTP(response, request)
//     assert.Equal(t, 200, response.StatusCode, "OK response is expected")
// }

func TestRequestsUrl(t *testing.T){
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/requests", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0")
	response, err := client.Do(request)
	if err != nil {
        fmt.Println("Request error: ", err.Error())
    }
    assert.Equal(t, 200, response.StatusCode, "OK response is expected")
}