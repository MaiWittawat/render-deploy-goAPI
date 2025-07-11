package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}
