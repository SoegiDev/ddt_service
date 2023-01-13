package schema

type SignUpJsonSchema struct {
	Role      []string `json:"role" binding:"required"`
	Email     string   `json:"email" binding:"required"`
	Username  string   `json:"username" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	CompanyId string   `json:"company_id" binding:"required"`
}

type SignInJsonSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
