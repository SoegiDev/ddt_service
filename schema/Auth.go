package schema

type SignUpJsonSchema struct {
	Role          []string `json:"role" binding:"required"`
	ApplicationId string   `json:"application_id"`
	CompanyId     string   `json:"company_id"`
	Email         string   `json:"email" binding:"required"`
	Username      string   `json:"username" binding:"required"`
	Password      string   `json:"password" binding:"required"`
}
