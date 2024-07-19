package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CreateTodoInput struct {
    Title string `json:"title" binding:"required"`
}

func CreateTodo(c *gin.Context) {
    var input CreateTodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    todo := models.Todo{UserID: userID.(uuid.UUID), Title: input.Title}
    if err := config.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    userID, _ := c.Get("userID")
    if err := config.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todo"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todo"})
        return
    }

    var input CreateTodoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo.Title = input.Title
    if err := config.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    var todo models.Todo
    if err := config.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todo"})
        return
    }

    if err := config.DB.Delete(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
