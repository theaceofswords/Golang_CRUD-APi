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
		json.NewEncoder(w).Encode(repository.ReadEmployee())

	case "POST":
		var emp models.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			fmt.Println("error")
		}
		repository.CreateEmployee(emp)
		fmt.Fprint(w, "addded")

	case "PUT":
		var emp models.Employee
		_ = json.NewDecoder(r.Body).Decode(&emp)
		repository.UpdateEmployee(emp)
		fmt.Fprint(w, "Updated")

	case "DELETE":
		id, _ := strconv.ParseInt(r.URL.Query().Get("userId"), 10, 64)
		repository.DeleteEmployee(id)
		fmt.Fprint(w, "deleted")

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}

// RequestHandler Exported
func RequestHandler() {
	http.HandleFunc("/employee", employeeCRUD)
	fmt.Println("Running")
	http.ListenAndServe(":8080", nil)
}
