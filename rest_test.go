package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildURL(t *testing.T) {
	host := "http://api.test.com"
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	testURL := BuildURL(host, queryParams)
	if testURL != "http://api.test.com?test=1&test2=2" {
		t.Error("Bad BuildURL result")
	}
}

func TestBuildRequest(t *testing.T) {
	method := "GET"
	baseURL := "http://api.test.com"
	key := "API_KEY"
	requestHeaders := make(map[string]string)
	requestHeaders["Content-Type"] = "application/json"
	requestHeaders["Authorization"] = "Bearer " + key
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	request := Request{
		Method:         method,
		BaseURL:        baseURL,
		RequestHeaders: requestHeaders,
		QueryParams:    queryParams,
	}
	req, e := BuildRequest(request)
	if e != nil {
		t.Errorf("Rest failed to BuildRequest. Returned error: %v", e)
	}
	if req == nil {
		t.Errorf("Failed to BuildRequest.")
	}

}

func TestBuildResponse(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	baseURL := fakeServer.URL
	method := "GET"
	request := Request{
		Method:  method,
		BaseURL: baseURL,
	}
	req, e := BuildRequest(request)
	res, e := MakeRequest(req)
	response, e := BuildResponse(res)
	if response.StatusCode != 200 {
		t.Error("Invalid status code in BuildResponse")
	}
	if len(response.ResponseBody) == 0 {
		t.Error("Invalid response body in BuildResponse")
	}
	if len(response.ResponseHeaders) == 0 {
		t.Error("Invalid response headers in BuildResponse")
	}
	if e != nil {
		t.Errorf("Rest failed to make a valid API request. Returned error: %v", e)
	}
}

func TestRest(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	host := fakeServer.URL
	endpoint := "/test_endpoint"
	baseURL := host + endpoint
	key := "API_KEY"
	requestHeaders := make(map[string]string)
	requestHeaders["Content-Type"] = "application/json"
	requestHeaders["Authorization"] = "Bearer " + key
	method := "GET"
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	request := Request{
		Method:         method,
		BaseURL:        baseURL,
		RequestHeaders: requestHeaders,
		QueryParams:    queryParams,
	}
	response, e := API(request)
	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if len(response.ResponseBody) == 0 {
		t.Error("Invalid response body")
	}
	if len(response.ResponseHeaders) == 0 {
		t.Error("Invalid response headers")
	}
	if e != nil {
		t.Errorf("Rest failed to make a valid API request. Returned error: %v", e)
	}
}
