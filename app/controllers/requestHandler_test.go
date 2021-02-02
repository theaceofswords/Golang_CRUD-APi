package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var err error

func TestEmployeeCRUD(t *testing.T) {
	request, _ := http.NewRequest("GET", "/employee", nil)
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder, request)
	//employeeCRUD(recorder, request)

	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected Ok, returned %v", response.Status)
	}
	//emp := recorder.Body.String()

	fmt.Println("GET Test complete")

	var jsonStr = []byte(` {
		        "FirstName": "john",
		        "LastName": "Mathew",
		        "EmpID": 32,
		        "Age": 29
			}`)

	request, _ = http.NewRequest("POST", "/employee", bytes.NewBuffer(jsonStr))
	//	recorder = httptest.NewRecorder()

	handler = http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder, request)
	//employeeCRUD(recorder, request)

	response = recorder.Result()
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

	fmt.Println("POST Test complete")

	jsonStr = []byte(` {
		"FirstName": "John",
		"LastName": "M",
		"EmpID": 32,
		"Age": 29
	}`)

	request, err = http.NewRequest("PUT", "/employee", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//	recorder = httptest.NewRecorder()
	handler = http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder, request)
	response = recorder.Result()

	if response.StatusCode != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v",
			response.StatusCode, http.StatusOK)
	}

	fmt.Println("PUT Test complete")

	request, err = http.NewRequest("DELETE", "/employee", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("EmpId", "32")
	request.URL.RawQuery = q.Encode()
	//recorder = httptest.NewRecorder()
	handler = http.HandlerFunc(employeeCRUD)
	handler.ServeHTTP(recorder, request)

	response = recorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("returned wrong status code: got %v want %v ", response.StatusCode, http.StatusOK)
	}

	fmt.Println("DELETE Test complete")

}
