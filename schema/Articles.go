package schema

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	Id        string `json:"id" gorm:"primaryKey;type:varchar(20)"`
	Title     string `gorm:"size:50;" json:"title"`
	UserID    string `json:"user_id" gorm:"primaryKey;type:varchar(100)"`
	Content   string `gorm:"size:255;type:text" json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsActive  bool `gorm:"type:bool;default:true" json:"status"`
}

func (Articles) TableName() string {
	return "articles"
}

type UpdateArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}
