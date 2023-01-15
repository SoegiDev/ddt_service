package schema

import (
	"time"

	"gorm.io/gorm"
)

type Division struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Code        string `gorm:"size:50;" json:"code"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool   `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool   `gorm:"type:bool;default:true" json:"status"`
	EstateId    string `gorm:"size:50;" json:"estate_id"`
	Gangs       []Gang `gorm:"foreignKey:DivisionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DivisionMap struct {
	Id   string `json:"id" gorm:"primaryKey;type:varchar(20)"`
	Name string `gorm:"size:50;" json:"name"`
}

func (Division) TableName() string {
	return "divisions"
}

type UpdateDivision struct {
	Code        string `gorm:"size:50;" json:"code"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	UpdatedAt   time.Time
}
