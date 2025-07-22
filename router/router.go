package router

import (
	"restfulapi/controllers"
	"restfulapi/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Signup)
    router.POST("/login", controllers.Login)
	router.POST("/promote", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.PromoteUser) 
	router.DELETE("/users/:id", middleware.AdminOnly(), controllers.DeleteUser)
	router.GET("/users",   controllers.GetUsers)

	protected := router.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", controllers.GetAllTasks) 
		protected.GET("/:id", controllers.GetTask)
	    protected.DELETE("/:id", middleware.AdminOnly(), controllers.DeleteTask)
	    protected.POST("/", middleware.AdminOnly(), controllers.AddTask)
	    protected.PUT("/:id", middleware.AdminOnly(), controllers.UpdateTask)
	}
	
	return router
}