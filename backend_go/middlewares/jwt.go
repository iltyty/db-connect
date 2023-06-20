package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/models"
	"github.com/iltyty/db_connect/backend_go/response"
	"github.com/iltyty/db_connect/backend_go/utils"
	"time"
)

const identityKey = "email"

var Day = 24 * time.Hour
var JWTMiddleware *jwt.GinJWTMiddleware

func init() {
	var err error
	JWTMiddleware, err = jwt.New(
		&jwt.GinJWTMiddleware{
			Realm:         "test zone",
			Key:           []byte("secret key"),
			Timeout:       15 * Day,
			MaxRefresh:    30 * Day,
			Authenticator: authenticator,
			PayloadFunc:   payloadFunc,
			Unauthorized:  unauthorizedFunc,
			TokenLookup:   "header: Authorization, query: token, cookie: jwt",
			TokenHeadName: "token",
			SendCookie:    true,
			CookieName:    "jwt",
		},
	)
	utils.CheckError(err)
	err = JWTMiddleware.MiddlewareInit()
	utils.CheckError(err)
}

func authenticator(c *gin.Context) (interface{}, error) {
	var dto dtos.UserLoginDTO
	if err := c.ShouldBind(&dto); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	if !models.AuthenticateUser(dto.Email, dto.Password) {
		return nil, jwt.ErrFailedAuthentication
	}

	user, err := models.GetUser(dto.Email)
	return &user, err
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.APIUser); ok {
		return jwt.MapClaims{
			identityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.APIUser{
		Email: claims[identityKey].(string),
	}
}

func unauthorizedFunc(c *gin.Context, code int, message string) {
	resp := response.Ctx{C: c}
	resp.Resp(code, 1, message, nil)
}
