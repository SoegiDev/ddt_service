package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"

	"gorm.io/gorm/clause"
)

type Company schema.Company

func (data Company) Save() (Company, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func CompanyFindAll() ([]Company, error) {
	var data []Company
	err := database.Database.Preload(clause.Associations).Where("is_deleted = ? AND is_active = ?", false, true).Find(&data).Error
	return data, err
}

func CompanyFindById(param string) (Company, error) {
	var data Company
	err := database.Database.Preload(clause.Associations).Where("id = ? AND is_deleted = ? AND is_active = ?", param, false, true).Or("code = ? AND is_deleted = ? AND is_active = ?", param, false, true).First(&data).Error
	return data, err
}

func CompanyFindByName(param string) (Company, error) {
	var data Company
	err := database.Database.Preload(clause.Associations).Where("name = ? AND is_deleted = ? AND is_active = ?", param, false, true).First(&data).Error
	return data, err
}

func CompanyMapFindById(params []string) ([]Company, error) {
	var data []Company
	err := database.Database.Preload("Employees").Preload("Estate").Where("id IN ?", params).Or("code IN ?", params).Or("name IN ?", params).Find(&data).Error
	return data, err
}
func (update_data *Company) CompanyChangeData(id string, ua schema.RequestCompanyUpdate) (Company, error) {
	err := database.Database.Model(Company{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := CompanyFindById(id)
	return res, nil
}
