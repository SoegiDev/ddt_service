package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Division schema.Division

func (emp *Division) Save() (*Division, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindDivisiAll() ([]Division, error) {
	var emp []Division
	err := database.Database.Find(&emp).Error
	return emp, err
}

func FindDivisiByName(name string) (Division, error) {
	var div Division
	err := database.Database.Where("name=?", name).First(&div).Error
	return div, err
}

func FindDivisiById(id string) (Division, error) {
	var div Division
	err := database.Database.Where("id=?", id).First(&div).Error
	return div, err
}

func FindDivisionMapById(params []string) ([]Division, error) {
	var div []Division
	err := database.Database.Where("id IN ?", params).Find(&div).Error
	return div, err
}

func FindDivisionMapByName(params []string) ([]Division, error) {
	var div []Division
	err := database.Database.Where("name IN ?", params).Find(&div).Error
	return div, err
}
