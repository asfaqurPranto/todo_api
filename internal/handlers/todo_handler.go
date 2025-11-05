package handlers

import (
	//"fmt"
	
	"net/http"
	"strconv"
	"todo_api/internal/config"
	"todo_api/internal/middleware"
	"todo_api/internal/models"
	"todo_api/internal/services"

	"github.com/gin-gonic/gin"
)

func Get_All(c *gin.Context){
	
	todos,err:=services.Get_ALL_Todo(c)
	if err!=nil{
		c.JSON(http.StatusAlreadyReported,gin.H{"message":err.Error()})
		return
	}

	c.JSON(http.StatusAlreadyReported,gin.H{
		"todos":todos,
	
	})


}

func Get_by_id(c *gin.Context){
	todos,err:=services.Get_ALL_Todo(c)
	if err!=nil{
		c.JSON(http.StatusAlreadyReported,gin.H{"message":err.Error()})
		return
	}

	input_id_string:=c.Param("todo_id")
	input_id,_:=strconv.Atoi(input_id_string)
	InputID:=uint(input_id)
	
	for _,todo :=range todos{
		if InputID==todo.ID{
			c.JSON(http.StatusOK,gin.H{
				"todo":todo,
			})
			return

		}
	}

	c.JSON(http.StatusNotFound,gin.H{
				"message":"Todo not found with this id",
		})
	
}

func Create_Todo(c *gin.Context){

	var Data models.CreateTodo
	err:=c.BindJSON(&Data)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	

	var NewTodo models.Todo
	NewTodo.Title=Data.Title
	NewTodo.Description=Data.Description
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	NewTodo.UserID=user.ID
	
	result:=config.DB.Create(&NewTodo)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":result.Error.Error()})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{"message":"todo Created Successfully"})
	

	
}

func Update_Todo(c *gin.Context){
	todos,err:=services.Get_ALL_Todo(c)
	if err!=nil{
		c.JSON(http.StatusAlreadyReported,gin.H{"message":err.Error()})
		return
	}

	input_id_string:=c.Param("todo_id")
	input_id,_:=strconv.Atoi(input_id_string)
	InputID:=uint(input_id)
	
	FoundStatus:=false 
	var todof models.Todo
	for _,todo :=range todos{
		if InputID==todo.ID{
			FoundStatus=true
			todof=todo
			break

		}
	}
	
	if !FoundStatus{
		c.JSON(http.StatusNotFound,gin.H{
				"message":"Todo not found with this id",
		})
	}

	var UpdateReq models.EditTodo
	err=c.BindJSON(&UpdateReq)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	if UpdateReq.Title!=""{
		todof.Title=UpdateReq.Title
		//fmt.Print(UpdateReq.Title)
	}
	if UpdateReq.Description!=""{
		todof.Description=UpdateReq.Description
	}
	
	todof.Completed=UpdateReq.Completed
	

	result:=config.DB.Save(&todof)
	if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }
	c.JSON(http.StatusAccepted,gin.H{"message":"todo updated successfully"})



	

}
