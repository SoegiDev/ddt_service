package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"fmt"
)

type Employee schema.Employee

func (emp *Employee) Save() (*Employee, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindEmployeeAll() ([]Employee, error) {
	var emp []Employee
	err := database.Database.Find(&emp).Error
	fmt.Println(emp)
	return emp, err
}

func FindEmployeeById(id string) (Employee, error) {
	var emp Employee
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}
