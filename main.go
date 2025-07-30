package main

import (
    "log"
    "restfulapi/Delivery/controllers"
    "restfulapi/Delivery/routers"
    domain "restfulapi/Domain"
    infrastructure "restfulapi/Infrastructure"
    repositories "restfulapi/Repositories"
    usecases "restfulapi/Usecases"
)

func main() {
    client, err := infrastructure.ConnectMongo("mongodb://localhost:27017")
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    taskCollection := client.Database("taskdb").Collection("tasks")
    userCollection := client.Database("userdb").Collection("users")

    var taskRepo domain.TaskRepository = repositories.NewMongoTaskRepository(taskCollection)
    var userRepo domain.UserRepository = repositories.NewMongoUserRepository(userCollection)

    taskUsecase := usecases.NewTaskUsecase(taskRepo)
    userUsecase := usecases.NewUserUsecase(userRepo)


	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	router := routers.SetupRouter(taskController, userController, userUsecase)

	router.Run(":8080")
}