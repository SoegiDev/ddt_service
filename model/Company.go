package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Company schema.Company
type UpdateCompany schema.UpdateCompany

func (data Company) Save() (Company, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func CompanyFindAll() ([]Company, error) {
	var data []Company
	err := database.Database.Find(&data).Error
	return data, err
}

func CompanyFindById(id string) (Company, error) {
	var data Company
	err := database.Database.Where("id=?", id).First(&data).Error
	return data, err
}

func CompanyFindByCode(code string) (Company, error) {
	var data Company
	err := database.Database.Where("code=?", code).First(&data).Error
	return data, err
}

func CompanyFindByName(name string) (Company, error) {
	var data Company
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func CompanyMapFindById(params []string) ([]Company, error) {
	var data []Company
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}
func CompanyMapFindByName(params []string) ([]Company, error) {
	var data []Company
	err := database.Database.Where("name IN ?", params).Find(&data).Error
	return data, err
}

func CompanyMapFindByCode(params []string) ([]Company, error) {
	var data []Company
	err := database.Database.Where("code IN ?", params).Find(&data).Error
	return data, err
}

func (update_data *Company) CompanyChangeData(id string, ua UpdateCompany) (Company, error) {
	err := database.Database.Model(Company{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := CompanyFindById(id)
	return res, nil
}
