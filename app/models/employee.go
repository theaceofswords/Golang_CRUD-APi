package models

import "fmt"

//EmployeeSlice Exported
var EmployeeSlice []Employee

// Employee Exported
type Employee struct {
	FirstName string 
	LastName  string
	ID        int
	Age       int
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

// Remove func Exported
func Remove(id int) {
	index := findByid(id)
	EmployeeSlice = append(EmployeeSlice[:index], EmployeeSlice[index+1:]...)
}

func (e *Employee) update(newEmp Employee) {
	e.FirstName = newEmp.FirstName
	e.LastName = newEmp.LastName
	e.Age = newEmp.Age
}

// UpdateEntry Exported
func UpdateEntry(empIP Employee) {

	index := findByid(empIP.ID)
	EmployeeSlice[index].update(empIP)
}
