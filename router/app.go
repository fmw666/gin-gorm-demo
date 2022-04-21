package router

import (
	"gin-ojsys.cn/config"
	"gin-ojsys.cn/docs"
	"gin-ojsys.cn/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwagger() {
	docs.SwaggerInfo.Title = "OJ SYS API"
	docs.SwaggerInfo.Description = "Gin API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.ServerSetting.PrefixUrl
	docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func Router() *gin.Engine {
	gin.SetMode(config.ServerSetting.RunMode)
	r := gin.Default()

	// Swagger 配置
	go initSwagger()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 路由规则
	apiv1 := r.Group("/api/v1")
	{
		problemAPI := apiv1.Group("/problems")
		problemAPI.Use()
		{
			problemAPI.GET("", service.GetProblemList)
			problemAPI.GET("/:id", service.GetProblemDetail)
		}
		authAPI := apiv1.Group("/auth")
		authAPI.Use()
		{
			authAPI.POST("/register", service.Register)
			authAPI.POST("/login", service.Login)
		}
	}

	return r
}
