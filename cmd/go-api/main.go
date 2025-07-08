package main

import (
	"go-api/internal/handler/api"
	user_handler "go-api/internal/handler/user"
	user_svc "go-api/internal/module/user"
	user_repo "go-api/internal/repository/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	userRepo := user_repo.NewUserRepository()
	userSvc := user_svc.NewUserService(userRepo)
	userHandler := user_handler.NewUserHandler(userSvc)

	api.RegisterUserRoutes(router, userHandler)

	router.Run(":8000")
}