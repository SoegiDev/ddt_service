package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"strconv"
)

type Role schema.Role
type UpdateRole schema.UpdateRole

func (data *Role) Save() (*Role, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func RoleFindByName(name string) (Role, error) {
	var data Role
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func FindRoleById(id string) (Role, error) {
	var data Role
	err := database.Database.Where("id=?", id).First(&data).Error
	return data, err
}

func FindRoleMapById(params []string) ([]Role, error) {
	dtParams := []int{}
	for _, element := range params {
		intVar, _ := strconv.Atoi(element)
		dtParams = append(dtParams, intVar)
	}
	var data []Role
	err := database.Database.Where("id IN ?", dtParams).Or("name IN ?", params).Find(&data).Error
	return data, err
}

func (data Role) RoleChangeData(id string, ua UpdateRole) (Role, error) {
	err := database.Database.Model(Role{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return data, err
	}
	res, _ := FindRoleById(id)
	return res, nil
}
