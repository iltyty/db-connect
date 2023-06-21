package response

import "github.com/gin-gonic/gin"

type Ctx struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Ctx) Resp(httpCode, errCode int, msg string, data interface{}) {
	c.C.JSON(httpCode, Response{errCode, msg, data})
}
