package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id           string `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Email        string `json:"email" gorm:"primaryKey;type:varchar(100)"`
	Username     string `json:"username" gorm:"primaryKey;type:varchar(100)"`
	NickName     string `gorm:"size:255;" json:"nickname"`
	CompanyCode  string `gorm:"size:10;" json:"company_code"`
	CompanyName  string `gorm:"size:10;" json:"company_name"`
	FullName     string `gorm:"size:255;" json:"fullname"`
	PhoneNumber  string `gorm:"size:13;" json:"phone_number"`
	Picture      string `gorm:"size:255;" json:"picture"`
	Address      string `gorm:"type:text;" json:"address"`
	Domain       string `gorm:"type:text;" json:"domain"`
	Department   string `gorm:"type:text;" json:"department"`
	OfficeNumber string `gorm:"type:text;" json:"office_number"`
	Expired      int    `gorm:"type:int;" json:"expired_time"`
	CreatedBy    string `gorm:"size:13;" json:"created_by"`
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Deleted      bool `gorm:"type:bool;default:false" json:"deleted"`
	UserId       string
	IsActive     bool `gorm:"type:bool;default:true" json:"status"`
}

type EmployeeAdd struct {
	Id          string `json:"id"`
	NickName    string `json:"nickname"`
	FullName    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Picture     string `json:"picture"`
	Address     string `json:"address"`
	UserId      string `json:"user_id"`
}

func (Employee) TableName() string {
	return "employees"
}
