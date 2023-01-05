package schema

import (
	"time"

	"gorm.io/gorm"
)

type Estate struct {
	Id          string `json:"id" gorm:"primaryKey;type:varchar(20)"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsActive    bool       `gorm:"type:bool;default:true" json:"status"`
	Divisions   []Division `gorm:"foreignKey:EstateId"`
}

func (Estate) TableName() string {
	return "estates"
}
