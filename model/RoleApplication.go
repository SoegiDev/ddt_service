package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type RoleApplication schema.RoleApplication
type UpdateRoleApplication schema.UpdateRoleApplication

func (data *RoleApplication) Save() (*RoleApplication, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func RoleAppAll(name string) (RoleApplication, error) {
	var data RoleApplication
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func RoleAppFindById(id string) (RoleApplication, error) {
	var data RoleApplication
	err := database.Database.Where("id=?", id).First(&data).Error
	return data, err
}

func RoleAppMapById(params []int) ([]RoleApplication, error) {
	var data []RoleApplication
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}
func RoleAppMapByName(params []string) ([]RoleApplication, error) {
	var data []RoleApplication
	err := database.Database.Where("name IN ?", params).Find(&data).Error
	return data, err
}

func (data RoleApplication) RoleAppChangeData(id string, ua UpdateRoleApplication) (RoleApplication, error) {
	err := database.Database.Model(Role{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return data, err
	}
	res, _ := RoleAppFindById(id)
	return res, nil
}
