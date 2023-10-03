package router

import (
	"github.com/gin-gonic/gin"
	"userAuthApp/handler"
)

func CreateAccountRoutes(router *gin.RouterGroup, handler *handler.AccountHandler) *gin.RouterGroup {
	r := router.Group("")
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	return r
}
