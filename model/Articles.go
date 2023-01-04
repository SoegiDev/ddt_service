package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
)

type Articles schema.Articles
type UpdateArticle schema.UpdateArticle

func (article *Articles) Save() (*Articles, error) {
	err := database.Database.Create(&article).Error
	return article, err
}

func (update_data *Articles) ChangeData(id string, ud UpdateArticle) (Articles, error) {
	err := database.Database.Model(Articles{}).Where("id = ?", id).Updates(ud).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindEntryById(id)
	return res, nil
}

func FindEntryById(id string) (Articles, error) {
	var article Articles
	err := database.Database.Where("id=?", id).First(&article).Error
	return article, err
}
