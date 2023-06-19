package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iltyty/db_connect/backend_go/middlewares"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/services"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	rootRouter := r.Group("api/v1")
	registerDataRouter(rootRouter)

	_ = r.Run("localhost:8000")
	return r
}

func registerDataRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("test")
	r.GET(
		"", func(c *gin.Context) {
			resp := response.Ctx{C: c}
			data, err := services.GetTestData()
			if err != nil {
				resp.Resp(http.StatusOK, 1, "get data failed", nil)
			}
			resp.Resp(http.StatusOK, 0, "get data success", data)
		},
	)
}
