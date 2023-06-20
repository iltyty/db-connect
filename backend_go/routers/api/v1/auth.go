package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/middlewares"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/services"
	"net/http"
)

var jwt = middlewares.JWTMiddleware
var LoginHandler = jwt.LoginHandler
var RefreshHandler = jwt.RefreshHandler
var LogoutHandler = jwt.LogoutHandler

func ResetPasswordHandler(c *gin.Context) {
	resp := response.Ctx{C: c}
	v, exists := c.Get("validate")
	if !exists {
		resp.Resp(http.StatusInternalServerError, 1, "validator middleware not found", nil)
	}

	var dto dtos.UserResetPasswordDTO
	if err := c.ShouldBind(&dto); err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err := v.(*validator.Validate).Struct(dto)
	if err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err = services.ResetPassword(dto)
	if err != nil {
		resp.Resp(http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.Resp(http.StatusOK, 0, "reset password success", nil)
}
