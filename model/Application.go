package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"

	"gorm.io/gorm/clause"
)

type Applications schema.Application

func (data *Applications) Save() (*Applications, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func AppsFindById(id string) (Applications, error) {
	var data Applications
	err := database.Database.Preload(clause.Associations).Where("id = ? AND is_deleted = ? AND is_active = ?", id, false, true).Or("code = ? AND is_deleted = ? AND is_active = ?", id, false, true).First(&data).Error
	return data, err
}

func AppsFindMapById(params []string) ([]Applications, error) {
	var data []Applications
	err := database.Database.Where("id IN ?", params).Or("name IN ?", params).Find(&data).Error
	return data, err
}

func AppsFindByAll() ([]Applications, error) {
	var data []Applications
	err := database.Database.Preload(clause.Associations).Where("is_deleted = ? AND is_active = ?", false, true).Find(&data).Error
	return data, err
}

func (update_data Applications) AppsChangeData(id string, ua schema.RequestApplicationUpdate) (Applications, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return update_data, err
	}
	res, _ := AppsFindById(id)
	return res, nil
}
