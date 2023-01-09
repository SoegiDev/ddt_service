package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"fmt"
)

type Company schema.Company
type UpdateCompany schema.UpdateCompany

func (emp Company) Save() (Company, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindCompanyAll() ([]Company, error) {
	var emp []Company
	err := database.Database.Find(&emp).Error
	fmt.Println(emp)
	return emp, err
}

func FindCompanyById(id string) (Company, error) {
	var emp Company
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}

func FindCompanyByName(id string) (Company, error) {
	var emp Company
	err := database.Database.Where("name=?", id).First(&emp).Error
	return emp, err
}

func FindCompanyByCode(id string) (Company, error) {
	var emp Company
	err := database.Database.Where("code=?", id).First(&emp).Error
	return emp, err
}

func FindCompanyMapById(params []string) ([]Company, error) {
	var corp []Company
	err := database.Database.Where("id IN ?", params).Find(&corp).Error
	fmt.Println(corp)
	return corp, err
}
func FindCompanyMapByName(params []string) ([]Company, error) {
	var corp []Company
	err := database.Database.Where("name IN ?", params).Find(&corp).Error
	fmt.Println(corp)
	return corp, err
}

func (update_data *Company) ChangeData(id string, ua UpdateCompany) (Company, error) {
	err := database.Database.Model(Company{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindCompanyById(id)
	return res, nil
}
