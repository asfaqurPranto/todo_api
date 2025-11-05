package main

import (
	"github.com/gin-gonic/gin"
	"todo_api/internal/handlers"
	"todo_api/internal/config"
	"todo_api/internal/models"
	"todo_api/internal/middleware"
)

func init(){
	//initialize the database and the schema
	config.Connect_Mysql_Server()
	config.DB.AutoMigrate(&models.User{},&models.Todo{})
	
}

func main(){
	router:=gin.Default()

	user_route:=router.Group("/user")
	{
		user_route.POST("/login",handlers.Login)
		user_route.POST("/register",handlers.Register)
		user_route.GET("/info",middleware.Login_Required,handlers.Info)		//will need to add middleware
	}
	todo_route:=router.Group("/todo")
	{
		todo_route.GET("/:todo_id",middleware.Login_Required,handlers.Get_by_id)		//need to add middleware(loginrequired) for all func 
		todo_route.GET("/",middleware.Login_Required,handlers.Get_All)
		todo_route.POST("/",middleware.Login_Required,handlers.Create_Todo)
		todo_route.PUT("/:todo_id",middleware.Login_Required,handlers.Update_Todo)
		todo_route.DELETE("/:todo_id",middleware.Login_Required,handlers.Delete_Todo)
		
	}
	
	router.Run()


}