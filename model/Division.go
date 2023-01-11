package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Division schema.Division
type UpdateDivision schema.UpdateDivision

func (data *Division) Save() (*Division, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func DivisionFindAll() ([]Division, error) {
	var data []Division
	err := database.Database.Find(&data).Error
	return data, err
}

func DivisionFindByName(name string) (Division, error) {
	var data Division
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func DivisionFindById(id string) (Division, error) {
	var data Division
	err := database.Database.Where("id=?", id).First(&data).Error
	return data, err
}

func DivisionFindByCode(code string) (Division, error) {
	var data Division
	err := database.Database.Where("code=?", code).First(&data).Error
	return data, err
}

func DivisonFindMapById(params []string) ([]Division, error) {
	var data []Division
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}

func DivisonFindMapByName(params []string) ([]Division, error) {
	var data []Division
	err := database.Database.Where("name IN ?", params).Find(&data).Error
	return data, err
}

func DivisonFindMapByCode(params []string) ([]Division, error) {
	var data []Division
	err := database.Database.Where("code IN ?", params).Find(&data).Error
	return data, err
}

func (data Division) DivisionChangeData(id string, ua UpdateDivision) (Division, error) {
	err := database.Database.Model(Article{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return data, err
	}
	res, _ := DivisionFindById(id)
	return res, nil
}
