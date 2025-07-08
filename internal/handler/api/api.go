package api

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userhandler handler.UserHandler) {
	public := router.Group("/users")
	public.GET("/", handler.MiddlewareAuth(), userhandler.GetAllUser)
	public.GET("/:id", handler.MiddlewareAuth(), userhandler.GetUser)
	public.POST("/:id", handler.MiddlewareAuth(), userhandler.CreateUser)
	public.PATCH("/:id", handler.MiddlewareAuth(), userhandler.UpdateUser)
	public.DELETE("/:id", handler.MiddlewareAuth(), userhandler.DeleteUser)
}
