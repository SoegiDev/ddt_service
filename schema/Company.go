package schema

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Code        string `json:"code" gorm:"size:10;unique"`
	Name        string `gorm:"size:100;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	PhoneNumber string `gorm:"size:15;" json:"phone_number"`
	Fax         string `gorm:"size:15;" json:"fax"`
	Sector      string `gorm:"type:text;" json:"sector"`
	Domain      string `gorm:"size:20;" json:"domain"`
	Address     string `gorm:"type:text;" json:"address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool       `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool       `gorm:"type:bool;default:true" json:"status"`
	Employees   []Employee `gorm:"foreignKey:CompanyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Estates     []Estate   `gorm:"foreignKey:CompanyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Company) TableName() string {
	return "companies"
}

type UpdateCompany struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Fax         string `json:"fax"`
	Sector      string `json:"sector"`
	Domain      string `json:"domain"`
	Address     string `json:"address"`
	UpdatedAt   time.Time
}
