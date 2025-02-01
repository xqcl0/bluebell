package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
	}

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
