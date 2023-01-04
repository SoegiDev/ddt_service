package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id          string `json:"id" gorm:"primaryKey;type:varchar(100)"`
	NickName    string `gorm:"size:255;" json:"nickname"`
	FullName    string `gorm:"size:255;" json:"fullname"`
	PhoneNumber string `gorm:"size:13;" json:"phone_number"`
	Picture     string `gorm:"size:255;" json:"picture"`
	Address     string `gorm:"type:text;" json:"address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Deleted     bool `gorm:"type:bool;default:false" json:"deleted"`
	UserId      string
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
}

func (Employee) TableName() string {
	return "employees"
}
