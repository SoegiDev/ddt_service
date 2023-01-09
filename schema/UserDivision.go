package schema

type UserDivision struct {
	DivisionId string `gorm:"primaryKey"`
	UserId     string `gorm:"primaryKey"`
}

func (UserDivision) TableName() string {
	return "user_divisions"
}
