package controllers

import (
	"encoding/json"
	"fmt"
	"golang-training/app/models"
	"golang-training/app/repository"
	"net/http"
)

func viewEntry(w http.ResponseWriter, r *http.Request) {

	fmt.Println(" get point hit")

	//fmt.Println(models.EmployeeSlice)
	json.NewEncoder(w).Encode(repository.GetAllEmployee())
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		fmt.Println("error")
	}

	// models.EmployeeSlice = append(models.EmployeeSlice, emp)
	json.NewEncoder(w).Encode(repository.AddEmployee(emp))

}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete function called")
	var id int
	_ = json.NewDecoder(r.Body).Decode(&id)
	fmt.Println("id:", id)

	repository.RemoveEmployee(id)
	//models.Remove(id)
	//	json.NewEncoder(w).Encode(models.EmployeeSlice)
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)

	// models.UpdateEntry(emp)
	json.NewEncoder(w).Encode(repository.UpdateEmployee(emp))
}

//RequestHandler Expoerted
func RequestHandler() {
	fmt.Println(models.EmployeeSlice)
	http.HandleFunc("/view", viewEntry)
	http.HandleFunc("/add", createEntry)
	http.HandleFunc("/delete", deleteEntry)
	http.HandleFunc("/update", updateEntry)
	http.ListenAndServe(":8080", nil)
}
