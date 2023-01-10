package public

import (
	"ddtservice_agri/controller"

	"github.com/gin-gonic/gin"
)

// @Host localhost:8000
// @Schemes

// @Title Authentication User
// @Version 1.0
// @Description Authentication User Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Fajar soegi
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

//@BasePath /auth
func Authentication(router *gin.Engine) {
	auth := router.Group("auth")
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
	//Corporation
	auth.POST("/company", controller.CompanyAddNew)
	auth.PATCH("/company/:ID/edit", controller.CompanyUpdate)

	//Application
	auth.POST("/application", controller.ApplicationAddNew)
	auth.PATCH("/application/:ID/edit", controller.ApplicationUpdate)
}
