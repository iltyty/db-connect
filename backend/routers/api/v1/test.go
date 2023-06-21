package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/services"
	"net/http"
)

func GetTestDataHandler(c *gin.Context) {
	resp := response.Ctx{C: c}
	data, err := services.GetTestData()
	if err != nil {
		resp.Resp(http.StatusOK, 1, "get data failed", nil)
	}
	resp.Resp(http.StatusOK, 0, "get data success", data)
}
