package schema

type UserRole struct {
	RoleId string `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
