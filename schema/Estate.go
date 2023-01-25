package schema

import (
	"time"

	"gorm.io/gorm"
)

type Estate struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Code        string `gorm:"size:20;unique" json:"code"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool       `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool       `gorm:"type:bool;default:true" json:"status"`
	CompanyCode string     `json:"company_code" gorm:"size:20;"`
	Company     Company    `gorm:"foreignKey:CompanyCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Divisions   []Division `gorm:"foreignKey:EstateCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Estate) TableName() string {
	return "estates"
}

type UpdateEstate struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CompanyCode string `json:"company_code"`
	Description string `json:"description"`
	UpdatedAt   time.Time
}
