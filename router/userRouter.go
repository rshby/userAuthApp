package router

import (
	"github.com/gin-gonic/gin"
	"userAuthApp/handler"
	"userAuthApp/middleware"
)

func CreateUserRoutes(route *gin.RouterGroup, handler *handler.UserHandler) *gin.RouterGroup {
	authMiddlerware := middleware.NewAuthMiddleware()

	r := route.Group("")
	r.POST("/user", handler.CreateUser)
	r.GET("/users", authMiddlerware.Auth(), handler.GetUser)
	return r
}
