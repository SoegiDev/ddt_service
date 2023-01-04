package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"fmt"
)

type Role schema.Role
type RoleMap schema.RoleMap

func (role *Role) Save() (*Role, error) {
	err := database.Database.Create(&role).Error
	return role, err
}

func FindRoleByName(name string) (Role, error) {
	var role Role
	err := database.Database.Where("name=?", name).First(&role).Error
	return role, err
}

func FindRoleById(id string) (Role, error) {
	var role Role
	err := database.Database.Where("id=?", id).First(&role).Error
	fmt.Println(role)
	return role, err
}

func FindRoleByIdMap(params []string) ([]Role, error) {
	var role []Role
	err := database.Database.Select("id").Where("id IN ?", params).Find(&role).Error
	fmt.Println(role)
	return role, err
}
