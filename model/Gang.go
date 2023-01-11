package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"strconv"
)

type Gang schema.Gang
type UpdateGang schema.UpdateGang

func (emp *Gang) Save() (*Gang, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func GangFindAll() ([]Gang, error) {
	var data []Gang
	err := database.Database.Find(&data).Error
	return data, err
}

func GangFindByName(name string) (Gang, error) {
	var data Gang
	err := database.Database.Where("name=?", name).First(&data).Error
	return data, err
}

func GangFindById(id string) (Gang, error) {
	intVar, _ := strconv.Atoi(id)
	var data Gang
	err := database.Database.Where("id=?", intVar).First(&data).Error
	return data, err
}
func GangFindByCode(code string) (Gang, error) {
	var data Gang
	err := database.Database.Where("code=?", code).First(&data).Error
	return data, err
}

func GangFindMapById(params []uint) ([]Gang, error) {
	var data []Gang
	err := database.Database.Where("id IN ?", params).Find(&data).Error
	return data, err
}

func GangFindMapByCode(params []string) ([]Gang, error) {
	var data []Gang
	err := database.Database.Where("code IN ?", params).Find(&data).Error
	return data, err
}

func (update_data Gang) GangChangeData(id string, ua UpdateEstate) (Gang, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return update_data, err
	}
	res, _ := GangFindById(id)
	return res, nil
}
