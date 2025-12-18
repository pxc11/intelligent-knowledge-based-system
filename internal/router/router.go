package router

import (
	"ikbs/internal/controller"
	"ikbs/lib/JWT"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/login", controller.Login)

	auth := api.Group("")
	auth.Use(JWT.JWTAuthMiddleware())
	auth.GET("/getUserInfo", controller.GetUserInfo)
}
