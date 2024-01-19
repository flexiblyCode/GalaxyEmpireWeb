package routes

import (
	"GalaxyEmpireWeb/api"
	"GalaxyEmpireWeb/api/user"
	"GalaxyEmpireWeb/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoutes()
}

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1.GET("/ping", api.Ping)
	u := v1.Group("/user")
	{
		u.GET("/:id", user.GetUser)
		u.POST("", user.CreateUser)
		u.DELETE("", user.DeleteUser)
	}
	return r
}
