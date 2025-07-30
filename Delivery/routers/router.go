package routers

import (
    "restfulapi/Delivery/controllers"
    infrastructure "restfulapi/Infrastructure"
    "github.com/gin-gonic/gin"
	usecases "restfulapi/Usecases"
)

func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController, userUsecase *usecases.UserUsecase) *gin.Engine {
    r := gin.Default()

    r.POST("/signup", userCtrl.Signup)
    r.POST("/login", userCtrl.Login)

    auth := r.Group("/")
    auth.Use(infrastructure.AuthMiddleware())

    auth.GET("/tasks", taskCtrl.GetAllTasks)
    auth.GET("/tasks/:id", taskCtrl.GetTask)

    admin := auth.Group("/")
    admin.Use(infrastructure.AdminOnly(userUsecase))

    admin.GET("/users", userCtrl.GetUsers)
    admin.POST("/users/promote", userCtrl.PromoteUser)
    admin.DELETE("/users/:id", userCtrl.DeleteUser)

    auth.PUT("/tasks/:id", taskCtrl.UpdateTask)
    auth.DELETE("/tasks/:id", taskCtrl.DeleteTask)
    auth.POST("/tasks", taskCtrl.AddTask)

    return r
}