package route

import (
	"api-user-login/app/controller/index"
	"api-user-login/app/util/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil, "json", 200)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil, "json", 200)
	})


	IndexRoute := engine.Group("/")
	{
		IndexRoute.GET("", index.Index)
	}
}
