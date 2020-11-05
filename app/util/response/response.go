package response

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}, responseType string, httpStatus int) {
	if responseType == "json" {
		g.Ctx.JSON(httpStatus, Response{
			Code:    code,
			Message: msg,
			Data:    data,
		})

		return
	}
}
