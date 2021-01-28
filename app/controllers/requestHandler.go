package controllers

import (
	"encoding/json"
	"fmt"
	"golang-training/app/models"
	"golang-training/app/repository"
	"net/http"
	"strconv"
)

func employeeCRUD(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(repository.GetAllEmployee())

	case "POST":
		var emp models.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			fmt.Println("error")
		}
		json.NewEncoder(w).Encode(repository.AddEmployee(emp))

	case "PUT":
		var emp models.Employee
		_ = json.NewDecoder(r.Body).Decode(&emp)
		json.NewEncoder(w).Encode(repository.UpdateEmployee(emp))

	case "DELETE":
		id, _ := strconv.ParseInt(r.URL.Query().Get("userId"), 10, 8)
		repository.RemoveEmployee(id)
	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

// RequestHandler Exported
func RequestHandler() {
	http.HandleFunc("/employee", employeeCRUD)
	http.ListenAndServe(":8080", nil)
}
