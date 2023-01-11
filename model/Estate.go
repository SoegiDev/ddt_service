package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Estate schema.Estate
type UpdateEstate schema.UpdateEstate

func (data *Estate) Save() (*Estate, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func EstateFindAll() ([]Estate, error) {
	var data []Estate
	err := database.Database.Find(&data).Error
	return data, err
}

func EstateFindByName(name string) (Estate, error) {
	var data Estate
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func EstateFindById(id string) (Estate, error) {
	var data Estate
	err := database.Database.Where("id=?", id).First(&data).Error
	return data, err
}

func EstateFindByCode(id string) (Estate, error) {
	var data Estate
	err := database.Database.Where("code=?", id).First(&data).Error
	return data, err
}

func EstateFindMapById(params []string) ([]Estate, error) {
	var data []Estate
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}

func EstateFindMapByCode(params []string) ([]Estate, error) {
	var data []Estate
	err := database.Database.Where("code IN ?", params).Find(&data).Error
	return data, err
}

func EstateFindMapByName(params []string) ([]Estate, error) {
	var div []Estate
	err := database.Database.Where("name IN ?", params).Find(&div).Error
	return div, err
}

func EstateFindDivisionById(params []string) ([]Estate, error) {
	var div []Estate
	err := database.Database.Preload("Company").Preload("Divisions").Preload("Divisions.Gangs").Where("id IN ?", params).Find(&div).Error
	return div, err
}

func (est *Estate) EstateSignDivision(id string, divUpdates []schema.Division) (Estate, error) {
	err := database.Database.Model(&est).Association("Divisions").Replace(divUpdates)
	if err != nil {
		return *est, err
	}
	res, _ := EstateFindById(id)
	return res, nil
}

func (update_data *Estate) EstateChangeData(id string, ua UpdateEstate) (Estate, error) {
	err := database.Database.Model(Estate{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := EstateFindById(id)
	return res, nil
}
