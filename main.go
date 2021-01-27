package main

import (
	"fmt"
	"golang-training/app/controllers"
	"golang-training/app/models"
)

func main() {

	e1 := models.Employee{"james", "bond", 12, 45}
	e2 := models.Employee{"Bruce", "wayne", 13, 32}

	models.EmployeeSlice = []models.Employee{e1, e2}

	//config.DBConn()

	controllers.RequestHandler()

	fmt.Println(models.EmployeeSlice)
}
