package schema

type Register struct {
	Role      []string `json:"role" binding:"required"`
	Email     string   `json:"email" binding:"required"`
	Username  string   `json:"username" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	CompanyId string   `json:"company_id" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
