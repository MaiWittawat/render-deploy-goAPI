package api

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler handler.UserHandler) {
	public := router.Group("/users")
	public.GET("/", handler.GetAllUser)
	public.GET("/:id", handler.GetUser)
	public.POST("/:id", handler.CreateUser)
	public.PATCH("/:id", handler.UpdateUser)
	public.DELETE("/:id", handler.DeleteUser)
}
