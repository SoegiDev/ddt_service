package schema

import (
	"time"

	"gorm.io/gorm"
)

type AccountRoleApplications struct {
	RoleId    string `gorm:"primaryKey"`
	AccountId string `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (AccountRoleApplications) TableName() string {
	return "account_role_applications"
}
