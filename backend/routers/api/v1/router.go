package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iltyty/db_connect/backend_go/conf"
	"github.com/iltyty/db_connect/backend_go/middlewares"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.ValidatorMiddleware())

	rootRouter := r.Group("api/v1")
	registerAuthRouter(rootRouter)
	registerDataRouter(rootRouter)
	registerUserRouter(rootRouter)

	s := conf.ServerConf
	_ = r.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
	return r
}

func registerAuthRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("auth")
	r.POST("login", LoginHandler)
	r.POST("logout", jwtMD.MiddlewareFunc(), LogoutHandler)
	r.GET("check", CheckTokenHandler)
	r.GET("refresh_token", RefreshHandler)
	r.POST("reset_password", ResetPasswordHandler)
	r.POST("email_code", SendEmailCodeHandler)
}

func registerDataRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("test")
	r.GET("", jwtMD.MiddlewareFunc(), GetTestDataHandler)
}

func registerUserRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("user")
	r.POST("", RegisterUserHandler)
	r.GET("", jwtMD.MiddlewareFunc(), GetUserHandler)
}
