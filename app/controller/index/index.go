package index

import (
	"github.com/gin-gonic/gin"
	"golang-http-example/app/util/response"
)

func Index(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(200, "你好 世界", nil, "json", 200)
}
