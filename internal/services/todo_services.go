package services

import (
	"github.com/gin-gonic/gin"
	"todo_api/internal/middleware"
	"todo_api/internal/models"
	//"net/http"
	"errors"
)

func Get_ALL_Todo(c *gin.Context) ([]models.Todo, error){
	user, err := middleware.UserInfo(c)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized access"})
		return[]models.Todo{}, errors.New("unauth access")
	}
	err = middleware.UserWithTodos(&user)
	if err != nil {
		//c.JSON(200, gin.H{"message": "You have no todos"})
		return []models.Todo{}, errors.New("empty todo")
	}
	
	return user.Todos,nil
}