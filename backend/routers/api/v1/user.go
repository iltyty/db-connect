package v1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/services"
	"net/http"
)

const identityKey = "email"

func RegisterUserHandler(c *gin.Context) {
	resp := response.Ctx{C: c}
	v, exists := c.Get("validate")
	if !exists {
		resp.Resp(http.StatusInternalServerError, 1, "validator middleware not found", nil)
	}

	var dto dtos.UserRegistrationDTO
	if err := c.ShouldBind(&dto); err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err := v.(*validator.Validate).Struct(dto)
	if err != nil {
		resp.Resp(http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := services.CreateUser(dto)
	if err != nil {
		resp.Resp(http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.Resp(http.StatusOK, 0, "register success", user)
}

func GetUserHandler(c *gin.Context) {
	resp := response.Ctx{C: c}

	claims := jwt.ExtractClaims(c)
	if email, ok := claims[identityKey]; ok {
		resp.Resp(http.StatusOK, 0, "get user success", email)
		return
	}
	resp.Resp(http.StatusOK, 1, "get user failed", nil)
}
