package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"strconv"

	"gorm.io/gorm/clause"
)

type Gang schema.Gang

func (emp *Gang) Save() (*Gang, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func GangFindAll() ([]Gang, error) {
	var data []Gang
	err := database.Database.Preload(clause.Associations).Where("is_deleted = ? AND is_active = ?", false, true).Find(&data).Error
	return data, err
}

func GangFindByName(param string) (Gang, error) {
	var data Gang
	err := database.Database.Preload(clause.Associations).Where("name=? AND is_deleted = ? AND is_active = ?", param, false, true).First(&data).Error
	return data, err
}

func GangFindById(id string) (Gang, error) {
	intVar, _ := strconv.Atoi(id)
	var data Gang
	err := database.Database.Preload(clause.Associations).Where("id=? AND is_deleted = ? AND is_active = ?", intVar, false, true).Or("code=? AND is_deleted = ? AND is_active = ?", id, false, true).First(&data).Error
	return data, err
}

func GangFindMapById(params []uint) ([]Gang, error) {
	var data []Gang
	err := database.Database.Where("id IN ?", params).Or("code IN ?", params).Or("name IN ?", params).Find(&data).Error
	return data, err
}

func (update_data Gang) GangChangeData(id string, ua schema.RequestGangUpdate) (Gang, error) {
	intVar, _ := strconv.Atoi(id)
	err := database.Database.Model(&update_data).Where("id = ?", intVar).Or("code = ?", id).Updates(ua).Error
	if err != nil {
		return update_data, err
	}
	res, _ := GangFindById(id)
	return res, nil
}
