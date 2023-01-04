package schema

import (
	"time"

	"gorm.io/gorm"
)

type UserDivision struct {
	DivisionId string `gorm:"primaryKey"`
	UserId     string `gorm:"primaryKey"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (UserDivision) TableName() string {
	return "user_divisions"
}
