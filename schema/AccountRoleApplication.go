package schema

type AccountRoleApplication struct {
	RoleId    string `gorm:"primaryKey"`
	AccountId string `gorm:"primaryKey"`
}

func (AccountRoleApplication) TableName() string {
	return "account_role_applications"
}
