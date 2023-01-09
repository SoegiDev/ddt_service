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

func FindEmployeeByUserId(id string) (Employee, error) {
	var emp Employee
	err := database.Database.Where("user_id=?", id).Where("is_active=?", true).First(&emp).Error
	return emp, err
}

func FindEmployeeById(id string) (Employee, error) {
	var emp Employee
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}

func FindEmployeeByName(id string) (Employee, error) {
	var emp Employee
	err := database.Database.Where("username=?", id).First(&emp).Error
	return emp, err
}

func FindEmployeeByEmail(id string) (Employee, error) {
	var emp Employee
	err := database.Database.Where("email=?", id).First(&emp).Error
	return emp, err
}
