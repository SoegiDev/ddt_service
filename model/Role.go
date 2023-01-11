package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"fmt"
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

func FindRoleMapById(params []int) ([]Role, error) {
	var data []Role
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}

func FindRoleMapByName(params []string) ([]Role, error) {
	var data []Role
	err := database.Database.Where("name IN ?", params).Find(&data).Error
	fmt.Println(data)
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
