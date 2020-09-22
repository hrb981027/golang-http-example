package index

import (
	"api-user-login/app/util/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(200, "Hello World", nil, "json", 200)
}
