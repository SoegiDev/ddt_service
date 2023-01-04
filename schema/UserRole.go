package schema

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	RoleId    string `gorm:"primaryKey"`
	UserId    string `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (UserRole) TableName() string {
	return "user_roles"
}
