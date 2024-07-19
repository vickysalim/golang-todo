package routes

import (
	"todo-app/controllers"
	"todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/register", controllers.Register)
        authRoutes.POST("/login", controllers.Login)
    }

    todoRoutes := r.Group("/todos")
    todoRoutes.Use(middlewares.JWTAuthMiddleware())
    {
        todoRoutes.POST("/", controllers.CreateTodo)
        todoRoutes.GET("/", controllers.GetTodos)
        todoRoutes.GET("/:id", controllers.GetTodoByID)
        todoRoutes.PUT("/:id", controllers.UpdateTodo)
        todoRoutes.DELETE("/:id", controllers.DeleteTodo)
    }

    return r
}
