package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var err error

var handler = http.HandlerFunc(employeeCRUD)

//func TestEmployeeCRUD(t *testing.T) {

func TestGET(t *testing.T) {
	request, _ := http.NewRequest("GET", "/employee", nil)

	q := request.URL.Query()
	q.Add("EmpId", "1")
	request.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)
	//employeeCRUD(recorder, request)

	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected Ok, returned %v", response.Status)
	}

	request, _ = http.NewRequest("GET", "/employee", nil)

	q = request.URL.Query()
	q.Add("EmpId", "a")
	request.URL.RawQuery = q.Encode()
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	response = recorder.Result()

	// expected := `{"error":"strconv.ParseInt: parsing \"a\": invalid syntax","status":422,"message":"Invalid parameter type"}`

	// if recorder.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		recorder.Body.String(), expected)
	// }

	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, returned %v", response.StatusCode)
	}

	request, _ = http.NewRequest("GET", "/employee", nil)

	q = request.URL.Query()
	q.Add("EmpId", "7")
	request.URL.RawQuery = q.Encode()
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	response = recorder.Result()

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, returned %v", response.StatusCode)
	}

}

func TestPOST(t *testing.T) {
	var jsonStr = []byte(` {
		        "FirstName": "George",
		        "LastName": "H",
		        "EmpID": 32,
		        "Age": 29
			}`)

	request, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)
	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected Ok, returned %v", response.Status)
	}

	// expected := `{
	// 	        "FirstName": "john",
	// 	        "LastName": "Mathew",
	// 	        "EmpID": 32,
	// 	        "Age": 29
	// 	    }`

	// if recorder.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		recorder.Body.String(), expected)
	// }

	fmt.Println("test2")

	jsonStr = []byte(` {
		"FirstName": "John",
		"LastName": "H",
		"EmpID": 32,
		"Age": 49
	}`)

	request, _ = http.NewRequest("POST", "/employee", bytes.NewBuffer(jsonStr))
	recorder3 := httptest.NewRecorder()

	handler.ServeHTTP(recorder3, request)
	response = recorder3.Result()
	if response.StatusCode != http.StatusConflict {
		t.Errorf("Expected 409, returned %v", response.Status)
	}

	scrStr := []byte(`{"somesentence"}`)

	request, _ = http.NewRequest("POST", "/employee", bytes.NewBuffer(scrStr))
	recorder2 := httptest.NewRecorder()

	handler = http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder2, request)
	response = recorder2.Result()
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, returned %v", response.StatusCode)
	}

}

func TestPUT(t *testing.T) {
	jsonStr := []byte(` {
		"FirstName": "Jeff",
		"LastName": "H",
		"EmpID": 32,
		"Age": 29
	}`)

	request, err := http.NewRequest("PUT", "/employee", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	response := recorder.Result()

	if response.StatusCode != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v",
			response.StatusCode, http.StatusOK)
	}

	scrStr := []byte(`{"somesentence"}`)

	request, _ = http.NewRequest("PUT", "/employee", bytes.NewBuffer(scrStr))
	recorder2 := httptest.NewRecorder()

	handler = http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder2, request)
	response = recorder2.Result()
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, returned %v", response.StatusCode)
	}

	jsonStr = []byte(` {
		"FirstName": "record",
		"LastName": "non existant",
		"EmpID": 112,
		"Age": 49
	}`)

	request, _ = http.NewRequest("PUT", "/employee", bytes.NewBuffer(jsonStr))
	recorder3 := httptest.NewRecorder()

	handler.ServeHTTP(recorder3, request)
	response = recorder3.Result()
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, returned %v", response.Status)
	}

	fmt.Println("PUT Test complete")
}

func TestDELETE(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/employee", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("EmpId", "32")
	request.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("returned wrong status code: got %v want %v ", response.StatusCode, http.StatusOK)
	}

	request, _ = http.NewRequest("DELETE", "/employee", nil)

	q = request.URL.Query()
	q.Add("EmpId", "t")
	request.URL.RawQuery = q.Encode()
	recorder2 := httptest.NewRecorder()
	handler.ServeHTTP(recorder2, request)
	response = recorder2.Result()

	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, returned %v", response.StatusCode)
	}

	request, _ = http.NewRequest("DELETE", "/employee", nil)

	q = request.URL.Query()
	q.Add("EmpId", "7")
	request.URL.RawQuery = q.Encode()
	recorder3 := httptest.NewRecorder()
	handler.ServeHTTP(recorder3, request)
	response = recorder3.Result()

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, returned %v", response.StatusCode)
	}

	fmt.Println("DELETE Test complete")

}
