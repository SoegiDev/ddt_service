package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Account schema.Account
type UpdateAccount schema.UpdateAccount

func (emp *Account) Save() (*Account, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func AccountFindAll() ([]Account, error) {
	var data []Account
	err := database.Database.Find(&data).Error
	return data, err
}

func AccountFindById(id string, active bool) (Account, error) {
	var data Account
	err := database.Database.Where("id=?", id).Where("is_active=?", active).First(&data).Error
	return data, err
}

func AccountFindByUserId(id string) ([]Account, error) {
	var data []Account
	err := database.Database.Where("user_id=?", id).Find(&data).Error
	return data, err
}

func (account Account) AccountSignRoleApp(id string, corps []schema.RoleApplication) (Account, error) {
	err := database.Database.Model(&account).Association("RoleApplications").Replace(corps)
	if err != nil {
		return account, err
	}
	res, _ := AccountFindById(id, true)
	return res, nil
}

func (update_data Account) AccountChangeData(id string, ua UpdateAccount) (Account, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return update_data, err
	}
	res, _ := AccountFindById(id, true)
	return res, nil
}
