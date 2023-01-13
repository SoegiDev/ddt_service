package protect

import (
	"ddtservice_agri/controller"
	"ddtservice_agri/middleware"

	"github.com/gin-gonic/gin"
)

// @Host localhost:8080

// @Title Swagger User Service
// @Version 1.0
// @Description This is User Service Test.
// @termsOfService http://swagger.io/terms/

// @contact.name Fajar soegi
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @BasePath /api
func Protected(router *gin.Engine) {
	protectedRoutes := router.Group("/api")

	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	//User
	protectedRoutes.POST("/user/employee", controller.EmployeeAddNew)
	protectedRoutes.GET("/user/:ID", controller.UserGetProfiles)
	// Article
	protectedRoutes.POST("/article", controller.ArticleAddNew)
	protectedRoutes.GET("/article/:ID", controller.ArticleFindById)
	protectedRoutes.GET("/article", controller.AllArticles)
	protectedRoutes.PATCH("/article/:ID/edit", controller.ArticleUpdate)

	// Estate
	protectedRoutes.POST("/estate", controller.EstateAddNew)
	protectedRoutes.PATCH("/estate/:ID/edit", controller.EstateUpdate)

	// Estate
	protectedRoutes.POST("/application", controller.ApplicationAddNew)
	protectedRoutes.PATCH("/application/:ID/edit", controller.ApplicationUpdate)

	// Gang
	protectedRoutes.POST("/gang", controller.GangAddNew)

	// Division
	protectedRoutes.POST("/division", controller.DivisionAddNew)

}
