package controllers

import (
	"encoding/json"
	"fmt"
	"golang-training/app/models"
	"golang-training/app/repository"
	"net/http"
	"strconv"
)

type messageErr struct {
	ErrError   string `json:"error"`
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
}

func employeeCRUD(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		id, err := strconv.ParseInt(r.URL.Query().Get("EmpId"), 10, 64)
		if err != nil {
			msg := messageErr{err.Error(), http.StatusUnprocessableEntity, "Invalid parameter type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
			//fmt.Println(msg)
		} else {
			res, err := repository.ReadById(id)
			if err != nil {
				msg := messageErr{err.Error(), http.StatusNotFound, "Record not found"}
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(msg)
			} else {
				json.NewEncoder(w).Encode(res)
			}

		}

		//json.NewEncoder(w).Encode(repository.ReadEmployee())

	case "POST":
		var emp models.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			msg := messageErr{err.Error(), http.StatusUnprocessableEntity, "Invalid Body type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
			//fmt.Println("error")
		} else {
			err = repository.CreateEmployee(emp)
			if err != nil {
				msg := messageErr{err.Error(), http.StatusConflict, "ID already exists"}
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(msg)
			} else {
				json.NewEncoder(w).Encode(emp)
			}

		}

	case "PUT":
		var emp models.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			msg := messageErr{err.Error(), http.StatusUnprocessableEntity, "Invalid Body type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
			//fmt.Println("error")
		} else {
			err = repository.UpdateEmployee(emp)
			if err != nil {
				msg := messageErr{err.Error(), http.StatusNotFound, "Record does not exist"}
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(msg)
			} else {
				fmt.Fprint(w, "Updated")
			}
		}

	case "DELETE":
		id, err := strconv.ParseInt(r.URL.Query().Get("EmpId"), 10, 64)
		if err != nil {
			msg := messageErr{err.Error(), http.StatusUnprocessableEntity, "Invalid parameter type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
		} else {
			err := repository.DeleteEmployee(id)
			if err != nil {
				msg := messageErr{err.Error(), http.StatusNotFound, "Record not found"}
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(msg)
			} else {
				fmt.Fprint(w, "deleted")
			}

		}

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
