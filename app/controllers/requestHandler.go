package controllers

import (
	"encoding/json"
	"fmt"
	"golang-training/app/repository"
	"net/http"
	"strconv"
)

func employeeCRUD(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(repository.OrmRead())

	case "POST":
		var emp repository.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			fmt.Println("error")
		}
		repository.OrmCreate(emp)
		fmt.Fprint(w, "addded")

	case "PUT":
		var emp repository.Employee
		_ = json.NewDecoder(r.Body).Decode(&emp)
		repository.OrmUpdate(emp)
		fmt.Println(w, "Updated")

	case "DELETE":
		id, _ := strconv.ParseInt(r.URL.Query().Get("userId"), 10, 64)
		repository.OrmDelete(id)
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
