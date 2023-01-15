package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"

	"gorm.io/gorm/clause"
)

type Estate schema.Estate

func (data *Estate) Save() (*Estate, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func EstateFindAll() ([]Estate, error) {
	var data []Estate
	err := database.Database.Preload(clause.Associations).Where("is_deleted = ? AND is_active = ?", false, true).Find(&data).Error
	return data, err
}

func EstateFindByName(param string) (Estate, error) {
	var data Estate
	err := database.Database.Preload(clause.Associations).Where("name = ? AND is_deleted = ? AND is_active = ?", param, false, true).First(&data).Error
	return data, err
}

func EstateFindById(param string) (Estate, error) {
	var data Estate
	err := database.Database.Preload(clause.Associations).Where("id = ? AND is_deleted = ? AND is_active = ?", param, false, true).Or("code = ? AND is_deleted = ? AND is_active = ?", param, false, true).First(&data).Error
	return data, err
}

func EstateFindMapById(params []string) ([]Estate, error) {
	var data []Estate
	err := database.Database.Where("id IN ?", params).Or("code IN ?", params).Or("name IN ?", params).Find(&data).Error
	return data, err
}

func EstateFindDivisionById(params []string) ([]Estate, error) {
	var div []Estate
	err := database.Database.Preload("Company").Preload("Divisions").Preload("Divisions.Gangs").Where("id IN ?", params).Or("code IN ?", params).Find(&div).Error
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

func (update_data *Estate) EstateChangeData(id string, ua schema.RequestEstateUpdate) (Estate, error) {
	err := database.Database.Model(Estate{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := EstateFindById(id)
	return res, nil
}
