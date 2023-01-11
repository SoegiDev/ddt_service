package schema

import (
	"time"

	"gorm.io/gorm"
)

type Estate struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Code        string `gorm:"size:50;" json:"code"`
	Name        string `gorm:"size:50;" json:"name"`
	CompanyId   string `json:"company_id" gorm:"size:50;"`
	Company     Company
	Description string `gorm:"type:text;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool       `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool       `gorm:"type:bool;default:true" json:"status"`
	Divisions   []Division `gorm:"foreignKey:EstateId"`
}

func (Estate) TableName() string {
	return "estates"
}

type UpdateEstate struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CompanyId   string `json:"company_id"`
	Description string `json:"description"`
	UpdatedAt   time.Time
}
