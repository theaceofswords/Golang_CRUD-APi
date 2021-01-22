package models

import "fmt"

var EmployeeSlice []Employee

type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        int    `json:"id"`
	Age       int    `json:"age"`
}

func findByid(id int) int {
	var index int
	for i, emp := range EmployeeSlice {
		if emp.ID == id {
			index = i
			fmt.Println(i, emp.ID)
			break
		}
	}
	return index
}

func Remove(id int) {
	index := findByid(id)
	//var index int
	// for i, emp := range EmployeeSlice {
	// 	if emp.ID == id {
	// 		index = i
	// 		fmt.Println("removing", i, emp.ID)
	// 		break
	// 	}
	// }
	EmployeeSlice = append(EmployeeSlice[:index], EmployeeSlice[index+1:]...)
}

func (e *Employee) update(newEmp Employee) {
	e.FirstName = newEmp.FirstName
	e.LastName = newEmp.LastName
	e.Age = newEmp.Age
}

func UpdateEntry(empIp Employee) {

	index := findByid(empIp.ID)

	// var index int
	// for i, emp := range EmployeeSlice {
	// 	if emp.ID == emp.ID {
	// 		index = i
	// 		break
	// 	}
	// }
	EmployeeSlice[index].update(empIp)
}
