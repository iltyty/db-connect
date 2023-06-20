package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iltyty/db_connect/backend_go/middlewares"
)

var jwt = middlewares.JWTMiddleware

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.ValidatorMiddleware())

	rootRouter := r.Group("api/v1")
	registerAuthRouter(rootRouter)
	registerDataRouter(rootRouter)
	registerUserRouter(rootRouter)

	_ = r.Run("localhost:8000")
	return r
}

func registerAuthRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("auth")
	r.POST("/login", jwt.LoginHandler)
	r.GET("/refresh_token", jwt.RefreshHandler)
	r.POST("/logout", jwt.MiddlewareFunc(), jwt.LogoutHandler)
}

func registerDataRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("test")
	r.GET("", jwt.MiddlewareFunc(), GetTestDataHandler)
}

func registerUserRouter(rootRouter *gin.RouterGroup) {
	r := rootRouter.Group("users")
	r.POST("", RegisterUserHandler)
}
