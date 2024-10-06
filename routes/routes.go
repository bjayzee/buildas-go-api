package routes

import (
	"go-crud-app/controllers"
	"go-crud-app/services"
	"go-crud-app/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userService := services.NewUserService()

	userController := controllers.NewUserController(userService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", userController.ListUsers)
	}

	return r
}
