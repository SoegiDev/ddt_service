package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"

	"gorm.io/gorm/clause"
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

func EmployeeFindById(id string) (Employee, error) {
	var data Employee
	err := database.Database.Preload(clause.Associations).Where("id = ? AND is_active= ?", id, true).Or("code= ? AND is_active= ?", id, true).First(&data).Error
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
