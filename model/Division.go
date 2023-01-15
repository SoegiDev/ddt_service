package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"

	"gorm.io/gorm/clause"
)

type Division schema.Division

func (data *Division) Save() (*Division, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func DivisionFindAll() ([]Division, error) {
	var data []Division
	err := database.Database.Preload(clause.Associations).Where("is_deleted = ? AND is_active = ?", false, true).Find(&data).Find(&data).Error
	return data, err
}

func DivisionFindByName(param string) (Division, error) {
	var data Division
	err := database.Database.Preload(clause.Associations).Where("name=? AND is_deleted = ? AND is_active = ?", param, false, true).Find(&data).First(&data).Error
	return data, err
}

func DivisionFindById(param string) (Division, error) {
	var data Division
	err := database.Database.Preload(clause.Associations).Where("id=? AND is_deleted = ? AND is_active = ?", param, false, true).Or("code=? AND is_deleted = ? AND is_active = ?", param, false, true).Find(&data).First(&data).Error
	return data, err
}

func DivisonFindMapById(params []string) ([]Division, error) {
	var data []Division
	err := database.Database.Where("id IN ?", params).Or("code IN ?", params).Or("name IN ?", params).Find(&data).Error
	return data, err
}

func (data Division) DivisionChangeData(id string, ua schema.RequestDivisionUpdate) (Division, error) {
	err := database.Database.Model(Article{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return data, err
	}
	res, _ := DivisionFindById(id)
	return res, nil
}
