package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Division schema.Division

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
