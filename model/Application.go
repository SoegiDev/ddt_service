package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Applications schema.Application
type UpdateApplication schema.UpdateApplication

func (data *Applications) Save() (*Applications, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func (update_data Applications) AppsChangeData(id string, ua UpdateApplication) (Applications, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return update_data, err
	}
	res, _ := AppsFindById(id)
	return res, nil
}

func AppsFindById(id string) (Applications, error) {
	var data Applications
	err := database.Database.Where("id = ?", id).First(&data).Error
	return data, err
}

func AppsFindMapById(params []string) ([]Applications, error) {
	var data []Applications
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}

func AppsFindMapByName(params []string) ([]Applications, error) {
	var data []Applications
	err := database.Database.Where("name IN ?", params).Find(&data).Error
	return data, err
}
