package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/middlewares"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/services"
	"github.com/iltyty/db_connect/backend_go/utils"
	"net/http"
)

const expireTime = 5 * 60 // expires after 5 minutes
var jwtMD = middlewares.JWTMiddleware
var LoginHandler = jwtMD.LoginHandler
var RefreshHandler = jwtMD.RefreshHandler
var LogoutHandler = jwtMD.LogoutHandler

func SendEmailCodeHandler(c *gin.Context) {
	resp := response.Ctx{C: c}
	v, exists := c.Get("validate")
	if !exists {
		resp.Resp(http.StatusInternalServerError, 1, "validator middleware not found", nil)
	}

	var dto dtos.UserSendEmailCodeDTO
	if err := c.ShouldBind(&dto); err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err := v.(*validator.Validate).Struct(dto)
	if err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	code, err := services.SendEmailCode(dto)
	if err != nil {
		resp.Resp(http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}

	rdb := utils.RDB
	err = rdb.Set(c, dto.Email+"-code", code, 300).Err()
	if err != nil {
		resp.Resp(http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.Resp(http.StatusOK, 0, "send email success", code)
}

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

func CheckTokenHandler(c *gin.Context) {
	resp := response.Ctx{C: c}
	_, err := jwtMD.CheckIfTokenExpire(c)
	if err != nil {
		resp.Resp(http.StatusOK, 1, "token invalid", nil)
		return
	}
	resp.Resp(http.StatusOK, 0, "token valid", nil)
}
