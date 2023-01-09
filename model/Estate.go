package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Estate schema.Estate
type UpdateEstate schema.UpdateEstate

func (emp *Estate) Save() (*Estate, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindEstateAll() ([]Estate, error) {
	var emp []Estate
	err := database.Database.Find(&emp).Error
	return emp, err
}

func FindEstateByName(name string) (Estate, error) {
	var div Estate
	err := database.Database.Where("name=?", name).First(&div).Error
	return div, err
}

func FindEstateById(id string) (Estate, error) {
	var div Estate
	err := database.Database.Where("id=?", id).First(&div).Error
	return div, err
}

func FindEstateMapById(params []string) ([]Estate, error) {
	var div []Estate
	err := database.Database.Where("id IN ?", params).Find(&div).Error
	return div, err
}

func (est *Estate) EstateAssignDivision(id string, divUpdates []schema.Division) (Estate, error) {
	err := database.Database.Model(&est).Association("Divisions").Replace(divUpdates)
	if err != nil {
		return *est, err
	}
	res, _ := FindEstateById(id)
	return res, nil
}
func FindEstateMapByName(params []string) ([]Estate, error) {
	var div []Estate
	err := database.Database.Where("name IN ?", params).Find(&div).Error
	return div, err
}

func (update_data *Estate) ChangeData(id string, ua UpdateEstate) (Estate, error) {
	err := database.Database.Model(Estate{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindEstateById(id)
	return res, nil
}
