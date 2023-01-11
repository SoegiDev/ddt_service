package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Employee schema.Employee

func (data *Employee) Save() (*Employee, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func EmployeeFindAll() ([]Employee, error) {
	var data []Employee
	err := database.Database.Find(&data).Error
	return data, err
}

func EmployeeFindByUserId(user_id string) (Employee, error) {
	var data Employee
	err := database.Database.Where("user_id=?", user_id).Where("is_active=?", true).First(&data).Error
	return data, err
}

func EmployeeFindById(id string) (Employee, error) {
	var data Employee
	err := database.Database.Where("id=?", id).Where("is_active=?", true).First(&data).Error
	return data, err
}

func EmployeeFindByUsername(username string) (Employee, error) {
	var data Employee
	err := database.Database.Where("username=?", username).Where("is_active=?", true).First(&data).Error
	return data, err
}

func EmployeeFindByEmail(email string) (Employee, error) {
	var data Employee
	err := database.Database.Where("email=?", email).First(&data).Error
	return data, err
}
